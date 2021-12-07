package main

import (
	"api/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	// create router
	router := gin.Default()

	// add route with handler
	router.GET("/", handlers.Index)
	router.GET("/page/:page", handlers.Page)

	router.POST("/recipes", handlers.NewRecipe)
	router.GET("/recipes/:id", handlers.GetRecipe)
	router.GET("/recipes", handlers.GetRecipes)
	router.PUT("/recipes/:id", handlers.UpdateRecipe)
	router.DELETE("/recipes/:id", handlers.DeleteRecipe)

	// run server
	err := router.Run(":5000")
	if err != nil {
		return
	}
}
