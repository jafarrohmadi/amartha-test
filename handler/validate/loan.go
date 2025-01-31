package validate

import (
	"fmt"

	"github.com/amartha-test/generated"
	"github.com/go-playground/validator/v10"
)

// ValidateLoanRequest validates the fields of LoanRequest
func ValidateLoanRequest(request generated.LoanRequest) error {
	validate := validator.New()

	// Validation rules
	rules := map[string]string{
		"Amount":               "required",
		"NumberOfInstallments": "required,min=1",
		"InterestRate":         "required,min=0,max=100",
		"UserId":               "required",
	}

	// Validate each field
	for field, rule := range rules {
		value := map[string]interface{}{"Amount": request.Amount, "NumberOfInstallments": request.NumberOfInstallments, "InterestRate": request.InterestRate, "UserId": request.UserId}[field]
		if err := validate.Var(value, rule); err != nil {
			return fmt.Errorf("Field '%s' failed validation: %s", field, err.Error())
		}
	}

	return nil
}
