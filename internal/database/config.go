package database

import (
    "context"
    "log"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func ConnectMongo() {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    client, err := mongo.Connect(
        ctx,
        options.Client().ApplyURI(
            "mongodb+srv://phanminh:phanminhdat11@cluster0.jxjisrt.mongodb.net/service_shortener_url?retryWrites=true&w=majority",
        ),
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

