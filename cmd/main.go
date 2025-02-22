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
	shutdownTimeout = 5
)

func main() {
	configPath := flag.String("config", "config.yaml", "config path")
	initMapPath := flag.String("init-data", "init_data.json", "init map path")
	flag.Parse()

	config, err := config.LoadConfig(*configPath)
	if err != nil {
		panic(err)
	}

	err = logger.InitLogger(config.Logger.Level, config.Logger.Development)
	if err != nil {
		panic(err)
	}
	defer func() {
		err = logger.GlobalLogger.Sync()
		if err != nil {
			panic(err)
		}
	}()

	service, err := graphmap.New(*initMapPath)
	if err != nil {
		panic(err)
	}

	controller := controller.New(service)

	server, err := api.NewServer(controller)
	if err != nil {
		panic(err)
	}

	httpServer := &http.Server{
		Addr:              fmt.Sprintf(":%d", config.Server.Port),
		ReadHeaderTimeout: time.Duration(config.Server.ReadHeaderTimeout) * time.Second,
		Handler:           server,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		logger.GlobalLogger.Info("cервер запущен", zap.Int("port", config.Server.Port))
		err = httpServer.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}()

	<-done
	logger.GlobalLogger.Info("cервер завершает работу...")

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout*time.Second)
	defer cancel()

	err = httpServer.Shutdown(ctx)
	if err != nil {
		panic(err)
	}
}
