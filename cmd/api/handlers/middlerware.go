package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/exercise/internal/pkg/pokemon"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// errorHandler process errors generated in the handler.
func errorHandler(c *gin.Context) {
	c.Next()
	if len(c.Errors) > 0 {
		for _, err := range c.Errors {
			code := getStatusError(err)
			c.JSON(code, Error{
				Code:    http.StatusText(code),
				Message: err.Error(),
			})
		}
	}
}

// getStatusError get the status code according to error.
func getStatusError(err error) int {
	var statusCode int
	var validationErr validator.ValidationErrors
	switch {
	case errors.As(err, &validationErr),
		errors.Is(err, strconv.ErrSyntax):
		statusCode = http.StatusUnprocessableEntity
	case errors.Is(err, pokemon.ErrNotFound):
		statusCode = http.StatusNotFound
	default:
		statusCode = http.StatusInternalServerError
	}

	return statusCode
}
