package config

import (
	"fmt"

	eventsenderrmq "github.com/reaport/ground-control/internal/service/event-sender-rmq"
	graphmap "github.com/reaport/ground-control/internal/service/graph-map"
	"github.com/spf13/viper"
)

type Config struct {
	Logger   *LoggerConfig          `mapstructure:"logger"`
	Server   *ServerConfig          `mapstructure:"server"`
	Map      *graphmap.Config       `mapstructure:"map"`
	RabbitMQ *eventsenderrmq.Config `mapstructure:"rabbitmq"`
}

type LoggerConfig struct {
	Level       string `mapstructure:"level"`
	Development bool   `mapstructure:"development"`
}

type ServerConfig struct {
	Port              int `mapstructure:"port"`
	ReadHeaderTimeout int `mapstructure:"read_header_timeout"`
}

func LoadConfig(path string) (*Config, error) {
	viper.SetConfigFile(path)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return &Config{}, fmt.Errorf("viper.ReadInConfig: %w", err)
	}

	config := &Config{}

	err = viper.Unmarshal(config)
	if err != nil {
		return &Config{}, fmt.Errorf("viper.Unmarshal: %w", err)
	}

	return config, nil
}
