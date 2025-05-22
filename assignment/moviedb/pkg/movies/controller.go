package movies

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type MoviesController struct {
	store movieStore
}

func NewMoviesController(s movieStore) *MoviesController {
	return &MoviesController{
		store: s,
	}
}

type movieStore interface {
	Add(movie CreateMovieReq) error
	Update(id uuid.UUID, movie UpdateMovieReq) error
	Delete(id uuid.UUID) error
	List() ([]Movie, error)
}

func (h MoviesController) CreateMovie(c *gin.Context) {
	fmt.Println("Running add movie")

	var movie CreateMovieReq
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.store.Add(movie)
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (h MoviesController) DeleteMovie(c *gin.Context) {
	movieId := c.Param("id")
	if _, err := uuid.Parse(movieId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	parsedMovieId, _ := uuid.Parse(movieId)
	h.store.Delete(parsedMovieId)
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (h MoviesController) UpdateMovie(c *gin.Context) {
	movieId := c.Param("id")
	if _, err := uuid.Parse(movieId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var movie UpdateMovieReq
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	parsedMovieId, _ := uuid.Parse(movieId)

	h.store.Update(parsedMovieId, movie)
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (h MoviesController) ListMovies(c *gin.Context) {
	r, err := h.store.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(200, r)
}

func (h MoviesController) Head(c *gin.Context) {
	c.Header("X-Custom-Header", "Test")
	c.Status(http.StatusOK)
}

func (h MoviesController) Options(c *gin.Context) {
	c.Header("Allow", "GET, POST, DELETE, PUT, OPTIONS, HEAD")
	c.Status(http.StatusOK)
}
