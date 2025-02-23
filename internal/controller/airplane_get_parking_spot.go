package controller

import (
	"context"
	"errors"

	"github.com/reaport/ground-control/pkg/api"
)

// AirplaneGetParkingSpot implements airplane_getParkingSpot operation.
//
// В зависимости от загрузки парковок отдает нужный узел.
//
// GET /airplane/{id}/parking
func (c *Controller) AirplaneGetParkingSpot(
	_ context.Context,
	_ api.AirplaneGetParkingSpotParams,
) (api.AirplaneGetParkingSpotRes, error) {
	return nil, errors.New("not implemented")
}
