package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/business"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/business/prometheushook/models"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/common/utils"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/log"
)

func AddRESTEndpoint(router *gin.Engine, businessServices *business.Services) {
	router.POST("/prometheus/webhook", func(c *gin.Context) {
		// Get logger from request
		reqLogger := log.GetLoggerFromGin(c)

		var alerts models.PrometheusAlertHook
		// Try to map data
		err := c.ShouldBindJSON(&alerts)
		// Check if error exists
		if err != nil {
			reqLogger.Error(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

			return
		}

		err = businessServices.PrometheusHookSvc.ManageHook(&alerts)
		// Check error
		if err != nil {
			reqLogger.Error(err)
			utils.AnswerWithError(c, err)

			return
		}

		c.Status(http.StatusNoContent)
	})
}
