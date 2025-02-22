package entity

const (
	nodeCapacity = 2
)

type AirportMap struct {
	Nodes []*Node `json:"nodes"`
	Edges []*Edge `json:"edges"`
}

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

func (node *Node) AddVehicle(vehicle *Vehicle) {
	node.Vehicles = append(node.Vehicles, vehicle)
}

type Vehicle struct {
	ID   string
	Type VehicleType
}

func NewVehicle(id string, vehicleType VehicleType) *Vehicle {
	return &Vehicle{
		ID:   id,
		Type: vehicleType,
	}
}

type Edge struct {
	From     string  `json:"from"`
	To       string  `json:"to"`
	Distance float64 `json:"distance"`
}

func NewEdge(from string, to string, distance float64) *Edge {
	return &Edge{
		From:     from,
		To:       to,
		Distance: distance,
	}
}

type VehicleType string

const (
	VehicleTypeAirplane  VehicleType = "airplane"
	VehicleTypeCatering  VehicleType = "catering"
	VehicleTypeRefueling VehicleType = "refueling"
	VehicleTypeCleaning  VehicleType = "cleaning"
	VehicleTypeBaggage   VehicleType = "baggage"
	VehicleTypeFollowMe  VehicleType = "follow-me"
	VehicleTypeCharging  VehicleType = "charging"
	VehicleTypeBus       VehicleType = "bus"
	VehicleTypeRamp      VehicleType = "ramp"
)
