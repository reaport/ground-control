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
		response        interface{}
		validationErr   *validate.Error
		decodeErr       *ogenerrors.DecodeRequestError
		decodeParamsErr *ogenerrors.DecodeParamsError
		securityErr     *ogenerrors.SecurityError
	)

	switch {
	case errors.As(err, &validationErr):
		response, httpStatusCode = getValidationErrorResponse(validationErr)
	case errors.As(err, &decodeErr):
		response, httpStatusCode = getDecodeRequestError()
	case errors.As(err, &decodeParamsErr):
		response, httpStatusCode = getDecodeParamsError(decodeParamsErr)
	case errors.As(err, &securityErr):
		response, httpStatusCode = getSecurityError()
	default:
		logInternalError(ctx, r, err)
		httpStatusCode = http.StatusInternalServerError
	}

	if response != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(httpStatusCode)
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			logger.GlobalLogger.Warn("failed to encode error response", zap.Error(err))
		}
		return
	}

	w.WriteHeader(httpStatusCode)
}

func logInternalError(_ context.Context, r *http.Request, err error) {
	logger.GlobalLogger.Error(
		"internal error",
		zap.Error(err),
		zap.String("requestURI", r.RequestURI),
		zap.String("method", r.Method),
	)
}
