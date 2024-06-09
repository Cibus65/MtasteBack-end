package model

import (
	"back-end/config"
	"context"
	"math/rand"
	"slices"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Recipe struct {
	Name         string            `json:"name"`
	Kitchen      string            `json:"kitchen"`
	Type         string            `json:"type"`
	Description  map[int]string    `json:"description"`
	Ingredients  map[string]string `json:"ingredients"`
	ID           int               `json:"ID"`
	ImgWindowUrl string            `json:"imgwindowurl"`
	UnixTime     int               `json:"unixTime"`
	Synopsis     string            `json:"synopsis"`
	IsFavourite  bool              `json:"isFavourite"`
}

func (r *Recipe) GetByPage(page int, userid int) ([]Recipe, error) {
	collection := config.MongoClient.Database("RecipeBook").Collection("recipes")
	//collection := config.MongoClient.Database("RecipeBook").Collection("recipe")

	var recipes []Recipe
	max_id, err := getMaxID(collection)
	if err != nil {
		return []Recipe{}, err
	}
	cursor, err := collection.Find(context.TODO(), bson.D{{"id", bson.D{{"$gte", max_id + 1 - page*20}}}})
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
	recipes, err, _ = checkIsFavourite(recipes, userid)
	if err != nil {
		return []Recipe{}, err
	}
	return recipes, nil
}
func (r *Recipe) GetByIDclassic(id int) (Recipe, error) {
	var recipe Recipe
	collection := config.MongoClient.Database("RecipeBook").Collection("recipes")
	//collection := config.MongoClient.Database("RecipeBook").Collection("recipe")
	filter := bson.D{{"id", id}}
	cursor := collection.FindOne(context.TODO(), filter)
	err := cursor.Decode(&recipe)
	if err != nil {
		return Recipe{}, err
	}
	return recipe, nil
}

func (r *Recipe) GetByID(id int, userid int) (Recipe, error) {
	var recipe Recipe
	collection := config.MongoClient.Database("RecipeBook").Collection("recipes")
	//collection := config.MongoClient.Database("RecipeBook").Collection("recipe")
	filter := bson.D{{"id", id}}
	cursor := collection.FindOne(context.TODO(), filter)
	err := cursor.Decode(&recipe)
	if err != nil {
		return Recipe{}, err
	}
	recipes, err, _ := checkIsFavourite([]Recipe{recipe}, userid)
	if err != nil {
		return Recipe{}, err
	}
	return recipes[0], nil
}

// func (r *Recipe) getByCategory(category string, page int) {
//
// }
func (r *Recipe) FindRecipe(words string, userid int) ([]Recipe, error) {
	collection := config.MongoClient.Database("RecipeBook").Collection("recipes")
	//collection := config.MongoClient.Database("RecipeBook").Collection("recipe")
	filter := bson.M{"name": bson.M{"$regex": words, "$options": "i"}}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return []Recipe{}, err
	}
	var recipes []Recipe

	for cursor.Next(context.TODO()) {
		var recipe Recipe

		if err = cursor.Decode(&recipe); err != nil {
			return []Recipe{}, err
		}
		recipes = append(recipes, recipe)

	}
	recipes, err, _ = checkIsFavourite(recipes, userid)
	if err != nil {
		return []Recipe{}, err
	}
	return recipes, nil
}

func (r *Recipe) GetRandomRecipe(userid int) ([]Recipe, error) {
	collection := config.MongoClient.Database("RecipeBook").Collection("recipes")
	maxID, err := getMaxID(collection)
	if err != nil {
		return []Recipe{}, err
	}
	var randomNumArray []int
	for len(randomNumArray) != 3 {
		randNum := rand.Intn(maxID) + 1
		if !slices.Contains(randomNumArray, randNum) {
			randomNumArray = append(randomNumArray, randNum)
		}
	}
	var recipes []Recipe
	for index := range randomNumArray {
		doc := collection.FindOne(context.TODO(), bson.D{{"id", randomNumArray[index]}})
		var recipe Recipe
		err = doc.Decode(&recipe)
		if err != nil {
			return []Recipe{}, err
		}
		recipes = append(recipes, recipe)

	}
	recipes, err, _ = checkIsFavourite(recipes, userid)
	if err != nil {
		return []Recipe{}, err
	}
	return recipes, nil

}

func getMaxID(collection *mongo.Collection) (int, error) {
	// collection := config.MongoClient.Database("RecipeBook").Collection("recipe")
	filter := bson.D{}
	opts := options.Find().SetSort(bson.D{{"unixtime", -1}, {"id", -1}})
	cursor, err := collection.Find(context.TODO(), filter, opts)
	if err != nil {
		return 0, err
	}
	max_id := 0

	for cursor.Next(context.TODO()) {
		var recipe Recipe

		if err = cursor.Decode(&recipe); err != nil {
			return 0, err
		}
		max_id = recipe.ID
		break
	}
	return max_id, nil

}

func checkIsFavourite(recipes []Recipe, userid int) ([]Recipe, error, int) {
	user := User{UserId: userid}
	favourite, err, code := user.GetFavouriteRecipes()

	for index, recipe := range recipes {
		if contain(favourite, &recipe) {
			recipes[index].IsFavourite = true
		} else {
			recipes[index].IsFavourite = false
		}
	}
	return recipes, err, code

}

func contain(recipes []Recipe, recipe *Recipe) bool {
	for _, rec := range recipes {
		if rec.ID == recipe.ID {
			return true
		}
	}
	return false
}
