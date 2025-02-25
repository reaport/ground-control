package graphmap

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/reaport/ground-control/internal/entity"
)

func (s *Service) loadInitData() error {
	initDataFile, err := os.Open(s.mapFilePath)
	if err != nil {
		return fmt.Errorf("os.Open: %w", err)
	}

	var initData entity.AirportMap
	err = json.NewDecoder(initDataFile).Decode(&initData)
	if err != nil {
		return fmt.Errorf("json.NewDecoder.Decode: %w", err)
	}

	s.airportMap = &initData

	return nil
}

func (s *Service) saveInitData() error {
	initDataFile, err := os.Create(s.mapFilePath)
	if err != nil {
		return fmt.Errorf("os.Create: %w", err)
	}

	err = json.NewEncoder(initDataFile).Encode(s.airportMap)
	if err != nil {
		return fmt.Errorf("json.NewEncoder.Encode: %w", err)
	}

	return nil
}
