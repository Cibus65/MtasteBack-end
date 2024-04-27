package controllers

import (
	"back-end/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"strconv"
)

type RecipeController struct{}

func (_ *RecipeController) GetRecipes(c *gin.Context) {
	file, err := os.OpenFile("../app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Printf("ERROR: %s", err)
		c.JSON(http.StatusNotFound, nil)
	}
	log.SetOutput(file)
	var recipe model.Recipe
	page, err := strconv.Atoi(c.Param("page"))
	if err != nil {
		log.Println("Page must be integer")
		c.JSON(http.StatusNotFound, nil)

	}
	recipes, err := recipe.GetByPage(page)
	if err != nil {
		log.Printf("Failed to get info into %d page\n\tERROR: %s", page, err)
		c.JSON(http.StatusNotFound, nil)

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
