package main

import (
	"back-end/controllers"

	"github.com/gin-gonic/gin"
)

func RunServer() {
	r := gin.Default()
	//recipe breakpoints
	r.GET("Mtaste/API/getRecipeByPage/:page", (&controllers.RecipeController{}).GetRecipes)
	r.GET("Mtaste/API/getRecipeByID/:id", (&controllers.RecipeController{}).GetRecipe)
	r.GET("Mtaste/API/findRecipe/:words", (&controllers.RecipeController{}).FindRecipe)
	r.GET("Mtaste/API/getRandomRecipe", (&controllers.RecipeController{}).GetRandomRecipe)

	//auth breakpoints
	r.POST("Mtaste/API/auth/signUp", (&controllers.AuthController{}).SignUp)
	r.POST("Mtaste/API/auth/signIn", (&controllers.AuthController{}).SignIn)

	//user breakpoints

	r.POST("Mtaste/API/user/addToFavourite", (&controllers.UserController{}).AddToFavourite)
	r.Run("0.0.0.0:8082")
}
func main() {
	RunServer()
}
