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

// AirplaneTakeOff implements airplane_takeOff operation.
//
// Удаляется самолет с карты.
//
// POST /airplane/{id}/take-off
func (c *Controller) AirplaneTakeOff(
	ctx context.Context,
	params api.AirplaneTakeOffParams,
) (api.AirplaneTakeOffRes, error) {
	err := c.mapService.TakeOffAirplane(ctx, params.ID)
	if err != nil {
		err = fmt.Errorf("c.mapService.TakeOffAirplane: %w", err)
		if errors.Is(err, entity.ErrVehicleNotFound) {
			logger.GlobalLogger.Error(
				"airplane not found in airstrip",
				zap.String("error", err.Error()),
				zap.String("airplane_id", params.ID),
			)
			return &api.AirplaneTakeOffNotFound{}, nil
		}

		logger.GlobalLogger.Error(
			"failed to take off airplane",
			zap.String("error", err.Error()),
			zap.String("airplane_id", params.ID),
		)

		return nil, err
	}

	logger.GlobalLogger.Info(
		"airplane take off",
		zap.String("airplane_id", params.ID),
	)

	return &api.AirplaneTakeOffOK{}, nil
}
