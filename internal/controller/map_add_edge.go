package controller

import (
	"context"
	"errors"

	"github.com/reaport/ground-control/pkg/api"
)

// MapAddEdge implements map_addEdge operation.
//
// Добавляет новое ребро между узлами на карте
// аэропорта.
//
// POST /map/edges
func (c *Controller) MapAddEdge(_ context.Context, _ *api.Edge) (api.MapAddEdgeRes, error) {
	return nil, errors.New("not implemented")
}
