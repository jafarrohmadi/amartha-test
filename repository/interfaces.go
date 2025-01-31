package repository

import (
	"context"

	"github.com/amartha-test/model/aggregate"
)

// RepositoryInterface defines methods for interacting with loan-related data
type RepositoryInterface interface {
	// CreateLoan creates a new loan record
	CreateLoan(ctx context.Context, loan aggregate.Loan) error

	// CreateLoanSchedule creates a new loan payment schedule
	CreateLoanSchedule(ctx context.Context, loanSchedule aggregate.LoanSchedule) error

	// GetLoanActive retrieves a loan
	GetLoanActive(ctx context.Context) ([]aggregate.Loan, error)

	// GetLoanByID retrieves a loan by its ID
	GetLoanByID(ctx context.Context, loanID string) (aggregate.Loan, error)

	// GetLastPayment retrieves the most recent payment for a loan
	GetLastPayment(ctx context.Context, loanID string) (aggregate.LoanTransaction, error)

	// CreatePayment creates a new payment transaction for a loan
	CreatePayment(ctx context.Context, transaction aggregate.LoanTransaction) error

	// UpdateLoan updates an existing loan record
	UpdateLoan(ctx context.Context, loan aggregate.Loan) error

	// UpdateLoanSchedule updates an existing loan record
	UpdateLoanSchedule(ctx context.Context, loan aggregate.LoanSchedule) error

	GetLoanSchedulesNotPaid(ctx context.Context, loanID string) ([]aggregate.LoanSchedule, error)
}
