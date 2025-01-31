package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/amartha-test/generated"
	"github.com/amartha-test/model/aggregate"
	"github.com/amartha-test/model/enum"
)

// CreateLoan processes loan creation
func (s *LoanUseCaseImpl) CreateLoan(ctx context.Context, now time.Time, request generated.LoanRequest) (*aggregate.Loan, error) {
	loan := aggregate.BuildLoan(now, request)
	fmt.Println(loan)
	err := s.Repository.CreateLoan(ctx, loan)
	if err != nil {
		return nil, err
	}

	disbursement := aggregate.BuildLoanTransaction(loan.ID, request.Amount, "DISBURSEMENT", now)
	err = s.Repository.CreatePayment(ctx, disbursement)
	if err != nil {
		return nil, err
	}

	for i := 1; i <= request.NumberOfInstallments; i++ {
		loanSchedule := aggregate.BuildLoanSchedule(loan.ID, i, (loan.DuePrincipal+loan.DueInterest)/float32(request.NumberOfInstallments), now)
		err = s.Repository.CreateLoanSchedule(ctx, loanSchedule)
		if err != nil {
			return nil, err
		}
	}

	return &loan, nil
}

// GetLoanID returns the current outstanding balance of a loan.
func (s *LoanUseCaseImpl) GetLoanID(ctx context.Context, loanID string) (*aggregate.Loan, error) {
	loan, err := s.Repository.GetLoanByID(ctx, loanID)
	if err != nil {
		return nil, err
	}

	loan.GetOutstanding()
	return &loan, nil
}

// GetDelinquentLoans fetches loans that have missed 2 or more payments.
func (s *LoanUseCaseImpl) GetDelinquentLoans(ctx context.Context) ([]aggregate.Loan, error) {
	// Fetch all loans
	loans, err := s.Repository.GetLoanActive(ctx)
	if err != nil {
		return nil, err
	}

	// Initialize the slice to hold delinquent loans
	var delinquentLoans []aggregate.Loan

	// Check each loan's schedule
	for _, loan := range loans {
		// Fetch the loan schedules that are pending or overdue
		schedules, err := s.Repository.GetLoanSchedulesNotPaid(ctx, loan.ID.String())
		if err != nil {
			return nil, err
		}

		for _, schedule := range schedules {
			if time.Since(schedule.DueDate).Hours() > 14*24 {
				delinquentLoans = append(delinquentLoans, loan)
				break
			}
		}
	}

	return delinquentLoans, nil
}

// MakePayment processes a loan payment and updates the outstanding balance.
func (s *LoanUseCaseImpl) MakePayment(ctx context.Context, loanID string, amount float32, now time.Time) error {
	// Fetch the loan by its ID
	loan, err := s.Repository.GetLoanByID(ctx, loanID)
	if err != nil {
		return err
	}

	// Validate the payment amount
	if amount <= 0 {
		return errors.New("invalid payment amount")
	}

	// Apply the interest payment
	remainingAmount := loan.ApplyInterestPayment(amount)

	// Apply the principal payment
	remainingAmount = loan.ApplyPrincipalPayment(remainingAmount)

	// Build the loan transaction (record the payment)
	repayment := aggregate.BuildLoanTransaction(loan.ID, amount, "REPAYMENT", now)

	// Create the payment transaction record
	err = s.Repository.CreatePayment(ctx, repayment)
	if err != nil {
		return err
	}

	// Check if the loan is fully paid off and close the loan if so
	if loan.IsFullyPaid() {
		loan.Status = enum.LoanStatusClosed
	}

	// Update the loan in the repository
	err = s.Repository.UpdateLoan(ctx, loan)
	if err != nil {
		return err
	}

	// Fetch and update loan schedules
	schedules, err := s.Repository.GetLoanSchedulesNotPaid(ctx, loanID)
	if err != nil {
		return err
	}

	remainingAmountForSchedule := amount
	for _, schedule := range schedules {
		remainingAmountForSchedule = schedule.UpdateStatus(remainingAmountForSchedule)

		// Update the schedule in the repository
		err := s.Repository.UpdateLoanSchedule(ctx, schedule)
		if err != nil {
			return err
		}

		// Break if all payment has been allocated
		if remainingAmountForSchedule == 0 {
			break
		}
	}

	return nil
}
