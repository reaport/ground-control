package graphmap

import (
	"context"

	"github.com/reaport/ground-control/internal/entity"
)

func (s *Service) GetAirportMapConfig(ctx context.Context) (*entity.AirportMapConfig, error) {
	for _, node := range s.airportMap.Nodes {
		if node.ID == airstripNodeID {
			return &entity.AirportMapConfig{
				AirstripNodeId: node.ID,
			}, nil
		}
	}

	return nil, entity.ErrAirstripNotFound
}
