package handlers

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"framework-api/models"
	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"

	"net/http"
	"os"
	"time"
)

type AuthHandler struct {
	collection *mongo.Collection
	ctx        context.Context
}
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type JWTOutput struct {
	Token   string    `json:"token"`
	Expires time.Time `json:"expires"`
}

func NewAuthHandler(ctx context.Context, collection *mongo.Collection) *AuthHandler {
	return &AuthHandler{
		collection: collection,
		ctx:        ctx,
	}
}

func (h *AuthHandler) SignInHandler(c *gin.Context) {
	var userCreds models.UserCreds
	if err := c.ShouldBindJSON(&userCreds); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("Processing credentials for user : %v", userCreds.Username)
	//log.Printf("UserCreds: %v", userCreds)
	//if userCreds.Username != "admin" || userCreds.Password != "password" {
	//	log.Printf("Invalid credentials")
	//	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	//	return
	//}

	//Check user credentials
	hasher := sha256.New()
	hasher.Write([]byte(userCreds.Password))
	givenHashedPassword := hex.EncodeToString(hasher.Sum(nil))
	curr := h.collection.FindOne(h.ctx, bson.M{"username": userCreds.Username, "password": givenHashedPassword})
	if curr.Err() != nil {
		log.Printf("User not found or invalid credentials: %v", curr.Err())
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	expirationTime := time.Now().Add(30 * time.Minute)
	claims := &Claims{
		Username: userCreds.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	//log.Printf("Claims contained in AuthHandler: %v", claims)
	log.Printf("Generating token for user: %v", userCreds.Username)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//log.Printf("Token Go struct: %v", token)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		log.Printf("Error generating token: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}
	//log.Printf("TokenString: %v", tokenString)
	jwtOutput := &JWTOutput{
		Token:   tokenString,
		Expires: expirationTime,
	}
	//log.Printf("JWTOutput: %v", jwtOutput)
	c.JSON(http.StatusOK, jwtOutput)

}

func (h *AuthHandler) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenValue := c.GetHeader("Authorization")
		log.Printf("TokenValue Parsed from Header: %v", tokenValue)
		claims := &Claims{}
		tokenStr, err := jwt.ParseWithClaims(tokenValue, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil {
			log.Printf("Error parsing token: %v", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}
		//log.Printf("TokenStr in AuthHandler: %v", tokenStr)
		if !tokenStr.Valid {
			log.Printf("Invalid token: %v", tokenStr)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}
		c.Next()
	}
}

func (h *AuthHandler) SignUpHandler(c *gin.Context) {
	log.Println("Processing SignUpHandler")
	var user models.UserProfile
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	curr := h.collection.FindOne(h.ctx, bson.M{"username": user.Username})
	if curr.Err() == nil {
		log.Printf("User already exists: %v", user.Username)
		c.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
		return
	}

	hasher := sha256.New()
	hasher.Write([]byte(user.Password))
	hashedPassword := hex.EncodeToString(hasher.Sum(nil))

	newUserData := bson.M{
		"name":     user.Name,
		"username": user.Username,
		"password": hashedPassword,
		"email":    user.Email,
		"age":      user.Age,
		"gender":   user.Gender,
	}

	_, err := h.collection.InsertOne(h.ctx, newUserData)
	if err != nil {
		log.Print("Error inserting new user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	msg := fmt.Sprintf("User %s created successfully", user.Username)
	c.JSON(http.StatusOK, gin.H{"message": msg})
}
