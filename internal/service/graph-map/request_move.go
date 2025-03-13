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
	withAirplane *string,
) (float64, error) {
	if vehicleType != entity.VehicleTypeFollowMe && withAirplane != nil {
		return 0, fmt.Errorf("%w: withAirplane is not supported for vehicle type %s", entity.ErrInvalidVehicleType, vehicleType)
	}

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

	if withAirplane != nil {
		if !nodeTo.IsValidType(entity.VehicleTypeAirplane) {
			return 0, fmt.Errorf(
				"%w: node %s does not support vehicle type %s",
				entity.ErrInvalidVehicleType,
				nodeIDTo,
				entity.VehicleTypeAirplane,
			)
		}
	}

	if !nodeFrom.ContainsVehicle(vehicleID) {
		return 0, fmt.Errorf("%w: vehicle %s not found in node %s", entity.ErrVehicleNotFound, vehicleID, nodeIDFrom)
	}

	if withAirplane != nil {
		if !nodeFrom.ContainsVehicle(*withAirplane) {
			return 0, fmt.Errorf("%w: vehicle %s not found in node %s", entity.ErrVehicleNotFound, *withAirplane, nodeIDFrom)
		}
	}

	edge := s.getEdge(nodeIDFrom, nodeIDTo)
	if edge == nil {
		return 0, fmt.Errorf("%w: no edge between %s and %s", entity.ErrEdgeNotFound, nodeIDFrom, nodeIDTo)
	}

	canMove, isDuplicate := s.canMoveToNode(nodeTo, vehicleType, vehicleID)
	if !canMove {
		return 0, fmt.Errorf("%w: cannot move vehicle %s to node %s", entity.ErrMoveNotAllowed, vehicleID, nodeIDTo)
	}

	
	nodeFrom.RemoveVehicle(vehicleID)
	
	if !isDuplicate {
		nodeTo.AddVehicle(entity.NewVehicle(vehicleID, vehicleType))
	}

	if withAirplane != nil {
		canMove, isDuplicate = s.canMoveToNode(nodeTo, entity.VehicleTypeAirplane, *withAirplane)
		if !canMove {
			return 0, fmt.Errorf("%w: cannot move vehicle %s to node %s", entity.ErrMoveNotAllowed, *withAirplane, nodeIDTo)
		}

		nodeFrom.RemoveVehicle(*withAirplane)

		if !isDuplicate {
			nodeTo.AddVehicle(entity.NewVehicle(*withAirplane, entity.VehicleTypeAirplane))
		}
	}

	return edge.Distance, nil
}

func (s *Service) getEdge(nodeIDFrom, nodeIDTo string) *entity.Edge {
	for _, edge := range s.airportMap.Edges {
		if edge.From == nodeIDFrom && edge.To == nodeIDTo {
			return edge
		}
	}
	return nil
}

func (s *Service) canMoveToNode(
	node *entity.Node,
	vehicleType entity.VehicleType,
	vehicleID string,
) (canMove bool, isDuplicate bool) {
	if len(node.Vehicles) == 0 {
		return true, false
	}

	existingVehicle := node.Vehicles[0]
	switch {
	case existingVehicle.Type == entity.VehicleTypeAirplane && vehicleType == entity.VehicleTypeFollowMe,
		existingVehicle.Type == entity.VehicleTypeFollowMe && vehicleType == entity.VehicleTypeAirplane:
		return true, false
	case existingVehicle.ID == vehicleID:
		return true, true
	default:
		return false, false
	}
}
