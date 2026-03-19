package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"framework-api/models"
	"net/http"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v9"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type RecipeHandler struct {
	collection    *mongo.Collection
	ctx           context.Context
	redisClient   *redis.Client
	elasticClient *elasticsearch.Client
}

//Constructor

func NewRecipesHandler(ctx context.Context, collection *mongo.Collection, redisClient *redis.Client, elasticClient *elasticsearch.Client) *RecipeHandler {
	return &RecipeHandler{
		collection:    collection,
		ctx:           ctx,
		redisClient:   redisClient,
		elasticClient: elasticClient,
	}
}

func (h *RecipeHandler) HomePageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", nil)
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
	zap.L().Info("Fetching all recipes")
	//Check Redis first
	val, err := h.redisClient.Get(h.ctx, "recipes").Result()
	//val is string returned by redis so this would need to be unmarshalled.

	if err == redis.Nil {
		zap.L().Info("Request sent to MongoDB")
		cur, err := h.collection.Find(h.ctx, bson.M{})
		if err != nil {
			zap.L().Error("Failed to fetch recipes from redis", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch recipes"})
			return
		}
		defer cur.Close(h.ctx)

		//Hold Results
		dbRecipes := make([]models.Recipe, 0)
		for cur.Next(h.ctx) {
			var recipe models.Recipe
			if err := cur.Decode(&recipe); err != nil {
				zap.L().Error("Failed to decode recipe", zap.Error(err))
				continue
			}
			dbRecipes = append(dbRecipes, recipe)
		}
		if len(dbRecipes) == 0 {
			zap.L().Warn("No recipes found")
			c.JSON(http.StatusNotFound, gin.H{"error": "No recipes found"})
			return
		}
		//update redis cache
		data, _ := json.Marshal(dbRecipes)
		zap.L().Info("Storing recipes in redis")
		h.redisClient.Set(h.ctx, "recipes", string(data), 0)
		c.JSON(http.StatusOK, dbRecipes)
	} else if err != nil {
		zap.L().Error("Failed to fetch recipes", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch recipes"})
		return
	} else {
		zap.L().Info("Found recipes in redis")
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
	zap.L().Info("Fetching recipe by id", zap.String("recipe_id", recipeId))

	val, err := h.redisClient.Get(h.ctx, "recipe:"+recipeId).Result()
	if err == nil {
		zap.L().Info("Found recipe in redis", zap.String("recipe_id", recipeId))
		var recipe models.Recipe
		json.Unmarshal([]byte(val), &recipe)
		c.JSON(http.StatusOK, recipe)
		return
	}
	zap.L().Info("Recipe not found in redis, fetching from DB", zap.String("recipe_id", recipeId))
	var recipe models.Recipe
	objectId, err := bson.ObjectIDFromHex(recipeId)
	if err != nil {
		zap.L().Error("Failed to convert ID to ObjectID", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	err = h.collection.FindOne(h.ctx, bson.M{"_id": objectId}).Decode(&recipe)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			zap.L().Warn("Recipe not found", zap.String("recipe_id", recipeId))
			c.JSON(http.StatusNotFound, gin.H{"error": "Recipe not found"})
		} else {
			zap.L().Error("Failed to find recipe, database error", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find recipe"})
		}
		return
	}
	//update redis cache
	data, _ := json.Marshal(recipe)
	zap.L().Info("Storing recipe in redis", zap.String("recipe_id", recipeId))
	h.redisClient.Set(h.ctx, "recipe:"+recipeId, string(data), 0)
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
		zap.L().Error("Failed to insert recipe", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert recipe"})
		return
	}
	//Invalidate cache
	h.redisClient.Del(h.ctx, "recipes")
	//Add recipe to elastic store
	err = h.insertRecipeInElasticstore(Recipe)
	if err != nil {
		zap.L().Error("Failed to insert recipe in elastic store", zap.Error(err))
	}
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
	objectId, err := bson.ObjectIDFromHex(recipeId)
	if err != nil {
		zap.L().Error("Failed to convert id", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var updateData bson.M
	zap.L().Info("Updating recipe", zap.String("recipe_id", recipeId))
	if err := c.ShouldBindJSON(&updateData); err != nil {
		zap.L().Error("Failed to bind JSON", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//recipe.PublishedAt = time.Now()
	delete(updateData, "_id")
	delete(updateData, "publishedAt")

	updatedRecipe := bson.M{
		"$set": updateData,
	}

	//Execute update
	_, err = h.collection.UpdateOne(h.ctx, bson.M{"_id": objectId}, updatedRecipe)
	if err != nil {
		zap.L().Error("Failed to update recipe", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update the recipe"})
		return
	}
	msg := fmt.Sprintf("Recipe Successfully Updated %v", recipeId)
	//After update invalidate cache
	h.redisClient.Del(h.ctx, "recipe:"+recipeId)
	h.redisClient.Del(h.ctx, "recipes") //Invalidate all recipes cache

	//Update in Elastic store
	var recipe models.Recipe
	err = h.collection.FindOne(h.ctx, bson.M{"_id": objectId}).Decode(&recipe)
	if err != nil {
		zap.L().Error("Failed to find recipe in DB", zap.Error(err))
	}
	//Update recipe in elastic store
	err = h.insertRecipeInElasticstore(recipe)
	if err != nil {
		zap.L().Error("Failed to insert recipe in elastic store", zap.Error(err))
	}
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
	zap.L().Info("Deleting recipe", zap.String("recipe_id", recipeId))
	objectId, err := bson.ObjectIDFromHex(recipeId)
	if err != nil {
		zap.L().Error("Failed to parse recipe id", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Recipe not found"})
		return
	}
	res, err := h.collection.DeleteOne(h.ctx, bson.M{"_id": objectId})
	if err != nil {
		zap.L().Error("Failed to delete recipe", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Recipe not found"})
		return
	}
	if res.DeletedCount == 0 {
		zap.L().Warn("Recipe not found to delete", zap.String("recipe_id", recipeId))
		c.JSON(http.StatusNotFound, gin.H{"error": "Recipe not found"})
		return
	}
	//After delete - invalidate cache
	h.redisClient.Del(h.ctx, "recipe:"+recipeId)
	h.redisClient.Del(h.ctx, "recipes")
	//Delete recipe from elastic store
	err = h.deleteRecipeInElasticStore(recipeId)
	if err != nil {
		zap.L().Error("Failed to delete recipe from elastic store", zap.Error(err))
	}
	c.JSON(http.StatusOK, gin.H{"message": "Recipe deleted successfully"})
}

// Search recipe in elasticsearch
func (h *RecipeHandler) insertRecipeInElasticstore(recipe models.Recipe) error {
	zap.L().Info("Inserting recipe in elastic store", zap.String("recipe_id", recipe.ID.Hex()))
	data, err := json.Marshal(recipe)
	if err != nil {
		zap.L().Error("Failed to marshal recipe", zap.Error(err))
		return err
	}

	res, err := h.elasticClient.Index(
		"recipe",
		bytes.NewReader(data),
		h.elasticClient.Index.WithDocumentID(recipe.ID.Hex()),
		h.elasticClient.Index.WithRefresh("true"),
	)

	if err != nil {
		zap.L().Error("Failed to insert recipe in elastic", zap.Error(err))
		return err
	}
	defer res.Body.Close()
	if res.IsError() {
		zap.L().Error("Failed to insert recipe in elastic", zap.String("response", res.String()))
		return errors.New("Failed to insert recipe in elastic")
	}
	zap.L().Info("Recipe inserted in elastic", zap.String("recipe_id", recipe.ID.Hex()))
	return nil

}

func (h *RecipeHandler) deleteRecipeInElasticStore(recpieId string) error {
	zap.L().Info("Deleting recipe from elastic store", zap.String("recipe_id", recpieId))
	res, err := h.elasticClient.Delete(
		"recipe",
		recpieId,
		h.elasticClient.Delete.WithRefresh("true"),
	)
	if err != nil {
		zap.L().Error("Failed to delete recipe from elastic store", zap.Error(err))
		return errors.New("Failed to delete recipe from elastic store")
	}
	defer res.Body.Close()
	if res.IsError() {
		zap.L().Error("Failed to delete recipe in elastic", zap.String("response", res.String()))
		return errors.New("Failed to delete recipe from elastic store")
	}
	zap.L().Info("Recipe deleted from elastic", zap.String("recipe_id", recpieId))
	return nil

}

func (h *RecipeHandler) SearchRecipeInElasticStore(c *gin.Context) {
	q := strings.TrimSpace(c.Query("q"))
	tag := strings.TrimSpace(c.Query("tag"))
	zap.L().Info("Searching recipes in elastic store", zap.String("q", q), zap.String("tag", tag))
	if q == "" && tag == "" {
		zap.L().Warn("Search query is required")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Search query is required"})
		return
	}

	should := make([]interface{}, 0)
	filter := make([]interface{}, 0)
	if q != "" {
		should = append(should, map[string]interface{}{
			"match": map[string]interface{}{
				"name": map[string]interface{}{
					"query":     q,
					"fuzziness": "AUTO",
				},
			},
		},
			map[string]interface{}{
				"match": map[string]interface{}{
					"tags": q,
				},
			},
		)
	}

	if tag != "" {
		filter = append(filter, map[string]interface{}{
			"term": map[string]interface{}{
				"tags.keyword": tag,
			},
		})
	}

	boolQuery := map[string]interface{}{}
	if len(should) > 0 {
		boolQuery["should"] = should
		boolQuery["minimum_should_match"] = 1
	}
	if len(filter) > 0 {
		boolQuery["filter"] = filter
	}
	zap.S().Infof("Search recipe query in elastic store: %v", boolQuery)

	searchBody := map[string]interface{}{
		"_source": []string{"id", "name", "tags", "imageUrl"},
		"query": map[string]interface{}{
			"bool": boolQuery,
		},
	}

	bodyBytes, err := json.Marshal(searchBody)
	if err != nil {
		zap.L().Error("Failed to marshal search body", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search recipes"})
		return
	}

	res, err := h.elasticClient.Search(
		h.elasticClient.Search.WithContext(h.ctx),
		h.elasticClient.Search.WithIndex("recipe"),
		h.elasticClient.Search.WithBody(bytes.NewReader(bodyBytes)),
	)
	if err != nil {
		zap.L().Error("Failed to search recipes", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search recipes"})
		return
	}
	defer res.Body.Close()

	if res.IsError() {
		zap.L().Error("Failed to search recipes in elastic", zap.String("response", res.String()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search recipes"})
	}

	var searchResp struct {
		Hits struct {
			Hits []struct {
				Source models.RecipeSearchResult `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}
	if err := json.NewDecoder(res.Body).Decode(&searchResp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse search response"})
		return
	}

	results := make([]models.RecipeSearchResult, 0, len(searchResp.Hits.Hits))
	for _, hit := range searchResp.Hits.Hits {
		results = append(results, hit.Source)
	}
	zap.L().Info("Found recipes", zap.Int("count", len(results)))
	c.JSON(http.StatusOK, results)

}
