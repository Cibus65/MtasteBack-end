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
	// Устанавливаем параметры API сервера MongoDB
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	// Создаем параметры подключения к MongoDB
	opts := options.Client().ApplyURI(fmt.Sprintf("mongodb+srv://ilyasuseinov3301:abc2024@recipebook.xxu8dre.mongodb.net/?retryWrites=true&w=majority&appName=RecipeBook"))
	// Добавляем параметры API сервера к параметрам подключения
	opts.SetServerAPIOptions(serverAPI)

	// Подключаемся к MongoDB
	client, err := mongo.Connect(context.TODO(), opts)
	fmt.Println(err, "1")
	if err != nil {
		return err
	}

	// Сохраняем клиент для дальнейшего использования
	MongoClient = client
	return nil
}
