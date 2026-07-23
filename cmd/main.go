package main

import (
	"fmt"
	"log"
	"net/http"

	"URL-Shortner/internal/config"
	"URL-Shortner/internal/db"
	"URL-Shortner/internal/handler"
	"URL-Shortner/internal/repository"
	"URL-Shortner/internal/service"

	_ "github.com/lib/pq"
)

func main() {
	cfg := config.LoadConfig()

	// Initialize DBs
	postgresDB, err := db.InitPostgres(cfg.PostgresDSN)
	if err != nil {
		log.Fatalf("Failed to connect to Postgres: %v", err)
	}
	defer postgresDB.Close()

	redisClient, err := db.InitRedis(cfg.RedisAddr, cfg.RedisPass, 0)
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	defer redisClient.Close()

	// Setup Repositories, Services, and Handlers
	pgRepo := repository.NewPostgresURLRepository(postgresDB)
	redisRepo := repository.NewRedisCacheRepository(redisClient)
	urlService := service.NewURLService(pgRepo, redisRepo)
	urlHandler := handler.NewURLHandler(urlService)

	// Routes
	mux := http.NewServeMux()
	mux.HandleFunc("/shorten", urlHandler.Shorten)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			handler.ServeHome(w, r)
			return
		}
		urlHandler.Redirect(w, r)
	})

	fmt.Printf("Server running on port %s...\n", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
