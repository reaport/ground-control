// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
	"net/url"
	"strings"
	"time"

	"github.com/go-faster/errors"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"go.opentelemetry.io/otel/trace"

	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
)

func trimTrailingSlashes(u *url.URL) {
	u.Path = strings.TrimRight(u.Path, "/")
	u.RawPath = strings.TrimRight(u.RawPath, "/")
}

// Invoker invokes operations described by OpenAPI v3 specification.
type Invoker interface {
	// AirplaneGetParkingSpot invokes airplane_getParkingSpot operation.
	//
	// В зависимости от загрузки парковок отдает нужный узел.
	//
	// GET /airplane/{id}/parking
	AirplaneGetParkingSpot(ctx context.Context, params AirplaneGetParkingSpotParams) (AirplaneGetParkingSpotRes, error)
	// MapGetAirportMap invokes map_getAirportMap operation.
	//
	// Возвращает полную карту аэропорта в виде графа.
	//
	// GET /map
	MapGetAirportMap(ctx context.Context) (*AirportMap, error)
	// MapRefreshAirportMap invokes map_refreshAirportMap operation.
	//
	// Возвращает карту к исходному состоянию.
	//
	// POST /map/refresh
	MapRefreshAirportMap(ctx context.Context) error
	// MapUpdateAirportMap invokes map_updateAirportMap operation.
	//
	// Обновляет карту аэропорта.
	//
	// PUT /map
	MapUpdateAirportMap(ctx context.Context, request *AirportMap) (MapUpdateAirportMapRes, error)
	// MovingGetRoute invokes moving_getRoute operation.
	//
	// Запрашивает маршрут из точки А в точку Б.
	//
	// POST /route
	MovingGetRoute(ctx context.Context, request *MovingGetRouteReq) (MovingGetRouteRes, error)
	// MovingNotifyArrival invokes moving_notifyArrival operation.
	//
	// Уведомляет вышку о прибытии транспорта в узел.
	//
	// POST /arrived
	MovingNotifyArrival(ctx context.Context, request *MovingNotifyArrivalReq) (MovingNotifyArrivalRes, error)
	// MovingRegisterVehicle invokes moving_registerVehicle operation.
	//
	// В зависимости от типа транспорта отдает нужную
	// начальную точку и id.
	//
	// POST /register-vehicle/{type}
	MovingRegisterVehicle(ctx context.Context, params MovingRegisterVehicleParams) (MovingRegisterVehicleRes, error)
	// MovingRequestMove invokes moving_requestMove operation.
	//
	// Запрашивает разрешение на перемещение из одного узла
	// в другой.
	//
	// POST /move
	MovingRequestMove(ctx context.Context, request *MovingRequestMoveReq) (MovingRequestMoveRes, error)
}

// Client implements OAS client.
type Client struct {
	serverURL *url.URL
	baseClient
}

var _ Handler = struct {
	*Client
}{}

// NewClient initializes new Client defined by OAS.
func NewClient(serverURL string, opts ...ClientOption) (*Client, error) {
	u, err := url.Parse(serverURL)
	if err != nil {
		return nil, err
	}
	trimTrailingSlashes(u)

	c, err := newClientConfig(opts...).baseClient()
	if err != nil {
		return nil, err
	}
	return &Client{
		serverURL:  u,
		baseClient: c,
	}, nil
}

type serverURLKey struct{}

// WithServerURL sets context key to override server URL.
func WithServerURL(ctx context.Context, u *url.URL) context.Context {
	return context.WithValue(ctx, serverURLKey{}, u)
}

func (c *Client) requestURL(ctx context.Context) *url.URL {
	u, ok := ctx.Value(serverURLKey{}).(*url.URL)
	if !ok {
		return c.serverURL
	}
	return u
}

// AirplaneGetParkingSpot invokes airplane_getParkingSpot operation.
//
// В зависимости от загрузки парковок отдает нужный узел.
//
// GET /airplane/{id}/parking
func (c *Client) AirplaneGetParkingSpot(ctx context.Context, params AirplaneGetParkingSpotParams) (AirplaneGetParkingSpotRes, error) {
	res, err := c.sendAirplaneGetParkingSpot(ctx, params)
	return res, err
}

