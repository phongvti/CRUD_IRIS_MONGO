package configs

import (
	"context"
	"fmt"
	"log"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func ConnectDB() *mongo.Client{
	client, err:= mongo.NewClient(options.Client().ApplyURI(EnvMongoURI()))
	if err!=nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err= client.Connect(ctx)

	if err!=nil{
		log.Fatal(err)
	}

	//ping the db
	err= client.Ping(ctx, nil)

	if err!=nil{
		log.Fatal(err)
	}

	fmt.Println("Connected DB")

	return client
}

//Client instance
var DB *mongo.Client= ConnectDB()

//get database collections
func GetCollection(collectionName string) *mongo.Collection {
	collection:= DB.Database("todoapp").Collection(collectionName)
	return collection
}


