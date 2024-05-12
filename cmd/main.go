package main

import (
	"back-end/controllers"

	"github.com/gin-gonic/gin"
)

func RunServer() {
	r := gin.Default()

	r.GET("Mtaste/API/getRecipeByPage/:page", (&controllers.RecipeController{}).GetRecipes)
	r.GET("Mtaste/API/getRecipeByID/:id", (&controllers.RecipeController{}).GetRecipe)
	r.GET("Mtaste/API/findRecipe/:words", (&controllers.RecipeController{}).FindRecipe)
	r.GET("Mtaste/API/getRandomRecipe", (&controllers.RecipeController{}).GetRandomRecipe)
	r.POST("Mtaste/API/user/signUp/", (&controllers.UserController{}).SignUp)
	r.POST("Mtaste/API/user/signIn/", (&controllers.UserController{}).SignIn)

	r.Run("0.0.0.0:8082")
}
func main() {
	RunServer()
}
