package graphmap

import (
	"context"
	"fmt"
	"strings"

	"github.com/reaport/ground-control/internal/entity"
)

const (
	garrageNodeID  = "garrage"
	airportNodeID  = "airport"
	closeDistance  = 10
	midDistance    = 100
	airstripNodeID = "airstrip"
)

func (s *Service) RegisterVehicle(
	_ context.Context,
	vehicleType entity.VehicleType,
) (*entity.VehicleInitInfo, error) {
	s.vehicleSequenceMutex.Lock()
	defer s.vehicleSequenceMutex.Unlock()

	s.mapMutex.Lock()
	defer s.mapMutex.Unlock()

	if vehicleType == entity.VehicleTypeAirplane {
		nodeID, vehicleID, err := s.RegisterAirplane()
		if err != nil {
			return nil, fmt.Errorf("RegisterAirplane: %w", err)
		}

		return entity.NewVehicleInitInfo(vehicleID, nodeID, nil), nil
	}

	var vehicleInitInfo entity.VehicleInitInfo

	vehicleSequence := s.vehicleSequenceMap[vehicleType]
	vehicleSequence++
	vehicleInitInfo.VehicleID = fmt.Sprintf("%s_%d", vehicleType, vehicleSequence)
	s.vehicleSequenceMap[vehicleType] = vehicleSequence

	vehicleGarrageNodeID := fmt.Sprintf("%s_%s", garrageNodeID, vehicleInitInfo.VehicleID)
	vehicleGarrageNode := entity.NewNode(vehicleGarrageNodeID, []entity.VehicleType{vehicleType})
	vehicleInitInfo.GarrageNodeID = vehicleGarrageNodeID
	vehicleGarrageNode.AddVehicle(entity.NewVehicle(vehicleInitInfo.VehicleID, vehicleType))
	s.airportMap.Nodes = append(s.airportMap.Nodes, vehicleGarrageNode)

	s.airportMap.Edges = append(s.airportMap.Edges, entity.NewEdge(garrageNodeID, vehicleGarrageNodeID, closeDistance))
	s.airportMap.Edges = append(s.airportMap.Edges, entity.NewEdge(vehicleGarrageNodeID, garrageNodeID, closeDistance))

	s.airportMap.Edges = append(s.airportMap.Edges, entity.NewEdge(vehicleGarrageNodeID, airportNodeID, closeDistance))

	for _, parkingNode := range s.airportMap.Nodes {
		if !strings.HasPrefix(parkingNode.ID, parkingPrefix) || len(strings.Split(parkingNode.ID, "_")) != 2 {
			continue
		}

		var garrageToParkingNode *entity.Node
		var garrageFromParkingNode *entity.Node

		for _, node := range s.airportMap.Nodes {
			if node == parkingNode || !strings.HasSuffix(node.ID, parkingNode.ID) || !strings.HasPrefix(node.ID, garrageNodeID) {
				continue
			}

			trimmed := strings.TrimSuffix(node.ID, "_"+parkingNode.ID)
			parts := strings.Split(trimmed, "_")
			direction := parts[len(parts)-1]
			if direction != "to" && direction != "from" {
				return nil, fmt.Errorf("%w: invalid direction %s", entity.ErrInvalidDirection, direction)
			}

			switch direction {
			case "to":
				garrageToParkingNode = s.findNodeByID(node.ID)
			case "from":
				garrageFromParkingNode = s.findNodeByID(node.ID)
			}

			if garrageToParkingNode != nil && garrageFromParkingNode != nil {
				break
			}
		}

		if garrageToParkingNode == nil {
			return nil, fmt.Errorf("%w: from garrage to %s", entity.ErrNodeNotFound, parkingNode.ID)
		}

		if garrageFromParkingNode == nil {
			return nil, fmt.Errorf("%w: from %s to garrage", entity.ErrNodeNotFound, parkingNode.ID)
		}

		serviceSpotNodeID := fmt.Sprintf("%s_%s", parkingNode.ID, vehicleInitInfo.VehicleID)
		serviceSpotNode := entity.NewNode(serviceSpotNodeID, []entity.VehicleType{vehicleType})
		s.airportMap.Nodes = append(s.airportMap.Nodes, serviceSpotNode)

		s.airportMap.Edges = append(s.airportMap.Edges, entity.NewEdge(garrageToParkingNode.ID, serviceSpotNodeID, midDistance))
		s.airportMap.Edges = append(s.airportMap.Edges, entity.NewEdge(serviceSpotNodeID, garrageFromParkingNode.ID, midDistance))

		s.airportMap.Edges = append(s.airportMap.Edges, entity.NewEdge(serviceSpotNodeID, parkingNode.ID, closeDistance))

		if vehicleInitInfo.ServiceSpots == nil {
			vehicleInitInfo.ServiceSpots = make(map[string]string)
		}

		vehicleInitInfo.ServiceSpots[parkingNode.ID] = serviceSpotNodeID
	}

	return &vehicleInitInfo, nil
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
