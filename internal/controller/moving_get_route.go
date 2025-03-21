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
func (c *Controller) MovingGetRoute( //nolint:funlen // a lof of logs
	ctx context.Context,
	req *api.MovingGetRouteReq,
) (api.MovingGetRouteRes, error) {
	vehicleType, err := convert.VehicleTypeFromAPI(req.Type)
	if err != nil {
		err = fmt.Errorf("VehicleTypeFromAPI: %w", err)
		logger.GlobalLogger.Error(
			"failed to convert vehicle type from API",
			zap.String("error", err.Error()),
			zap.String("type", string(req.Type)),
			zap.String("from", req.From),
			zap.String("to", req.To),
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
				zap.String("error", err.Error()),
				zap.String("from", req.From),
				zap.String("to", req.To),
				zap.String("vehicle_type", string(req.Type)),
			)
			return &api.MovingGetRouteNotFound{}, nil

		case errors.Is(err, entity.ErrNodeNotFound):
			logger.GlobalLogger.Error(
				"one or both nodes not found",
				zap.String("error", err.Error()),
				zap.String("from", req.From),
				zap.String("to", req.To),
				zap.String("vehicle_type", string(req.Type)),
			)
			return &api.MovingGetRouteNotFound{}, nil

		case errors.Is(err, entity.ErrInvalidVehicleType):
			logger.GlobalLogger.Error(
				"invalid vehicle type for destination",
				zap.String("error", err.Error()),
				zap.String("from", req.From),
				zap.String("to", req.To),
				zap.String("vehicle_type", string(req.Type)),
			)
			return &api.MovingGetRouteNotFound{}, nil

		case errors.Is(err, entity.ErrRouteNotFound):
			logger.GlobalLogger.Error(
				"route not found",
				zap.String("error", err.Error()),
				zap.String("from", req.From),
				zap.String("to", req.To),
				zap.String("vehicle_type", string(req.Type)),
			)
			return &api.MovingGetRouteNotFound{}, nil

		default:
			logger.GlobalLogger.Error(
				"failed to get route",
				zap.String("error", err.Error()),
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

	err = c.eventSender.SendEvent(ctx, &entity.Event{
		Type: entity.RouteCalculatedEventType,
		Data: entity.EventData{
			"from":  req.From,
			"to":    req.To,
			"type":  req.Type,
			"route": route,
		},
	})
	if err != nil {
		logger.GlobalLogger.Error(
			"failed to send event",
			zap.Error(fmt.Errorf("c.eventSender.SendEvent: %w", err)),
			zap.String("event_type", string(entity.RouteCalculatedEventType)),
			zap.String("from", req.From),
			zap.String("to", req.To),
			zap.String("vehicle_type", string(req.Type)),
			zap.Any("route", route),
		)
	}

	return &apiRoute, nil
}
