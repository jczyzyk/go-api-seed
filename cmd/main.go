package main

import (
	"api-seed/pkg/config"
	"api-seed/pkg/server"

	log "github.com/rs/zerolog"
)

func main() {
	log.TimeFieldFormat = log.TimeFormatUnix
	config.Init()
	server.Init()
}
