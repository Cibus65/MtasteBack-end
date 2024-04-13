package model

import (
	"back-end/config"
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"math/rand"
	"os"
	"time"
)

type User struct {
	Login    string `json:"login" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	ID       int    `json:"id"`
}

type SignInUser struct {
	Login    string `json: "login"`
	Password string `json: "password"`
}
type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"userId"`
}

func (siu *SignInUser) GenerateJWT() (string, error) {
	user, err := GetUser(siu.Login, siu.Password)
	if err != nil {
		fmt.Println(err)
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 + time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID,
	})
	return token.SignedString([]byte(os.Getenv("signingKey")))
}

func GetUser(login string, password string) (User, error) {
	collection := config.MongoClient.Database("Users").Collection("users")
	filter := bson.D{{"login", login}}
	var exist User
	_ = collection.FindOne(context.TODO(), filter).Decode(&exist)
	if exist.Login == "" && exist.Password == "" && exist.Email == "" {
		return User{}, fmt.Errorf("Пользователя с логином %s не существует", login)
	} else if exist.Password != password {
		return User{}, fmt.Errorf("Пароль неверный")
	} else {
		return exist, nil
	}
}

func (u *User) CreateUser() int {
	_, err := GetUser(u.Login, u.Password)
	if err == nil {
		fmt.Println("Пользователь с таким логином уже зарегистрирован")
		return 0
	}
	if u.Login == "" {
		fmt.Println("Логин не должен быть пустым")
		return 0
	} else if u.Password == "" {

		fmt.Println("Пароль не должен быть пустым")
		return 0
	} else if u.Email == "" {
		fmt.Println("Почта не должна быть пустым")
		return 0
	}

	fmt.Printf("user:%v\n", u.ID)
	return u.ID
}

//func (u *User) SignIn()

//func (u *User) AddFavourite(userID int, link string, recipeID int) (string, error) {
//	collection := config.MongoClient.Database("Users").Collection("users")
//	var user User
//	filter := bson.D{{"id", userID}}
//	fmt.Println(userID)
//	us := collection.FindOne(context.TODO(), filter)
//	err := us.Decode(&user)
//	if err != nil {
//		return "", err
//	}
//	if user.ID == 0 {
//		return "Такого пользователя не существует", nil
//	}
//	//result := collection.FindOne(context.TODO(), filter)
//	//var recipe Recipe
//
//	newFav := make(map[string]interface{})
//	newFav["link"] = link
//	newFav["recipeID"] = recipeID
//	flag := false
//	for _, elem := range user.SavedRecipes {
//		left_link := elem["link"].(string)
//		left_RID := elem["recipeID"].(int32)
//		if left_link == link || int(left_RID) == recipeID {
//			flag = true
//		}
//	}
//	if flag {
//		return "Этот рецепт уже добавлен в избранное", nil
//	}
//	user.SavedRecipes = append(user.SavedRecipes, newFav)
//	upd, err := collection.ReplaceOne(context.TODO(), filter, user)
//	fmt.Println(upd)
//	if err != nil {
//		return "", err
//	}
//	return fmt.Sprintf("Рецепт с ID[ %d ] был добавлен", recipeID), nil
//}

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
