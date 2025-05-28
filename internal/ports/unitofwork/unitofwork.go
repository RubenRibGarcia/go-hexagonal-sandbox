package unitofwork

import (
	"context"
	"fmt"
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

	fmt.Printf("Atomic operation result: %v, error: %v\n", rvalue, err)

	if err == nil {
		if err = uow.Commit(ctx); err != nil {
			return nil, err
		}
	} else {
		if err = uow.Rollback(ctx); err != nil {
			return nil, err
		}
	}

	return rvalue, err
}
