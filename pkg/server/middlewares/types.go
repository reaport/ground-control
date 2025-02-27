package middlewares

import (
	"errors"
	"net/http"
	"strings"

	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/validate"
)

type ErrorResponse struct {
	Code    ErrorCode `json:"code"`
	Details any       `json:"details,omitempty"`
}

type ErrorCode string

const (
	ErrorCodeValidationError    ErrorCode = "VALIDATION_ERROR"
	ErrorCodeDecodeRequestError ErrorCode = "DECODE_REQUEST_ERROR"
	ErrorCodeDecodeParamsError  ErrorCode = "DECODE_PARAMS_ERROR"
	ErrorCodeSecurityError      ErrorCode = "SECURITY_ERROR"
)

func ToValidationErrorResponse(err *validate.Error) *ErrorResponse {
	rawDetails := processValidationErrors(err)
	details := convertValidationErrorDetailToSimplified(rawDetails)

	return &ErrorResponse{
		Code:    ErrorCodeValidationError,
		Details: details,
	}
}

func getValidationErrorResponse(validationErr *validate.Error) (*ErrorResponse, int) {
	return ToValidationErrorResponse(validationErr), http.StatusUnprocessableEntity
}

func getDecodeRequestError() (*ErrorResponse, int) {
	return &ErrorResponse{Code: ErrorCodeDecodeRequestError}, http.StatusBadRequest
}

func getDecodeParamsError(dpsErr *ogenerrors.DecodeParamsError) (*ErrorResponse, int) {
	var dpErr *ogenerrors.DecodeParamError
	var details errorDetails

	if errors.As(dpsErr, &dpErr) {
		details = make(errorDetails)
		details[dpErr.Name] = dpErr.Err.Error()
	}
	return &ErrorResponse{Code: ErrorCodeDecodeParamsError, Details: details}, dpsErr.Code()
}

type ValidationErrorCode string

const (
	ValidationErrorCodeUnexpected            ValidationErrorCode = "UNEXPECTED"
	ValidationErrorCodeNoRegexMatch          ValidationErrorCode = "NO_REGEX_MATCH"
	ValidationErrorCodeMinLength             ValidationErrorCode = "MIN_LENGTH"
	ValidationErrorCodeMaxLength             ValidationErrorCode = "MAX_LENGTH"
	ValidationErrorCodeMinProperties         ValidationErrorCode = "MIN_PROPERTIES"
	ValidationErrorCodeMaxProperties         ValidationErrorCode = "MAX_PROPERTIES"
	ValidationErrorCodeDuplicateElement      ValidationErrorCode = "DUPLICATE_ELEMENT"
	ValidationErrorCodeFieldRequired         ValidationErrorCode = "FIELD_REQUIRED"
	ValidationErrorCodeNotANumber            ValidationErrorCode = "NOT_A_NUMBER"
	ValidationErrorCodeInfiniteNumber        ValidationErrorCode = "INFINITE_NUMBER"
	ValidationErrorCodeNotMultiple           ValidationErrorCode = "NOT_MULTIPLE"
	ValidationErrorCodeSpaceCharacter        ValidationErrorCode = "SPACE_CHARACTER"
	ValidationErrorCodeNotPrintableCharacter ValidationErrorCode = "NOT_PRINTABLE_CHARACTER"
	ValidationErrorCodeInvalidCharacter      ValidationErrorCode = "INVALID_CHARACTER"
	ValidationErrorCodeInvalidEmailFormat    ValidationErrorCode = "INVALID_EMAIL_FORMAT"
	ValidationErrorCodeBlank                 ValidationErrorCode = "BLANK"
	ValidationErrorCodeTooLong               ValidationErrorCode = "TOO_LONG"
)

type validationErrorDetail map[string]validationErrorDetailItem

type validationErrorDetailItem struct {
	Type                  validationErrorDetailItemType `json:"omit"`
	String                string
	ValidationErrorDetail validationErrorDetail
}

type (
	validationErrorDetailItemType string
	errorDetails                  map[string]interface{}
)

