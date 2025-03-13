package middlewares

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/validate"
	"go.uber.org/zap"

	"github.com/reaport/ground-control/pkg/logger"
)

func ErrorHandler(ctx context.Context, w http.ResponseWriter, r *http.Request, err error) {
	var (
		httpStatusCode  int
		response        *ErrorResponse
		validationErr   *validate.Error
		decodeErr       *ogenerrors.DecodeRequestError
		decodeParamsErr *ogenerrors.DecodeParamsError
		securityErr     *ogenerrors.SecurityError
	)

	switch {
	case errors.As(err, &validationErr):
		response, httpStatusCode = getValidationErrorResponse(validationErr)
		logger.GlobalLogger.Debug("validation error", zap.String("error", err.Error()), zap.Any("response", response))
	case errors.As(err, &decodeErr):
		response, httpStatusCode = getDecodeRequestError()
		logger.GlobalLogger.Error("decode request error", zap.String("error", err.Error()), zap.Any("response", response))
	case errors.As(err, &decodeParamsErr):
		response, httpStatusCode = getDecodeParamsError(decodeParamsErr)
		logger.GlobalLogger.Error("decode params error", zap.String("error", err.Error()), zap.Any("response", response))
	case errors.As(err, &securityErr):
		response, httpStatusCode = getSecurityError()
		logger.GlobalLogger.Error("security error", zap.String("error", err.Error()), zap.Any("response", response))
	default:
		logInternalError(ctx, r, err)
		httpStatusCode = http.StatusInternalServerError
	}

	if response != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(httpStatusCode)
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			logger.GlobalLogger.Warn("failed to encode error response", zap.String("error", err.Error()))
		}
		return
	}

	w.WriteHeader(httpStatusCode)
}

func logInternalError(_ context.Context, r *http.Request, err error) {
	logger.GlobalLogger.Error(
		"internal error",
		zap.String("error", err.Error()),
		zap.String("requestURI", r.RequestURI),
		zap.String("method", r.Method),
	)
}
