package controllers

import (
	"back-end/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type RecipeController struct{}

func (_ *RecipeController) GetRecipes(c *gin.Context) {
	var recipe model.Recipe
	page, err := strconv.Atoi(c.Param("page"))
	if err != nil {
		c.Error(err)
	}
	recipes, err := recipe.Get(page)
	if err != nil {
		c.Error(err)
	}
	c.JSON(http.StatusOK, recipes)

}
