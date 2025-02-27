package graphmap

import (
	"context"
	"fmt"

	"github.com/reaport/ground-control/internal/entity"
)

func (s *Service) RefreshAirportMap(_ context.Context) error {
	s.mapMutex.Lock()
	defer s.mapMutex.Unlock()

	s.vehicleSequenceMutex.Lock()
	defer s.vehicleSequenceMutex.Unlock()

	err := s.loadInitData()
	if err != nil {
		return fmt.Errorf("loadInitData: %w", err)
	}

	s.vehicleSequenceMap = map[entity.VehicleType]int{}

	return nil
}
