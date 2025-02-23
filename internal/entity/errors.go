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
)
