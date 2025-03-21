// Code generated by ogen, DO NOT EDIT.

package api

import (
	"github.com/go-faster/errors"
)

// AirplaneGetParkingSpotBadRequest is response for AirplaneGetParkingSpot operation.
type AirplaneGetParkingSpotBadRequest struct{}

func (*AirplaneGetParkingSpotBadRequest) airplaneGetParkingSpotRes() {}

// AirplaneGetParkingSpotConflict is response for AirplaneGetParkingSpot operation.
type AirplaneGetParkingSpotConflict struct{}

func (*AirplaneGetParkingSpotConflict) airplaneGetParkingSpotRes() {}

type AirplaneGetParkingSpotOK struct {
	// ID узла.
	NodeId string `json:"nodeId"`
}

// GetNodeId returns the value of NodeId.
func (s *AirplaneGetParkingSpotOK) GetNodeId() string {
	return s.NodeId
}

// SetNodeId sets the value of NodeId.
func (s *AirplaneGetParkingSpotOK) SetNodeId(val string) {
	s.NodeId = val
}

func (*AirplaneGetParkingSpotOK) airplaneGetParkingSpotRes() {}

// AirplaneTakeOffNotFound is response for AirplaneTakeOff operation.
type AirplaneTakeOffNotFound struct{}

func (*AirplaneTakeOffNotFound) airplaneTakeOffRes() {}

// AirplaneTakeOffOK is response for AirplaneTakeOff operation.
type AirplaneTakeOffOK struct{}

func (*AirplaneTakeOffOK) airplaneTakeOffRes() {}

// Ref: #/components/schemas/AirportMap
type AirportMap struct {
	Nodes []Node `json:"nodes"`
	Edges []Edge `json:"edges"`
}

// GetNodes returns the value of Nodes.
func (s *AirportMap) GetNodes() []Node {
	return s.Nodes
}

// GetEdges returns the value of Edges.
func (s *AirportMap) GetEdges() []Edge {
	return s.Edges
}

// SetNodes sets the value of Nodes.
func (s *AirportMap) SetNodes(val []Node) {
	s.Nodes = val
}

// SetEdges sets the value of Edges.
func (s *AirportMap) SetEdges(val []Edge) {
	s.Edges = val
}

// Ref: #/components/schemas/AirportMapConfig
type AirportMapConfig struct {
	// ID узла взлетно-посадочной полосы.
	AirstripNodeId string `json:"airstripNodeId"`
}

// GetAirstripNodeId returns the value of AirstripNodeId.
func (s *AirportMapConfig) GetAirstripNodeId() string {
	return s.AirstripNodeId
}

// SetAirstripNodeId sets the value of AirstripNodeId.
func (s *AirportMapConfig) SetAirstripNodeId(val string) {
	s.AirstripNodeId = val
}

// Ref: #/components/schemas/Edge
type Edge struct {
	// ID начального узла.
	From string `json:"from"`
	// ID конечного узла.
	To string `json:"to"`
	// Расстояние между узлами.
	Distance float64 `json:"distance"`
}

// GetFrom returns the value of From.
func (s *Edge) GetFrom() string {
	return s.From
}

// GetTo returns the value of To.
func (s *Edge) GetTo() string {
	return s.To
}

// GetDistance returns the value of Distance.
func (s *Edge) GetDistance() float64 {
	return s.Distance
}

// SetFrom sets the value of From.
func (s *Edge) SetFrom(val string) {
	s.From = val
}

// SetTo sets the value of To.
func (s *Edge) SetTo(val string) {
	s.To = val
}

// SetDistance sets the value of Distance.
func (s *Edge) SetDistance(val float64) {
	s.Distance = val
}

// Ref: #/components/schemas/ErrorResponse
type ErrorResponse struct {
	Code ErrorResponseCode `json:"code"`
	// Описание ошибки.
	Message OptString `json:"message"`
}

// GetCode returns the value of Code.
func (s *ErrorResponse) GetCode() ErrorResponseCode {
	return s.Code
}

// GetMessage returns the value of Message.
func (s *ErrorResponse) GetMessage() OptString {
	return s.Message
}

// SetCode sets the value of Code.
func (s *ErrorResponse) SetCode(val ErrorResponseCode) {
	s.Code = val
}

// SetMessage sets the value of Message.
func (s *ErrorResponse) SetMessage(val OptString) {
	s.Message = val
}

func (*ErrorResponse) mapUpdateAirportMapRes() {}
func (*ErrorResponse) movingNotifyArrivalRes() {}
func (*ErrorResponse) movingRequestMoveRes()   {}

type ErrorResponseCode string

