package config

import (
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	Server *Server
	Db     *Database `mapstructure:"database"`
}

type Server struct {
	Port int
}

type Database struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
	SSLMode  string
}

var (
	once           sync.Once
	configInstance *Config
)

func GetConfig() *Config {
	once.Do(
		func() {
			viper.SetConfigName("config")
			viper.SetConfigType("yaml")
			viper.AddConfigPath(".")

			if err := viper.ReadInConfig(); err != nil {
				panic(err)
			}

			if err := viper.Unmarshal(&configInstance); err != nil {
				panic(err)
			}
		},
	)
	return configInstance
}
