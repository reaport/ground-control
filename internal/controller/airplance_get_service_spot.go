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

// AirplaneIDServiceTypeGet implements GET /airplane/{id}/service/{type} operation.
//
// В зависимости от типа транспорта отдает нужный узел
// для парковки.
//
// GET /airplane/{id}/service/{type}
func (c *Controller) AirplaneIDServiceTypeGet(
	ctx context.Context,
	req api.AirplaneIDServiceTypeGetParams,
) (api.AirplaneIDServiceTypeGetRes, error) {
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

	nodeID, err := c.mapService.GetAirplaneServiceSpot(ctx, req.ID, vehicleType)
	if err != nil {
		err = fmt.Errorf("c.mapService.GetAirplaneServiceSpot: %w", err)

		switch {
		case errors.Is(err, entity.ErrAirplaneServiceSpotIsFull):
			logger.GlobalLogger.Error(
				"airplane service spot is full",
				zap.Error(err),
				zap.String("airplane_id", req.ID),
				zap.String("vehicle_type", string(req.Type)),
			)
			return &api.AirplaneIDServiceTypeGetConflict{}, nil

		case errors.Is(err, entity.ErrAirplaneParkingSpotNotFound):
			logger.GlobalLogger.Error(
				"airplane parking spot not found",
				zap.Error(err),
				zap.String("airplane_id", req.ID),
				zap.String("vehicle_type", string(req.Type)),
			)
			return &api.AirplaneIDServiceTypeGetNotFound{}, nil

		default:
			logger.GlobalLogger.Error(
				"failed to get airplane service spot",
				zap.Error(err),
				zap.String("airplane_id", req.ID),
				zap.String("vehicle_type", string(req.Type)),
			)
			return nil, err
		}
	}

	logger.GlobalLogger.Info(
		"airplane service spot found",
		zap.String("airplane_id", req.ID),
		zap.String("vehicle_type", string(req.Type)),
		zap.String("node_id", nodeID),
	)

	return &api.AirplaneIDServiceTypeGetOK{
		NodeId: nodeID,
	}, nil
}
