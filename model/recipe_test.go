package model

import (
	"back-end/config"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"testing"
)

var mongoClient *mongo.Client

func TestRecipe_GetByPage(t *testing.T) {
	var page_slice []int = []int{1, 10, 100, 1000, 10000}
	for _, page := range page_slice {
		actual, actual_err := (&Recipe{}).GetByPage(page)

		collection := mongoClient.Database("RecipeBook").Collection("recipes")
		filter := bson.D{}
		opts := options.Find().SetSort(bson.D{{"unixtime", -1}, {"id", -1}})
		cursor, expected_err := collection.Find(context.TODO(), filter, opts)
		if expected_err != actual_err {
			t.Errorf("Ожидал следующая ошибка:\n\t%s\nВозникла следующая:\n\t%s", expected_err, actual_err)
		}
		var expected []Recipe
		max_id := 0

		for cursor.Next(context.TODO()) {
			var recipe Recipe
			if expected_err = cursor.Decode(&recipe); expected_err != nil {
				if expected_err != actual_err {
					t.Errorf("Ожидал следующая ошибка:\n\t%s\nВозникла следующая:\n\t%s", expected_err, actual_err)
				}
			}

			max_id = recipe.ID
			break
		}
		cursor, expected_err = collection.Find(context.TODO(), bson.D{{"id", bson.D{{"$gte", max_id + 1 - page*20}}}})
		if expected_err != nil {
			t.Error(expected_err)
		}
		for cursor.Next(context.TODO()) {
			var recipe Recipe
			if expected_err != actual_err {
				t.Errorf("Ожидал следующая ошибка:\n\t%s\nВозникла следующая:\n\t%s", expected_err, actual_err)
			}

			if expected_err = cursor.Decode(&recipe); expected_err != nil {
				if expected_err != actual_err {
					t.Errorf("Ожидал следующая ошибка:\n\t%s\nВозникла следующая:\n\t%s", expected_err, actual_err)
				}
			}
			if len(expected) < 20 {
				expected = append(expected, recipe)
			} else {
				break
			}
		}
		if actual[0].Name != expected[0].Name {
			t.Fail()
		}

	}
}
func TestRecipe_GetByID(t *testing.T) {
	var id_slice []int = []int{1, 10, 100, 1000, 10000}

	for _, id := range id_slice {
		var expected Recipe
		collection := mongoClient.Database("RecipeBook").Collection("recipes")
		filter := bson.D{{"id", id}}
		cursor := collection.FindOne(context.TODO(), filter)
		expected_err := cursor.Decode(&expected)

		actual, actual_err := (&Recipe{}).GetByID(id)
		if expected_err != actual_err {
			t.Errorf("Ожидал следующая ошибка:\n\t%s\nВозникла следующая:\n\t%s", expected_err, actual_err)
		}
		if expected.Name != actual.Name {
			t.Fail()
		}

	}
}
func TestMain(m *testing.M) {

	key := config.KEY
	connectToDB(key)
	code := m.Run()
	os.Exit(code)
}

func connectToDB(KEY string) {

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(KEY).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	mongoClient = client
}
