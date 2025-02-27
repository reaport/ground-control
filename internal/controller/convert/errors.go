package convert

import "errors"

var (
	ErrEmptyVehicleType   = errors.New("empty vehicle type")
	ErrInvalidVehicleType = errors.New("invalid vehicle type")
)
