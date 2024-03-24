package main

import (
	"back-end/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("Mtaste/API/getRecipeByPage/:page", (&controllers.RecipeController{}).GetRecipes)
	r.GET("Mtaste/API/getRecipeByID/:id", (&controllers.RecipeController{}).GetRecipe)
	r.Run()
}
