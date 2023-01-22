package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	// database = "promotions"
)

func GetCollection(collection string) *mongo.Collection {
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
	host := os.Getenv("MONGO_HOST")
	if host == "" {
		log.Fatal("Should define .env MONGO_HOST")
	}
	port := os.Getenv("MONGO_PORT")
	if port == "" {
		log.Fatal("Should define .env MONGO_PORT")
	}
	usr := os.Getenv("MONGO_USER")
	if usr == "" {
		log.Fatal("Should define .env MONGO_USER")
	}
	pwd := os.Getenv("MONGO_PASS")
	if pwd == "" {
		log.Fatal("Should define .env MONGO_PASS")
	}
	database := os.Getenv("MONGO_DATABASE")
	if database == ""{
		log.Fatal("Should define .env MONGO_DATABASE")
	}

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s", usr, pwd, host, port)
	log.Println("URI: " + uri)

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))

	if err != nil {
		log.Fatal("Error in connection!: " + uri)
		panic(err.Error())
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("Error: client.Connect() Failed!")
		panic(err.Error())
	}
	return client.Database(database).Collection(collection)
}
