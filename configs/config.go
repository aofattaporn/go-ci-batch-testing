package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

type IConfig interface {
}

type config struct {
}

func LoadConfig(path string) (IConfig, error) {

	if path == "" {
		return nil, fmt.Errorf("config file is empty")
	}

	viper.SetConfigFile(path)
	viper.AutomaticEnv()

	// read config file
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed read config file: %v", err)
	}

	config := &config{}

	return config, nil
}
