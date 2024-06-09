package controllers

import (
	"back-end/model"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (_ *UserController) Favourite(c *gin.Context) {
	var user_interface interface{}
	c.Bind(&user_interface)
	new := user_interface.(map[string]interface{})
	if new["userID"] == nil {
		new["userID"] = 162758239.
	}
	var user = model.User{
		UserId:   int(new["userID"].(float64)),
		RecipeID: int(new["recipeID"].(float64)),
	}
	result, flag, err, code := user.Favourite()

	if err != nil {
		c.JSON(404, map[string]interface{}{
			"result":    result,
			"flag":      flag,
			"userID":    user.UserId,
			"error":     "",
			"errorCode": code,
		})
	} else {

		c.JSON(http.StatusOK, map[string]interface{}{
			// "token": token,
			"result":    result,
			"flag":      flag,
			"userID":    user.UserId,
			"error":     fmt.Sprintf("%s", err),
			"errorCode": code,
		})
	}
}

func (_ *UserController) GetFavouriteRecipes(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		c.JSON(404, map[string]interface{}{
			"error":     err,
			"errorCode": 100,
			"recipes":   []model.Recipe{},
		})
	}
	var user = model.User{
		UserId:   userId,
		RecipeID: 0,
	}
	recipes, err, code := user.GetFavouriteRecipes()
	if err != nil {
		c.JSON(404, map[string]interface{}{
			"error":     err,
			"errorCode": code,
			"recipes":   []model.Recipe{},
		})
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"error":     err,
			"errorCode": code,
			"recipes":   recipes,
		})
	}

}
