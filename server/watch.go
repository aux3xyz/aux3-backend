package server

import (
	"context"
	"log"
	"time"

	"aux3.xyz/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dbName         = "aux3"
	collectionName = "events"
	retryDelay     = time.Second
)

func Watch() error {
	client := common.GetMongoClient()
	ctx := context.Background()

	collection := client.Database(dbName).Collection(collectionName)
	opts := options.ChangeStream().SetFullDocument(options.UpdateLookup)

	for {
		changeStream, err := collection.Watch(ctx, mongo.Pipeline{}, opts)
		if err != nil {
			log.Printf("Failed to start change stream: %v", err)
			time.Sleep(retryDelay)
			continue
		}

		if err := watchChanges(ctx, changeStream); err != nil {
			log.Printf("Change stream error: %v", err)
			changeStream.Close(ctx)
			time.Sleep(retryDelay)
			continue
		}
	}
}

func watchChanges(ctx context.Context, changeStream *mongo.ChangeStream) error {
	defer changeStream.Close(ctx)

	for changeStream.Next(ctx) {
		var event bson.M
		if err := changeStream.Decode(&event); err != nil {
			log.Printf("Error decoding change event: %v", err)
			continue
		}

		go Process() // TODO: Setup Process function; pass decoded event?

		if err := changeStream.Err(); err != nil {
			return err
		}
	}

	return nil
}
