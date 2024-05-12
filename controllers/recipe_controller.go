package controllers

import (
	"back-end/model"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type RecipeController struct{}

func (_ *RecipeController) GetRecipes(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")                   // Разрешаем все домены
	c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS") // Разрешаем определенные методы
	c.Header("Access-Control-Allow-Headers", "Content-Type")       //Разрешаем определенные заголовки

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
	c.Header("Access-Control-Allow-Origin", "*")                   // Разрешаем все домены
	c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS") // Разрешаем определенные методы
	c.Header("Access-Control-Allow-Headers", "Content-Type")
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
	c.Header("Access-Control-Allow-Origin", "*")                   // Разрешаем все домены
	c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS") // Разрешаем определенные методы
	c.Header("Access-Control-Allow-Headers", "Content-Type")
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

	c.Header("Access-Control-Allow-Origin", "*")                   // Разрешаем все домены
	c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS") // Разрешаем определенные методы
	c.Header("Access-Control-Allow-Headers", "Content-Type")
	recipe, err := (&model.Recipe{}).GetRandomRecipe()
	if err != nil {
		log.Printf("Failed to get random recipe: \n\tERROR: %s", err)
		c.JSON(http.StatusNotFound, nil)

	}
	c.JSON(http.StatusOK, recipe)
}
