package model

import (
	"back-end/config"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
	Login        string `json:"login"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	ID           int    `json:"id"`
	SavedRecipes []map[string]interface{}
}

func (u *User) AddFavourite(link string, recipeID int) (string, error) {
	collection := config.MongoClient.Database("Users").Collection("users")
	var user User
	filter := bson.D{{"login", u.Login}, {"id", u.ID}}
	us := collection.FindOne(context.TODO(), filter)
	//result := collection.FindOne(context.TODO(), filter)
	//var recipe Recipe
	err := us.Decode(&user)
	if err != nil {
		return "", err
	}

	newFav := make(map[string]interface{})
	newFav["link"] = link
	newFav["recipeID"] = recipeID
	flag := false
	for _, elem := range user.SavedRecipes {
		left_link := elem["link"].(string)
		left_RID := elem["recipeID"].(int32)
		if left_link == link || int(left_RID) == recipeID {
			flag = true
		}
	}
	if flag {
		return "Этот рецепт уже добавлен в избранное", nil
	}
	user.SavedRecipes = append(user.SavedRecipes, newFav)
	upd, err := collection.ReplaceOne(context.TODO(), filter, user)
	fmt.Println(upd)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Рецепт с ID[ %d ] был добавлен", recipeID), nil
}

//func (u *User) AddUser(login ,email,password  string){
//	filter := bson.D{{"login", u.Login}, {"id", u.ID}}
//
//}
//
//func (u *User)
