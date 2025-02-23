package graphmap

import (
	"context"

	"github.com/reaport/ground-control/internal/entity"
)

func (s *Service) GetAirportMap(_ context.Context) (*entity.AirportMap, error) {
	return s.airportMap, nil
}
