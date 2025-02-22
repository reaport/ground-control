package graphmap

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"github.com/reaport/ground-control/internal/entity"
)

type Service struct {
	airportMap *entity.AirportMap
	mapMutex   *sync.RWMutex

	vehicleSequenceMap   map[entity.VehicleType]int
	vehicleSequenceMutex *sync.RWMutex
}

func New(initDataFilePath string) (*Service, error) {
	initDataFile, err := os.Open(initDataFilePath)
	if err != nil {
		return nil, fmt.Errorf("os.Open: %w", err)
	}

	var initData entity.AirportMap
	err = json.NewDecoder(initDataFile).Decode(&initData)
	if err != nil {
		return nil, fmt.Errorf("json.NewDecoder.Decode: %w", err)
	}

	return &Service{
		airportMap:           &initData,
		mapMutex:             &sync.RWMutex{},
		vehicleSequenceMap:   map[entity.VehicleType]int{},
		vehicleSequenceMutex: &sync.RWMutex{},
	}, nil
}
