package controller

import (
	"context"
	"errors"

	"github.com/reaport/ground-control/pkg/api"
)

// GET /map.
func (c *Controller) MapGetAirportMap(_ context.Context) (*api.AirportMap, error) {
	return nil, errors.New("not implemented")
}
