package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/reaport/ground-control/internal/config"
	"github.com/reaport/ground-control/internal/controller"
	"github.com/reaport/ground-control/internal/entity"
	eventsenderrmq "github.com/reaport/ground-control/internal/service/event-sender-rmq"
	graphmap "github.com/reaport/ground-control/internal/service/graph-map"
	"github.com/reaport/ground-control/pkg/api"
	"github.com/reaport/ground-control/pkg/logger"
	"github.com/reaport/ground-control/pkg/server/middlewares"
	"go.uber.org/zap"
)

const (
	shutdownTimeout = 5 * time.Second
)

func main() {
	configPath := flag.String("config", "config.yaml", "config path")
	flag.Parse()

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		panic(err)
	}

	err = logger.InitLogger(cfg.Logger.Level, cfg.Logger.Development)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = logger.GlobalLogger.Sync()
	}()

	service, err := graphmap.New(cfg.Map)
	if err != nil {
		logger.GlobalLogger.Fatal("failed to initialize graph map service", zap.String("error", err.Error()))
	}

	eventSender, err := eventsenderrmq.New(cfg.RabbitMQ)
	if err != nil {
		logger.GlobalLogger.Fatal("failed to initialize event sender", zap.String("error", err.Error()))
	}

	err = eventSender.SendEvent(context.Background(), &entity.Event{
		Type: entity.MapRefreshedEventType,
	})
	if err != nil {
		logger.GlobalLogger.Error(
			"failed to send event",
			zap.Error(fmt.Errorf("c.eventSender.SendEvent: %w", err)),
			zap.String("event_type", string(entity.MapRefreshedEventType)),
		)
	}

	ctrl := controller.New(service, eventSender)

	server, err := api.NewServer(ctrl, api.WithErrorHandler(middlewares.ErrorHandler))
	if err != nil {
		logger.GlobalLogger.Fatal("failed to create server", zap.String("error", err.Error()))
	}

	httpServer := &http.Server{
		Addr:              fmt.Sprintf(":%d", cfg.Server.Port),
		ReadHeaderTimeout: time.Duration(cfg.Server.ReadHeaderTimeout) * time.Second,
		Handler:           server,
	}

	httpServer.Handler = CORSMiddleware(httpServer.Handler)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		logger.GlobalLogger.Info("server is starting", zap.Int("port", cfg.Server.Port))
		err = httpServer.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.GlobalLogger.Fatal("server failed to start", zap.String("error", err.Error()))
		}
	}()

	<-done
	logger.GlobalLogger.Info("server is shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	err = httpServer.Shutdown(ctx)
	if err != nil {
		logger.GlobalLogger.Error("failed to shutdown server gracefully", zap.String("error", err.Error()))
	} else {
		logger.GlobalLogger.Info("server shutdown completed")
	}
}

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
