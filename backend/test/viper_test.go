package test

import (
	"github.com/spf13/viper"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	viper.SetConfigFile("./config.yaml")
	endpoint := viper.GetString("server.endpoint")
}
