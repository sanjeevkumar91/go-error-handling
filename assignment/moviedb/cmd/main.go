package main

import (
	"moviedb/pkg/movies"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	router.Use(movies.LoggerMiddleware())
	store := movies.NewMemStore()
	moviesController := movies.NewMoviesController(store)

	authGroup := router.Group("/api")
	authGroup.Use(movies.AuthMiddleware())
	{
		authGroup.GET("/movies", moviesController.ListMovies)
		authGroup.POST("/movies", moviesController.CreateMovie)
		authGroup.DELETE("/movies/:id", moviesController.DeleteMovie)
		authGroup.PUT("/movies/:id", moviesController.UpdateMovie)
		authGroup.HEAD("/movies", moviesController.Head)
		authGroup.OPTIONS("/movies", moviesController.Options)
	}

	router.Run(":8080")
}
