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
	graphmap "github.com/reaport/ground-control/internal/service/graph-map"
	"github.com/reaport/ground-control/pkg/api"
	"github.com/reaport/ground-control/pkg/logger"
	"go.uber.org/zap"
)

const (
	shutdownTimeout = 5 * time.Second
)

func main() {
	configPath := flag.String("config", "config.yaml", "config path")
	initMapPath := flag.String("init-data", "init_data.json", "init map path")
	flag.Parse()

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		logger.GlobalLogger.Fatal("failed to load config", zap.Error(err))
	}

	err = logger.InitLogger(cfg.Logger.Level, cfg.Logger.Development)
	if err != nil {
		logger.GlobalLogger.Fatal("failed to initialize logger", zap.Error(err))
	}
	defer func() {
		_ = logger.GlobalLogger.Sync()
	}()

	service, err := graphmap.New(*initMapPath)
	if err != nil {
		logger.GlobalLogger.Fatal("failed to initialize graph map service", zap.Error(err))
	}

	ctrl := controller.New(service)

	server, err := api.NewServer(ctrl)
	if err != nil {
		logger.GlobalLogger.Fatal("failed to create server", zap.Error(err))
	}

	httpServer := &http.Server{
		Addr:              fmt.Sprintf(":%d", cfg.Server.Port),
		ReadHeaderTimeout: time.Duration(cfg.Server.ReadHeaderTimeout) * time.Second,
		Handler:           server,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		logger.GlobalLogger.Info("server is starting", zap.Int("port", cfg.Server.Port))
		err = httpServer.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.GlobalLogger.Fatal("server failed to start", zap.Error(err))
		}
	}()

	<-done
	logger.GlobalLogger.Info("server is shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	err = httpServer.Shutdown(ctx)
	if err != nil {
		logger.GlobalLogger.Error("failed to shutdown server gracefully", zap.Error(err))
	} else {
		logger.GlobalLogger.Info("server shutdown completed")
	}
}
