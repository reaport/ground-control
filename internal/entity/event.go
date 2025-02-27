package entity

type EventType string

const (
	GroundControlStartedEventType   EventType = "ground_control_started"
	AirplaneGetParkingSpotEventType EventType = "airplane_get_parking_spot"
	MapRefreshedEventType           EventType = "map_refreshed"
	MapUpdatedEventType             EventType = "map_updated"
	RouteCalculatedEventType        EventType = "route_calculated"
	VehicleArrivedEventType         EventType = "vehicle_arrived"
	VehicleRegisteredEventType      EventType = "vehicle_registered"
	VehicleLeftNodeEventType        EventType = "vehicle_left_node"
)

type EventData map[string]interface{}

type Event struct {
	Type EventType `json:"type"`
	Data EventData `json:"data"`
}
