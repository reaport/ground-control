package controller

import (
	"context"
	"errors"

	"github.com/reaport/ground-control/pkg/api"
)

// POST /arrived.
func (c *Controller) MovingNotifyArrival(
	_ context.Context,
	_ *api.MovingNotifyArrivalReq,
) (api.MovingNotifyArrivalRes, error) {
	return nil, errors.New("not implemented")
}
