package config

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

var MongoClient *mongo.Client

func init() {
	file, err := os.OpenFile("../app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	log.SetOutput(file)
	if err != nil {
		panic(err)
	}
	err = ConnectToDB()
	if err != nil {
		log.Fatalf("Failed connection to MongoDB:\n\t", err)
	}
}
func ConnectToDB() error {

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb://127.0.0.1").SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return err
	}
	MongoClient = client
	return nil
}
