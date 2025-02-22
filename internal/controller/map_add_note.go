package controller

import (
	"context"
	"errors"

	"github.com/reaport/ground-control/pkg/api"
)

// POST /map/nodes.
func (c *Controller) MapAddNode(_ context.Context, _ *api.Node) (api.MapAddNodeRes, error) {
	return nil, errors.New("not implemented")
}
