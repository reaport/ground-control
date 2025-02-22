package controller

import (
	"context"
	"errors"

	"github.com/reaport/ground-control/pkg/api"
)

// POST /map/edges.
func (c *Controller) MapAddEdge(_ context.Context, _ *api.Edge) (api.MapAddEdgeRes, error) {
	return nil, errors.New("not implemented")
}
