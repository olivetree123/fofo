package config

import (
	. "fofo/common"
	"github.com/spf13/viper"
)

var Config *viper.Viper

func init() {
	Config = viper.New()
	Config.SetConfigFile("/etc/fofo/config.toml")
	err := Config.ReadInConfig()
	if err != nil {
		Logger.Error(err)
		return
	}
}
