package graphmap

import (
	"context"
	"strings"

	"github.com/reaport/ground-control/internal/entity"
)

func (s *Service) GetAirplaneServiceSpot(
	_ context.Context,
	airplaneID string,
	vehicleType entity.VehicleType,
) (string, error) {
	s.mapMutex.Lock()
	defer s.mapMutex.Unlock()

	var airplaneParkingNode *entity.Node

	for _, node := range s.airportMap.Nodes {
		if strings.HasPrefix(node.ID, airplaneParkingPrefix) && node.ContainsVehicle(airplaneID) {
			airplaneParkingNode = node
			break
		}
	}

	if airplaneParkingNode == nil {
		return "", entity.ErrAirplaneParkingSpotNotFound
	}

	for _, edge := range s.airportMap.Edges {
		if edge.From == airplaneParkingNode.ID {
			serviceNode := s.findNodeByID(edge.To)
			if serviceNode.IsValidType(vehicleType) {
				return serviceNode.ID, nil
			}
		}
		if edge.To == airplaneParkingNode.ID {
			serviceNode := s.findNodeByID(edge.From)
			if serviceNode.IsValidType(vehicleType) {
				return serviceNode.ID, nil
			}
		}
	}

	return "", entity.ErrAirplaneServiceSpotIsFull
}
