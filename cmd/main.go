package main

import (
	"back-end/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("Mtaste/API/getRecipe/:page", (&controllers.RecipeController{}).GetRecipes)
	r.Run()
}
