package graphmap

import (
	"context"
	"math"

	"github.com/reaport/ground-control/internal/entity"
)

func (s *Service) GetRoute(ctx context.Context, nodeIDFrom, nodeIDTo string, vehicleType entity.VehicleType) ([]string, error) {
	if s.airportMap == nil {
		return nil, entity.ErrAirportMapIsNotInitialized
	}

	if nodeIDFrom == nodeIDTo {
		return nil, entity.ErrSameNodes
	}

	fromNode, toNode := s.findNodeByID(nodeIDFrom), s.findNodeByID(nodeIDTo)
	if fromNode == nil || toNode == nil {
		return nil, entity.ErrNodeNotFound
	}

	if !toNode.IsValidType(vehicleType) {
		return nil, entity.ErrInvalidVehicleType
	}

	// Инициализация данных для алгоритма Дейкстры
	distances := make(map[string]float64) // Расстояния от начального узла
	previous := make(map[string]string)   // Предыдущий узел в кратчайшем пути
	nodes := make(map[string]bool)        // Все узлы

	// Инициализация расстояний и множества узлов
	for _, node := range s.airportMap.Nodes {
		distances[node.ID] = math.Inf(1) // Бесконечность
		previous[node.ID] = ""
		nodes[node.ID] = true
	}
	distances[nodeIDFrom] = 0 // Расстояние до начального узла равно 0

	// Основной цикл алгоритма Дейкстры
	for len(nodes) > 0 {
		// Находим узел с минимальным расстоянием
		minNode := s.findMinDistanceNode(nodes, distances)
		if minNode == "" {
			break // Все оставшиеся узлы недостижимы
		}

		// Удаляем узел из множества непосещённых
		delete(nodes, minNode)

		// Если достигли конечного узла, строим путь
		if minNode == nodeIDTo {
			return s.buildPath(previous, nodeIDTo), nil
		}

		// Обновляем расстояния до соседей, учитывая vehicleType
		for _, edge := range s.airportMap.Edges {
			// Проверяем оба направления ребра (неориентированный граф)
			if edge.From == minNode || edge.To == minNode {
				// Определяем целевой узел
				targetNodeID := edge.To
				if edge.To == minNode {
					targetNodeID = edge.From
				}

				// Проверяем, поддерживает ли целевой узел vehicleType
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

	return nil, entity.ErrRouteNotFound
}

// Вспомогательный метод для поиска узла по ID
func (s *Service) findNodeByID(nodeID string) *entity.Node {
	for _, node := range s.airportMap.Nodes {
		if node.ID == nodeID {
			return node
		}
	}
	return nil
}

// Вспомогательный метод для поиска узла с минимальным расстоянием
func (s *Service) findMinDistanceNode(nodes map[string]bool, distances map[string]float64) string {
	minNode := ""
	minDist := math.Inf(1)
	for node := range nodes {
		if distances[node] < minDist {
			minDist = distances[node]
			minNode = node
		}
	}
	return minNode
}

// Вспомогательный метод для построения пути
func (s *Service) buildPath(previous map[string]string, nodeIDTo string) []string {
	path := []string{}
	for node := nodeIDTo; node != ""; node = previous[node] {
		path = append([]string{node}, path...)
	}
	return path
}
