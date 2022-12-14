// Package app configures and runs application.
package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/evrone/go-clean-template/internal/usecase/repo/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	grpc_v1 "github.com/evrone/go-clean-template/internal/controller/grpc/v1"
	grpcserver "github.com/evrone/go-clean-template/pkg/grpc/server"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	"github.com/evrone/go-clean-template/config"
	v1 "github.com/evrone/go-clean-template/internal/controller/http/v1"
	"github.com/evrone/go-clean-template/internal/usecase/botanical_garden"
	"github.com/evrone/go-clean-template/internal/usecase/lemonade"
	"github.com/evrone/go-clean-template/pkg/httpserver"
	"github.com/evrone/go-clean-template/pkg/logger"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)
	var err error
	// Repository
	//pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	//if err != nil {
	//	l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	//}
	//defer pg.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.MONGO.Timeout)*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().SetMaxPoolSize(cfg.MONGO.PoolMax).ApplyURI(cfg.MONGO.URL))
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - mongo.Connect: %w", err))
	}
	defer func() {
		err := client.Disconnect(ctx)
		l.Warn(fmt.Errorf("app - Run - mongo.Disconnect: %w", err).Error())
	}()

	// Use case
	lemonadeGameUsecase := usecase.New(
		mongodb.NewGameRepository(client),
	)
	botanicalGardenUsecase := botanical_garden.New(
		mongodb.NewGameRepository(client),
	)
	// RabbitMQ RPC Server
	//rmqRouter := amqprpc.NewRouter(translationUseCase)

	//rmqServer, err := server.New(cfg.RMQ.URL, cfg.RMQ.ServerExchange, rmqRouter, l)
	//if err != nil {
	//	l.Fatal(fmt.Errorf("app - Run - rmqServer - server.New: %w", err))
	//}

	// HTTP Server
	handler := gin.New()
	v1.NewRouter(handler, l, *lemonadeGameUsecase, *botanicalGardenUsecase)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// GRPC Server
	grpcServer := grpc.NewServer()
	gameGRPCServer := grpcserver.NewGameLemonadeGRPCServer(
		l,
		grpcServer,
		*grpc_v1.NewLemonadeRoutes(*lemonadeGameUsecase, l),
		*grpc_v1.NewBotanicalGardenRoutes(*botanicalGardenUsecase, l),
	)
	go func() {
		_ = gameGRPCServer.StartGRPCServer(cfg.GRPC.Port)
	}()

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
