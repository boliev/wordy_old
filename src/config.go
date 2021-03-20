package main

import (
	"fmt"

	"github.com/spf13/viper"
)

// NewConfig initialisation
func NewConfig() {
	viper.SetConfigName("config.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s ", err))
	}
}
