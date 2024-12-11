package database

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


var Client *mongo.Client

func GetDatabase() *mongo.Database {

	if Client == nil {
		log.Fatalf("Mongo DB is not initialized")
	}

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Failed to load from .env")
	}

	DB_NAME := os.Getenv("DB_NAME")

	return Client.Database(DB_NAME)
}

func ConnectDatabase() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Failed to load from .env")
	}

	MONGO_URI := os.Getenv("MONGO_URI")

	clientOption := options.Client().ApplyURI(MONGO_URI)

	res, err := mongo.Connect(context.Background(), clientOption)
	if err != nil {
		log.Fatalf("Error while connecting to MongoDB: %s", err)
	}

	Client = res

	err = Client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error while ping to MongoDB: %s", err)
	}

	log.Printf("Successfully connected to MongoDB!")
}

func DisconnectDatabase() {

	if Client == nil {
		log.Fatalf("MongoDB isn't initialized.")
	}

	err := Client.Disconnect(context.Background())
	if err != nil {
		log.Fatalf("Error while disconnecting to MongoDB: %s", err)
	}

	log.Printf("Successfully disconnected from MongoDB!")
}