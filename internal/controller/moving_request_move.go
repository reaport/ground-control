package controller

import (
	"context"
	"errors"

	"github.com/reaport/ground-control/pkg/api"
)

// POST /move.
func (c *Controller) MovingRequestMove(
	_ context.Context,
	_ *api.MovingRequestMoveReq,
) (api.MovingRequestMoveRes, error) {
	return nil, errors.New("not implemented")
}
