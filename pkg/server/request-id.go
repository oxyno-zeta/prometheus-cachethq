package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

const requestIDHeader = "X-Request-Id"
const requestIDContextKey = "RequestID"

func requestID(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get request id from request
		requestID := c.Request.Header.Get(requestIDHeader)

		// Check if request id exists
		if requestID == "" {
			// Generate uuid
			uuid, err := uuid.NewV4()
			if err != nil {
				// Log error
				logger.Errorln(err)
				// Send response
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			// Save it in variable
			requestID = uuid.String()
		}

		// Store it in context
		c.Set(requestIDContextKey, requestID)

		// Put it on header
		c.Writer.Header().Set(requestIDHeader, requestID)

		// Next
		c.Next()
	}
}
