package controllers

import (
	"back-end/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct{}

func (_ *AuthController) SignUp(c *gin.Context) {

	var user model.Auth = model.Auth{
		Login:         c.Query("login"),
		Password:      c.Query("password"),
		RetryPassword: c.Query("retry_password"),
	}
	result, err, code := user.CreateUser()

	if err != nil {
		c.JSON(http.StatusOK, map[string]interface{}{
			// "token": token,
			"result":    result,
			"login":     user.Login,
			"error":     fmt.Sprintf("%s", err),
			"errorCode": code,
		})
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			// "token": token,
			"result":    result,
			"login":     user.Login,
			"error":     "",
			"errorCode": code,
		})
	}
}
func (_ *AuthController) SignIn(c *gin.Context) {
	var user model.Auth = model.Auth{
		Login:    c.Query("login"),
		Password: c.Query("password"),
	}
	token, err, code := user.SignIn()
	if err != nil {
		c.JSON(http.StatusOK, map[string]interface{}{
			"token":     token,
			"login":     user.Login,
			"error":     fmt.Sprintf("%s", err),
			"errorCode": code,
		})
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"token":     token,
			"login":     user.Login,
			"error":     "",
			"errorCode": code,
		})
	}

}

// func (_ *AuthController) AddNewUser(c *gin.Context) {
// 	var user model.Auth = model.Auth{
// 		Login:    c.Query("login"),
// 		Email:    c.Query("email"),
// 		Password: c.Query("password"),
// 	}
// 	user.ID = 1000_000_000 + rand.Intn(1000_000_000)%1000_000_000

// 	user.Password = hashpasswd(user.Password)
// 	user.CreateUser()

// }
