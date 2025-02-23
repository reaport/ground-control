package graphmap

import (
	"context"
	"fmt"

	"github.com/reaport/ground-control/internal/entity"
)

func (s *Service) NotifyArrival(_ context.Context, nodeID string, vehicleID string) error {
	s.mapMutex.Lock()
	defer s.mapMutex.Unlock()

	node := s.findNodeByID(nodeID)
	if node == nil {
		return fmt.Errorf("%w: node %s not found", entity.ErrNodeNotFound, nodeID)
	}

	if !node.ContainsVehicle(vehicleID) {
		return fmt.Errorf("%w: vehicle %s not found in node %s", entity.ErrVehicleNotFound, vehicleID, nodeID)
	}

	return nil
}
