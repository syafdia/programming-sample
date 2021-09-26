package http

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/syafdia/sb/chal002/internal/domain/entity"
	"github.com/syafdia/sb/chal002/internal/domain/usecase"
)

type MovieHandler interface {
	FindOne(c *gin.Context)
	FindMultiple(c *gin.Context)
}

type movieHandler struct {
	findOneMovieUseCase       usecase.FindOneMovieUseCase
	findMultipleMoviesUseCase usecase.FindMultipleMoviesUseCase
}

func NewMovieHandler(
	findOneMovieUseCase usecase.FindOneMovieUseCase,
	findMultipleMoviesUseCase usecase.FindMultipleMoviesUseCase,
) MovieHandler {
	return &movieHandler{
		findOneMovieUseCase:       findOneMovieUseCase,
		findMultipleMoviesUseCase: findMultipleMoviesUseCase,
	}
}

func (m *movieHandler) FindOne(c *gin.Context) {
	result, err := m.findOneMovieUseCase.Execute(c.Request.Context(), c.Param("id"))
	if err != nil {
		m.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (m *movieHandler) FindMultiple(c *gin.Context) {
	searchWord := c.Query("search_word")
	pagination, err := strconv.Atoi(c.Query("pagination"))
	if err != nil {
		pagination = 1
	}

	results, err := m.findMultipleMoviesUseCase.Execute(c.Request.Context(), searchWord, pagination)
	if err != nil {
		m.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, results)
}

func (m *movieHandler) handleError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, entity.ErrUnauthorized):
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
	case errors.Is(err, entity.ErrBadGateway):
		c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}
}
