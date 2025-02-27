package entity

import "errors"

var (
	ErrAirstripIsFull   = errors.New("airstrip is full")
	ErrAirstripNotFound = errors.New("airstrip not found")

	ErrAirportMapIsNotInitialized = errors.New("airport map is not initialized")
	ErrSameNodes                  = errors.New("provided nodes are the same")
	ErrNodeNotFound               = errors.New("node not found")
	ErrRouteNotFound              = errors.New("route not found")
	ErrInvalidVehicleType         = errors.New("invalid vehicle type")

	ErrVehicleNotFound = errors.New("vehicle not found")
	ErrEdgeNotFound    = errors.New("edge not found")
	ErrMoveNotAllowed  = errors.New("move is not allowed")

	ErrAirplaneParkingSpotIsFull   = errors.New("airplane parking spot is full")
	ErrAirplaneServiceSpotIsFull   = errors.New("airplane service spot is full")
	ErrAirplaneParkingSpotNotFound = errors.New("airplane parking spot not found")

	ErrMapHasVehicles   = errors.New("map has vehicles")
	ErrInvalidDirection = errors.New("invalid direction")
)
