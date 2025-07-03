package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerAddress string `env:"SERVER_ADDRESS" env-default:":12345"`
	// DBAddr        string `env:"DB_ADDR"`
}

func Load() *Config {
	var cfg Config
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file!: ", err.Error())
	}

	cfg.ServerAddress, _ = os.LookupEnv("SERVER_ADDR")

	return &cfg
}
