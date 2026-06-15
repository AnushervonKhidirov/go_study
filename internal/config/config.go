package config

import (
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type AppConfig struct {
	Env              string `yaml:"env" env-default:"local"`
	HttpServerConfig `yaml:"http_server"`
	DatabaseConfig   `yaml:"database"`
}

type HttpServerConfig struct {
	Address     string `yaml:"address" env-default:"localhost:8080"`
	Timeout     string `yaml:"timeout" env-default:"4s"`
	IdleTimeout string `yaml:"idle_timeout" env-default:"20s"`
}

type DatabaseConfig struct {
	DBHost     string `yaml:"host" env-required:"true"`
	DBPort     uint   `yaml:"port" env-required:"true"`
	DBUser     string `yaml:"user" env-required:"true"`
	DBPassword string `yaml:"password" env-required:"true"`
	DBName     string `yaml:"dbname" env-required:"true"`
}

func InitConfig() *AppConfig {
	err := godotenv.Load()

	if err != nil {
		panic("Enable load .env file")
	}

	configPath := os.Getenv("CONFIG_PATH")

	var cfg AppConfig
	err = cleanenv.ReadConfig(configPath, &cfg)

	if err != nil {
		panic("Enable to load config")
	}

	return &cfg
}
