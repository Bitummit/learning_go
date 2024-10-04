package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Env string `yaml:"env" env-default:"local"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer `yaml:"http_server"`
}

type HTTPServer struct {
	Address string `yaml:"address" env-default:"localhost:8000"`
	Timeout string `yaml:"timeout" env-default:"2s"`
	IdleTimeout string `yaml:"idle_timeout" env-default:"60s"`

}

func MustLoad() *Config{
	if err := godotenv.Load(); err != nil {
        log.Print("No .env file found", err)
    }

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is empty")
	}

	if _, err := os.Stat(configPath); err != nil {
		log.Fatal("No config file found")
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalln("cannot read config file")
	}

	return &cfg
}