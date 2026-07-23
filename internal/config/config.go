package config

import (
	"os"
)

type Config struct {
	Port        string
	PostgresDSN string
	RedisAddr   string
	RedisPass   string
}

func LoadConfig() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	dsn := os.Getenv("POSTGRES_DSN")
	if dsn == "" {
		dsn = "postgres://postgres:postgres@localhost:5432/urlshortener?sslmode=disable"
	}

	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		redisAddr = "localhost:6379"
	}

	redisPass := os.Getenv("REDIS_PASS")

	return &Config{
		Port:        port,
		PostgresDSN: dsn,
		RedisAddr:   redisAddr,
		RedisPass:   redisPass,
	}
}
