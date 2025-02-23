package controller

import (
	"context"
	"errors"

	"github.com/reaport/ground-control/pkg/api"
)

// MovingRequestMove implements moving_requestMove operation.
//
// Запрашивает разрешение на перемещение из одного узла
// в другой.
//
// POST /move
func (c *Controller) MovingRequestMove(
	_ context.Context,
	_ *api.MovingRequestMoveReq,
) (api.MovingRequestMoveRes, error) {
	return nil, errors.New("not implemented")
}
