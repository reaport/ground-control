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
		return s.registerAirplane()
	}

	return s.registerGroundVehicle(vehicleType)
}

func (s *Service) registerAirplane() (*entity.VehicleInitInfo, error) {
	nodeID, vehicleID, err := s.findAirstripForAirplane()
	if err != nil {
		return nil, fmt.Errorf("registerAirplane: %w", err)
	}

	return entity.NewVehicleInitInfo(vehicleID, nodeID, nil), nil
}

func (s *Service) registerGroundVehicle(vehicleType entity.VehicleType) (*entity.VehicleInitInfo, error) {
	vehicleID := s.generateVehicleID(vehicleType)
	vehicleGarrageNodeID := fmt.Sprintf("%s_%s", garrageNodeID, vehicleID)
	vehicleGarrageNode := entity.NewNode(vehicleGarrageNodeID, []entity.VehicleType{vehicleType})

	s.addVehicleToGarage(vehicleGarrageNode, vehicleID, vehicleType)
	s.createGarageEdges(vehicleGarrageNodeID)

	serviceSpots, err := s.createServiceSpots(vehicleType, vehicleID)
	if err != nil {
		return nil, err
	}

	return &entity.VehicleInitInfo{
		VehicleID:     vehicleID,
		GarrageNodeID: vehicleGarrageNodeID,
		ServiceSpots:  serviceSpots,
	}, nil
}

func (s *Service) generateVehicleID(vehicleType entity.VehicleType) string {
	vehicleSequence := s.vehicleSequenceMap[vehicleType] + 1
	s.vehicleSequenceMap[vehicleType] = vehicleSequence
	return fmt.Sprintf("%s_%d", vehicleType, vehicleSequence)
}

func (s *Service) addVehicleToGarage(node *entity.Node, vehicleID string, vehicleType entity.VehicleType) {
	node.AddVehicle(entity.NewVehicle(vehicleID, vehicleType))
	s.airportMap.Nodes = append(s.airportMap.Nodes, node)
}

func (s *Service) createGarageEdges(vehicleGarrageNodeID string) {
	s.airportMap.Edges = append(s.airportMap.Edges,
		entity.NewEdge(garrageNodeID, vehicleGarrageNodeID, closeDistance),
		entity.NewEdge(vehicleGarrageNodeID, garrageNodeID, closeDistance),
		entity.NewEdge(vehicleGarrageNodeID, airportNodeID, closeDistance),
	)
}

func (s *Service) createServiceSpots(vehicleType entity.VehicleType, vehicleID string) (map[string]string, error) {
	serviceSpots := make(map[string]string)

	for _, parkingNode := range s.airportMap.Nodes {
		if !isParkingNode(parkingNode.ID) {
			continue
		}

		garrageToParking, garrageFromParking, err := s.findGarageConnections(parkingNode.ID)
		if err != nil {
			return nil, err
		}

		serviceSpotNodeID := fmt.Sprintf("%s_%s", parkingNode.ID, vehicleID)
		serviceSpotNode := entity.NewNode(serviceSpotNodeID, []entity.VehicleType{vehicleType})
		s.airportMap.Nodes = append(s.airportMap.Nodes, serviceSpotNode)

		s.airportMap.Edges = append(s.airportMap.Edges,
			entity.NewEdge(garrageToParking.ID, serviceSpotNodeID, midDistance),
			entity.NewEdge(serviceSpotNodeID, garrageFromParking.ID, midDistance),
			entity.NewEdge(serviceSpotNodeID, parkingNode.ID, closeDistance),
		)

		serviceSpots[parkingNode.ID] = serviceSpotNodeID
	}

	return serviceSpots, nil
}

func isParkingNode(nodeID string) bool {
	return strings.HasPrefix(nodeID, "parking_") && len(strings.Split(nodeID, "_")) == 2
}

func (s *Service) findGarageConnections(parkingNodeID string) (*entity.Node, *entity.Node, error) {
	var garrageToParking, garrageFromParking *entity.Node

	for _, node := range s.airportMap.Nodes {
		if node == nil || !strings.HasSuffix(node.ID, parkingNodeID) {
			continue
		}

		if strings.HasPrefix(node.ID, garrageNodeID) {
			parts := strings.Split(strings.TrimSuffix(node.ID, "_"+parkingNodeID), "_")
			direction := parts[len(parts)-1]
			switch direction {
			case "to":
				garrageToParking = node
			case "from":
				garrageFromParking = node
			}
		}
	}

	if garrageToParking == nil || garrageFromParking == nil {
		return nil, nil, fmt.Errorf("garage connection not found for %s", parkingNodeID)
	}

	return garrageToParking, garrageFromParking, nil
}

func (s *Service) findAirstripForAirplane() (string, string, error) {
	for _, node := range s.airportMap.Nodes {
		if node.ID == airstripNodeID && len(node.Vehicles) == 0 {
			vehicleID := s.generateVehicleID(entity.VehicleTypeAirplane)
			node.AddVehicle(entity.NewVehicle(vehicleID, entity.VehicleTypeAirplane))
			return node.ID, vehicleID, nil
		}
	}
	return "", "", entity.ErrAirstripNotFound
}
