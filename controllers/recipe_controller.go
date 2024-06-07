package controllers

import (
	"back-end/model"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type RecipeController struct{}

func (_ *RecipeController) GetRecipes(c *gin.Context) {

	var recipe model.Recipe
	page, err := strconv.Atoi(c.Param("page"))
	if err != nil {
		c.JSON(http.StatusNotFound, err)

	}
	recipes, err := recipe.GetByPage(page)
	if err != nil {
		c.JSON(http.StatusNotFound, err)

	}
	c.JSON(http.StatusOK, recipes)

}
func (_ *RecipeController) GetRecipe(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(err)
	}
	recipe, err := (&model.Recipe{}).GetByID(id)
	if err != nil {
		c.Error(err)
		c.Status(404)
		return
	}
	c.JSON(http.StatusOK, recipe)

}
func (_ *RecipeController) FindRecipe(c *gin.Context) {
	words := c.Param("words")
	words = strings.Replace(words, "+", " ", -1)
	recipes, err := (&model.Recipe{}).FindRecipe(words)
	if err != nil {
		log.Printf("Failed to find recipe with words: %s \n\tERROR: %s", words, err)
		c.JSON(http.StatusNotFound, nil)

	}
	c.JSON(http.StatusOK, recipes)
}

func (_ *RecipeController) GetRandomRecipe(c *gin.Context) {

	recipes, err := (&model.Recipe{}).GetRandomRecipe()
	if err != nil {
		log.Printf("Failed to get random recipe: \n\tERROR: %s", err)
		c.JSON(http.StatusNotFound, nil)

	}
	c.JSON(http.StatusOK, recipes)
}
