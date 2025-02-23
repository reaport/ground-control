package graphmap

import (
	"context"
	"fmt"

	"github.com/reaport/ground-control/internal/entity"
)

const (
	garrageNodeID               = "garrage"
	airportNodeID               = "airport"
	garrageToNewVehicleDistance = 0
	airportToNewVehicleDistance = 0
	airstripNodeID              = "airstrip"
)

func (s *Service) RegisterVehicle(
	_ context.Context,
	vehicleType entity.VehicleType,
) (nodeID string, vehicleID string, err error) {
	s.vehicleSequenceMutex.Lock()
	defer s.vehicleSequenceMutex.Unlock()

	s.mapMutex.Lock()
	defer s.mapMutex.Unlock()

	if vehicleType == entity.VehicleTypeAirplane {
		nodeID, vehicleID, err = s.RegisterAirplane()
		if err != nil {
			return "", "", fmt.Errorf("RegisterAirplane: %w", err)
		}

		return nodeID, vehicleID, nil
	}

	vehicleSequence := s.vehicleSequenceMap[vehicleType]
	vehicleSequence++
	vehicleID = fmt.Sprintf("%s_%d", vehicleType, vehicleSequence)
	s.vehicleSequenceMap[vehicleType] = vehicleSequence

	nodeID = fmt.Sprintf("%s_%s", garrageNodeID, vehicleID)
	newNode := entity.NewNode(nodeID, []entity.VehicleType{vehicleType})
	newNode.AddVehicle(entity.NewVehicle(vehicleID, vehicleType))
	s.airportMap.Nodes = append(s.airportMap.Nodes, newNode)

	newGarageEdge := entity.NewEdge(garrageNodeID, nodeID, garrageToNewVehicleDistance)
	s.airportMap.Edges = append(s.airportMap.Edges, newGarageEdge)

	newAirportEdge := entity.NewEdge(airportNodeID, nodeID, airportToNewVehicleDistance)
	s.airportMap.Edges = append(s.airportMap.Edges, newAirportEdge)

	return nodeID, vehicleID, nil
}

func (s *Service) RegisterAirplane() (nodeID string, vehicleID string, err error) {
	for _, node := range s.airportMap.Nodes {
		if node.ID == airstripNodeID {
			if len(node.Vehicles) != 0 {
				return "", "", entity.ErrAirstripIsFull
			}

			airplaneSequence := s.vehicleSequenceMap[entity.VehicleTypeAirplane]
			airplaneSequence++
			airplaneID := fmt.Sprintf("%s_%d", entity.VehicleTypeAirplane, airplaneSequence)
			s.vehicleSequenceMap[entity.VehicleTypeAirplane] = airplaneSequence

			node.AddVehicle(entity.NewVehicle(airplaneID, entity.VehicleTypeAirplane))

			return node.ID, airplaneID, nil
		}
	}

	return "", "", entity.ErrAirstripNotFound
}
