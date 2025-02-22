package controller

import (
	"context"
	"errors"

	"github.com/reaport/ground-control/pkg/api"
)

// POST /register-vehicle/{type}.
func (c *Controller) MovingRegisterVehicle(
	_ context.Context,
	_ api.MovingRegisterVehicleParams,
) (api.MovingRegisterVehicleRes, error) {
	return nil, errors.New("not implemented")
}
