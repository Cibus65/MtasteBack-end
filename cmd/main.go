package main

import (
	"back-end/controllers"

	"github.com/gin-gonic/gin"
)

func RunServer() {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})
	//recipe breakpoints
	r.GET("Mtaste/API/getRecipeByPage/:page/:userID", (&controllers.RecipeController{}).GetRecipes)
	r.GET("Mtaste/API/getRecipeByID/:id/:userID", (&controllers.RecipeController{}).GetRecipe)
	r.GET("Mtaste/API/findRecipe/:words/:userID", (&controllers.RecipeController{}).FindRecipe)
	r.GET("Mtaste/API/getRandomRecipe/:userID", (&controllers.RecipeController{}).GetRandomRecipe)

	//auth breakpoints
	r.POST("Mtaste/API/auth/signUp", (&controllers.AuthController{}).SignUp)
	r.POST("Mtaste/API/auth/signIn", (&controllers.AuthController{}).SignIn)

	//user breakpoints

	r.POST("Mtaste/API/user/addToFavourite", (&controllers.UserController{}).AddToFavourite)
	r.POST("Mtaste/API/user/deleteFromFavourite", (&controllers.UserController{}).DeleteFromFavourite)
	r.GET("Mtaste/API/user/getFavouriteRecipes/:userID", (&controllers.UserController{}).GetFavouriteRecipes)
	r.Run("0.0.0.0:8082")
}
func main() {
	RunServer()
}
