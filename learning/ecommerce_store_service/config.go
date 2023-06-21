package main

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	ConnectionString string `mapstructure:"connection_string"`
	Port             int    `mapstructure:"port"`
	SecretKey        string `mapstructure:"secret_key"`
}

var AppConfig *Config

func LoadConfig() {
	log.Println("Loading Server Configuration ...")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Fatal()
	}
	log.Println("Server Configuration has been Loaded successfully!")
}
