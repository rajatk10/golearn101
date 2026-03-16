package main

import (
	"context"
	"framework-api/handlers"
	"os"

	_ "framework-api/docs"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Recipe struct {
	ID           bson.ObjectID `json:"id" bson:"_id"`
	Name         string        `json:"name" bson:"name"`
	Tags         []string      `json:"tags" bson:"tags"`
	Ingredients  []string      `json:"ingredients" bson:"ingredients"`
	Instructions []string      `json:"instructions" bson:"instructions"`
	PublishedAt  time.Time     `json:"publishedAt" bson:"publishedAt"`
}

// mongodb connection
var ctx context.Context
var err error
var client *mongo.Client
var collectionRecipes *mongo.Collection
var collectionUsers *mongo.Collection
var redisClient *redis.Client

// From RecipeHandler
var recipeHandler *handlers.RecipeHandler

// From AuthHandler
var authHandler *handlers.AuthHandler

func init() {
	log.Println("Initializing the init() function...")
	ctx = context.Background()
	//MONGODB_HOST := "mongodb://localhost:27017"
	mongoDBUri := os.Getenv("MONGODB_URI")
	if mongoDBUri == "" {
		log.Fatal("MONGODB_URI is not set")
		os.Exit(1)
	}
	//client, err = mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017"))
	client, err = mongo.Connect(options.Client().ApplyURI(mongoDBUri))
	if err != nil {
		log.Fatal("Failed to connect to mongodb", err)
	}
	if err = client.Ping(ctx, nil); err != nil {
		log.Fatal("Failed to ping mongodb", err)
	}
	log.Print("Connected to mongodb client")
	collectionRecipes = client.Database("recipeDB").Collection("recipes")
	log.Print("Connected to mongodb collection recipes")

	//Setup Redis
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	if err := redisClient.Ping(ctx).Err(); err != nil {
		log.Fatal("Failed to connect to redis", err)
	}
	//status := redisClient.Ping
	//log.Printf("Redis Status: %v", status)
	log.Print("Connected to redis client")
	recipeHandler = handlers.NewRecipesHandler(ctx, collectionRecipes, redisClient)
	log.Println("Initialize Authentication Handler")
	collectionUsers = client.Database("recipeDB").Collection("users")
	authHandler = handlers.NewAuthHandler(ctx, collectionUsers)
}

// Swagger Documentation
// @title Recipe API
// @version 1.0
// @description This is a sample recipe management API.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8088
// @BasePath /

func main() {
	log.Println("Initializing server...")
	engine := gin.Default()

	//Check Server API status
	engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Welcome to Recipe APIs")
	})
	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	//RECIPE APIs
	engine.GET("/recipes", recipeHandler.GetRecipes)
	engine.GET("/recipe/:id", recipeHandler.GetRecipeById)
	engine.POST("/signup", authHandler.SignUpHandler)
	//engine.POST("/recipe", recipeHandler.InsertRecipe)
	//engine.PATCH("/recipe/:id", recipeHandler.UpdateRecipeById)
	//engine.DELETE("/recipe/:id", recipeHandler.DeleteRecipeById)

	//Swagger Route
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//AUTH APIs
	engine.POST("/signin", authHandler.SignInHandler)

	//AUTH Middleware
	engine.Use(authHandler.AuthMiddleware())
	engine.POST("/recipe", recipeHandler.InsertRecipe)
	engine.PATCH("/recipe/:id", recipeHandler.UpdateRecipeById)
	engine.DELETE("/recipe/:id", recipeHandler.DeleteRecipeById)

	//start the server
	if err := engine.Run(":8088"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	//engine.RunTLS(":443", "certs/localhost.crt", "certs/localhost.key")

}
