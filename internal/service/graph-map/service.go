package graphmap

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/reaport/ground-control/internal/entity"
)

type Service struct {
	airportMap *entity.AirportMap
}

func New(initDataFilePath string) (*Service, error) {
	initDataFile, err := os.Open(initDataFilePath)
	if err != nil {
		return nil, fmt.Errorf("os.Open: %w", err)
	}

	var initData entity.AirportMap
	if err := json.NewDecoder(initDataFile).Decode(&initData); err != nil {
		return nil, fmt.Errorf("json.NewDecoder.Decode: %w", err)
	}

	return &Service{airportMap: &initData}, nil
}
