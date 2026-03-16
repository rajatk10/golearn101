package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func main() {
	users := map[string]string{
		"admin": "password",
		"root":  "redh8123ace",
		"user":  "userDempo321",
	}
	ctx := context.Background()
	mongoDBUri := os.Getenv("MONGODB_URI")
	client, _ := mongo.Connect(options.Client().ApplyURI(mongoDBUri))
	collection := client.Database("recipeDB").Collection("users")
	//hashedPassword := sha256.New()
	for username, password := range users {
		h := sha256.New()
		h.Write([]byte(password))
		h.Sum(nil)
		hashedPassword := hex.EncodeToString(h.Sum(nil))
		collection.InsertOne(ctx, bson.M{
			"username": username,
			"password": hashedPassword,
		})
		fmt.Println("User: ", username, "Password: ", hashedPassword)
	}
}
