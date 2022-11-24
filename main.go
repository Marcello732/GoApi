package main

import (
	"net/http"

	"github.com/Marcello732/goApi/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := setupRouter()
	_ = r.Run(":8080")
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	userRepo := controllers.NewUserRepo()
	r.POST("/users", userRepo.CreateUser)
	r.GET("/users", userRepo.GetUsers)
	r.GET("/users/:id", userRepo.GetUser)
	r.PUT("/users/:id", userRepo.UpdateUser)
	r.DELETE("/users/:id", userRepo.DeleteUser)

	recipeRepo := controllers.NewRecipeRepo()
	r.POST("/recipes", recipeRepo.CreateRecipe)
	r.GET("/recipes", recipeRepo.GetRecipes)
	r.GET("/recipes/:id", recipeRepo.GetRecipe)
	r.PUT("/recipes/:id", recipeRepo.UpdateRecipe)
	r.DELETE("/recipes/:id", recipeRepo.DeleteRecipe)

	return r
}