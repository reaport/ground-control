// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// AirplaneGetParkingSpot implements airplane_getParkingSpot operation.
	//
	// В зависимости от загрузки парковок отдает нужный узел.
	//
	// GET /airplane/{id}/parking
	AirplaneGetParkingSpot(ctx context.Context, params AirplaneGetParkingSpotParams) (AirplaneGetParkingSpotRes, error)
	// AirplaneTakeOff implements airplane_takeOff operation.
	//
	// Удаляется самолет с карты.
	//
	// POST /airplane/{id}/take-off
	AirplaneTakeOff(ctx context.Context, params AirplaneTakeOffParams) (AirplaneTakeOffRes, error)
	// MapGetAirportMap implements map_getAirportMap operation.
	//
	// Возвращает полную карту аэропорта в виде графа.
	//
	// GET /map
	MapGetAirportMap(ctx context.Context) (*AirportMap, error)
	// MapGetAirportMapConfig implements map_getAirportMapConfig operation.
	//
	// Возвращает конфигурацию аэропорта.
	//
	// GET /map/config
	MapGetAirportMapConfig(ctx context.Context) (*AirportMapConfig, error)
	// MapRefreshAirportMap implements map_refreshAirportMap operation.
	//
	// Возвращает карту к исходному состоянию.
	//
	// POST /map/refresh
	MapRefreshAirportMap(ctx context.Context) error
	// MapUpdateAirportMap implements map_updateAirportMap operation.
	//
	// Обновляет карту аэропорта.
	//
	// PUT /map
	MapUpdateAirportMap(ctx context.Context, req *AirportMap) (MapUpdateAirportMapRes, error)
	// MovingGetRoute implements moving_getRoute operation.
	//
	// Запрашивает маршрут из точки А в точку Б.
	//
	// POST /route
	MovingGetRoute(ctx context.Context, req *MovingGetRouteReq) (MovingGetRouteRes, error)
	// MovingNotifyArrival implements moving_notifyArrival operation.
	//
	// Уведомляет вышку о прибытии транспорта в узел.
	//
	// POST /arrived
	MovingNotifyArrival(ctx context.Context, req *MovingNotifyArrivalReq) (MovingNotifyArrivalRes, error)
	// MovingRegisterVehicle implements moving_registerVehicle operation.
	//
	// В зависимости от типа транспорта отдает нужную
	// начальную точку и id.
	//
	// POST /register-vehicle/{type}
	MovingRegisterVehicle(ctx context.Context, params MovingRegisterVehicleParams) (MovingRegisterVehicleRes, error)
	// MovingRequestMove implements moving_requestMove operation.
	//
	// Запрашивает разрешение на перемещение из одного узла
	// в другой.
	//
	// POST /move
	MovingRequestMove(ctx context.Context, req *MovingRequestMoveReq) (MovingRequestMoveRes, error)
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h Handler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		baseServer: s,
	}, nil
}
