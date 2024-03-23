package config

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

var MongoClient *mongo.Client

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	KEY := os.Getenv("KEY")
	fmt.Println(KEY)
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
