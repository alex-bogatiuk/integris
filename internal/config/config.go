package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Database DatabaseConfig
}

type DatabaseConfig struct {
	Type     string
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

var Cfg *Config

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	err = viper.Unmarshal(&Cfg)
	if err != nil {
		log.Fatalf("Unable to decode into struct %s", err)
	}
}
