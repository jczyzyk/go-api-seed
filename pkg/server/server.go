package server

import (
	"api-seed/pkg/constants"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	log "github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func Init() {
	if isProd() {
		gin.SetMode(gin.ReleaseMode)
	}

	server := gin.New()

	server.Use(gin.LoggerWithWriter(gin.DefaultWriter))
	server.Use(gin.Recovery())
	SetupRoutes(server)

	err := server.Run(fmt.Sprintf(":%s", viper.GetString(constants.Port)))
	if err != nil {
		log.Error().Err(err).Msg(fmt.Sprintf("error while running server: %v", err.Error()))
		defer os.Exit(1)
	}
}

func isProd() bool {
	return os.Getenv(constants.Env) == constants.Prod
}
