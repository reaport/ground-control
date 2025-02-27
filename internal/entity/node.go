package entity

type Node struct {
	ID       string        `json:"id"`
	Types    []VehicleType `json:"types"`
	Vehicles []*Vehicle    `json:"-"`
}

func NewNode(id string, types []VehicleType) *Node {
	return &Node{
		ID:       id,
		Types:    types,
		Vehicles: make([]*Vehicle, 0, nodeCapacity),
	}
}

func (n *Node) AddVehicle(vehicle *Vehicle) {
	n.Vehicles = append(n.Vehicles, vehicle)
}

func (n *Node) ContainsVehicle(vehicleID string) bool {
	for _, vehicle := range n.Vehicles {
		if vehicle.ID == vehicleID {
			return true
		}
	}
	return false
}

func (n *Node) RemoveVehicle(vehicleID string) {
	for i, vehicle := range n.Vehicles {
		if vehicle.ID == vehicleID {
			n.Vehicles = append(n.Vehicles[:i], n.Vehicles[i+1:]...)
			break
		}
	}
}

func (n *Node) IsValidType(vehicleType VehicleType) bool {
	for _, t := range n.Types {
		if t == vehicleType {
			return true
		}
	}

	return false
}
