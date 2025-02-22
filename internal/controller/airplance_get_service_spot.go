package controller

import (
	"context"
	"errors"

	"github.com/reaport/ground-control/pkg/api"
)

// GET /airplane/{id}/service/{type}.
func (c *Controller) AirplaneIDServiceTypeGet(
	_ context.Context,
	_ api.AirplaneIDServiceTypeGetParams,
) (api.AirplaneIDServiceTypeGetRes, error) {
	return nil, errors.New("not implemented")
}
