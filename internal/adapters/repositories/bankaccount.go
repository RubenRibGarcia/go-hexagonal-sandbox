package repositories

import (
	"context"
	"errors"
	"github.com/RubenRibGarcia/go-hexagonal-sandbox/internal/core/domain"
	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5"
	"time"
)

type BankAccountRepositoryImpl struct {
	tx pgx.Tx
}

func NewBankAccountRepository(tx pgx.Tx) BankAccountRepositoryImpl {
	return BankAccountRepositoryImpl{
		tx: tx,
	}
}

func (cori BankAccountRepositoryImpl) Get(ctx context.Context, id uuid.UUID) (domain.BankAccount, error) {
	rows, err := cori.tx.Query(ctx, "SELECT * FROM bank_accounts WHERE id = $1", id)
	if err != nil {
		return domain.BankAccount{}, err
	}

	bankAccount, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByNameLax[domain.BankAccount])
	if err != nil {
		return domain.BankAccount{}, err
	}

	rows, err = cori.tx.Query(ctx, "SELECT * FROM transactions WHERE bank_account_id = $1", id)
	if err != nil {
		return domain.BankAccount{}, err
	}
	transactions, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[domain.Transaction])
	if err != nil {
		return domain.BankAccount{}, err
	}

	bankAccount.Transactions = transactions

	return bankAccount, nil
}

func (cori BankAccountRepositoryImpl) Create(ctx context.Context, entity domain.BankAccount) (domain.BankAccount, error) {
	id := uuid.Must(uuid.NewV7())
	now := time.Now()
	ct, err := cori.tx.Exec(
		ctx,
		"INSERT INTO bank_accounts (id, created_at, updated_at, balance) VALUES ($1, $2, $3, $4)",
		id,
		now,
		now,
		entity.Balance,
	)

	if err != nil {
		return domain.BankAccount{}, err
	}
	if ct.RowsAffected() == 0 {
		return domain.BankAccount{}, errors.New("no row inserted")
	}

	entity.ID = &id
	entity.CreatedAt = &now
	entity.UpdatedAt = &now

	return entity, nil
}

func (cori BankAccountRepositoryImpl) Update(ctx context.Context, entity domain.BankAccount) (domain.BankAccount, error) {
	now := time.Now()
	ct, err := cori.tx.Exec(
		ctx,
		"UPDATE bank_accounts SET updated_at = $1, balance = $2 WHERE id = $3",
		now,
		entity.Balance,
		entity.ID,
	)

	if err != nil {
		return domain.BankAccount{}, err
	}
	if ct.RowsAffected() == 0 {
		return domain.BankAccount{}, errors.New("no row updated")
	}

	for _, transaction := range entity.Transactions {
		if transaction.ID == nil {
			id := uuid.Must(uuid.NewV7())
			transaction.ID = &id
			transaction.CreatedAt = &now

			ct, err = cori.tx.Exec(
				ctx,
				"INSERT INTO transactions (id, created_at, bank_account_id, amount, kind) VALUES ($1, $2, $3, $4, $5)",
				id,
				now,
				entity.ID,
				transaction.Amount,
				transaction.Kind,
			)

			if err != nil {
				return domain.BankAccount{}, err
			}
		}
	}

	return entity, nil
}
