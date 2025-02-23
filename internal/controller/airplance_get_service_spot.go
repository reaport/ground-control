package controller

import (
	"context"
	"errors"

	"github.com/reaport/ground-control/pkg/api"
)

// AirplaneIDServiceTypeGet implements GET /airplane/{id}/service/{type} operation.
//
// В зависимости от типа транспорта отдает нужный узел
// для парковки.
//
// GET /airplane/{id}/service/{type}
func (c *Controller) AirplaneIDServiceTypeGet(
	_ context.Context,
	_ api.AirplaneIDServiceTypeGetParams,
) (api.AirplaneIDServiceTypeGetRes, error) {
	return nil, errors.New("not implemented")
}
