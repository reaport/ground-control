package graphmap

import (
	"context"

	"github.com/reaport/ground-control/internal/entity"
)

func (s *Service) TakeOffAirplane(_ context.Context, id string) error {
	s.mapMutex.Lock()
	defer s.mapMutex.Unlock()

	for _, node := range s.airportMap.Nodes {
		if node.ContainsVehicle(id) && node.ID == airstripNodeID {
			node.RemoveVehicle(id)
			return nil
		}
	}

	return entity.ErrVehicleNotFound
}
