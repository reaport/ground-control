package controller

import (
	"context"
	"fmt"

	"github.com/reaport/ground-control/internal/controller/convert"
	"github.com/reaport/ground-control/pkg/api"
)

// MapGetAirportMap implements map_getAirportMap operation.
//
// Возвращает полную карту аэропорта в виде графа.
//
// GET /map
func (c *Controller) MapGetAirportMap(ctx context.Context) (*api.AirportMap, error) {
	airportMap, err := c.mapService.GetAirportMap(ctx)
	if err != nil {
		return nil, fmt.Errorf("MapService.GetAirportMap: %w", err)
	}

	apiAirportMap, err := convert.AirportMapToAPI(airportMap)
	if err != nil {
		return nil, fmt.Errorf("convert.AirportMapToAPI: %w", err)
	}

	return apiAirportMap, nil
}
