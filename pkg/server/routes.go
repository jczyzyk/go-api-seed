package server

import (
	"api-seed/pkg/services/health"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(server *gin.Engine) chan struct{} {
	ready := make(chan struct{})

	health.Config(server, ready)
	return ready
}
