package app

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ramonmedeiros/iba/internal/pkg/recipes"
)

const (
	IDParam = "id"

	FieldParam = "field"
	ValueParam = "value"
)

func (s *Server) setupEndpoints() {
	usersEndpoint := s.router.Group("/cocktails")
	usersEndpoint.GET("", s.getCocktails)
}

func (s *Server) getCocktails(c *gin.Context) {
	results, err := s.recipeService.Search(
		c.QueryArray("term"),
		c.QueryArray("notIncluded"),
	)
	if errors.Is(err, recipes.ErrNotFound) {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.IndentedJSON(http.StatusOK, results)
}
