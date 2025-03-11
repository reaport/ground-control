package controller

import (
	"context"
	"fmt"

	"github.com/reaport/ground-control/pkg/api"
	"github.com/reaport/ground-control/pkg/logger"
	"go.uber.org/zap"
)

// MapGetAirportMapConfig implements map_getAirportMapConfig operation.
//
// Возвращает конфигурацию аэропорта.
//
// GET /map/config
func (c *Controller) MapGetAirportMapConfig(ctx context.Context) (*api.AirportMapConfig, error) {
	airportMapConfig, err := c.mapService.GetAirportMapConfig(ctx)
	if err != nil {
		err = fmt.Errorf("c.mapService.GetAirportMapConfig: %w", err)
		logger.GlobalLogger.Error(
			"failed to get airport map config",
			zap.Error(err),
		)
		return nil, err
	}

	return &api.AirportMapConfig{
		AirstripNodeId: airportMapConfig.AirstripNodeId,
	}, nil
}
