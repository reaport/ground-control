package graphmap

import (
	"context"
	"fmt"
	"strings"

	"github.com/reaport/ground-control/internal/entity"
)

const (
	garrageNodeID          = "garrage"
	airportNodeID          = "airport"
	closeDistance          = 10
	airstripNodeID         = "airstrip"
	serviceCrossroadPrefix = "service_crossroad"
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

	newGarageEdge := entity.NewEdge(garrageNodeID, vehicleGarrageNodeID, closeDistance)
	s.airportMap.Edges = append(s.airportMap.Edges, newGarageEdge)

	newAirportEdge := entity.NewEdge(airportNodeID, vehicleGarrageNodeID, closeDistance)
	s.airportMap.Edges = append(s.airportMap.Edges, newAirportEdge)

	for _, airplaneParkingNode := range s.airportMap.Nodes {
		if !strings.HasPrefix(airplaneParkingNode.ID, airplaneParkingPrefix) {
			continue
		}

		airplaneParkingNodeParts := strings.Split(airplaneParkingNode.ID, "_")
		if len(airplaneParkingNodeParts) != 3 {
			continue
		}

		airplaneParkingNumber := airplaneParkingNodeParts[len(airplaneParkingNodeParts)-1]

		var serviceCrossroadNode *entity.Node

		for _, node := range s.airportMap.Nodes {
			if !strings.HasPrefix(node.ID, serviceCrossroadPrefix) {
				continue
			}

			parts := strings.Split(node.ID, "_")
			serviceCrossroadNumber := parts[len(parts)-1]

			if serviceCrossroadNumber == airplaneParkingNumber {
				serviceCrossroadNode = s.findNodeByID(node.ID)
				break
			}
		}

		if serviceCrossroadNode == nil {
			return nil, fmt.Errorf("%w: for airplane parking %s", entity.ErrServiceCrossroadNotFound, airplaneParkingNode.ID)
		}

		serviceSpotNodeID := fmt.Sprintf("%s_%s", airplaneParkingNode.ID, vehicleInitInfo.VehicleID)
		serviceSpotNode := entity.NewNode(serviceSpotNodeID, []entity.VehicleType{vehicleType})
		s.airportMap.Nodes = append(s.airportMap.Nodes, serviceSpotNode)

		newAirplaneParkingEdge := entity.NewEdge(airplaneParkingNode.ID, serviceSpotNodeID, closeDistance)
		s.airportMap.Edges = append(s.airportMap.Edges, newAirplaneParkingEdge)

		newServiceCrossroadEdge := entity.NewEdge(serviceCrossroadNode.ID, serviceSpotNodeID, closeDistance)
		s.airportMap.Edges = append(s.airportMap.Edges, newServiceCrossroadEdge)

		if vehicleInitInfo.ServiceSpots == nil {
			vehicleInitInfo.ServiceSpots = make(map[string]string)
		}

		vehicleInitInfo.ServiceSpots[airplaneParkingNode.ID] = serviceSpotNodeID
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
