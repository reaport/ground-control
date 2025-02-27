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

// MapUpdateAirportMap implements map_updateAirportMap operation.
//
// Обновляет карту аэропорта.
//
// PUT /map
func (c *Controller) MapUpdateAirportMap(ctx context.Context, req *api.AirportMap) (api.MapUpdateAirportMapRes, error) {
	airportMap, err := convert.AirportMapFromAPI(req)
	if err != nil {
		err = fmt.Errorf("convert.AirportMapFromAPI: %w", err)
		logger.GlobalLogger.Error(
			"failed to convert airport map from API",
			zap.Error(err),
			zap.Any("map", req),
		)
		return nil, err
	}

	err = c.mapService.UpdateAirportMap(ctx, airportMap)
	if err != nil {
		err = fmt.Errorf("c.mapService.UpdateAirportMap: %w", err)
		if errors.Is(err, entity.ErrMapHasVehicles) {
			logger.GlobalLogger.Error(
				"map has vehicles",
				zap.Error(err),
				zap.Any("map", req),
			)
			return &api.ErrorResponse{
				Code: api.ErrorResponseCodeMAPHASVEHICLES,
			}, nil
		}
		logger.GlobalLogger.Error(
			"failed to update map",
			zap.Error(err),
			zap.Any("map", req),
		)
		return nil, err
	}

	logger.GlobalLogger.Info(
		"map updated",
		zap.Any("map", req),
	)

	err = c.eventSender.SendEvent(ctx, &entity.Event{
		Type: entity.MapUpdatedEventType,
		Data: entity.EventData{
			"map": req,
		},
	})
	if err != nil {
		logger.GlobalLogger.Error(
			"failed to send event",
			zap.Error(fmt.Errorf("c.eventSender.SendEvent: %w", err)),
			zap.String("event_type", string(entity.MapUpdatedEventType)),
			zap.Any("map", req),
		)
	}

	return &api.MapUpdateAirportMapOK{}, nil
}
