package controller

import (
	"context"
	"errors"

	"github.com/reaport/ground-control/pkg/api"
)

// GET /airplane/{id}/parking.
func (c *Controller) AirplaneGetParkingSpot(
	_ context.Context,
	_ api.AirplaneGetParkingSpotParams,
) (api.AirplaneGetParkingSpotRes, error) {
	return nil, errors.New("not implemented")
}
