package controller

import (
	"context"
	"errors"

	"github.com/reaport/ground-control/pkg/api"
)

// MapAddNode implements map_addNode operation.
//
// Добавляет новый узел на карту аэропорта.
//
// POST /map/nodes
func (c *Controller) MapAddNode(_ context.Context, _ *api.Node) (api.MapAddNodeRes, error) {
	return nil, errors.New("not implemented")
}
