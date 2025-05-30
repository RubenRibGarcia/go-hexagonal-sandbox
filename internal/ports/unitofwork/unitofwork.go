package unitofwork

import (
	"context"
	"github.com/RubenRibGarcia/go-hexagonal-sandbox/internal/ports/repositories"
)

type UnitOfWorkFactory interface {
	NewUnitOfWork(ctx context.Context) (UnitOfWork, error)
}

type UnitOfWork interface {
	BankAccounts() repositories.BankAccountRepository
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}

func Atomic[R any](ctx context.Context, uowf UnitOfWorkFactory, fw func(uow UnitOfWork) (*R, error)) (*R, error) {
	uow, err := uowf.NewUnitOfWork(ctx)

	if err != nil {
		return nil, err
	}

	rvalue, err := fw(uow)

	if err == nil {
		if commitErr := uow.Commit(ctx); commitErr != nil {
			return nil, commitErr
		}
	} else {
		if rollbackErr := uow.Rollback(ctx); rollbackErr != nil {
			return nil, rollbackErr
		}
	}

	return rvalue, err
}
