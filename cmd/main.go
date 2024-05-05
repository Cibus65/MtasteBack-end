package main

import (
	"back-end/controllers"
	"github.com/gin-gonic/gin"
)

func RunServer() {
	r := gin.Default()

	r.GET("Mtaste/API/getRecipeByPage/:page", (&controllers.RecipeController{}).GetRecipes)
	r.GET("Mtaste/API/getRecipeByID/:id", (&controllers.RecipeController{}).GetRecipe)
	r.Run("0.0.0.0:10080")
}
func main() {
	RunServer()
}
