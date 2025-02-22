package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/reaport/ground-control/internal/config"
	"github.com/reaport/ground-control/internal/controller"
	"github.com/reaport/ground-control/pkg/api"
	"github.com/reaport/ground-control/pkg/logger"
	"go.uber.org/zap"
)

func main() {
	configPath := flag.String("config", "config/config.yaml", "config path")
	flag.Parse()

	config, err := config.LoadConfig(*configPath)
	if err != nil {
		panic(err)
	}

	err = logger.InitLogger(config.Logger.Level, config.Logger.Development)
	if err != nil {
		panic(err)
	}
	defer logger.GlobalLogger.Sync()

	controller := controller.New(nil)
	server, err := api.NewServer(controller)
	if err != nil {
		panic(err)
	}

	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Server.Port),
		Handler: server,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		logger.GlobalLogger.Info("cервер запущен", zap.Int("port", config.Server.Port))
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	<-done
	logger.GlobalLogger.Info("cервер завершает работу...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		panic(err)
	}
}
