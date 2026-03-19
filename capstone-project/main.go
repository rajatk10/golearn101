package main

import (
	"context"
	"framework-api/utils"

	//"crypto/tls"
	"framework-api/handlers"
	"os"

	_ "framework-api/docs"
	"net/http"
	"time"

	"github.com/MicahParks/keyfunc/v2"
	"github.com/elastic/go-elasticsearch/v9"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
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
var redisClient *redis.Client

// From RecipeHandler
var recipeHandler *handlers.RecipeHandler

// From AuthHandler
var authHandler *handlers.AuthHandler

// For AWS Service
var region string
var userPoolID string
var clientID string
var issuer string
var jwks *keyfunc.JWKS
var logger *zap.Logger
var loggerCleanup func()

func init() {
	logger, loggerCleanup, err = utils.InitLogger()
	if err != nil {
		panic(err)
	}
	zap.ReplaceGlobals(logger)
	logger.Info("Initializing the init() function...")
	ctx = context.Background()
	//MONGODB_HOST := "mongodb://localhost:27017"
	mongoDBUri := os.Getenv("MONGODB_URI")
	if mongoDBUri == "" {
		logger.Fatal("MONGODB_URI is not set")
	}
	//client, err = mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017"))
	client, err = mongo.Connect(options.Client().ApplyURI(mongoDBUri))
	if err != nil {
		logger.Fatal("Failed to connect to mongodb", zap.Error(err))
	}
	if err = client.Ping(ctx, nil); err != nil {
		logger.Fatal("Failed to ping mongodb", zap.Error(err))
	}
	logger.Info("Connected to mongodb client")
	collectionRecipes = client.Database("recipeDB").Collection("recipes")
	logger.Info("Connected to mongodb collection recipes")

	//Setup Redis
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	if err := redisClient.Ping(ctx).Err(); err != nil {
		logger.Fatal("Failed to connect to redis", zap.Error(err))
	}
	//status := redisClient.Ping
	//log.Printf("Redis Status: %v", status)
	logger.Info("Connected to redis client")
	//Setup Elasticsearch
	elasticSearchUri := os.Getenv("ELASTICSEARCH_URI")
	elasticSearchCfg := elasticsearch.Config{
		Addresses: []string{
			elasticSearchUri,
		},
	}
	elasticsearchClient, err := elasticsearch.NewClient(elasticSearchCfg)
	if err != nil {
		logger.Fatal("Failed to connect to elasticsearch", zap.Error(err))
	} else {
		logger.Info("Connected to elasticsearch")
	}

	//Setup AWS
	region = os.Getenv("AWS_REGION")
	userPoolID = os.Getenv("AWS_USER_POOL_ID")
	clientID = os.Getenv("AWS_CLIENT_ID")
	issuer = os.Getenv("AWS_ISSUER")
	jwks, err = handlers.InitCognitoJWKS(region, userPoolID)
	if err != nil {
		logger.Fatal("Failed to initialize AWS Cognito JWKS", zap.Error(err))
	}
	recipeHandler = handlers.NewRecipesHandler(ctx, collectionRecipes, redisClient, elasticsearchClient)
	logger.Info("Initialize Authentication Handler")
	authHandler = handlers.NewAuthHandler()
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
	if loggerCleanup != nil {
		defer loggerCleanup()
	}
	logger.Info("Initializing server-main now...")
	engine := gin.Default()
	engine.LoadHTMLGlob("static/*.html")
	engine.Static("/static", "static")
	engine.StaticFile("/favicon.ico", "static/images/cooking.png")
	//Setting up CORS
	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "OPTIONS", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))
	//Check Server API status
	engine.GET("/", recipeHandler.HomePageHandler)
	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	//RECIPE APIs
	engine.GET("/recipes", recipeHandler.GetRecipes)
	engine.GET("/recipe/:id", recipeHandler.GetRecipeById)
	engine.GET("/recipes/search", recipeHandler.SearchRecipeInElasticStore)

	//Swagger Route
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//AUTH Middleware (Protects the routes below)
	//engine.Use(authHandler.AuthMiddleware(jwks, issuer, clientID))
	engine.POST("/recipe", recipeHandler.InsertRecipe)
	engine.PATCH("/recipe/:id", recipeHandler.UpdateRecipeById)
	engine.DELETE("/recipe/:id", recipeHandler.DeleteRecipeById)
	//Setting up CORS

	//start the server
	if err := engine.Run(":8088"); err != nil {
		logger.Fatal("Failed to start server", zap.Error(err))
	}

	//engine.RunTLS(":443", "certs/localhost.crt", "certs/localhost.key")

}