const (
	ErrorResponseCodeVEHICLENOTFOUNDINNODE ErrorResponseCode = "VEHICLE_NOT_FOUND_IN_NODE"
	ErrorResponseCodeEDGENOTFOUND          ErrorResponseCode = "EDGE_NOT_FOUND"
	ErrorResponseCodeMAPHASVEHICLES        ErrorResponseCode = "MAP_HAS_VEHICLES"
)

// AllValues returns all ErrorResponseCode values.
func (ErrorResponseCode) AllValues() []ErrorResponseCode {
	return []ErrorResponseCode{
		ErrorResponseCodeVEHICLENOTFOUNDINNODE,
		ErrorResponseCodeEDGENOTFOUND,
		ErrorResponseCodeMAPHASVEHICLES,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s ErrorResponseCode) MarshalText() ([]byte, error) {
	switch s {
	case ErrorResponseCodeVEHICLENOTFOUNDINNODE:
		return []byte(s), nil
	case ErrorResponseCodeEDGENOTFOUND:
		return []byte(s), nil
	case ErrorResponseCodeMAPHASVEHICLES:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *ErrorResponseCode) UnmarshalText(data []byte) error {
	switch ErrorResponseCode(data) {
	case ErrorResponseCodeVEHICLENOTFOUNDINNODE:
		*s = ErrorResponseCodeVEHICLENOTFOUNDINNODE
		return nil
	case ErrorResponseCodeEDGENOTFOUND:
		*s = ErrorResponseCodeEDGENOTFOUND
		return nil
	case ErrorResponseCodeMAPHASVEHICLES:
		*s = ErrorResponseCodeMAPHASVEHICLES
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

// MapRefreshAirportMapOK is response for MapRefreshAirportMap operation.
type MapRefreshAirportMapOK struct{}

// MapUpdateAirportMapOK is response for MapUpdateAirportMap operation.
type MapUpdateAirportMapOK struct{}

func (*MapUpdateAirportMapOK) mapUpdateAirportMapRes() {}

// MovingGetRouteNotFound is response for MovingGetRoute operation.
type MovingGetRouteNotFound struct{}

func (*MovingGetRouteNotFound) movingGetRouteRes() {}

type MovingGetRouteOKApplicationJSON []string

func (*MovingGetRouteOKApplicationJSON) movingGetRouteRes() {}

type MovingGetRouteReq struct {
	// ID начального узла.
	From string `json:"from"`
	// ID конечного узла.
	To string `json:"to"`
	// Тип транспорта.
	Type VehicleType `json:"type"`
}

// GetFrom returns the value of From.
func (s *MovingGetRouteReq) GetFrom() string {
	return s.From
}

// GetTo returns the value of To.
func (s *MovingGetRouteReq) GetTo() string {
	return s.To
}

// GetType returns the value of Type.
func (s *MovingGetRouteReq) GetType() VehicleType {
	return s.Type
}

// SetFrom sets the value of From.
func (s *MovingGetRouteReq) SetFrom(val string) {
	s.From = val
}

// SetTo sets the value of To.
func (s *MovingGetRouteReq) SetTo(val string) {
	s.To = val
}

// SetType sets the value of Type.
func (s *MovingGetRouteReq) SetType(val VehicleType) {
	s.Type = val
}

// MovingNotifyArrivalNotFound is response for MovingNotifyArrival operation.
type MovingNotifyArrivalNotFound struct{}

func (*MovingNotifyArrivalNotFound) movingNotifyArrivalRes() {}

// MovingNotifyArrivalOK is response for MovingNotifyArrival operation.
type MovingNotifyArrivalOK struct{}

func (*MovingNotifyArrivalOK) movingNotifyArrivalRes() {}

type MovingNotifyArrivalReq struct {
	// ID транспорта.
	VehicleId   string      `json:"vehicleId"`
	VehicleType VehicleType `json:"vehicleType"`
	// ID узла.
	NodeId string `json:"nodeId"`
}

// GetVehicleId returns the value of VehicleId.
func (s *MovingNotifyArrivalReq) GetVehicleId() string {
	return s.VehicleId
}

// GetVehicleType returns the value of VehicleType.
func (s *MovingNotifyArrivalReq) GetVehicleType() VehicleType {
	return s.VehicleType
}

// GetNodeId returns the value of NodeId.
func (s *MovingNotifyArrivalReq) GetNodeId() string {
	return s.NodeId
}

// SetVehicleId sets the value of VehicleId.
func (s *MovingNotifyArrivalReq) SetVehicleId(val string) {
	s.VehicleId = val
}

// SetVehicleType sets the value of VehicleType.
func (s *MovingNotifyArrivalReq) SetVehicleType(val VehicleType) {
	s.VehicleType = val
}

// SetNodeId sets the value of NodeId.
func (s *MovingNotifyArrivalReq) SetNodeId(val string) {
	s.NodeId = val
}

// MovingRegisterVehicleBadRequest is response for MovingRegisterVehicle operation.
type MovingRegisterVehicleBadRequest struct{}

func (*MovingRegisterVehicleBadRequest) movingRegisterVehicleRes() {}

// MovingRegisterVehicleConflict is response for MovingRegisterVehicle operation.
type MovingRegisterVehicleConflict struct{}

func (*MovingRegisterVehicleConflict) movingRegisterVehicleRes() {}

type MovingRegisterVehicleOK struct {
	// ID узла.
	GarrageNodeId string `json:"garrageNodeId"`
	// ID транспорта.
	VehicleId string `json:"vehicleId"`
	// Id узлов парковочных мест для обслуживания самолетов
	// (парковка\_самолета:парковка\_сервисной\_машинки).
	ServiceSpots MovingRegisterVehicleOKServiceSpots `json:"serviceSpots"`
}

// GetGarrageNodeId returns the value of GarrageNodeId.
func (s *MovingRegisterVehicleOK) GetGarrageNodeId() string {
	return s.GarrageNodeId
}

// GetVehicleId returns the value of VehicleId.
func (s *MovingRegisterVehicleOK) GetVehicleId() string {
	return s.VehicleId
}

// GetServiceSpots returns the value of ServiceSpots.
func (s *MovingRegisterVehicleOK) GetServiceSpots() MovingRegisterVehicleOKServiceSpots {
	return s.ServiceSpots
}

// SetGarrageNodeId sets the value of GarrageNodeId.
func (s *MovingRegisterVehicleOK) SetGarrageNodeId(val string) {
	s.GarrageNodeId = val
}

// SetVehicleId sets the value of VehicleId.
func (s *MovingRegisterVehicleOK) SetVehicleId(val string) {
	s.VehicleId = val
}

// SetServiceSpots sets the value of ServiceSpots.
func (s *MovingRegisterVehicleOK) SetServiceSpots(val MovingRegisterVehicleOKServiceSpots) {
	s.ServiceSpots = val
}

func (*MovingRegisterVehicleOK) movingRegisterVehicleRes() {}

// Id узлов парковочных мест для обслуживания самолетов
// (парковка\_самолета:парковка\_сервисной\_машинки).
type MovingRegisterVehicleOKServiceSpots map[string]string

func (s *MovingRegisterVehicleOKServiceSpots) init() MovingRegisterVehicleOKServiceSpots {
	m := *s
	if m == nil {
		m = map[string]string{}
		*s = m
	}
	return m
}

// MovingRequestMoveConflict is response for MovingRequestMove operation.
type MovingRequestMoveConflict struct{}

func (*MovingRequestMoveConflict) movingRequestMoveRes() {}

// MovingRequestMoveForbidden is response for MovingRequestMove operation.
type MovingRequestMoveForbidden struct{}

func (*MovingRequestMoveForbidden) movingRequestMoveRes() {}

// MovingRequestMoveNotFound is response for MovingRequestMove operation.
type MovingRequestMoveNotFound struct{}

func (*MovingRequestMoveNotFound) movingRequestMoveRes() {}

type MovingRequestMoveOK struct {
	// Расстояние до следующего узла.
	Distance float64 `json:"distance"`
}

// GetDistance returns the value of Distance.
func (s *MovingRequestMoveOK) GetDistance() float64 {
	return s.Distance
}

// SetDistance sets the value of Distance.
func (s *MovingRequestMoveOK) SetDistance(val float64) {
	s.Distance = val
}

func (*MovingRequestMoveOK) movingRequestMoveRes() {}

type MovingRequestMoveReq struct {
	// ID транспорта.
	VehicleId   string      `json:"vehicleId"`
	VehicleType VehicleType `json:"vehicleType"`
	// ID текущего узла.
	From string `json:"from"`
	// ID следующего узла.
	To string `json:"to"`
	// ID самолета, который следует за follow-me.
	WithAirplane OptString `json:"withAirplane"`
}

// GetVehicleId returns the value of VehicleId.
func (s *MovingRequestMoveReq) GetVehicleId() string {
	return s.VehicleId
}

// GetVehicleType returns the value of VehicleType.
func (s *MovingRequestMoveReq) GetVehicleType() VehicleType {
	return s.VehicleType
}

// GetFrom returns the value of From.
func (s *MovingRequestMoveReq) GetFrom() string {
	return s.From
}

// GetTo returns the value of To.
func (s *MovingRequestMoveReq) GetTo() string {
	return s.To
}

// GetWithAirplane returns the value of WithAirplane.
func (s *MovingRequestMoveReq) GetWithAirplane() OptString {
	return s.WithAirplane
}

// SetVehicleId sets the value of VehicleId.
func (s *MovingRequestMoveReq) SetVehicleId(val string) {
	s.VehicleId = val
}

// SetVehicleType sets the value of VehicleType.
func (s *MovingRequestMoveReq) SetVehicleType(val VehicleType) {
	s.VehicleType = val
}

// SetFrom sets the value of From.
func (s *MovingRequestMoveReq) SetFrom(val string) {
	s.From = val
}

// SetTo sets the value of To.
func (s *MovingRequestMoveReq) SetTo(val string) {
	s.To = val
}

// SetWithAirplane sets the value of WithAirplane.
func (s *MovingRequestMoveReq) SetWithAirplane(val OptString) {
	s.WithAirplane = val
}

// Ref: #/components/schemas/Node
type Node struct {
	// Уникальный идентификатор узла.
	ID       string        `json:"id"`
	Types    []VehicleType `json:"types"`
	Vehicles []Vehicle     `json:"vehicles"`
}

// GetID returns the value of ID.
func (s *Node) GetID() string {
	return s.ID
}

// GetTypes returns the value of Types.
func (s *Node) GetTypes() []VehicleType {
	return s.Types
}

// GetVehicles returns the value of Vehicles.
func (s *Node) GetVehicles() []Vehicle {
	return s.Vehicles
}

// SetID sets the value of ID.
func (s *Node) SetID(val string) {
	s.ID = val
}

// SetTypes sets the value of Types.
func (s *Node) SetTypes(val []VehicleType) {
	s.Types = val
}

// SetVehicles sets the value of Vehicles.
func (s *Node) SetVehicles(val []Vehicle) {
	s.Vehicles = val
}

// NewOptString returns new OptString with value set to v.
func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}

// OptString is optional string.
type OptString struct {
	Value string
	Set   bool
}

// IsSet returns true if OptString was set.
func (o OptString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// Ref: #/components/schemas/Vehicle
type Vehicle struct {
	// Уникальный идентификатор транспорта.
	ID   string      `json:"id"`
	Type VehicleType `json:"type"`
}

// GetID returns the value of ID.
func (s *Vehicle) GetID() string {
	return s.ID
}

// GetType returns the value of Type.
func (s *Vehicle) GetType() VehicleType {
	return s.Type
}

// SetID sets the value of ID.
func (s *Vehicle) SetID(val string) {
	s.ID = val
}

// SetType sets the value of Type.
func (s *Vehicle) SetType(val VehicleType) {
	s.Type = val
}

// Тип транспорта.
// Ref: #/components/schemas/VehicleType
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

// AllValues returns all VehicleType values.
func (VehicleType) AllValues() []VehicleType {
	return []VehicleType{
		VehicleTypeAirplane,
		VehicleTypeCatering,
		VehicleTypeRefueling,
		VehicleTypeCleaning,
		VehicleTypeBaggage,
		VehicleTypeFollowMe,
		VehicleTypeCharging,
		VehicleTypeBus,
		VehicleTypeRamp,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s VehicleType) MarshalText() ([]byte, error) {
	switch s {
	case VehicleTypeAirplane:
		return []byte(s), nil
	case VehicleTypeCatering:
		return []byte(s), nil
	case VehicleTypeRefueling:
		return []byte(s), nil
	case VehicleTypeCleaning:
		return []byte(s), nil
	case VehicleTypeBaggage:
		return []byte(s), nil
	case VehicleTypeFollowMe:
		return []byte(s), nil
	case VehicleTypeCharging:
		return []byte(s), nil
	case VehicleTypeBus:
		return []byte(s), nil
	case VehicleTypeRamp:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *VehicleType) UnmarshalText(data []byte) error {
	switch VehicleType(data) {
	case VehicleTypeAirplane:
		*s = VehicleTypeAirplane
		return nil
	case VehicleTypeCatering:
		*s = VehicleTypeCatering
		return nil
	case VehicleTypeRefueling:
		*s = VehicleTypeRefueling
		return nil
	case VehicleTypeCleaning:
		*s = VehicleTypeCleaning
		return nil
	case VehicleTypeBaggage:
		*s = VehicleTypeBaggage
		return nil
	case VehicleTypeFollowMe:
		*s = VehicleTypeFollowMe
		return nil
	case VehicleTypeCharging:
		*s = VehicleTypeCharging
		return nil
	case VehicleTypeBus:
		*s = VehicleTypeBus
		return nil
	case VehicleTypeRamp:
		*s = VehicleTypeRamp
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}
