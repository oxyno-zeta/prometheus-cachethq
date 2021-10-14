package middlewares

import (
	"context"

	"github.com/gin-gonic/gin"
	uuid "github.com/gofrs/uuid"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/common/utils"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/log"
	"github.com/pkg/errors"
)

type contextKey struct {
	name string
}

var reqCtxKey = &contextKey{name: "request-id"}

const requestIDHeader = "X-Request-Id"
const requestIDContextKey = "RequestID"

func RequestID(logger log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get request id from request
		requestID := c.Request.Header.Get(requestIDHeader)

		// Check if request id exists
		if requestID == "" {
			// Generate uuid
			uuid, err := uuid.NewV4()
			// Check error
			if err != nil {
				// Add stack trace to error
				err2 := errors.WithStack(err)
				// Log error
				logger.Error(err2)
				// Send response
				utils.AnswerWithError(c, err2)

				return
			}
			// Save it in variable
			requestID = uuid.String()
		}

		// Store it in context
		c.Set(requestIDContextKey, requestID)
		// Update request with new context
		c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), reqCtxKey, requestID))

		// Put it on header
		c.Writer.Header().Set(requestIDHeader, requestID)

		// Next
		c.Next()
	}
}

func GetRequestIDFromGin(c *gin.Context) string {
	requestIDObj, requestIDExists := c.Get(requestIDContextKey)
	if requestIDExists {
		// return request id
		return requestIDObj.(string)
	}

	return ""
}

func GetRequestIDFromContext(ctx context.Context) string {
	requestIDObj := ctx.Value(reqCtxKey)
	if requestIDObj != nil {
		return requestIDObj.(string)
	}

	return ""
}
