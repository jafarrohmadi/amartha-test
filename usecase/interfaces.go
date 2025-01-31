package usecase

import (
	"context"
	"time"

	"github.com/amartha-test/generated"
	"github.com/amartha-test/model/aggregate"
)

type UseCaseInterface interface {
	CreateLoan(ctx context.Context, now time.Time, request generated.LoanRequest) (*aggregate.Loan, error)

	GetLoanID(ctx context.Context, loanID string) (*aggregate.Loan, error)

	// IsDelinquent checks if the loan has more than 2 weeks of non-payment.
	GetDelinquentLoans(ctx context.Context) ([]aggregate.Loan, error)

	// MakePayment processes a loan payment and updates the outstanding balance.
	MakePayment(ctx context.Context, loanID string, amount float32, paymentDate time.Time) error
}
