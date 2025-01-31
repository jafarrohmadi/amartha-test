package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/amartha-test/model/aggregate"
	"github.com/amartha-test/model/enum"
)

// CreateLoan stores a new loan in the database
func (l *LoanRepository) CreateLoan(ctx context.Context, loan aggregate.Loan) error {
	// Create a new loan record in the database
	if err := l.db.Create(&loan).Error; err != nil {
		return fmt.Errorf("failed to create loan: %v", err)
	}
	return nil
}

// CreateLoanSchedule stores a new loan schedule in the database
func (l *LoanRepository) CreateLoanSchedule(ctx context.Context, loanSchedule aggregate.LoanSchedule) error {
	// Create a new loan schedule record
	if err := l.db.Create(&loanSchedule).Error; err != nil {
		return fmt.Errorf("failed to create loan schedule: %v", err)
	}
	return nil
}

func (l *LoanRepository) GetLoanActive(ctx context.Context) ([]aggregate.Loan, error) {
	var loans []aggregate.Loan

	err := l.db.Where("status = ?", enum.LoanStatusActive).Find(&loans).Error
	if err != nil {
		return nil, err
	}

	return loans, nil

}

// GetLoanByID retrieves a loan by its ID
func (l *LoanRepository) GetLoanByID(ctx context.Context, loanID string) (aggregate.Loan, error) {
	var loan aggregate.Loan
	if err := l.db.Where("id = ?", loanID).First(&loan).Error; err != nil {
		return loan, fmt.Errorf("failed to get loan by ID: %v", err)
	}
	return loan, nil
}

// GetLastPayment retrieves the last payment for a given loan
func (l *LoanRepository) GetLastPayment(ctx context.Context, loanID string) (aggregate.LoanTransaction, error) {
	var transaction aggregate.LoanTransaction
	if err := l.db.Where("loan_id = ?", loanID).Order("payment_date desc").First(&transaction).Error; err != nil {
		return transaction, fmt.Errorf("failed to get last payment for loan ID %s: %v", loanID, err)
	}
	return transaction, nil
}

// CreatePayment records a new loan repayment transaction
func (l *LoanRepository) CreatePayment(ctx context.Context, transaction aggregate.LoanTransaction) error {
	// Create a new payment transaction
	if err := l.db.Create(&transaction).Error; err != nil {
		return fmt.Errorf("failed to create loan payment: %v", err)
	}
	return nil
}

// UpdateLoan updates an existing loan record in the database
func (l *LoanRepository) UpdateLoan(ctx context.Context, loan aggregate.Loan) error {
	// Update the loan record in the database
	if err := l.db.Save(&loan).Error; err != nil {
		return fmt.Errorf("failed to update loan: %v", err)
	}
	return nil
}

// UpdateLoanSchedule updates a loan schedule in the repository.
func (l *LoanRepository) UpdateLoanSchedule(ctx context.Context, schedule aggregate.LoanSchedule) error {
	// Assuming you are using GORM, use it to update the schedule in the database.
	err := l.db.Model(&aggregate.LoanSchedule{}).
		Where("id = ?", schedule.ID).
		Updates(map[string]interface{}{
			"status":     schedule.Status,
			"amount":     schedule.Amount,
			"updated_at": time.Now(), // Optional: You can update the timestamp if necessary
		}).Error

	if err != nil {
		return fmt.Errorf("could not update loan schedule: %v", err)
	}

	return nil
}

// GetLoanSchedulesNotPaid retrieves loan schedules for a specific loan that are not yet paid.
func (l *LoanRepository) GetLoanSchedulesNotPaid(ctx context.Context, loanID string) ([]aggregate.LoanSchedule, error) {
	var schedules []aggregate.LoanSchedule

	// Assuming you're using GORM, query for the loan schedules that are not paid
	err := l.db.Where("loan_id = ? AND status != ?", loanID, enum.LoanScheduleStatusPaid).
		Find(&schedules).Error

	if err != nil {
		return nil, fmt.Errorf("could not retrieve loan schedules: %v", err)
	}

	return schedules, nil
}
