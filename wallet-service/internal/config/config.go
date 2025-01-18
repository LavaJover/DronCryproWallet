package config

import (
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type WalletConfig struct{
	APIKey string `yaml:"api_key"`
}

func MustLoad(configPath string) *WalletConfig{
	if configPath == ""{
		panic("Config path is empty!")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err){
		panic("Config file does not exist " + configPath)
	}

	var cfg WalletConfig

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil{
		panic("config path is empty " + err.Error())
	}

	return &cfg
}