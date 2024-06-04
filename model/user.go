package model

import (
	"back-end/config"
	"context"
	"crypto/sha1"
	"fmt"
	"math/rand"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	salt       = "tklw12hfoiv3pjihu5u521jofc29urji"
	signingKey = "gag2rp1jkr21fvi0jio2jqfwcpkkngjy2t0tfp"
)

type User struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
	ID       int    `json:"id"`
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"userId"`
}

func (u *User) GenerateJWT(user User) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 + time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID,
	})
	return token.SignedString([]byte(signingKey))
}

func getUser(login string, password string) (User, error,int) {
	if  login == "" || password == "" {
		return User{}, fmt.Errorf("Все данные должны быть заполнены!"),1
	}
	collection := config.MongoClient.Database("RecipeBook").Collection("users")
	filter := bson.D{{"login", login}}
	var exist User
	_ = collection.FindOne(context.TODO(), filter).Decode(&exist)
	if exist.Login == "" && exist.Password == ""  {
		return User{}, fmt.Errorf("Пользователя не существует: %s", login),3
	} else if exist.Password != password {
		return User{}, fmt.Errorf("Пароль неверный"),4
	} else {
		return exist, nil,0
	}
}

func (u *User) CreateUser() (bool, error,int) {
	if  u.Login == "" || u.Password == "" {
		return false, fmt.Errorf("Все данные должны быть заполнены!"),1
	}
	u.Password = hashpasswd(u.Password)
	
	_, flag, err := findUser( u.Login)
	if err != nil {
		return false, err,100
	}
	if !flag {
		collection := config.MongoClient.Database("RecipeBook").Collection("users")
		u.ID = rand.Intn(1000_000_00) + 1000_000_00
		collection.InsertOne(context.TODO(), u)
		return true, nil,0
	} else {
		return false, fmt.Errorf("Такой пользователь уже существует"),2
	}
	return true, nil,0
}

func findUser(login string) (User, bool, error) {
	collection := config.MongoClient.Database("RecipeBook").Collection("users")

	cursor, err := collection.Find(context.TODO(), bson.D{{"login", login}})
	if err != nil {
		return User{}, false, err
	}
 
	var user User
	for cursor.Next(context.TODO()) {
		if err = cursor.Decode(&user); err != nil {
			return User{}, false, err
		}
	}
	if user.Login != "" {
		return user, true, nil
	} else {
		return User{}, false, err
	}
}

func (u *User) SignIn() (string, error,int) {
	u.Password = hashpasswd(u.Password)
	user, err,code := getUser(u.Login, u.Password)
	if err != nil {
		return "", err,code
	}
	jwt,err := u.GenerateJWT(user)
	if err!=nil{
		return "",err,100
	}
	return jwt,err,code

}

func hashpasswd(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
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
	user.Password = password
	collection.InsertOne(context.TODO(), user)
	return "Пользователь зарегистрирован"

}
