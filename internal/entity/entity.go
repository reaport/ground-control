package entity

const (
	nodeCapacity = 2
)

type AirportMap struct {
	Nodes []*Node `json:"nodes"`
	Edges []*Edge `json:"edges"`
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

type VehicleInitInfo struct {
	VehicleID     string
	GarrageNodeID string
	ServiceSpots  map[string]string
}

func NewVehicleInitInfo(
	vehicleID string,
	garrageNodeID string,
	serviceSpots map[string]string,
) *VehicleInitInfo {
	return &VehicleInitInfo{
		VehicleID:     vehicleID,
		GarrageNodeID: garrageNodeID,
		ServiceSpots:  serviceSpots,
	}
}
