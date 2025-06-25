package config

import (
	"log"

	"github.com/spf13/viper"
)

type config struct {
	PORT        string
	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
}

var ENV *config

func LoadConfig() {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic("Error reading config file: "+ err.Error())
	}

	ENV = &config{
		PORT: viper.GetString("PORT"),
		DB_HOST: viper.GetString("DB_HOST"),
		DB_PORT: viper.GetString("DB_PORT"),
		DB_USER: viper.GetString("DB_USER"),
		DB_PASSWORD: viper.GetString("DB_PASSWORD"),
		DB_NAME: viper.GetString("DB_NAME"),
	}

	if ENV.PORT == "" {
		log.Fatal("port is required but not found")
	}
	if ENV.DB_HOST == "" {
		log.Fatal("DB_HOST is required but not set")
	}

	if err := viper.Unmarshal(&ENV); err != nil {
		panic("Error unmarshalling config: " + err.Error())
	}
}