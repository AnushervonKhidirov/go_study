package config

import (
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type AppConfig struct {
	Env              string `yaml:"env" env-default:"local"`
	HttpServerConfig `yaml:"http_server"`
}

type HttpServerConfig struct {
	Address     string `yaml:"address" env-default:"localhost:8080"`
	Timeout     string `yaml:"timeout" env-default:"4s"`
	IdleTimeout string `yaml:"idle_timeout" env-default:"20s"`
}

func InitConfig() *AppConfig {
	err := godotenv.Load()

	if err != nil {
		panic("Enable load .env file")
	}

	var cfg AppConfig

	configPath := os.Getenv("CONFIG_PATH")

	fmt.Println("configPath", configPath)

	err = cleanenv.ReadConfig(configPath, &cfg)

	if err != nil {
		panic("Enable to load config")
	}

	return &cfg
}
