package controller

import (
	"context"
	"fmt"

	"github.com/reaport/ground-control/pkg/logger"
	"go.uber.org/zap"
)

// MapRefreshAirportMap implements map_refreshAirportMap operation.
//
// Возвращает карту к исходному состоянию.
//
// POST /map/refresh
func (c *Controller) MapRefreshAirportMap(ctx context.Context) error {
	err := c.mapService.RefreshAirportMap(ctx)
	if err != nil {
		err = fmt.Errorf("c.mapService.RefreshAirportMap: %w", err)
		logger.GlobalLogger.Error(
			"failed to refresh map",
			zap.Error(err),
		)
		return err
	}

	logger.GlobalLogger.Info(
		"map refreshed",
	)

	return nil
}
