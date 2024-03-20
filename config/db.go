package config

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const KEY = "mongodb+srv://ilyasuseinov3301:mishka_2023@recipebook.xxu8dre.mongodb.net/?retryWrites=true&w=majority&appName=RecipeBook"

var MongoClient *mongo.Client

func init() {
	ConnectToDB(KEY)
}
func ConnectToDB(KEY string) {

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(KEY).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	MongoClient = client
}
