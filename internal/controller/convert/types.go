package convert

import (
	"fmt"

	"github.com/reaport/ground-control/internal/entity"
	"github.com/reaport/ground-control/pkg/api"
)

func AirportMapToAPI(airportMap *entity.AirportMap) (*api.AirportMap, error) {
	apiNodes, err := NodesToAPI(airportMap.Nodes)
	if err != nil {
		return nil, fmt.Errorf("NodesToAPI: %w", err)
	}

	apiEdges := EdgesToAPI(airportMap.Edges)

	return &api.AirportMap{
		Nodes: apiNodes,
		Edges: apiEdges,
	}, nil
}

func AirportMapFromAPI(airportMap *api.AirportMap) (*entity.AirportMap, error) {
	nodes, err := NodesFromAPI(airportMap.Nodes)
	if err != nil {
		return nil, fmt.Errorf("NodesFromAPI: %w", err)
	}

	edges := EdgesFromAPI(airportMap.Edges)

	return &entity.AirportMap{
		Nodes: nodes,
		Edges: edges,
	}, nil
}

func NodesToAPI(nodes []*entity.Node) ([]api.Node, error) {
	apiNodes := make([]api.Node, len(nodes))
	for i, node := range nodes {
		apiNode, err := NodeToAPI(node)
		if err != nil {
			return nil, fmt.Errorf("NodeToAPI: %w", err)
		}
		apiNodes[i] = apiNode
	}

	return apiNodes, nil
}

func NodesFromAPI(apiNodes []api.Node) ([]*entity.Node, error) {
	nodes := make([]*entity.Node, len(apiNodes))
	for i, apiNode := range apiNodes {
		node, err := NodeFromAPI(apiNode)
		if err != nil {
			return nil, fmt.Errorf("NodeFromAPI: %w", err)
		}
		nodes[i] = node
	}

	return nodes, nil
}

func NodeToAPI(node *entity.Node) (api.Node, error) {
	apiVehicleTypes, err := VehicleTypesToAPI(node.Types)
	if err != nil {
		return api.Node{}, fmt.Errorf("VehicleTypesToAPI: %w", err)
	}

	apiVehicles, err := VehiclesToAPI(node.Vehicles)
	if err != nil {
		return api.Node{}, fmt.Errorf("VehiclesToAPI: %w", err)
	}

	return api.Node{
		ID:       node.ID,
		Types:    apiVehicleTypes,
		Vehicles: apiVehicles,
	}, nil
}

func NodeFromAPI(apiNode api.Node) (*entity.Node, error) {
	vehicleTypes, err := VehicleTypesFromAPI(apiNode.Types)
	if err != nil {
		return nil, fmt.Errorf("VehicleTypesFromAPI: %w", err)
	}

	vehicles, err := VehiclesFromAPI(apiNode.Vehicles)
	if err != nil {
		return nil, fmt.Errorf("VehiclesFromAPI: %w", err)
	}

	return &entity.Node{
		ID:       apiNode.ID,
		Types:    vehicleTypes,
		Vehicles: vehicles,
	}, nil
}

func VehicleTypesToAPI(vehicleTypes []entity.VehicleType) ([]api.VehicleType, error) {
	apiVehicleTypes := make([]api.VehicleType, len(vehicleTypes))
	for i, vehicleType := range vehicleTypes {
		apiVehicleType, err := VehicleTypeToAPI(vehicleType)
		if err != nil {
			return nil, fmt.Errorf("VehicleTypeToAPI: %w", err)
		}
		apiVehicleTypes[i] = apiVehicleType
	}

	return apiVehicleTypes, nil
}

func VehicleTypesFromAPI(apiVehicleTypes []api.VehicleType) ([]entity.VehicleType, error) {
	vehicleTypes := make([]entity.VehicleType, len(apiVehicleTypes))
	for i, apiVehicleType := range apiVehicleTypes {
		vehicleType, err := VehicleTypeFromAPI(apiVehicleType)
		if err != nil {
			return nil, fmt.Errorf("VehicleTypeFromAPI: %w", err)
		}
		vehicleTypes[i] = vehicleType
	}

	return vehicleTypes, nil
}

func VehicleTypeToAPI(vehicleType entity.VehicleType) (api.VehicleType, error) {
	if vehicleType == "" {
		return "", ErrEmptyVehicleType
	}

	switch vehicleType {
	case entity.VehicleTypeAirplane:
		return api.VehicleTypeAirplane, nil
	case entity.VehicleTypeCatering:
		return api.VehicleTypeCatering, nil
	case entity.VehicleTypeRefueling:
		return api.VehicleTypeRefueling, nil
	case entity.VehicleTypeCleaning:
		return api.VehicleTypeCleaning, nil
	case entity.VehicleTypeBaggage:
		return api.VehicleTypeBaggage, nil
	case entity.VehicleTypeFollowMe:
		return api.VehicleTypeFollowMe, nil
	case entity.VehicleTypeCharging:
		return api.VehicleTypeCharging, nil
	case entity.VehicleTypeBus:
		return api.VehicleTypeBus, nil
	case entity.VehicleTypeRamp:
		return api.VehicleTypeRamp, nil
	default:
		return "", ErrInvalidVehicleType
	}
}

