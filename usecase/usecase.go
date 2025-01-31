package usecase

import (
	"github.com/amartha-test/repository"
)

type LoanUseCaseImpl struct {
	Repository repository.RepositoryInterface
}

type NewUseCaseOptions struct {
	Repository repository.RepositoryInterface
}

func NewUseCase(opts NewUseCaseOptions) *LoanUseCaseImpl {
	return &LoanUseCaseImpl{
		Repository: opts.Repository,
	}
}
