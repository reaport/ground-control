package graphmap

import (
	"context"
	"strings"

	"github.com/reaport/ground-control/internal/entity"
)

const (
	airplaneParkingPrefix = "airplane_parking"
)

func (s *Service) GetAirplaneParkingSpot(_ context.Context, airplaneID string) (string, error) {
	s.mapMutex.Lock()
	defer s.mapMutex.Unlock()

	for _, node := range s.airportMap.Nodes {
		if strings.HasPrefix(node.ID, airplaneParkingPrefix) {
			if len(node.Vehicles) == 0 || node.Vehicles[0].Type == entity.VehicleTypeFollowMe {
				node.AddVehicle(entity.NewVehicle(airplaneID, entity.VehicleTypeAirplane))
				return node.ID, nil
			}
		}
	}

	return "", entity.ErrAirplaneParkingSpotIsFull
}