func EdgesToAPI(edges []*entity.Edge) []api.Edge {
	apiEdges := make([]api.Edge, len(edges))
	for i, edge := range edges {
		apiEdges[i] = EdgeToAPI(edge)
	}

	return apiEdges
}

func EdgesFromAPI(apiEdges []api.Edge) []*entity.Edge {
	edges := make([]*entity.Edge, len(apiEdges))
	for i, apiEdge := range apiEdges {
		edges[i] = EdgeFromAPI(apiEdge)
	}

	return edges
}

func EdgeToAPI(edge *entity.Edge) api.Edge {
	return api.Edge{
		From:     edge.From,
		To:       edge.To,
		Distance: edge.Distance,
	}
}

func EdgeFromAPI(apiEdge api.Edge) *entity.Edge {
	return &entity.Edge{
		From:     apiEdge.From,
		To:       apiEdge.To,
		Distance: apiEdge.Distance,
	}
}

func VehiclesToAPI(vehicles []*entity.Vehicle) ([]api.Vehicle, error) {
	apiVehicles := make([]api.Vehicle, len(vehicles))
	for i, vehicle := range vehicles {
		apiVehicle, err := VehicleToAPI(vehicle)
		if err != nil {
			return nil, fmt.Errorf("VehicleToAPI: %w", err)
		}
		apiVehicles[i] = apiVehicle
	}

	return apiVehicles, nil
}

func VehiclesFromAPI(apiVehicles []api.Vehicle) ([]*entity.Vehicle, error) {
	vehicles := make([]*entity.Vehicle, len(apiVehicles))
	for i, apiVehicle := range apiVehicles {
		vehicle, err := VehicleFromAPI(apiVehicle)
		if err != nil {
			return nil, fmt.Errorf("VehicleFromAPI: %w", err)
		}
		vehicles[i] = vehicle
	}

	return vehicles, nil
}

func VehicleToAPI(vehicle *entity.Vehicle) (api.Vehicle, error) {
	vehicleType, err := VehicleTypeToAPI(vehicle.Type)
	if err != nil {
		return api.Vehicle{}, fmt.Errorf("VehicleTypeToAPI: %w", err)
	}

	return api.Vehicle{
		ID:   vehicle.ID,
		Type: vehicleType,
	}, nil
}

func VehicleFromAPI(apiVehicle api.Vehicle) (*entity.Vehicle, error) {
	vehicleType, err := VehicleTypeFromAPI(apiVehicle.Type)
	if err != nil {
		return nil, fmt.Errorf("VehicleTypeFromAPI: %w", err)
	}

	return &entity.Vehicle{
		ID:   apiVehicle.ID,
		Type: vehicleType,
	}, nil
}

func VehicleTypeFromAPI(vehicleType api.VehicleType) (entity.VehicleType, error) {
	if vehicleType == "" {
		return "", ErrEmptyVehicleType
	}

	switch vehicleType {
	case api.VehicleTypeAirplane:
		return entity.VehicleTypeAirplane, nil
	case api.VehicleTypeCatering:
		return entity.VehicleTypeCatering, nil
	case api.VehicleTypeRefueling:
		return entity.VehicleTypeRefueling, nil
	case api.VehicleTypeCleaning:
		return entity.VehicleTypeCleaning, nil
	case api.VehicleTypeBaggage:
		return entity.VehicleTypeBaggage, nil
	case api.VehicleTypeFollowMe:
		return entity.VehicleTypeFollowMe, nil
	case api.VehicleTypeCharging:
		return entity.VehicleTypeCharging, nil
	case api.VehicleTypeBus:
		return entity.VehicleTypeBus, nil
	case api.VehicleTypeRamp:
		return entity.VehicleTypeRamp, nil
	default:
		return "", ErrInvalidVehicleType
	}
}

func VehicleInitInfoToAPI(vehicleInitInfo *entity.VehicleInitInfo) *api.MovingRegisterVehicleOK {
	return &api.MovingRegisterVehicleOK{
		GarrageNodeId: vehicleInitInfo.GarrageNodeID,
		VehicleId:     vehicleInitInfo.VehicleID,
		ServiceSpots:  vehicleInitInfo.ServiceSpots,
	}
}
