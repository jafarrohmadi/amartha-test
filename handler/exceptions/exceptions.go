package exceptions

import (
	"context"
	"errors"
	"strings"

	"github.com/amartha-test/model/response"
	"github.com/go-playground/validator/v10"
)

var IDAlreadyExistsError = errors.New("ID already exists")
var NotFoundError = errors.New("not found error")
var BadRequestError = errors.New("bad request")

func HandleError(ctx context.Context, err error) *response.ApiResponse {
	// Handle validation errors
	if _, ok := err.(validator.ValidationErrors); ok {
		return response.BuildErrorResponse(response.ValidationError, nil)
	}

	// Handle ID already exists error
	if strings.Contains(err.Error(), IDAlreadyExistsError.Error()) {
		message := err.Error()
		return response.BuildErrorResponse(response.ValidationIDAlreadyExistsError, &message)
	}

	// Handle generic not found error
	if strings.Contains(err.Error(), NotFoundError.Error()) {
		message := err.Error()
		return response.BuildErrorResponse(response.GenericResourceNotFound, &message)
	}

	// Handle bad request errors
	if strings.Contains(err.Error(), BadRequestError.Error()) {
		message := err.Error()
		return response.BuildErrorResponse(response.ValidationError, &message)
	}

	// Handle other errors as internal server errors
	return response.BuildErrorResponse(response.GenericServerError, nil)
}
