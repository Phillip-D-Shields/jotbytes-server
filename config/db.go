package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const DB_URI = "mongodb://localhost:27017"

// DB variable to share the MongoDB client across the application
var DB *mongo.Client

// DatabaseName is the name of the database
const DatabaseName = "jotbytes"

// ConnectDB initializes the database connection
func ConnectDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Replace your_mongodb_uri with your actual MongoDB URI
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(DB_URI))
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Ping the database to verify connection is established
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	DB = client
	log.Println("Connected to MongoDB!")
}

// GetCollection returns a handle to a collection in your MongoDB
func GetCollection(collectionName string) *mongo.Collection {
	return DB.Database(DatabaseName).Collection(collectionName)
}
