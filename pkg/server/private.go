package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/errors"
)

func handleError(c *gin.Context, err error) {
	// Try to match a business error
	errMatch, ok := err.(*errors.GeneralError)
	if ok {
		// Manage general error cases
		// Bad input case
		if errMatch.ErrorType == errors.BadInputErrorType {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// Not found case
		if errMatch.ErrorType == errors.NotFoundErrorType {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		// Default case
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}
