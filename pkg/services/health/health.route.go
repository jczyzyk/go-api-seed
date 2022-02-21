package health

import "github.com/gin-gonic/gin"

func Config(server *gin.Engine, ready <-chan struct{}) {
	server.GET("/health", GetHandler(ready))
}
