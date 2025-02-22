package controller

import (
	"context"
	"errors"

	"github.com/reaport/ground-control/pkg/api"
)

// POST /route.
func (c *Controller) MovingGetRoute(_ context.Context, _ *api.MovingGetRouteReq) (api.MovingGetRouteRes, error) {
	return nil, errors.New("not implemented")
}
