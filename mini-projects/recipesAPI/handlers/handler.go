package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"framework-api/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type RecipeHandler struct {
	collection  *mongo.Collection
	ctx         context.Context
	redisClient *redis.Client
}

//Constructor

func NewRecipesHandler(ctx context.Context, collection *mongo.Collection, redisClient *redis.Client) *RecipeHandler {
	return &RecipeHandler{
		collection:  collection,
		ctx:         ctx,
		redisClient: redisClient,
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
func (h *RecipeHandler) GetRecipes(c *gin.Context) {
	log.Print("Fetching all recipes")
	//Check Redis first
	val, err := h.redisClient.Get(h.ctx, "recipes").Result()
	//val is string returned by redis so this would need to be unmarshalled.

	if err == redis.Nil {
		log.Println("Request sent to MongoDB")
		cur, err := h.collection.Find(h.ctx, bson.M{})
		if err != nil {
			log.Printf("Failed to fetch recipes from redis: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch recipes"})
			return
		}
		defer cur.Close(h.ctx)

		//Hold Results
		dbRecipes := make([]models.Recipe, 0)
		for cur.Next(h.ctx) {
			var recipe models.Recipe
			if err := cur.Decode(&recipe); err != nil {
				log.Printf("Failed to decode recipe: %v", err)
				continue
			}
			dbRecipes = append(dbRecipes, recipe)
		}
		if len(dbRecipes) == 0 {
			log.Printf("Failed to fetch recipes as no recipes found")
			c.JSON(http.StatusNotFound, gin.H{"error": "No recipes found"})
			return
		}
		//update redis cache
		data, _ := json.Marshal(dbRecipes)
		log.Println("Storing recipes in redis")
		h.redisClient.Set(h.ctx, "recipes", string(data), 0)
		c.JSON(http.StatusOK, dbRecipes)
	} else if err != nil {
		log.Printf("Failed to fetch recipes: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch recipes"})
		return
	} else {
		log.Println("Found recipes in redis ")
		recipes := make([]models.Recipe, 0)
		json.Unmarshal([]byte(val), &recipes)
		c.JSON(http.StatusOK, recipes)
	}
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
func (h *RecipeHandler) GetRecipeById(c *gin.Context) {
	recipeId := c.Param("id")
	log.Printf("Updating recipe with id: %v", recipeId)
	var recipe models.Recipe
	objectId, err := bson.ObjectIDFromHex(recipeId)
	if err != nil {
		log.Printf("Failed to convert ID to ObjectID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	err = h.collection.FindOne(h.ctx, bson.M{"_id": objectId}).Decode(&recipe)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Printf("Failed to find recipe with id: %v", recipeId)
			c.JSON(http.StatusNotFound, gin.H{"error": "Recipe not found"})
		} else {
			log.Printf("Failed to find recpie, database error")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find recipe"})
		}
		return
	}
	c.JSON(http.StatusOK, recipe)
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
func (h *RecipeHandler) InsertRecipe(c *gin.Context) {
	var Recipe models.Recipe
	if err := c.ShouldBindJSON(&Recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	Recipe.ID = bson.NewObjectID()
	Recipe.PublishedAt = time.Now()
	_, err := h.collection.InsertOne(h.ctx, Recipe)
	if err != nil {
		log.Printf("Failed to insert recipe: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert recipe"})
		return
	}
	//Invalidate cache
	h.redisClient.Del(h.ctx, "recipes")
	c.JSON(http.StatusCreated, Recipe)
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
func (h *RecipeHandler) UpdateRecipeById(c *gin.Context) {
	recipeId := c.Param("id")
	var recipe models.Recipe
	log.Printf("Updating recipe with id: %v", recipeId)
	if err := c.ShouldBindJSON(&recipe); err != nil {
		log.Printf("Failed to bind JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//recipe.PublishedAt = time.Now()
	objectId, err := bson.ObjectIDFromHex(recipeId)
	if err != nil {
		log.Printf("Failed to convert id: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	updatedRecipe := bson.M{
		"$set": bson.M{
			"name":         recipe.Name,
			"tags":         recipe.Tags,
			"ingredients":  recipe.Ingredients,
			"instructions": recipe.Instructions,
		},
	}

	//Execute update
	_, err = h.collection.UpdateOne(h.ctx, bson.M{"_id": objectId}, updatedRecipe)
	if err != nil {
		log.Printf("Failed to update the recipe %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update the recipe"})
		return
	}
	msg := fmt.Sprintf("Recipe Successfully Updated %v", recipeId)
	//After update invalidate cache
	h.redisClient.Del(h.ctx, "recipes")
	c.JSON(http.StatusOK, gin.H{"message": msg})
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
func (h *RecipeHandler) DeleteRecipeById(c *gin.Context) {
	recipeId := c.Param("id")
	log.Printf("Deleting recipe with id: %v", recipeId)
	objectId, err := bson.ObjectIDFromHex(recipeId)
	if err != nil {
		log.Printf("Failed to fetch valid objectID from id %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Recipe not found"})
		return
	}
	res, err := h.collection.DeleteOne(h.ctx, bson.M{"_id": objectId})
	if err != nil {
		log.Printf("Failed to delete recipe %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Recipe not found"})
		return
	}
	if res.DeletedCount == 0 {
		log.Printf("Failed to delete recipe")
		c.JSON(http.StatusNotFound, gin.H{"error": "Recipe not found"})
		return
	}
	//After delete - invalidate cache
	h.redisClient.Del(h.ctx, "recipes")
	c.JSON(http.StatusOK, gin.H{"message": "Recipe deleted successfully"})
}
