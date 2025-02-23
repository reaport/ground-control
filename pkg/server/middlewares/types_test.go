package middlewares

import (
	"errors"
	"testing"

	"github.com/ogen-go/ogen/validate"
	"github.com/stretchr/testify/require"
)

func TestToValidationErrorResponse(t *testing.T) {
	tests := []struct {
		name     string
		inputErr *validate.Error
		expected *ErrorResponse
	}{
		{
			name: "No Regex Match",
			inputErr: &validate.Error{
				Fields: []validate.FieldError{
					{
						Name:  "username",
						Error: &validate.NoRegexMatchError{},
					},
				},
			},
			expected: &ErrorResponse{
				Code: ErrorCodeValidationError,
				Details: errorDetails{
					"username": string(ValidationErrorCodeNoRegexMatch),
				},
			},
		},
		{
			name: "Min Length",
			inputErr: &validate.Error{
				Fields: []validate.FieldError{
					{
						Name:  "password",
						Error: &validate.MinLengthError{},
					},
				},
			},
			expected: &ErrorResponse{
				Code: ErrorCodeValidationError,
				Details: errorDetails{
					"password": string(ValidationErrorCodeMinLength),
				},
			},
		},
		{
			name: "Max Length",
			inputErr: &validate.Error{
				Fields: []validate.FieldError{
					{
						Name:  "password",
						Error: &validate.MaxLengthError{},
					},
				},
			},
			expected: &ErrorResponse{
				Code: ErrorCodeValidationError,
				Details: errorDetails{
					"password": string(ValidationErrorCodeMaxLength),
				},
			},
		},
		{
			name: "Field Required",
			inputErr: &validate.Error{
				Fields: []validate.FieldError{
					{
						Name:  "email",
						Error: validate.ErrFieldRequired,
					},
				},
			},
			expected: &ErrorResponse{
				Code: ErrorCodeValidationError,
				Details: errorDetails{
					"email": string(ValidationErrorCodeFieldRequired),
				},
			},
		},
		{
			name: "Duplicate Element",
			inputErr: &validate.Error{
				Fields: []validate.FieldError{
					{
						Name:  "tags",
						Error: errors.New("duplicate element"),
					},
				},
			},
			expected: &ErrorResponse{
				Code: ErrorCodeValidationError,
				Details: errorDetails{
					"tags": string(ValidationErrorCodeDuplicateElement),
				},
			},
		},
		{
			name: "Not a Number",
			inputErr: &validate.Error{
				Fields: []validate.FieldError{
					{
						Name:  "age",
						Error: errors.New("is not a number"),
					},
				},
			},
			expected: &ErrorResponse{
				Code: ErrorCodeValidationError,
				Details: errorDetails{
					"age": string(ValidationErrorCodeNotANumber),
				},
			},
		},
		{
			name: "Min Properties",
			inputErr: &validate.Error{
				Fields: []validate.FieldError{
					{
						Name:  "profile",
						Error: errors.New("object properties less than"),
					},
				},
			},
			expected: &ErrorResponse{
				Code: ErrorCodeValidationError,
				Details: errorDetails{
					"profile": string(ValidationErrorCodeMinProperties),
				},
			},
		},
		{
			name: "Max Properties",
			inputErr: &validate.Error{
				Fields: []validate.FieldError{
					{
						Name:  "profile",
						Error: errors.New("object properties greater than"),
					},
				},
			},
			expected: &ErrorResponse{
				Code: ErrorCodeValidationError,
				Details: errorDetails{
					"profile": string(ValidationErrorCodeMaxProperties),
				},
			},
		},
		{
			name: "Not Multiple",
			inputErr: &validate.Error{
				Fields: []validate.FieldError{
					{
						Name:  "count",
						Error: errors.New("not multiple"),
					},
				},
			},
			expected: &ErrorResponse{
				Code: ErrorCodeValidationError,
				Details: errorDetails{
					"count": string(ValidationErrorCodeNotMultiple),
				},
			},
		},
		{
			name: "Space Character",
			inputErr: &validate.Error{
				Fields: []validate.FieldError{
					{
						Name:  "username",
						Error: errors.New("space character"),
					},
				},
			},
			expected: &ErrorResponse{
				Code: ErrorCodeValidationError,
				Details: errorDetails{
					"username": string(ValidationErrorCodeSpaceCharacter),
				},
			},
		},
		{
			name: "Not Printable Character",
			inputErr: &validate.Error{
				Fields: []validate.FieldError{
					{
						Name:  "password",
						Error: errors.New("not printable character"),
					},
				},
			},
			expected: &ErrorResponse{
				Code: ErrorCodeValidationError,
				Details: errorDetails{
					"password": string(ValidationErrorCodeNotPrintableCharacter),
				},
			},
		},
		{
			name: "Invalid Character",
			inputErr: &validate.Error{
				Fields: []validate.FieldError{
					{
						Name:  "username",
						Error: errors.New("invalid character"),
					},
				},
			},
			expected: &ErrorResponse{
				Code: ErrorCodeValidationError,
				Details: errorDetails{
					"username": string(ValidationErrorCodeInvalidCharacter),
				},
			},
		},
		{
			name: "Invalid Email Format",
			inputErr: &validate.Error{
				Fields: []validate.FieldError{
					{
						Name:  "email",
						Error: errors.New("got @ multiple times"),
					},
				},
			},
			expected: &ErrorResponse{
				Code: ErrorCodeValidationError,
				Details: errorDetails{
					"email": string(ValidationErrorCodeInvalidEmailFormat),
				},
			},
		},
		{
			name: "Blank",
			inputErr: &validate.Error{
				Fields: []validate.FieldError{
					{
						Name:  "description",
						Error: errors.New("blank"),
					},
				},
			},
			expected: &ErrorResponse{
				Code: ErrorCodeValidationError,
				Details: errorDetails{
					"description": string(ValidationErrorCodeBlank),
				},
			},
		},
		{
			name: "Too Long",
			inputErr: &validate.Error{
				Fields: []validate.FieldError{
					{
						Name:  "bio",
						Error: errors.New("too long"),
					},
				},
			},
			expected: &ErrorResponse{
				Code: ErrorCodeValidationError,
				Details: errorDetails{
					"bio": string(ValidationErrorCodeTooLong),
				},
			},
		},
		{
			name: "Unexpected",
			inputErr: &validate.Error{
				Fields: []validate.FieldError{
					{
						Name:  "unknown",
						Error: errors.New("unexpected error"),
					},
				},
			},
			expected: &ErrorResponse{
				Code: ErrorCodeValidationError,
				Details: errorDetails{
					"unknown": string(ValidationErrorCodeUnexpected),
				},
			},
		},
		{
			name: "Nested Errors",
			inputErr: &validate.Error{
				Fields: []validate.FieldError{
					{
						Name: "address",
						Error: &validate.Error{
							Fields: []validate.FieldError{
								{
									Name:  "street",
									Error: &validate.MinLengthError{},
								},
								{
									Name: "city",
									Error: &validate.Error{
										Fields: []validate.FieldError{
											{
												Name:  "name",
												Error: errors.New("invalid character"),
											},
										},
									},
								},
							},
						},
					},
				},
			},
			expected: &ErrorResponse{
				Code: ErrorCodeValidationError,
				Details: errorDetails{
					"address": errorDetails{
						"street": string(ValidationErrorCodeMinLength),
						"city": errorDetails{
							"name": string(ValidationErrorCodeInvalidCharacter),
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToValidationErrorResponse(tt.inputErr)
			require.Equal(t, tt.expected, result)
		})
	}
}
