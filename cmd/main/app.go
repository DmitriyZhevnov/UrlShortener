package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/DmitriyZhevnov/UrlShortener/internal/config"
	"github.com/DmitriyZhevnov/UrlShortener/internal/handler"
	"github.com/DmitriyZhevnov/UrlShortener/internal/repository"
	"github.com/DmitriyZhevnov/UrlShortener/internal/server"
	"github.com/DmitriyZhevnov/UrlShortener/internal/service"
	"github.com/DmitriyZhevnov/UrlShortener/pkg/client/postgresql"
	"github.com/DmitriyZhevnov/UrlShortener/pkg/client/redis"
	"github.com/DmitriyZhevnov/UrlShortener/pkg/logging"
	"github.com/DmitriyZhevnov/UrlShortener/pkg/utils"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	maxAttemptsForConnectPostgres = 5
)

// @title URL Shortener
// @version 1.0
// @description REST API for URL Shortener

// @host localhost:8080
// @BasePath /
func main() {
	log := logging.GetLogger()

	ctx := context.Background()

	router := mux.NewRouter()

	cfg := config.GetConfig()

	linkHasher := utils.NewLinkHasher(cfg.Domain)

	postgresClient, err := postgresql.NewClient(ctx, maxAttemptsForConnectPostgres, cfg.Storage.Postgresql)
	if err != nil {
		panic(err)
	}

	redisClient, err := redis.NewClient(ctx, cfg.Storage.Redis)
	if err != nil {
		panic(err)
	}

	storage := repository.NewRepository(postgresClient, redisClient)

	service := service.NewService(log, storage, linkHasher)

	handler := handler.NewHandler(service, cfg.Domain)
	handler.Register(router)

	srv := server.NewServer(router, cfg.HTTP)

	go func() {
		if err := srv.Run(log); !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(fmt.Sprintf("error occurred while running http server: %s\n", err.Error()), nil)
		}
	}()

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	postgresClient.Close()

	if err := redisClient.Close(); err != nil {
		log.Fatal(err.Error(), nil)
	}

	if err := srv.Stop(ctx); err != nil {
		log.Fatal(fmt.Sprintf("failed to stop server: %v", err), nil)
	}
}
