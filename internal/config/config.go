package config

import (
	"os"
)

type Config struct {
	ServerAddress string //`env:"SERVER_ADDR" env-default:":12345"`
	// DBAddr        string `env:"DB_ADDR"`
}

func Load() *Config {
	var cfg Config
	//вне контейнера подгружает, в контейнере не нужно
	// if err := godotenv.Load(); err != nil {
	// 	log.Fatal("No .env file!: ", err.Error())
	// }

	cfg.ServerAddress, _ = os.LookupEnv("SERVER_ADDR")
	cfg.ServerAddress = ":" + cfg.ServerAddress

	return &cfg
}
