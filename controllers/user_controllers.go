package controllers

import (
	"back-end/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserController struct{}

func (_ *UserController) AddToFavourite(c *gin.Context) {
	UserID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		panic(err)
	}
	link := c.Param("link")
	recipeID, err := strconv.Atoi(c.Param("recipeID"))
	if err != nil {
		c.Error(err)
	}
	message, err := (&model.User{}).AddFavourite(UserID, link, recipeID)
	if err != nil {
		c.Error(err)
	}
	c.JSON(http.StatusOK, message)

}
func (_ *UserController) AddNewUser(c *gin.Context) {
	login, email, password := c.Param("login"), c.Param("email"), c.Param("password")
	(&model.User{}).AddUser(login, email, password)
}
