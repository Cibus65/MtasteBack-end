package main

import (
	"back-end/controllers"
	"github.com/gin-gonic/gin"
)

func RunServer() {
	var engine *gin.Engine
	engine = gin.Default()
	engine.GET("Mtaste/API/getRecipe/:page", (&controllers.RecipeController{}).GetRecipes)
	engine.Run("0.0.0.0:8080")
}

func main() {
	RunServer()
	//r.GET("Mtaste/API/user/addTofavourite/:userID/:link/:recipeID", (&controllers.UserController{}).AddToFavourite)
	//r.POST("Mtaste/API/user/addUser", (&controllers.UserController{}).AddNewUser)
	//r.POST("Mtaste/API/user/SignIn", (&controllers.UserController{}).SignIn)
}
