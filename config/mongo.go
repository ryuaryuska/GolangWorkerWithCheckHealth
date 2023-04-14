
package config

import (
	"context"
	"os"
	"WorkerWithCheckHealth/exception"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConnection() *mongo.Database {
	mongoURL := os.Getenv("MONGO_HOST")

	clientOptions := options.Client()
	clientOptions.ApplyURI(mongoURL)
	client, err := mongo.NewClient(clientOptions)
	exception.PanicIfNeeded(err)

	err = client.Connect(context.Background())
	exception.PanicIfNeeded(err)

	dbName := os.Getenv("MONGO_NAME")
	return client.Database(dbName)
}
