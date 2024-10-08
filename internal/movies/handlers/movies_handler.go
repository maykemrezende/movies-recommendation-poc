package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/maykemrezende/movies-recommendation-poc/internal/movies/entities"
	"go.uber.org/zap"
)

type MoviesHandler struct {
	Logger *zap.Logger
}

func NewMoviesHandler(logger *zap.Logger) *MoviesHandler {
	return &MoviesHandler{Logger: logger}
}

// Create godoc
// @Summary Get movies
// @Description Get all movies
// @Tags movies
// @Accept application/json
// @Produce application/json
// @Success 200
// @Router /movies [get]
func (h *MoviesHandler) GetMovies(c *gin.Context) {

	var movies []entities.Movie

	movies = append(movies, entities.Movie{
		ID:       1,
		Title:    "The Shawshank Redemption",
		Year:     1994,
		Genre:    "Drama",
		Director: "Frank Darabont",
	})

	c.JSON(200, movies)
}
