package controllers

import (
	"back-end/model"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (_ *UserController) AddToFavourite(c *gin.Context) {
	userid, _ := strconv.Atoi(c.Query("userID"))
	recipeid, _ := strconv.Atoi(c.Query("recipeID"))
	var user = model.User{
		UserId:   userid,
		RecipeID: recipeid,
	}
	result, flag, err, code := user.AddToFavourite()
	if err != nil {
		c.JSON(http.StatusOK, map[string]interface{}{
			// "token": token,
			"result":    result,
			"flag":      flag,
			"userID":    user.UserId,
			"error":     fmt.Sprintf("%s", err),
			"errorCode": code,
		})
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"result":    result,
			"flag":      flag,
			"userID":    user.UserId,
			"error":     "",
			"errorCode": code,
		})
	}
}
