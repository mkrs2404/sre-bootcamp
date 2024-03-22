package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mkrs2404/sre-bootcamp/pkg/errors"
	"github.com/mkrs2404/sre-bootcamp/pkg/service"
)

type healthCheckHandler struct {
	HealthCheckService service.HealthCheckService
}

func NewHealthCheckHandler(hs service.HealthCheckService) *healthCheckHandler {
	return &healthCheckHandler{
		HealthCheckService: hs,
	}
}

func (h *healthCheckHandler) RegisterRoutes(router *gin.RouterGroup) {
	// route to get basic application health check (to check whether the app is up or not)
	router.GET("/health", h.getHealth)

	// route to perform deep health check that does a ping to the underlying db
	router.GET("/healthz", h.getHealthz)
}

// getHealth handles a request to get health.
func (h *healthCheckHandler) getHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

// getHealthz handles a request to get health of the service and other dependencies like database.
func (h *healthCheckHandler) getHealthz(c *gin.Context) {
	err := h.HealthCheckService.GetHealthz()
	if err != nil {
		c.JSON(errors.StatusCode(err), gin.H{
			"status": "unable to connect to db",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
