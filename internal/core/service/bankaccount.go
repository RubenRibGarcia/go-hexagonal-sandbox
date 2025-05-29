package service

import (
	"context"
	"github.com/RubenRibGarcia/go-hexagonal-sandbox/internal/core/domain"
	"github.com/RubenRibGarcia/go-hexagonal-sandbox/internal/ports/unitofwork"
	"github.com/gofrs/uuid"
	"github.com/shopspring/decimal"
)

type BankAccountService struct {
	uowf unitofwork.UnitOfWorkFactory
}

type DepositRequest struct {
	BankAccountID uuid.UUID
	Amount        decimal.Decimal
}

type WithdrawRequest struct {
	BankAccountID uuid.UUID
	Amount        decimal.Decimal
}

type TransferRequest struct {
	FromBankAccountID uuid.UUID
	ToBankAccountID   uuid.UUID
	Amount            decimal.Decimal
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

func (cos *BankAccountService) Deposit(ctx context.Context, request DepositRequest) (domain.BankAccount, error) {
	bankAccount, err := unitofwork.Atomic(ctx, cos.uowf, func(uow unitofwork.UnitOfWork) (*domain.BankAccount, error) {
		entity, err := uow.BankAccounts().Get(ctx, request.BankAccountID)
		if err != nil {
			return nil, err
		}

		if err = entity.Deposit(request.Amount); err != nil {
			return nil, err
		}

		if entity, err = uow.BankAccounts().Update(ctx, entity); err != nil {
			return nil, err
		}

		return &entity, nil
	})

	if err != nil {
		return domain.BankAccount{}, err
	}

	return *bankAccount, nil
}

func (cos *BankAccountService) Withdraw(ctx context.Context, request WithdrawRequest) (domain.BankAccount, error) {
	bankAccount, err := unitofwork.Atomic(ctx, cos.uowf, func(uow unitofwork.UnitOfWork) (*domain.BankAccount, error) {
		entity, err := uow.BankAccounts().Get(ctx, request.BankAccountID)
		if err != nil {
			return nil, err
		}

		if err = entity.Withdraw(request.Amount); err != nil {
			return nil, err
		}

		if entity, err = uow.BankAccounts().Update(ctx, entity); err != nil {
			return nil, err
		}

		return &entity, nil
	})

	if err != nil {
		return domain.BankAccount{}, err
	}

	return *bankAccount, nil
}

func (cos *BankAccountService) Transfer(ctx context.Context, request TransferRequest) (domain.BankAccount, error) {
	account, err := unitofwork.Atomic(ctx, cos.uowf, func(uow unitofwork.UnitOfWork) (*domain.BankAccount, error) {
		fromAccount, err := uow.BankAccounts().Get(ctx, request.FromBankAccountID)
		if err != nil {
			return nil, err
		}

		toAccount, err := uow.BankAccounts().Get(ctx, request.ToBankAccountID)
		if err != nil {
			return nil, err
		}

		if err = fromAccount.Withdraw(request.Amount); err != nil {
			return nil, err
		}

		if err = toAccount.Deposit(request.Amount); err != nil {
			return nil, err
		}

		if _, err = uow.BankAccounts().Update(ctx, fromAccount); err != nil {
			return nil, err
		}

		if _, err = uow.BankAccounts().Update(ctx, toAccount); err != nil {
			return nil, err
		}

		return &fromAccount, nil
	})

	if err != nil {
		return domain.BankAccount{}, err
	}

	return *account, nil
}