func (c *Client) sendAirplaneGetParkingSpot(ctx context.Context, params AirplaneGetParkingSpotParams) (res AirplaneGetParkingSpotRes, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("airplane_getParkingSpot"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/airplane/{id}/parking"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, AirplaneGetParkingSpotOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/airplane/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.StringToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/parking"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeAirplaneGetParkingSpotResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// MapGetAirportMap invokes map_getAirportMap operation.
//
// Возвращает полную карту аэропорта в виде графа.
//
// GET /map
func (c *Client) MapGetAirportMap(ctx context.Context) (*AirportMap, error) {
	res, err := c.sendMapGetAirportMap(ctx)
	return res, err
}

func (c *Client) sendMapGetAirportMap(ctx context.Context) (res *AirportMap, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("map_getAirportMap"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/map"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, MapGetAirportMapOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/map"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeMapGetAirportMapResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// MapRefreshAirportMap invokes map_refreshAirportMap operation.
//
// Возвращает карту к исходному состоянию.
//
// POST /map/refresh
func (c *Client) MapRefreshAirportMap(ctx context.Context) error {
	_, err := c.sendMapRefreshAirportMap(ctx)
	return err
}

func (c *Client) sendMapRefreshAirportMap(ctx context.Context) (res *MapRefreshAirportMapOK, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("map_refreshAirportMap"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/map/refresh"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, MapRefreshAirportMapOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/map/refresh"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeMapRefreshAirportMapResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// MapUpdateAirportMap invokes map_updateAirportMap operation.
//
// Обновляет карту аэропорта.
//
// PUT /map
func (c *Client) MapUpdateAirportMap(ctx context.Context, request *AirportMap) (MapUpdateAirportMapRes, error) {
	res, err := c.sendMapUpdateAirportMap(ctx, request)
	return res, err
}

func (c *Client) sendMapUpdateAirportMap(ctx context.Context, request *AirportMap) (res MapUpdateAirportMapRes, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("map_updateAirportMap"),
		semconv.HTTPRequestMethodKey.String("PUT"),
		semconv.HTTPRouteKey.String("/map"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, MapUpdateAirportMapOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/map"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "PUT", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeMapUpdateAirportMapRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeMapUpdateAirportMapResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// MovingGetRoute invokes moving_getRoute operation.
//
// Запрашивает маршрут из точки А в точку Б.
//
// POST /route
func (c *Client) MovingGetRoute(ctx context.Context, request *MovingGetRouteReq) (MovingGetRouteRes, error) {
	res, err := c.sendMovingGetRoute(ctx, request)
	return res, err
}

func (c *Client) sendMovingGetRoute(ctx context.Context, request *MovingGetRouteReq) (res MovingGetRouteRes, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("moving_getRoute"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/route"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, MovingGetRouteOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/route"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeMovingGetRouteRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeMovingGetRouteResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// MovingNotifyArrival invokes moving_notifyArrival operation.
//
// Уведомляет вышку о прибытии транспорта в узел.
//
// POST /arrived
func (c *Client) MovingNotifyArrival(ctx context.Context, request *MovingNotifyArrivalReq) (MovingNotifyArrivalRes, error) {
	res, err := c.sendMovingNotifyArrival(ctx, request)
	return res, err
}

func (c *Client) sendMovingNotifyArrival(ctx context.Context, request *MovingNotifyArrivalReq) (res MovingNotifyArrivalRes, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("moving_notifyArrival"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/arrived"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, MovingNotifyArrivalOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/arrived"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeMovingNotifyArrivalRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeMovingNotifyArrivalResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// MovingRegisterVehicle invokes moving_registerVehicle operation.
//
// В зависимости от типа транспорта отдает нужную
// начальную точку и id.
//
// POST /register-vehicle/{type}
func (c *Client) MovingRegisterVehicle(ctx context.Context, params MovingRegisterVehicleParams) (MovingRegisterVehicleRes, error) {
	res, err := c.sendMovingRegisterVehicle(ctx, params)
	return res, err
}

func (c *Client) sendMovingRegisterVehicle(ctx context.Context, params MovingRegisterVehicleParams) (res MovingRegisterVehicleRes, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("moving_registerVehicle"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/register-vehicle/{type}"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, MovingRegisterVehicleOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [2]string
	pathParts[0] = "/register-vehicle/"
	{
		// Encode "type" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "type",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.StringToString(string(params.Type)))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeMovingRegisterVehicleResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// MovingRequestMove invokes moving_requestMove operation.
//
// Запрашивает разрешение на перемещение из одного узла
// в другой.
//
// POST /move
func (c *Client) MovingRequestMove(ctx context.Context, request *MovingRequestMoveReq) (MovingRequestMoveRes, error) {
	res, err := c.sendMovingRequestMove(ctx, request)
	return res, err
}

func (c *Client) sendMovingRequestMove(ctx context.Context, request *MovingRequestMoveReq) (res MovingRequestMoveRes, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("moving_requestMove"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/move"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, MovingRequestMoveOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/move"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeMovingRequestMoveRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeMovingRequestMoveResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}
