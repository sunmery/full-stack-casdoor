package test

import (
	"fmt"
	"github.com/spf13/viper"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	viper.SetConfigType("yaml")
	viper.SetConfigFile("./app.yaml")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	endpoint := viper.GetString("server.endpoint")

	t.Log("endpoint:", endpoint)
}
