package server

import (
    "context"
    "log"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/bson"

    "aux3.xyz/common"
)

func Watch() {
    // Get the shared MongoDB client
    client := common.GetMongoClient()
    ctx := context.Background()

    // Get collection to watch
    // TODO: setup right database and collection names
    collection := client.Database("aux3").Collection("events")

    // Set up change stream
    pipeline := mongo.Pipeline{}
    opts := options.ChangeStream().SetFullDocument(options.UpdateLookup)
    
    // Start watching for changes
    changeStream, err := collection.Watch(ctx, pipeline, opts)
    if err != nil {
        log.Fatalf("Failed to start change stream: %v", err)
    }
    defer changeStream.Close(ctx)

    // Loop forever watching for changes
    for changeStream.Next(ctx) {
        var event bson.M
        if err := changeStream.Decode(&event); err != nil {
            log.Printf("Error decoding change event: %v", err)
            continue
        }

        // Send event to process
        // TODO: Setup Process function
        go Process()

        // Basic error handling with reconnection
        if err := changeStream.Err(); err != nil {
            log.Printf("Error in change stream: %v", err)
            time.Sleep(1 * time.Second)
            return Watch() // Recursive call to restart watching
        }
    }
}
