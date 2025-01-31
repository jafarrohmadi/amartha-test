package repository

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type LoanRepository struct {
	db *gorm.DB
}

type NewRepositoryOptions struct {
	Dsn string
}

// NewRepository initializes a new LoanRepository with database connection.
func NewRepository(options NewRepositoryOptions) *LoanRepository {
	// Connect to PostgreSQL database using DSN
	db, err := gorm.Open(postgres.Open(options.Dsn), &gorm.Config{})
	if err != nil {
		return nil
	}
	return &LoanRepository{db: db}
}
