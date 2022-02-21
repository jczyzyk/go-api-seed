package config

import (
	"api-seed/pkg/constants"

	"github.com/spf13/viper"
)

func loadDefaultConf() {
	viper.SetDefault(constants.Port, "8002")
}
