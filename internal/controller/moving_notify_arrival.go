package controller

import (
	"context"
	"errors"

	"github.com/reaport/ground-control/pkg/api"
)

// MovingNotifyArrival implements moving_notifyArrival operation.
//
// Уведомляет вышку о прибытии транспорта в узел.
//
// POST /arrived
func (c *Controller) MovingNotifyArrival(
	_ context.Context,
	_ *api.MovingNotifyArrivalReq,
) (api.MovingNotifyArrivalRes, error) {
	return nil, errors.New("not implemented")
}
