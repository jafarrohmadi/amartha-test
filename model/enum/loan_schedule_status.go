package enum

type LoanScheduleStatus string

// pending, paid,
const (
	LoanScheduleStatusPending LoanScheduleStatus = "PENDING"
	LoanScheduleStatusPaid    LoanScheduleStatus = "PAID"
)
