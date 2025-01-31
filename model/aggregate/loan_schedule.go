package aggregate

import (
	"time"

	"github.com/amartha-test/model/enum"
	"github.com/google/uuid"
)

// LoanSchedule represents the schedule for loan repayments
type LoanSchedule struct {
	ID        uuid.UUID               `json:"id" gorm:"primaryKey"`
	LoanID    uuid.UUID               `json:"loan_id" gorm:"index"`
	StartDate time.Time               `json:"start_date"`
	DueDate   time.Time               `json:"due_date"`
	Amount    float32                 `json:"amount"`
	Status    enum.LoanScheduleStatus `json:"status"` // pending, paid, overdue
	CreatedAt time.Time               `json:"created_at"`
	UpdatedAt time.Time               `json:"updated_at"`
}

func BuildLoanSchedule(loanID uuid.UUID, installment int, amount float32, now time.Time) LoanSchedule {
	dueDate := now.AddDate(0, 0, installment*7)
	return LoanSchedule{
		ID:        uuid.New(),
		LoanID:    loanID,
		StartDate: dueDate.AddDate(0, 0, -7),
		DueDate:   dueDate,
		Amount:    amount,
		Status:    enum.LoanScheduleStatusPending,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

// UpdateStatus updates the schedule status after a payment
func (s *LoanSchedule) UpdateStatus(paymentAmount float32) float32 {
	if s.Amount <= paymentAmount {
		s.Status = enum.LoanScheduleStatusPaid
		paymentAmount -= s.Amount
	} else {
		s.Amount -= paymentAmount
		paymentAmount = 0
	}

	return paymentAmount
}
