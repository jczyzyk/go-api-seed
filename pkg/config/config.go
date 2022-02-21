package config

import (
	"api-seed/pkg/constants"
	"os"

	"github.com/spf13/viper"
)

// Init is an exported method that takes the environment starts the viper
// (external lib) and returns the configuration struct.
func Init() {
	loadDefaultConf()

	switch env := os.Getenv(constants.Env); env {
	case constants.Dev:
		loadDevConf()
	case constants.Prod:
		loadProdConf()
	default:
		loadDevConf()
	}

	viper.AutomaticEnv()
}
