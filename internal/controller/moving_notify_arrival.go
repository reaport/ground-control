package controller

import (
	"context"
	"errors"
	"fmt"

	"github.com/reaport/ground-control/internal/entity"
	"github.com/reaport/ground-control/pkg/api"
	"github.com/reaport/ground-control/pkg/logger"
	"go.uber.org/zap"
)

// MovingNotifyArrival implements moving_notifyArrival operation.
//
// Уведомляет вышку о прибытии транспорта в узел.
//
// POST /arrived
func (c *Controller) MovingNotifyArrival(
	ctx context.Context,
	req *api.MovingNotifyArrivalReq,
) (api.MovingNotifyArrivalRes, error) {
	err := c.mapService.NotifyArrival(ctx, req.NodeId, req.VehicleId)
	if err != nil {
		err = fmt.Errorf("c.mapService.NotifyArrival: %w", err)

		switch {
		case errors.Is(err, entity.ErrNodeNotFound):
			logger.GlobalLogger.Error(
				"node not found",
				zap.Error(err),
				zap.String("node_id", req.NodeId),
				zap.String("vehicle_id", req.VehicleId),
			)
			return &api.MovingNotifyArrivalNotFound{}, nil

		case errors.Is(err, entity.ErrVehicleNotFound):
			logger.GlobalLogger.Error(
				"vehicle not found",
				zap.Error(err),
				zap.String("node_id", req.NodeId),
				zap.String("vehicle_id", req.VehicleId),
			)
			return &api.ErrorResponse{
				Code: api.ErrorResponseCodeVEHICLENOTFOUNDINNODE,
			}, nil

		default:
			logger.GlobalLogger.Error(
				"failed to notify arrival",
				zap.Error(err),
				zap.String("node_id", req.NodeId),
				zap.String("vehicle_id", req.VehicleId),
			)
			return nil, err
		}
	}

	logger.GlobalLogger.Info(
		"vehicle arrived",
		zap.String("node_id", req.NodeId),
		zap.String("vehicle_id", req.VehicleId),
	)

	return &api.MovingNotifyArrivalOK{}, nil
}
