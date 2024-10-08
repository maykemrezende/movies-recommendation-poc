package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/maykemrezende/movies-recommendation-poc/internal/movies/handlers"
)

func InitMoviesRoutes(r *gin.Engine, moviesHandler *handlers.MoviesHandler) {
	r.GET("/api/v1/movies", moviesHandler.GetMovies)
}
