package common

import (
    "context"
    "log"
    "os"
    "sync"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)


var (
    mongoClient *mongo.Client
    mongoOnce   sync.Once
)


func init() {
    // TODO load env
}

func GetConfig(name string) string {
    // TODO access env var from map?
    return os.Getenv(name)
}

func GetMongoClient() *mongo.Client {
    mongoOnce.Do(func() {
        // Get MongoDB URI from environment variable, fallback to default if not set
        mongoURI := GetConfig("MONGODB_URI")
        if mongoURI == "" {
            mongoURI = "mongodb://localhost:27017"
        }

        // Connect to MongoDB
        ctx := context.Background()
        client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
        if err != nil {
            log.Fatalf("Failed to connect to MongoDB: %v", err)
        }

        // Ping the database to verify connection
        if err := client.Ping(ctx, nil); err != nil {
            log.Fatalf("Failed to ping MongoDB: %v", err)
        }

        mongoClient = client
    })

    return mongoClient
}