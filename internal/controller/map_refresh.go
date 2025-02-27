package controller

import (
	"context"
	"fmt"

	"github.com/reaport/ground-control/internal/entity"
	"github.com/reaport/ground-control/pkg/logger"
	"go.uber.org/zap"
)

// MapRefreshAirportMap implements map_refreshAirportMap operation.
//
// Возвращает карту к исходному состоянию.
//
// POST /map/refresh
func (c *Controller) MapRefreshAirportMap(ctx context.Context) error {
	err := c.mapService.RefreshAirportMap(ctx)
	if err != nil {
		err = fmt.Errorf("c.mapService.RefreshAirportMap: %w", err)
		logger.GlobalLogger.Error(
			"failed to refresh map",
			zap.Error(err),
		)
		return err
	}

	logger.GlobalLogger.Info(
		"map refreshed",
	)
	err = c.eventSender.SendEvent(ctx, &entity.Event{
		Type: entity.MapRefreshedEventType,
	})
	if err != nil {
		logger.GlobalLogger.Error(
			"failed to send event",
			zap.Error(fmt.Errorf("c.eventSender.SendEvent: %w", err)),
			zap.String("event_type", string(entity.MapRefreshedEventType)),
		)
	}

	return nil
}
