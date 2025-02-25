package controller

import (
	"context"
	"fmt"

	"github.com/reaport/ground-control/internal/controller/convert"
	"github.com/reaport/ground-control/pkg/api"
	"github.com/reaport/ground-control/pkg/logger"
	"go.uber.org/zap"
)

// MapGetAirportMap implements map_getAirportMap operation.
//
// Возвращает полную карту аэропорта в виде графа.
//
// GET /map
func (c *Controller) MapGetAirportMap(ctx context.Context) (*api.AirportMap, error) {
	airportMap, err := c.mapService.GetAirportMap(ctx)
	if err != nil {
		err = fmt.Errorf("c.mapService.GetAirportMap: %w", err)
		logger.GlobalLogger.Error(
			"failed to get airport map",
			zap.Error(err),
		)
		return nil, err
	}

	apiAirportMap, err := convert.AirportMapToAPI(airportMap)
	if err != nil {
		err = fmt.Errorf("convert.AirportMapToAPI: %w", err)
		logger.GlobalLogger.Error(
			"failed to convert airport map to API",
			zap.Error(err),
		)
		return nil, err
	}

	logger.GlobalLogger.Info(
		"airport map got",
	)

	return apiAirportMap, nil
}
