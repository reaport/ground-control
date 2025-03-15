package graphmap

import (
	"context"

	"github.com/reaport/ground-control/internal/entity"
)

func (s *Service) GetAirportMapConfig(_ context.Context) (*entity.AirportMapConfig, error) {
	for _, node := range s.airportMap.Nodes {
		if node.ID == airstripNodeID {
			return &entity.AirportMapConfig{
				AirstripNodeID: node.ID,
			}, nil
		}
	}

	return nil, entity.ErrAirstripNotFound
}
