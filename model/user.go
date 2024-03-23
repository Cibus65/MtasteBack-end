package model

import (
	"back-end/config"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"math/rand"
)

type User struct {
	Login        string `json:"login"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	ID           int    `json:"id"`
	SavedRecipes []map[string]interface{}
}

func (u *User) AddFavourite(userID int, link string, recipeID int) (string, error) {
	collection := config.MongoClient.Database("Users").Collection("users")
	var user User
	filter := bson.D{{"id", userID}}
	fmt.Println(userID)
	us := collection.FindOne(context.TODO(), filter)
	err := us.Decode(&user)
	if err != nil {
		return "", err
	}
	if user.ID == 0 {
		return "Такого пользователя не существует", nil
	}
	//result := collection.FindOne(context.TODO(), filter)
	//var recipe Recipe

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

func (u *User) AddUser(login, email, password string) string {
	collection := config.MongoClient.Database("Users").Collection("users")
	filter := bson.D{{"login", login}}
	var user User
	us := collection.FindOne(context.TODO(), filter)
	us.Decode(&user)
	if user.ID != 0 {
		return "Пользователь с таким ником уже существует"
	}
	user.ID = rand.Intn(1000_000_000) + 1000_000_00
	user.Login = login
	user.Email = email
	user.Password = password
	collection.InsertOne(context.TODO(), user)
	return "Пользователь зарегистрирован"

}
