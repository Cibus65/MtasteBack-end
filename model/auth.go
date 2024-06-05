package model

import (
	"back-end/config"
	"context"
	"crypto/sha1"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	salt          = "tklw12hfoiv3pjihu5u521jofc29urji"
	signingKey    = "gag2rp1jkr21fvi0jio2jqfwcpkkngjy2t0tfp"
	valid_symbols = "abcdefghijklmnopqrstuvwxyz0123456789"
)

type Auth struct {
	Login         string `json:"login" binding:"required"`
	Password      string `json:"password" binding:"required"`
	RetryPassword string `json:"retrypassword" binding:"required"`
	ID            int    `json:"id"`
}

type tokenClaims struct {
	jwt.StandardClaims
	AuthId int `json:"AuthId"`
}

func (a *Auth) GenerateJWT(user Auth) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 + time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID,
	})
	return token.SignedString([]byte(signingKey))
}

func getUser(login string, password string) (Auth, error, int) {
	if login == "" || password == "" {
		return Auth{}, fmt.Errorf("Все данные должны быть заполнены!"), 1
	}
	collection := config.MongoClient.Database("RecipeBook").Collection("users")
	filter := bson.D{{"login", login}}
	var exist Auth
	_ = collection.FindOne(context.TODO(), filter).Decode(&exist)
	if exist.Login == "" && exist.Password == "" {
		return Auth{}, fmt.Errorf("Пользователя не существует: %s", login), 3
	} else if exist.Password != password {
		return Auth{}, fmt.Errorf("Пароль неверный"), 4
	} else {
		return exist, nil, 0
	}
}

func (a *Auth) CreateUser() (bool, error, int) {
	if a.Login == "" || a.Password == "" || a.RetryPassword == "" {
		return false, fmt.Errorf("Все данные должны быть заполнены!"), 1
	}
	if !password_valid(a.Password) || !password_valid(a.RetryPassword) {
		return false, fmt.Errorf("Пароль должен быть не короче 8 символов, не длиннее 50, содержать хотя бы одну цифру и латинскую букву, не содержать другие спец.символы кроме \"_\""), 8
	}

	if !login_valid(a.Login) {
		return false, fmt.Errorf("Логин должен быть не короче 8 символов, не длиннее 50, содержать хотя бы одну латинскую букву,не содержать другие спец.символы кроме \"_\" и цифр"), 9
	}
	if a.Password != a.RetryPassword {
		return false, fmt.Errorf("Пароли должны быть одинаковыми"), 5
	}
	a.Password = hashpasswd(a.Password)

	_, flag, err := findUser(a.Login)
	if err != nil {
		return false, err, 100
	}
	if !flag {
		collection := config.MongoClient.Database("RecipeBook").Collection("users")
		a.ID = rand.Intn(1000_000_00) + 1000_000_00
		collection.InsertOne(context.TODO(), a)
		return true, nil, 0
	} else {
		return false, fmt.Errorf("Такой пользователь уже существует"), 2
	}
}

func findUser(login string) (Auth, bool, error) {
	collection := config.MongoClient.Database("RecipeBook").Collection("users")

	cursor, err := collection.Find(context.TODO(), bson.D{{"login", login}})
	if err != nil {
		return Auth{}, false, err
	}

	var user Auth
	for cursor.Next(context.TODO()) {
		if err = cursor.Decode(&user); err != nil {
			return Auth{}, false, err
		}
	}
	if user.Login != "" {
		return user, true, nil
	} else {
		return Auth{}, false, err
	}
}

func (a *Auth) SignIn() (string, error, int) {
	a.Password = hashpasswd(a.Password)
	user, err, code := getUser(a.Login, a.Password)
	if err != nil {
		return "", err, code
	}
	jwt, err := a.GenerateJWT(user)
	if err != nil {
		return "", err, 100
	}
	return jwt, err, code

}

func hashpasswd(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

//func (a *Auth) SignIn()

//func (a *Auth) AddFavourite(userID int, link string, recipeID int) (string, error) {
//	collection := config.MongoClient.Database("Users").Collection("users")
//	var user Auth
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

func (a *Auth) AddUser(login, email, password string) string {
	collection := config.MongoClient.Database("Users").Collection("users")
	filter := bson.D{{"login", login}}
	var user Auth
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

func password_valid(password string) bool {
	countDigit, countChar, countUnderline := 0, 0, 0
	for index := range password {

		if strings.Contains("0123456789", string(password[index])) {
			countDigit++
		} else if strings.Contains(valid_symbols, string(password[index])) {
			countChar++
		} else if string(password[index]) == "_" {
			countUnderline++
		} else {
			return false
		}
	}
	if countChar+countDigit+countUnderline <= 50 && countChar+countDigit+countUnderline >= 8 && countDigit > 0 && countChar > 0 {
		return true
	} else {
		return false
	}
}
func login_valid(login string) bool {
	countDigit, countChar, countUnderline := 0, 0, 0
	for index := range login {

		if strings.Contains("0123456789", string(login[index])) {
			countDigit++
		} else if strings.Contains(valid_symbols, string(login[index])) {
			countChar++
		} else if string(login[index]) == "_" {
			countUnderline++
		} else {
			return false
		}
	}
	if countChar+countDigit+countUnderline <= 50 && countChar+countDigit+countUnderline >= 8 && countChar > 0 {
		return true
	} else {
		return false
	}
}
