package aggregate

import (
	"time"

	"github.com/amartha-test/generated"
	"github.com/amartha-test/model/enum"
	"github.com/google/uuid"
)

// Loan represents a loan entity
type Loan struct {
	ID                  uuid.UUID       `json:"id" gorm:"primaryKey"`
	UserID              string          `json:"user_id"`
	TotalAmount         float32         `json:"total_amount"`
	BalancePrincipal    float32         `json:"balance_principal"`
	DuePrincipal        float32         `json:"due_principal"`
	DueInterest         float32         `json:"due_interest"`
	PaidPrincipal       float32         `json:"paid_principal"`
	PaidInterest        float32         `json:"paid_interest"`
	InterestRate        int             `json:"interest_rate"`
	Status              enum.LoanStatus `json:"status"` // active, delinquent, paid
	NumberOfInstallment int             `json:"number_of_installment"`
	StartDate           time.Time       `json:"start_date"`
	EndDate             time.Time       `json:"end_date"`
	OutstandingBalance  float32         `json:"outstanding_balance" gorm:"-"`
	CreatedAt           time.Time       `json:"created_at"`
	UpdatedAt           time.Time       `json:"updated_at"`
}

func BuildLoan(now time.Time, loan generated.LoanRequest) Loan {
	endDate := now.AddDate(0, 0, int(loan.NumberOfInstallments)*7)
	dueInterest := loan.Amount * float32(loan.InterestRate) / 100
	return Loan{
		ID:                  uuid.New(),
		UserID:              loan.UserId,
		TotalAmount:         loan.Amount,
		BalancePrincipal:    0,
		DuePrincipal:        loan.Amount,
		DueInterest:         dueInterest,
		PaidPrincipal:       0,
		PaidInterest:        0,
		InterestRate:        loan.InterestRate,
		Status:              enum.LoanStatusActive,
		NumberOfInstallment: loan.NumberOfInstallments,
		StartDate:           now,
		EndDate:             endDate,
		CreatedAt:           now,
		UpdatedAt:           now,
	}
}

// ApplyInterestPayment applies the payment to the loan's interest and returns the remaining amount
func (l *Loan) ApplyInterestPayment(amount float32) float32 {
	remainingAmount := amount

	if l.DueInterest > 0 {
		interestPayment := l.DueInterest
		if remainingAmount >= interestPayment {
			l.PaidInterest += interestPayment
			l.DueInterest = 0
			remainingAmount -= interestPayment
		} else {
			l.PaidInterest += remainingAmount
			l.DueInterest -= remainingAmount
			remainingAmount = 0
		}
	}

	return remainingAmount
}

// ApplyPrincipalPayment applies the payment to the loan's principal and returns the remaining amount
func (l *Loan) ApplyPrincipalPayment(amount float32) float32 {
	remainingAmount := amount

	if l.DuePrincipal > 0 {
		principalPayment := l.DuePrincipal
		if remainingAmount >= principalPayment {
			l.PaidPrincipal += principalPayment
			l.DuePrincipal = 0
			remainingAmount -= principalPayment
		} else {
			l.PaidPrincipal += remainingAmount
			l.DuePrincipal -= remainingAmount
			remainingAmount = 0
		}
	}

	return remainingAmount
}

// IsFullyPaid checks if the loan is fully paid off
func (l *Loan) IsFullyPaid() bool {
	return l.DuePrincipal == 0 && l.DueInterest == 0
}

func (l *Loan) GetOutstanding() {
	l.OutstandingBalance = l.TotalAmount - l.BalancePrincipal
}
