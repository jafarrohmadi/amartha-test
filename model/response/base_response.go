package response

import (
	"net/http"
)

type ApiResponse struct {
	StatusCode int     `json:"status_code,omitempty"`
	Message    *string `json:"message"`
}

type ApiMessage struct {
	Message *string `json:"message"`
}

type Code struct {
	HttpStatusCode int
	ErrorCode      string
	Message        string
}

var Ok Code = Code{http.StatusOK, "", "Request performed successfully"}
var Created Code = Code{http.StatusCreated, "", "Resource created successfully"}
var ValidationError Code = Code{http.StatusBadRequest, "VALIDATION_ERROR", "Request validation error"}
var GenericResourceNotFound Code = Code{http.StatusNotFound, "RESOURCE_NOT_FOUND", "Requested resource not found"}
var GenericServerError Code = Code{http.StatusInternalServerError, "SERVER_ERROR", "There is a problem processing your request. Error has been logged."}
var ValidationIDAlreadyExistsError Code = Code{http.StatusBadRequest, "VALIDATION_ERROR", "ID already exists"}
var InvalidCoordinateError Code = Code{http.StatusBadRequest, "INVALID_COORDINATE_ERROR", "Invalid coordinate error"}

func BuildErrorResponse(responseCode Code, message *string) *ApiResponse {
	responseMessage := responseCode.Message
	if message != nil {
		responseMessage = *message
	}
	return &ApiResponse{
		StatusCode: responseCode.HttpStatusCode,
		Message:    &responseMessage,
	}
}

func BuildErrorMessage(message *string) *ApiMessage {
	return &ApiMessage{
		Message: message,
	}
}
