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

// MovingGetRoute implements moving_getRoute operation.
//
// Запрашивает маршрут из точки А в точку Б.
//
// POST /route
func (c *Controller) MovingGetRoute(ctx context.Context, req *api.MovingGetRouteReq) (api.MovingGetRouteRes, error) {
	vehicleType, err := convert.VehicleTypeFromAPI(req.Type)
	if err != nil {
		err = fmt.Errorf("VehicleTypeFromAPI: %w", err)
		logger.GlobalLogger.Error(
			"failed to convert vehicle type from API",
			zap.Error(err),
			zap.String("type", string(req.Type)),
		)
		return nil, err
	}

	route, err := c.mapService.GetRoute(ctx, req.From, req.To, vehicleType)
	if err != nil {
		err = fmt.Errorf("c.mapService.GetRoute: %w", err)

		switch {
		case errors.Is(err, entity.ErrSameNodes):
			logger.GlobalLogger.Error(
				"requested route for same nodes",
				zap.Error(err),
				zap.String("from", req.From),
				zap.String("to", req.To),
				zap.String("vehicle_type", string(req.Type)),
			)
			return &api.MovingGetRouteNotFound{}, nil

		case errors.Is(err, entity.ErrNodeNotFound):
			logger.GlobalLogger.Error(
				"one or both nodes not found",
				zap.Error(err),
				zap.String("from", req.From),
				zap.String("to", req.To),
				zap.String("vehicle_type", string(req.Type)),
			)
			return &api.MovingGetRouteNotFound{}, nil

		case errors.Is(err, entity.ErrInvalidVehicleType):
			logger.GlobalLogger.Error(
				"invalid vehicle type for destination",
				zap.Error(err),
				zap.String("from", req.From),
				zap.String("to", req.To),
				zap.String("vehicle_type", string(req.Type)),
			)
			return &api.MovingGetRouteNotFound{}, nil

		case errors.Is(err, entity.ErrRouteNotFound):
			logger.GlobalLogger.Error(
				"route not found",
				zap.Error(err),
				zap.String("from", req.From),
				zap.String("to", req.To),
				zap.String("vehicle_type", string(req.Type)),
			)
			return &api.MovingGetRouteNotFound{}, nil

		default:
			logger.GlobalLogger.Error(
				"failed to get route",
				zap.Error(err),
				zap.String("from", req.From),
				zap.String("to", req.To),
				zap.String("vehicle_type", string(req.Type)),
			)
			return nil, fmt.Errorf("c.mapService.GetRoute: %w", err)
		}
	}

	apiRoute := api.MovingGetRouteOKApplicationJSON(route)
	logger.GlobalLogger.Info(
		"got route",
		zap.String("from", req.From),
		zap.String("to", req.To),
		zap.String("vehicle_type", string(req.Type)),
		zap.Any("route", route),
	)

	return &apiRoute, nil
}
