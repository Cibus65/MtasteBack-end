package model

import (
	"back-end/config"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Recipe struct {
	Name        string            `json:"name"`
	Kitchen     string            `json:"kitchen"`
	Type        string            `json:"type"`
	Description map[int]string    `json:"description"`
	Ingredients map[string]string `json:"ingredients"`
	ID          int               `json:"ID"`
	Img         string            `json:"img"`
	UnixTime    int               `json:"unixTime"`
}

func (r *Recipe) GetByPage(page int) ([]Recipe, error) {
	collection := config.MongoClient.Database("RecipeBook").Collection("recipe")
	filter := bson.D{}
	opts := options.Find().SetSort(bson.D{{"unixtime", -1}, {"id", -1}})
	cursor, err := collection.Find(context.TODO(), filter, opts)
	if err != nil {
		return []Recipe{}, err
	}
	var recipes []Recipe
	max_id := 0

	for cursor.Next(context.TODO()) {
		var recipe Recipe

		if err = cursor.Decode(&recipe); err != nil {
			return []Recipe{}, err
		}
		max_id = recipe.ID
		break
	}
	cursor, err = collection.Find(context.TODO(), bson.D{{"id", bson.D{{"$gte", max_id + 1 - page*20}}}})
	if err != nil {
		return []Recipe{}, err
	}
	for cursor.Next(context.TODO()) {
		var recipe Recipe

		if err = cursor.Decode(&recipe); err != nil {
			return []Recipe{}, err
		}
		if len(recipes) < 20 {
			recipes = append(recipes, recipe)
		} else {
			break
		}
	}
	return recipes, nil
}
func (r *Recipe) GetByID(id int) (Recipe, error) {
	var recipe Recipe
	collection := config.MongoClient.Database("RecipeBook").Collection("recipe")
	filter := bson.D{{"id", id}}
	cursor := collection.FindOne(context.TODO(), filter)
	err := cursor.Decode(&recipe)
	if err != nil {
		return Recipe{}, err
	}

	return recipe, nil
}

//func (r *Recipe) getByCategory(category string, page int) {
//
//}
