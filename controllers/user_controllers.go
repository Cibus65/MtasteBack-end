package controllers

import (
	"back-end/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserController struct{}

func (_ *UserController) AddToFavourite(c *gin.Context) {
	link := c.Param("link")
	recipeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(err)
	}
	message, err := (&model.User{
		Login:        "xapsiel2",
		Email:        "xapsiel@mail.ru",
		Password:     "1234",
		ID:           1,
		SavedRecipes: nil,
	}).AddFavourite(link, recipeID)
	if err != nil {
		c.Error(err)
	}
	c.JSON(http.StatusOK, message)

}
