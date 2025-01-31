package aggregate

import (
	"time"

	"github.com/google/uuid"
)

// LoanTransaction represents a repayment transaction
type LoanTransaction struct {
	ID          uuid.UUID `json:"id" gorm:"primaryKey"`
	LoanID      uuid.UUID `json:"loan_id" gorm:"index"`
	AmountPaid  float32   `json:"amount_paid"`
	PaymentDate time.Time `json:"payment_date"`
	Type        string    `json:"type"` // REPAYMENT, DISBUSEMENT
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func BuildLoanTransaction(loanID uuid.UUID, amount float32, typeTransaction string, now time.Time) LoanTransaction {
	return LoanTransaction{
		ID:          uuid.New(),
		LoanID:      loanID,
		AmountPaid:  amount,
		PaymentDate: now,
		Type:        typeTransaction,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}
