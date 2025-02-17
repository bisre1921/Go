package controllers

import (
	"context"
	"fmt"
	"log"
	"mongoApi/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var connectionString = "mongodb+srv://bisrat:bisrat1234@mongoandgoapi.amequ.mongodb.net/?retryWrites=true&w=majority&appName=mongoAndGoApi"
var dbName = "movies"
var collectionName = "watchlist"

var collection *mongo.Collection

func init() {
	clientOption := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect((clientOption))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	collection = client.Database(dbName).Collection(collectionName)
	fmt.Println("Collection instance created!")

}

func insertOne(movie models.Movie) {
	insertResult, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a Single Record ", insertResult.InsertedID)
}

func updateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Updated a Single Record ", result.UpsertedID)
}
