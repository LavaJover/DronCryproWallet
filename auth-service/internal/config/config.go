package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct{
	Env string `yaml:"env" env-required:"true" env-default:"local"`
	Dsn string `yaml:"dsn" env-required:"true"`
	GRPCServer `yaml:"grpc_server" env-required:"true"`
}

type GRPCServer struct{
	Host string `yaml:"host" env-default:"localhost" env-required:"true"`
	Port string `yaml:"port" env-required:"true"`
}

func MustLoad() *Config{
	var cfg Config

	configPath := os.Getenv("AUTH_CONFIG_PATH")

	if _, err := os.Stat(configPath); err != nil{
		log.Fatal("auth-service config file not found")
	}

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil{
		log.Fatal("failed to read config")
	}

	return &cfg
}