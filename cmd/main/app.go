package main

import (
	"context"
	"errors"
	"log"
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
	"github.com/DmitriyZhevnov/UrlShortener/pkg/utils"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

const (
	maxAttemptsForConnectPostgres = 5
)

func main() {
	router := httprouter.New()

	cfg := config.GetConfig()

	linkHasher := utils.NewLinkHasher(cfg.HashSalt)

	postgresClient, err := postgresql.NewClient(context.Background(), maxAttemptsForConnectPostgres, cfg.Storage.Postgresql)
	if err != nil {
		panic(err)
	}

	redisClient, err := redis.NewClient(context.Background())
	if err != nil {
		panic(err)
	}

	storage := repository.NewRepository(postgresClient, redisClient)

	service := service.NewService(storage, linkHasher)

	handler := handler.NewHandler(service)
	handler.Register(router)

	srv := server.NewServer(router)

	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("error occurred while running http server: %s\n", err.Error())
		}
	}()

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := postgresClient.Close(); err != nil {
		log.Fatalf(err.Error())
	}

	if err := redisClient.Close(); err != nil {
		log.Fatalf(err.Error())
	}

	if err := srv.Stop(ctx); err != nil {
		log.Fatalf("failed to stop server: %v", err)
	}
}
