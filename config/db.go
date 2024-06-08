package config

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func init() {

	err := ConnectToDB()
	fmt.Println(err, "0")
	if err != nil {
		log.Fatalf("Failed connection to MongoDB:\n\t", err)
	}
}
func ConnectToDB() error {

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(fmt.Sprintf("mongodb+srv://ilyasuseinov3301:abc2024@recipebook.xxu8dre.mongodb.net/?retryWrites=true&w=majority&appName=RecipeBook")).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opts)
	fmt.Println(err, "1")
	if err != nil {
		return err
	}
	MongoClient = client
	return nil
}
