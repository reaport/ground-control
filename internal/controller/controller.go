package controller

import (
	"context"

	"github.com/reaport/ground-control/internal/entity"
)

type MapService interface {
	GetAirportMap(ctx context.Context) (*entity.AirportMap, error)
	RefreshAirportMap(ctx context.Context) error
	UpdateAirportMap(_ context.Context, airportMap *entity.AirportMap) error

	RegisterVehicle(ctx context.Context, vehicleType entity.VehicleType) (*entity.VehicleInitInfo, error)
	GetRoute(ctx context.Context, nodeIDFrom, nodeIDTo string, vehicleType entity.VehicleType) ([]string, error)
	RequestMove(
		ctx context.Context,
		vehicleID string,
		nodeIDFrom, nodeIDTo string,
		vehicleType entity.VehicleType,
	) (float64, error)
	NotifyArrival(ctx context.Context, nodeID string, vehicleID string) error

	GetAirplaneParkingSpot(ctx context.Context, airplaneID string) (string, error)
}

type Controller struct {
	mapService MapService
}

func New(mapService MapService) *Controller {
	return &Controller{
		mapService: mapService,
	}
}
