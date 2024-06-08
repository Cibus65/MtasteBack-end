package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	opts := options.Client().ApplyURI(fmt.Sprintf("mongodb://mongodb:27017")).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opts)
	fmt.Println(err, "1")
	if err != nil {
		return err
	}
	MongoClient = client
	return nil
}
