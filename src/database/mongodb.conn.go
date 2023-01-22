package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/HansBukerG/wm-back-end/src/app_env"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetCollection(collection string) *mongo.Collection {
	host := app_env.GetEnv("MONGO_HOST")
	port := app_env.GetEnv("MONGO_PORT")
	usr := app_env.GetEnv("MONGO_USER")
	pwd := app_env.GetEnv("MONGO_PASS")
	database := app_env.GetEnv("MONGO_DATABASE")

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
