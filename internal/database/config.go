package database

import (
    "context"
    "log"
	"os"
    "time"

    "github.com/joho/godotenv"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func ConnectMongo() {

	   // Load .env
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

	mongoURI := os.Getenv("MONGO_URI")
    if mongoURI == "" {
        log.Fatal("MONGO_URI is not set")
    }


    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    client, err := mongo.Connect(
        ctx,
        options.Client().ApplyURI(mongoURI),
    )
	

    if err != nil {
        log.Fatal("Mongo connect error:", err)
    }

    // Ping thử
    if err := client.Ping(ctx, nil); err != nil {
        log.Fatal("Mongo ping error:", err)
    }

    Client = client
    log.Println("✅ MongoDB connected")
}

