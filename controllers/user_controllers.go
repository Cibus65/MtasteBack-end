package controllers

import (
	"back-end/model"
	"crypto/sha1"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"math/rand"
	"net/http"
	"os"
)

func init() {
	godotenv.Load(".env")
}

type UserController struct{}

//	func (_ *UserController) AddToFavourite(c *gin.Context) {
//		UserID, err := strconv.Atoi(c.Param("userID"))
//		if err != nil {
//			panic(err)
//		}
//		link := c.Param("link")
//		recipeID, err := strconv.Atoi(c.Param("recipeID"))
//		if err != nil {
//			c.Error(err)
//		}
//		message, err := (&model.User{}).AddFavourite(UserID, link, recipeID)
//		if err != nil {
//			c.Error(err)
//		}
//		c.JSON(http.StatusOK, message)
//
// }
func (_ *UserController) SignIn(c *gin.Context) {
	var user model.SignInUser = model.SignInUser{
		Login:    c.Query("login"),
		Password: c.Query("password"),
	}
	user.Password = hashpasswd(user.Password)
	token, err := user.GenerateJWT()

	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

func (_ *UserController) AddNewUser(c *gin.Context) {
	var user model.User = model.User{
		Login:    c.Query("login"),
		Email:    c.Query("email"),
		Password: c.Query("password"),
	}
	user.ID = 1000_000_000 + rand.Intn(1000_000_000)%1000_000_000

	user.Password = hashpasswd(user.Password)
	user.CreateUser()

}

func hashpasswd(password string) string {
	salt := os.Getenv("salt")
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
