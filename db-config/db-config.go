package dbconfig

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func EnvMongoURI() string {

	var dbURL = "mongodb+srv://aniket_mauto:9TxEsbh9vZVzIz68@shared-cluster.tb9e4zs.mongodb.net/test?maxPoolSize=1500&w=majority"

	return dbURL
}

var client *mongo.Client

func ResolveClientDB() *mongo.Client {
	if client != nil {
		return client
	}

	var err error

	clientOptions := options.Client().ApplyURI(EnvMongoURI())
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// check the connection
	err = client.Ping(context.Background(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connection established...")
	return client
}

func CloseClientDB() {
	if client == nil {
		return
	}

	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	// TODO optional you can log your closed MongoDB client
	fmt.Println("Connection to MongoDB closed.")
}

// DB Client instance
var DB = ResolveClientDB()

// GetCollection getting database collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	// var local = "golangAPI"
	var remote = "mautodb"
	collection := client.Database(remote).Collection(collectionName)
	return collection
}
