package main

import (
	"back-end/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("Mtaste/API/getRecipe/:page", (&controllers.RecipeController{}).GetRecipes)
	r.GET("Mtaste/API/user/addTofavourite/:userID/:link/:recipeID", (&controllers.UserController{}).AddToFavourite)
	r.GET("Mtaste/API/user/addUser/:login/:email/:password", (&controllers.UserController{}).AddNewUser)

	r.Run()
}
