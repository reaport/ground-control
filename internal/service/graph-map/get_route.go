package graphmap

import (
	"context"
	"fmt"
	"math"

	"github.com/reaport/ground-control/internal/entity"
)

func (s *Service) GetRoute(
	_ context.Context,
	nodeIDFrom, nodeIDTo string,
	vehicleType entity.VehicleType,
) ([]string, error) {
	if s.airportMap == nil {
		return nil, fmt.Errorf("%w: airport map is not initialized", entity.ErrAirportMapIsNotInitialized)
	}

	if nodeIDFrom == nodeIDTo {
		return nil, fmt.Errorf("%w: requested route for same nodes", entity.ErrSameNodes)
	}

	fromNode, toNode := s.findNodeByID(nodeIDFrom), s.findNodeByID(nodeIDTo)
	if fromNode == nil {
		return nil, fmt.Errorf("%w: %s (from)", entity.ErrNodeNotFound, nodeIDFrom)
	}

	if toNode == nil {
		return nil, fmt.Errorf("%w: %s (to)", entity.ErrNodeNotFound, nodeIDTo)
	}

	if !toNode.IsValidType(vehicleType) {
		return nil, fmt.Errorf(
			"%w: node %s does not support vehicle type %s",
			entity.ErrInvalidVehicleType,
			nodeIDTo,
			vehicleType,
		)
	}

	distances := make(map[string]float64)
	previous := make(map[string]string)
	unvisited := make(map[string]bool)

	for _, node := range s.airportMap.Nodes {
		distances[node.ID] = math.Inf(1)
		previous[node.ID] = ""
		unvisited[node.ID] = true
	}
	distances[nodeIDFrom] = 0

	for len(unvisited) > 0 {
		minNode := s.findMinDistanceNode(unvisited, distances)
		if minNode == "" {
			break
		}

		delete(unvisited, minNode)

		if minNode == nodeIDTo {
			return s.buildPath(previous, nodeIDTo), nil
		}

		for _, edge := range s.airportMap.Edges {
			if edge.From == minNode {
				targetNodeID := edge.To
				targetNode := s.findNodeByID(targetNodeID)
				if targetNode != nil && targetNode.IsValidType(vehicleType) {
					alt := distances[minNode] + edge.Distance
					if alt < distances[targetNodeID] {
						distances[targetNodeID] = alt
						previous[targetNodeID] = minNode
					}
				}
			}
		}
	}

	return nil, fmt.Errorf("%w: route not found", entity.ErrRouteNotFound)
}

func (s *Service) findNodeByID(nodeID string) *entity.Node {
	for _, node := range s.airportMap.Nodes {
		if node.ID == nodeID {
			return node
		}
	}
	return nil
}

func (s *Service) findMinDistanceNode(unvisited map[string]bool, distances map[string]float64) string {
	var minNode string
	minDist := math.Inf(1)
	for node := range unvisited {
		if distances[node] < minDist {
			minDist = distances[node]
			minNode = node
		}
	}
	return minNode
}

func (s *Service) buildPath(previous map[string]string, nodeIDTo string) []string {
	path := []string{}
	for node := nodeIDTo; node != ""; node = previous[node] {
		path = append([]string{node}, path...)
	}
	return path
}
