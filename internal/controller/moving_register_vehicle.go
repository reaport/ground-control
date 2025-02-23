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

// MovingRegisterVehicle implements moving_registerVehicle operation.
//
// В зависимости от типа транспорта отдает нужную
// начальную точку и id.
//
// POST /register-vehicle/{type}
func (c *Controller) MovingRegisterVehicle(
	ctx context.Context,
	params api.MovingRegisterVehicleParams,
) (api.MovingRegisterVehicleRes, error) {
	vehicleType, err := convert.VehicleTypeFromAPI(params.Type)
	if err != nil {
		logger.GlobalLogger.Error(
			"failed to convert vehicle type from API",
			zap.Error(fmt.Errorf("VehicleTypeFromAPI: %w", err)),
		)
		return &api.MovingRegisterVehicleBadRequest{}, nil
	}

	nodeID, vehicleID, err := c.mapService.RegisterVehicle(ctx, vehicleType)
	if err != nil {
		err = fmt.Errorf("c.mapService.RegisterVehicle: %w", err)
		if errors.Is(err, entity.ErrAirstripIsFull) {
			logger.GlobalLogger.Error(
				"airstrip is full",
				zap.Error(err),
			)
			return &api.MovingRegisterVehicleConflict{}, nil
		}

		logger.GlobalLogger.Error(
			"failed to register vehicle",
			zap.Error(err),
		)
		return nil, err
	}

	return &api.MovingRegisterVehicleOK{
		NodeId:    nodeID,
		VehicleId: vehicleID,
	}, nil
}
