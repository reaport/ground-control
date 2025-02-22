package config

import (
	"fmt"

	graphmap "github.com/reaport/ground-control/internal/service/graph-map"
	"github.com/spf13/viper"
)

type Config struct {
	Logger *LoggerConfig    `mapstructure:"logger"`
	Server *ServerConfig    `mapstructure:"server"`
	Map    *graphmap.Config `mapstructure:"map"`
}

type LoggerConfig struct {
	Level       string `mapstructure:"level"`
	Development bool   `mapstructure:"development"`
}

type ServerConfig struct {
	Port int `mapstructure:"port"`
}

func LoadConfig(path string) (config *Config, err error) {
	viper.SetConfigFile(path)

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return &Config{}, fmt.Errorf("viper.ReadInConfig: %w", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		return &Config{}, fmt.Errorf("viper.Unmarshal: %w", err)
	}

	return config, nil
}
