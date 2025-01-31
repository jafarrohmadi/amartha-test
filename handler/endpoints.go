package handler

import (
	"net/http"
	"time"

	"github.com/amartha-test/generated"
	"github.com/amartha-test/handler/exceptions"
	"github.com/amartha-test/handler/validate"
	"github.com/amartha-test/model/response"
	"github.com/labstack/echo/v4"
)

// PostLoans creates a new loan
func (s *Server) PostLoans(ctx echo.Context) error {
	var (
		context      = ctx.Request().Context()
		requestData  generated.LoanRequest
		now          = time.Now()
		httpResponse *response.ApiResponse
		errorMessage *response.ApiMessage
	)

	if err := ctx.Bind(&requestData); err != nil {
		httpResponse = response.BuildErrorResponse(response.ValidationError, nil)
		errorMessage = response.BuildErrorMessage(httpResponse.Message)
		return ctx.JSON(httpResponse.StatusCode, errorMessage)
	}

	if err := validate.ValidateLoanRequest(requestData); err != nil {
		httpResponse = response.BuildErrorResponse(response.ValidationError, nil)
		errorMessage = response.BuildErrorMessage(httpResponse.Message)
		return ctx.JSON(httpResponse.StatusCode, errorMessage)
	}

	data, err := s.UseCase.CreateLoan(context, now, requestData)
	if err != nil {
		httpResponse = exceptions.HandleError(context, err)
		errorMessage = response.BuildErrorMessage(httpResponse.Message)
		return ctx.JSON(httpResponse.StatusCode, errorMessage)
	}

	return ctx.JSON(http.StatusCreated, data)
}

// GetLoansDelinquent retrieves delinquent loans
func (s *Server) GetLoansDelinquent(ctx echo.Context) error {
	context := ctx.Request().Context()

	data, err := s.UseCase.GetDelinquentLoans(context)
	if err != nil {
		httpResponse := exceptions.HandleError(context, err)
		errorMessage := response.BuildErrorMessage(httpResponse.Message)
		return ctx.JSON(httpResponse.StatusCode, errorMessage)
	}

	return ctx.JSON(http.StatusOK, data)
}

// GetLoansLoanId retrieves loan details
func (s *Server) GetLoansLoanId(ctx echo.Context, loanId string) error {
	context := ctx.Request().Context()

	data, err := s.UseCase.GetLoanID(context, loanId)
	if err != nil {
		httpResponse := exceptions.HandleError(context, err)
		errorMessage := response.BuildErrorMessage(httpResponse.Message)
		return ctx.JSON(httpResponse.StatusCode, errorMessage)
	}

	return ctx.JSON(http.StatusOK, data)
}

// PostLoansLoanIdRepayment handles loan repayment
func (s *Server) PostLoansLoanIdRepayment(ctx echo.Context, loanId string) error {
	var (
		context      = ctx.Request().Context()
		requestData  generated.PaymentRequest
		now          = time.Now()
		httpResponse *response.ApiResponse
		errorMessage *response.ApiMessage
	)

	if err := ctx.Bind(&requestData); err != nil {
		httpResponse = response.BuildErrorResponse(response.ValidationError, nil)
		errorMessage = response.BuildErrorMessage(httpResponse.Message)
		return ctx.JSON(httpResponse.StatusCode, errorMessage)
	}

	err := s.UseCase.MakePayment(context, loanId, requestData.Amount, now)
	if err != nil {
		httpResponse = exceptions.HandleError(context, err)
		errorMessage = response.BuildErrorMessage(httpResponse.Message)
		return ctx.JSON(httpResponse.StatusCode, errorMessage)
	}

	return ctx.JSON(http.StatusOK, nil)
}
