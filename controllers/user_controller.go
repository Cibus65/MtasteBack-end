package controllers

import (
	"back-end/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (_ *UserController) SignUp(c *gin.Context) {
	fmt.Println(1)
	var user model.User = model.User{
		Login:    c.Query("login"),
		Password: c.Query("password"),
	}
	result, err,code := user.CreateUser()
	fmt.Print(result)

	if err != nil {
		c.JSON(http.StatusOK, map[string]interface{}{
			// "token": token,
			"result": result,
			"login":  user.Login,
			"error":  fmt.Sprintf("%s", err),
			"errorCode":code,
		})
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			// "token": token,
			"result": result,
			"login":  user.Login,
			"error":  "",
			"errorCode":code,

		})
	}
}
func (_ *UserController) SignIn(c *gin.Context) {
	var user model.User = model.User{
		Login:    c.Query("login"),
		Password: c.Query("password"),
	}
	token, err,code := user.SignIn()
	if err != nil {
		c.JSON(http.StatusOK, map[string]interface{}{
			"token": token,
			"login": user.Login,
			"error": fmt.Sprintf("%s", err),
			"errorCode":code,

		})
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"token": token,
			"login": user.Login,
			"error": "",
			"errorCode":code,

		})
	}

}

// func (_ *UserController) AddNewUser(c *gin.Context) {
// 	var user model.User = model.User{
// 		Login:    c.Query("login"),
// 		Email:    c.Query("email"),
// 		Password: c.Query("password"),
// 	}
// 	user.ID = 1000_000_000 + rand.Intn(1000_000_000)%1000_000_000

// 	user.Password = hashpasswd(user.Password)
// 	user.CreateUser()

// }
