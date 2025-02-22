package entity

type AirportMap struct {
	Nodes []Node `json:"nodes"`
	Edges []Edge `json:"edges"`
}

type Node struct {
	ID       string        `json:"id"`
	Types    []VehicleType `json:"types"`
	Vehicles []*Vehicle    `json:"-"`
}

type Vehicle struct {
	ID   string
	Type VehicleType
}

type Edge struct {
	From     string  `json:"from"`
	To       string  `json:"to"`
	Distance float64 `json:"distance"`
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
