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
			zap.String("type", string(params.Type)),
		)
		return &api.MovingRegisterVehicleBadRequest{}, nil
	}

	vehicleInitInfo, err := c.mapService.RegisterVehicle(ctx, vehicleType)
	if err != nil {
		err = fmt.Errorf("c.mapService.RegisterVehicle: %w", err)
		if errors.Is(err, entity.ErrAirstripIsFull) {
			logger.GlobalLogger.Error(
				"airstrip is full",
				zap.Error(err),
				zap.String("type", string(params.Type)),
			)
			return &api.MovingRegisterVehicleConflict{}, nil
		}

		logger.GlobalLogger.Error(
			"failed to register vehicle",
			zap.Error(err),
			zap.String("type", string(params.Type)),
		)
		return nil, err
	}

	logger.GlobalLogger.Info(
		"vehicle registered",
		zap.String("garrage_node_id", vehicleInitInfo.GarrageNodeID),
		zap.String("vehicle_id", vehicleInitInfo.VehicleID),
		zap.String("type", string(params.Type)),
		zap.Any("service_spots", vehicleInitInfo.ServiceSpots),
	)

	return convert.VehicleInitInfoToAPI(vehicleInitInfo), nil
}