const (
	stringValidationErrorDetailItem                validationErrorDetailItemType = "string"
	validationErrorDetailValidationErrorDetailItem validationErrorDetailItemType = "ValidationErrorDetail"
)

func (item *validationErrorDetailItem) setString(v string) {
	item.Type = stringValidationErrorDetailItem
	item.String = v
}

func (item *validationErrorDetailItem) setValidationErrorDetail(v validationErrorDetail) {
	item.Type = validationErrorDetailValidationErrorDetailItem
	item.ValidationErrorDetail = v
}

func newValidationErrorDetailItem(v interface{}) validationErrorDetailItem {
	var item validationErrorDetailItem
	switch v := v.(type) {
	case ValidationErrorCode:
		item.setString(string(v))
	case validationErrorDetail:
		item.setValidationErrorDetail(v)
	}
	return item
}

func processValidationErrors(err *validate.Error) validationErrorDetail {
	details := make(validationErrorDetail)

	for _, field := range err.Fields {
		var nestedErr *validate.Error
		if errors.As(field.Error, &nestedErr) {
			details[field.Name] = newValidationErrorDetailItem(processValidationErrors(nestedErr))
		} else {
			validationCode := getValidationCode(field.Error)
			details[field.Name] = newValidationErrorDetailItem(validationCode)
		}
	}

	return details
}

func getValidationCode(err error) ValidationErrorCode {
	switch {
	case errors.Is(err, &validate.NoRegexMatchError{}):
		return ValidationErrorCodeNoRegexMatch

	case errors.Is(err, &validate.MinLengthError{}):
		return ValidationErrorCodeMinLength

	case errors.Is(err, &validate.MaxLengthError{}):
		return ValidationErrorCodeMaxLength

	case errors.Is(err, validate.ErrFieldRequired):
		return ValidationErrorCodeFieldRequired

	case strings.Contains(err.Error(), "duplicate element"):
		return ValidationErrorCodeDuplicateElement

	case strings.Contains(err.Error(), "is not a number"):
		return ValidationErrorCodeNotANumber

	case strings.Contains(err.Error(), "object properties") && strings.Contains(err.Error(), "less than"):
		return ValidationErrorCodeMinProperties

	case strings.Contains(err.Error(), "object properties") && strings.Contains(err.Error(), "greater than"):
		return ValidationErrorCodeMaxProperties

	case strings.Contains(err.Error(), "less than"):
		return ValidationErrorCodeMinLength

	case strings.Contains(err.Error(), "greater than"):
		return ValidationErrorCodeMaxLength

	case strings.Contains(err.Error(), "not multiple"):
		return ValidationErrorCodeNotMultiple

	case strings.Contains(err.Error(), "blank"):
		return ValidationErrorCodeBlank

	case strings.Contains(err.Error(), "too long"):
		return ValidationErrorCodeTooLong

	case strings.Contains(err.Error(), "space character"):
		return ValidationErrorCodeSpaceCharacter

	case strings.Contains(err.Error(), "not printable character"):
		return ValidationErrorCodeNotPrintableCharacter

	case strings.Contains(err.Error(), "invalid character"):
		return ValidationErrorCodeInvalidCharacter

	case strings.Contains(err.Error(), "@"):
		return ValidationErrorCodeInvalidEmailFormat

	default:
		return ValidationErrorCodeUnexpected
	}
}

func convertValidationErrorDetailToSimplified(detail validationErrorDetail) errorDetails {
	simplified := make(errorDetails)

	for key, item := range detail {
		switch item.Type {
		case stringValidationErrorDetailItem:
			simplified[key] = item.String
		case validationErrorDetailValidationErrorDetailItem:
			simplified[key] = convertValidationErrorDetailToSimplified(item.ValidationErrorDetail)
		}
	}

	return simplified
}

func getSecurityError() (*ErrorResponse, int) {
	return &ErrorResponse{Code: ErrorCodeSecurityError}, http.StatusForbidden
}
