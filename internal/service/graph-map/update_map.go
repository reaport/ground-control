package graphmap

import (
	"context"
	"fmt"

	"github.com/reaport/ground-control/internal/entity"
)

func (s *Service) UpdateAirportMap(_ context.Context, airportMap *entity.AirportMap) error {
	s.mapMutex.Lock()
	defer s.mapMutex.Unlock()

	for _, node := range airportMap.Nodes {
		if len(node.Vehicles) != 0 {
			return entity.ErrMapHasVehicles
		}
	}

	s.airportMap = airportMap

	err := s.saveInitData()
	if err != nil {
		return fmt.Errorf("s.saveInitData: %w", err)
	}

	return nil
}
