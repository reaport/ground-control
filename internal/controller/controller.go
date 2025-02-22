package controller

import (
	"context"

	"github.com/reaport/ground-control/internal/entity"
)

type MapService interface {
	GetAirportMap(ctx context.Context) (*entity.AirportMap, error)
	RegisterVehicle(ctx context.Context, vehicleType entity.VehicleType) (nodeID string, vehicleID string, err error)
}

type Controller struct {
	mapService MapService
}

func New(mapService MapService) *Controller {
	return &Controller{
		mapService: mapService,
	}
}
