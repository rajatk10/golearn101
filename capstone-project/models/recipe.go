package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Recipe struct {
	ID           bson.ObjectID `json:"id" bson:"_id"`
	Name         string        `json:"name" bson:"name"`
	Tags         []string      `json:"tags" bson:"tags"`
	Ingredients  []string      `json:"ingredients" bson:"ingredients"`
	Instructions []string      `json:"instructions" bson:"instructions"`
	PublishedAt  time.Time     `json:"publishedAt" bson:"publishedAt"`
	ImageURL     string        `json:"imageUrl" bson:"imageUrl"`
}

type RecipeSearchResult struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Tags     []string `json:"tags"`
	ImageURL string   `json:"imageUrl"`
}
