package service

import (
	"context"
	"fmt"
	"github.com/RubenRibGarcia/go-hexagonal-sandbox/internal/core/domain"
	"github.com/RubenRibGarcia/go-hexagonal-sandbox/internal/ports/unitofwork"
	"github.com/gofrs/uuid"
)

type BankAccountService struct {
	uowf unitofwork.UnitOfWorkFactory
}

func NewBankAccountService(
	uowf unitofwork.UnitOfWorkFactory,
) BankAccountService {
	return BankAccountService{
		uowf: uowf,
	}
}

func (cos *BankAccountService) CreateBankAccount(ctx context.Context) (domain.BankAccount, error) {
	bankAccount, err := unitofwork.Atomic(ctx, cos.uowf, func(uow unitofwork.UnitOfWork) (*domain.BankAccount, error) {
		entity, err := uow.BankAccounts().Create(ctx, domain.NewBankAccount())
		if err != nil {
			return nil, err
		}
		return &entity, nil
	})

	if err != nil {
		return domain.BankAccount{}, err
	}
	fmt.Printf("bank account created: %+v\n", bankAccount)

	return *bankAccount, nil
}

func (cos *BankAccountService) GetBankAccount(ctx context.Context, id uuid.UUID) (domain.BankAccount, error) {
	bankAccount, err := unitofwork.Atomic(ctx, cos.uowf, func(uow unitofwork.UnitOfWork) (*domain.BankAccount, error) {
		entity, err := uow.BankAccounts().Get(ctx, id)
		if err != nil {
			return nil, err
		}
		return &entity, nil
	})

	if err != nil {
		return domain.BankAccount{}, err
	}

	return *bankAccount, nil
}
