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

// AirplaneGetParkingSpot implements airplane_getParkingSpot operation.
//
// В зависимости от загрузки парковок отдает нужный узел.
//
// GET /airplane/{id}/parking
func (c *Controller) AirplaneGetParkingSpot(
	ctx context.Context,
	req api.AirplaneGetParkingSpotParams,
) (api.AirplaneGetParkingSpotRes, error) {
	nodeID, err := c.mapService.GetAirplaneParkingSpot(ctx, req.ID)
	if err != nil {
		err = fmt.Errorf("c.mapService.GetAirplaneParkingSpot: %w", err)
		if errors.Is(err, entity.ErrAirplaneParkingSpotIsFull) {
			logger.GlobalLogger.Error(
				"airplane parking spot is full",
				zap.String("error", err.Error()),
				zap.String("airplane_id", req.ID),
			)
			return &api.AirplaneGetParkingSpotConflict{}, nil
		}

		logger.GlobalLogger.Error(
			"failed to get airplane parking spot",
			zap.String("error", err.Error()),
			zap.String("airplane_id", req.ID),
		)
		return nil, err
	}

	logger.GlobalLogger.Info(
		"airplane get parking spot",
		zap.String("airplane_id", req.ID),
		zap.String("node_id", nodeID),
	)

	err = c.eventSender.SendEvent(ctx, &entity.Event{
		Type: entity.AirplaneGetParkingSpotEventType,
		Data: entity.EventData{
			"airplane_id": req.ID,
			"node_id":     nodeID,
		},
	})
	if err != nil {
		logger.GlobalLogger.Error(
			"failed to send event",
			zap.Error(fmt.Errorf("c.eventSender.SendEvent: %w", err)),
			zap.String("event_type", string(entity.AirplaneGetParkingSpotEventType)),
			zap.String("airplane_id", req.ID),
			zap.String("node_id", nodeID),
		)
	}

	return &api.AirplaneGetParkingSpotOK{
		NodeId: nodeID,
	}, nil
}
