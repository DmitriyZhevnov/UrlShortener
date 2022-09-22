package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/DmitriyZhevnov/UrlShortener/internal/config"
	"github.com/DmitriyZhevnov/UrlShortener/internal/handler"
	"github.com/DmitriyZhevnov/UrlShortener/internal/repository"
	"github.com/DmitriyZhevnov/UrlShortener/internal/service"
	"github.com/DmitriyZhevnov/UrlShortener/pkg/client/postgresql"
	"github.com/DmitriyZhevnov/UrlShortener/pkg/client/redis"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

const (
	maxAttemptsForConnectPostgres = 5
)

func main() {
	router := httprouter.New()

	cfg := config.GetConfig()

	postgresClient, err := postgresql.NewClient(context.Background(), maxAttemptsForConnectPostgres, cfg.Storage.Postgresql)
	if err != nil {
		panic(err)
	}

	redisClient, err := redis.NewClient(context.Background())
	if err != nil {
		panic(err)
	}

	storage := repository.NewRepository(postgresClient, redisClient)

	service := service.NewService(storage)

	handler := handler.NewHandler(service)
	handler.Register(router)

	startServer(router, cfg)
}

func startServer(router *httprouter.Router, cfg *config.Config) {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(server.Serve(listener))
}
