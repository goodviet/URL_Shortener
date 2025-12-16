package database

import "go.mongodb.org/mongo-driver/mongo"

func URLCollection() *mongo.Collection {
    return Client.Database("service_shortener_url").Collection("urls")
}
