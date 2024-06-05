package model

import (
	"back-end/config"
	"context"
	"fmt"
	"slices"

	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
	UserId   int `json:"userID"`
	RecipeID int `json:"recipeID"`
}
type FavouriteList struct {
	UserId    int   `json:"userID"`
	Favourite []int `json:"favourite"`
}

func (u *User) AddToFavourite() (FavouriteList, bool, error, int) {
	collection := config.MongoClient.Database("RecipeBook").Collection("favourite")
	userFavourite, err := findFavourite(u.UserId)
	if slices.Contains(userFavourite.Favourite, u.RecipeID) {
		return userFavourite, false, fmt.Errorf("Этот рецепт уже добавлен в список"), 6
	}
	if err != nil {
		return FavouriteList{}, false, err, 100
	}
	userFavourite.Favourite = append(userFavourite.Favourite, u.RecipeID)
	_ = collection.FindOneAndReplace(context.Background(), bson.D{{"userid", u.UserId}}, userFavourite)
	return userFavourite, true, nil, 0
}

func findFavourite(userID int) (FavouriteList, error) {
	collection := config.MongoClient.Database("RecipeBook").Collection("favourite")
	cursor, err := collection.Find(context.TODO(), bson.D{{"userid", userID}})

	if err != nil {
		return FavouriteList{}, err
	}
	var user FavouriteList
	for cursor.Next(context.TODO()) {
		if err = cursor.Decode(&user); err != nil {
			return FavouriteList{}, err
		}
	}
	if user.UserId == 0 {
		user = FavouriteList{UserId: userID, Favourite: []int{}}
		collection.InsertOne(context.Background(), user)
	}
	return user, nil

}
