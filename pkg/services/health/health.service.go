package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var readyResponse gin.H = gin.H{"message": "ok"}
var notReadyResponse gin.H = gin.H{"message": "not-ready"}

func GetHandler(ready <-chan struct{}) func(ctx *gin.Context) {
	isReady := true

	/*************** INFRASTRUCTURE ********************
	In case the endpoint /health is use to check
	the availability and the readiness of the service
	we can use a channel to inform that everything is
	ready so then	turn the "isReady" variable to true.

	isReady := false
	go func() {
		<-ready
		isReady = true
	}()
	**************************************************/

	return func(ctx *gin.Context) {
		if isReady {
			ctx.JSON(http.StatusOK, readyResponse)

			return
		}
		ctx.JSON(http.StatusTooEarly, notReadyResponse)
	}
}
