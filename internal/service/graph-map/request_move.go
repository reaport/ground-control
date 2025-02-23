package graphmap

import (
	"context"
	"fmt"

	"github.com/reaport/ground-control/internal/entity"
)

func (s *Service) RequestMove(
	_ context.Context,
	vehicleID string,
	nodeIDFrom, nodeIDTo string,
	vehicleType entity.VehicleType,
) (float64, error) {
	s.mapMutex.Lock()
	defer s.mapMutex.Unlock()

	nodeFrom := s.findNodeByID(nodeIDFrom)
	if nodeFrom == nil {
		return 0, fmt.Errorf("%w: node %s not found", entity.ErrNodeNotFound, nodeIDFrom)
	}

	nodeTo := s.findNodeByID(nodeIDTo)
	if nodeTo == nil {
		return 0, fmt.Errorf("%w: node %s not found", entity.ErrNodeNotFound, nodeIDTo)
	}

	if !nodeTo.IsValidType(vehicleType) {
		return 0, fmt.Errorf(
			"%w: node %s does not support vehicle type %s",
			entity.ErrInvalidVehicleType,
			nodeIDTo,
			vehicleType,
		)
	}

	if !nodeFrom.ContainsVehicle(vehicleID) {
		return 0, fmt.Errorf("%w: vehicle %s not found in node %s", entity.ErrVehicleNotFound, vehicleID, nodeIDFrom)
	}

	edge := s.getEdge(nodeIDFrom, nodeIDTo)
	if edge == nil {
		return 0, fmt.Errorf("%w: no edge between %s and %s", entity.ErrEdgeNotFound, nodeIDFrom, nodeIDTo)
	}

	if !s.canMoveToNode(nodeTo, vehicleType) {
		return 0, fmt.Errorf("%w: cannot move vehicle %s to node %s", entity.ErrMoveNotAllowed, vehicleID, nodeIDTo)
	}

	nodeFrom.RemoveVehicle(vehicleID)
	nodeTo.AddVehicle(entity.NewVehicle(vehicleID, vehicleType))

	return edge.Distance, nil
}

func (s *Service) getEdge(nodeIDFrom, nodeIDTo string) *entity.Edge {
	for _, edge := range s.airportMap.Edges {
		if (edge.From == nodeIDFrom && edge.To == nodeIDTo) || (edge.From == nodeIDTo && edge.To == nodeIDFrom) {
			return edge
		}
	}
	return nil
}

func (s *Service) canMoveToNode(node *entity.Node, vehicleType entity.VehicleType) bool {
	if len(node.Vehicles) == 0 {
		return true
	}

	existingVehicle := node.Vehicles[0]
	switch {
	case existingVehicle.Type == entity.VehicleTypeAirplane && vehicleType == entity.VehicleTypeFollowMe:
		return true
	case existingVehicle.Type == entity.VehicleTypeFollowMe && vehicleType == entity.VehicleTypeAirplane:
		return true
	default:
		return false
	}
}
