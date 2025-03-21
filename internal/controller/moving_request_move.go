package controller

import (
	"context"
	"errors"
	"fmt"

	"github.com/reaport/ground-control/internal/controller/convert"
	"github.com/reaport/ground-control/internal/entity"
	"github.com/reaport/ground-control/pkg/api"
	"github.com/reaport/ground-control/pkg/logger"
	"go.uber.org/zap"
)

// MovingRequestMove implements moving_requestMove operation.
//
// Запрашивает разрешение на перемещение из одного узла
// в другой.
//
// POST /move
func (c *Controller) MovingRequestMove( //nolint:funlen // a lot of logs
	ctx context.Context,
	req *api.MovingRequestMoveReq,
) (api.MovingRequestMoveRes, error) {
	vehicleType, err := convert.VehicleTypeFromAPI(req.VehicleType)
	if err != nil {
		err = fmt.Errorf("VehicleTypeFromAPI: %w", err)
		logger.GlobalLogger.Error(
			"failed to convert vehicle type from API",
			zap.String("error", err.Error()),
			zap.String("type", string(req.VehicleType)),
		)
		return nil, err
	}

	var withAirplane *string
	if req.WithAirplane.Set && req.WithAirplane.Value != "" {
		withAirplane = &req.WithAirplane.Value
	}

	distance, err := c.mapService.RequestMove(ctx, req.VehicleId, req.From, req.To, vehicleType, withAirplane)
	if err != nil {
		err = fmt.Errorf("c.mapService.RequestMove: %w", err)
		switch {
		case errors.Is(err, entity.ErrNodeNotFound):
			logger.GlobalLogger.Error(
				"one of both nodes not found",
				zap.String("error", err.Error()),
				zap.String("vehicle_id", req.VehicleId),
				zap.String("from", req.From),
				zap.String("to", req.To),
				zap.String("vehicle_type", string(req.VehicleType)),
				zap.String("with_airplane", req.WithAirplane.Value),
			)
			return &api.MovingRequestMoveNotFound{}, nil

		case errors.Is(err, entity.ErrInvalidVehicleType):
			logger.GlobalLogger.Error(
				"invalid vehicle type",
				zap.String("error", err.Error()),
				zap.String("vehicle_id", req.VehicleId),
				zap.String("from", req.From),
				zap.String("to", req.To),
				zap.String("vehicle_type", string(req.VehicleType)),
				zap.String("with_airplane", req.WithAirplane.Value),
			)
			return &api.MovingRequestMoveForbidden{}, nil

		case errors.Is(err, entity.ErrVehicleNotFound):
			logger.GlobalLogger.Error(
				"vehicle not found",
				zap.String("error", err.Error()),
				zap.String("vehicle_id", req.VehicleId),
				zap.String("from", req.From),
				zap.String("to", req.To),
				zap.String("vehicle_type", string(req.VehicleType)),
				zap.String("with_airplane", req.WithAirplane.Value),
			)
			return &api.ErrorResponse{
				Code: api.ErrorResponseCodeVEHICLENOTFOUNDINNODE,
			}, nil

		case errors.Is(err, entity.ErrEdgeNotFound):
			logger.GlobalLogger.Error(
				"edge not found",
				zap.String("error", err.Error()),
				zap.String("vehicle_id", req.VehicleId),
				zap.String("from", req.From),
				zap.String("to", req.To),
				zap.String("vehicle_type", string(req.VehicleType)),
				zap.String("with_airplane", req.WithAirplane.Value),
			)
			return &api.ErrorResponse{
				Code: api.ErrorResponseCodeEDGENOTFOUND,
			}, nil

		case errors.Is(err, entity.ErrMoveNotAllowed):
			logger.GlobalLogger.Error(
				"move not allowed",
				zap.String("error", err.Error()),
				zap.String("vehicle_id", req.VehicleId),
				zap.String("from", req.From),
				zap.String("to", req.To),
				zap.String("vehicle_type", string(req.VehicleType)),
				zap.String("with_airplane", req.WithAirplane.Value),
			)
			return &api.MovingRequestMoveConflict{}, nil

		default:
			logger.GlobalLogger.Error(
				"failed to request move",
				zap.String("error", err.Error()),
				zap.String("vehicle_id", req.VehicleId),
				zap.String("from", req.From),
				zap.String("to", req.To),
				zap.String("vehicle_type", string(req.VehicleType)),
				zap.String("with_airplane", req.WithAirplane.Value),
			)
			return nil, err
		}
	}

	logger.GlobalLogger.Info(
		"move requested",
		zap.String("vehicle_id", req.VehicleId),
		zap.String("from", req.From),
		zap.String("to", req.To),
		zap.String("vehicle_type", string(req.VehicleType)),
		zap.String("with_airplane", req.WithAirplane.Value),
		zap.Float64("distance", distance),
	)

	err = c.eventSender.SendEvent(ctx, &entity.Event{
		Type: entity.VehicleLeftNodeEventType,
		Data: entity.EventData{
			"vehicle_id":    req.VehicleId,
			"vehicle_type":  string(req.VehicleType),
			"from":          req.From,
			"to":            req.To,
			"distance":      distance,
			"with_airplane": req.WithAirplane.Value,
		},
	})
	if err != nil {
		logger.GlobalLogger.Error(
			"failed to send event",
			zap.String("error", err.Error()),
			zap.String("vehicle_id", req.VehicleId),
			zap.String("vehicle_type", string(req.VehicleType)),
			zap.String("from", req.From),
			zap.String("to", req.To),
			zap.Float64("distance", distance),
			zap.String("with_airplane", req.WithAirplane.Value),
		)
	}

	return &api.MovingRequestMoveOK{
		Distance: distance,
	}, nil
}
