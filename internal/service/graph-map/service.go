package graphmap

import (
	"fmt"
	"sync"

	"github.com/reaport/ground-control/internal/entity"
)

type Service struct {
	mapFilePath string

	airportMap *entity.AirportMap
	mapMutex   *sync.RWMutex

	vehicleSequenceMap   map[entity.VehicleType]int
	vehicleSequenceMutex *sync.RWMutex
}

func New(cfg *Config) (*Service, error) {
	service := &Service{
		mapFilePath:          cfg.MapFilePath,
		mapMutex:             &sync.RWMutex{},
		vehicleSequenceMap:   map[entity.VehicleType]int{},
		vehicleSequenceMutex: &sync.RWMutex{},
	}

	err := service.loadInitData()
	if err != nil {
		return nil, fmt.Errorf("loadInitData: %w", err)
	}

	return service, nil
}
