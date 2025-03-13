package controller

import (
	"context"

	"github.com/reaport/ground-control/internal/entity"
)

type (
	MapService interface {
		GetAirportMap(ctx context.Context) (*entity.AirportMap, error)
		RefreshAirportMap(ctx context.Context) error
		UpdateAirportMap(_ context.Context, airportMap *entity.AirportMap) error
		GetAirportMapConfig(ctx context.Context) (*entity.AirportMapConfig, error)

		RegisterVehicle(ctx context.Context, vehicleType entity.VehicleType) (*entity.VehicleInitInfo, error)
		GetRoute(ctx context.Context, nodeIDFrom, nodeIDTo string, vehicleType entity.VehicleType) ([]string, error)
		RequestMove(
			ctx context.Context,
			vehicleID string,
			nodeIDFrom, nodeIDTo string,
			vehicleType entity.VehicleType,
			withAirplane *string,
		) (float64, error)
		NotifyArrival(ctx context.Context, nodeID string, vehicleID string) error

		GetAirplaneParkingSpot(ctx context.Context, airplaneID string) (string, error)
	}

	EventSender interface {
		SendEvent(ctx context.Context, event *entity.Event) error
	}

	Controller struct {
		mapService  MapService
		eventSender EventSender
	}
)

func New(mapService MapService, eventSender EventSender) *Controller {
	return &Controller{
		mapService:  mapService,
		eventSender: eventSender,
	}
}
