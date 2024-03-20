package main

import (
	"back-end/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/getRecipe/:page", (&controllers.RecipeController{}).GetRecipes)
	r.GET("/user/add/favourite/:link/:id", (&controllers.UserController{}).AddToFavourite)
	r.Run()
}
