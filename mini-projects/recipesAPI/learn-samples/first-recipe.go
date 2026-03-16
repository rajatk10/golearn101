package main

import (
	_ "framework-api/docs"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Recipe struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Tags         []string  `json:"tags"`
	Ingredients  []string  `json:"ingredients"`
	Instructions []string  `json:"instructions"`
	PublishedAt  time.Time `json:"publishedAt"`
}

var recipes []Recipe

func init() {
	//It is loaded even before main() is called and a file can have multiple init() functions
	log.Println("Initializing the init() function...")
	recipes = make([]Recipe, 0)

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
	engine.GET("/recipes", getRecipes)
	engine.GET("/recipe/:id", getRecipeById)
	engine.POST("/recipe", insertRecipe)
	engine.PUT("/recipe/:id", updateRecipeById)
	engine.DELETE("/recipe/:id", deleteRecipeById)

	//Swagger Route
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//start the server
	if err := engine.Run(":8089"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}

// Swagger Documentation
// getRecipes godoc
// @Summary Get All Recipes
// @Description Gets all the recipes from the in-memory store
// @Tags recipes
// @Accept json
// @Produce json
// @Success 200 {array} main.Recipe
// @Failure 400 {object} map[string]string "error"
// @Router /recipes [get]
func getRecipes(c *gin.Context) {
	//c.JSON(http.StatusOK, recipes)
	c.JSON(http.StatusOK, recipes)

}

// Swagger Documentation
// getRecipesById godoc
// @Summary Get Recipe by ID
// @Description Gets the recipe by ID from the in-memory store
// @Tags recipes
// @Accept json
// @Produce json
// @Param id path string true "Recipe ID"
// @Success 200 {object} Recipe
// @Failure 400 {object} map[string]string "error"
// @Router /recipe/{id} [get]
func getRecipeById(c *gin.Context) {
	recipeId := c.Param("id")
	index := -1
	log.Printf("Fetching Recipe with id = %v\n", recipeId)
	for i, r := range recipes {
		if r.ID == recipeId {
			index = i
			break
		}
	}
	if index == -1 {
		log.Printf("Failed to fetch recipe as ID is invalid")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Recipe ID is invalid"})
		return
	}
	c.JSON(http.StatusOK, recipes[index])
}

// Swagger Documentation
// insertRecipe godoc
// @Summary Create a new recipe
// @Description Adds a new recipe to the in-memory store
// @Tags recipes
// @Accept json
// @Produce json
// @Param recipe body Recipe true "Recipe Data"
// @Success 201 {object} Recipe
// @Failure 400 {object} map[string]string "error"
// @Router /recipe [post]
func insertRecipe(c *gin.Context) {
	var recipe Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	recipe.ID = xid.New().String() //xid is package which generates unique id - more in line with mongodb id
	recipe.PublishedAt = time.Now()
	recipes = append(recipes, recipe)
	c.JSON(http.StatusCreated, recipe)
}

// Swagger Documentation
// updateRecipeById godoc
// @Summary UPDATE Recipe by ID
// @Description Update the recipe by ID from the in-memory store
// @Tags recipes
// @Accept json
// @Produce json
// @Param id path string true "Recipe ID"
// @Success 200 {object} Recipe
// @Failure 400 {object} map[string]string "error"
// @Router /recipe/{id} [put]
func updateRecipeById(c *gin.Context) {
	recipeId := c.Param("id")
	log.Printf("Updating recipe with id: %v", recipeId)
	index := -1
	for i, r := range recipes {
		if r.ID == recipeId {
			index = i
			break
		}
	}
	if index == -1 {
		log.Printf("Recipe not found with id: %v", recipeId)
		c.JSON(http.StatusNotFound, gin.H{"error": "Recipe not found"})
		return
	}
	var updateRecipe Recipe
	if err := c.ShouldBindJSON(&updateRecipe); err != nil {
		log.Printf("Error while binding id: %v with JSON: %v", recipeId, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updateRecipe.ID = recipeId
	updateRecipe.PublishedAt = time.Now()
	recipes[index] = updateRecipe
	c.JSON(http.StatusOK, updateRecipe)
}

// Swagger Documentation
// deleteRecipeById godoc
// @Summary Delete Recipe by ID from Recipes
// @Description Delete the recipe by ID from the in-memory store
// @Tags recipes
// @Accept json
// @Produce json
// @Param id path string true "Recipe ID"
// @Success 200 {object} Recipe
// @Failure 400 {object} map[string]string "error"
// @Router /recipe/{id} [delete]
func deleteRecipeById(c *gin.Context) {
	recipeId := c.Param("id")
	log.Printf("Deleting recipe with id: %v", recipeId)
	index := -1
	var recipeName string
	for i, r := range recipes {
		if r.ID == recipeId {
			index = i
			recipeName = r.Name
			break
		}
	}
	if index == -1 {
		log.Printf("Recipe not found with ID = %v", recipeId)
		c.JSON(http.StatusNotFound, gin.H{"error": "Recipe not found"})
		return
	}
	recipes = append(recipes[:index], recipes[index+1:]...)
	c.JSON(http.StatusOK, gin.H{
		"recipe":  recipeName,
		"message": "Recipe deleted successfully",
	})
}
