package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct{
	Env string	`yaml:"env" env-default:"local" env-required:"true"`
	HttpServer 	`yaml:"http_server"`
}

type HttpServer struct{
	Address string 			`yaml:"address" env-default:"localhost:8080"`
	Timeout time.Duration 	`yaml:"timeout" env-default:"4s"`
}

func MustLoad() *Config{
	configPath := os.Getenv("CONFIG_PATH")

	if configPath == ""{
		log.Fatal("config path required")
	}

	if _, err := os.Stat(configPath); err != nil{
		log.Fatalf("failed to load config file: %v", err)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil{
		log.Fatalf("failed to read config file: %v", err)
	}

	return &cfg
}