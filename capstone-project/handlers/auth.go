package handlers

import (
	"net/http"
	"strings"

	"github.com/MicahParks/keyfunc/v2"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

type AuthHandler struct{}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func InitCognitoJWKS(region, userPoolID string) (*keyfunc.JWKS, error) {
	jwksURL := "https://cognito-idp." + region + ".amazonaws.com/" + userPoolID + "/.well-known/jwks.json"
	//Fetch public key from AWS
	jwks, err := keyfunc.Get(jwksURL, keyfunc.Options{})
	if err != nil {
		zap.L().Error("Error fetching public key from AWS", zap.Error(err))
		return nil, err
	}
	zap.L().Info("Public key fetched from AWS")
	return jwks, nil
}

func (h *AuthHandler) AuthMiddleware(jwks *keyfunc.JWKS, expectedIssuer, expectedClientID string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		//Now parse the token using keys fetched from AWS Cognito
		token, err := jwt.Parse(tokenString, jwks.Keyfunc)
		if err != nil || !token.Valid {
			zap.L().Warn("Error parsing token", zap.Error(err))
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			zap.L().Warn("Invalid token claims")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}
		if claims["iss"] != expectedIssuer {
			zap.L().Warn("Invalid issuer")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid issuer"})
			return
		}
		if claims["client_id"] != expectedClientID && claims["aud"] != expectedClientID {
			zap.L().Warn("Invalid client_id")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid client_id"})
			return
		}

		//Ensure access token
		if claims["token_use"] != "access" {
			zap.L().Warn("Invalid token use")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token use"})
			return
		}
		userID := claims["sub"].(string)
		zap.L().Info("User authenticated", zap.String("user_id", userID))
		c.Set("userID", userID)
		c.Next()
	}
}
