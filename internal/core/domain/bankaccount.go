package domain

import (
	"errors"
	"time"

	"github.com/gofrs/uuid"
	"github.com/shopspring/decimal"
)

type TransactionKind string

const (
	CREDIT TransactionKind = "credit"
	DEBIT  TransactionKind = "debit"
)

type TransactionOperation string

const (
	DEPOSIT    TransactionOperation = "deposit"
	WITHDRAWAL TransactionOperation = "withdrawal"
	TRANSFER   TransactionOperation = "transfer"
)

type Transaction struct {
	ID        *uuid.UUID           `db:"id"`
	CreatedAt *time.Time           `db:"created_at"`
	Amount    decimal.Decimal      `db:"amount"`
	Kind      TransactionKind      `db:"kind"`
	Operation TransactionOperation `db:"operation"`
}

type BankAccount struct {
	ID           *uuid.UUID      `db:"id"`
	CreatedAt    *time.Time      `db:"created_at"`
	UpdatedAt    *time.Time      `db:"updated_at"`
	Transactions []*Transaction  `db:"-"`
	Balance      decimal.Decimal `db:"balance"`
}

func NewBankAccount() BankAccount {
	now := time.Now()
	return BankAccount{
		CreatedAt:    &now,
		UpdatedAt:    &now,
		Transactions: []*Transaction{},
		Balance:      decimal.Zero,
	}
}

func (bk *BankAccount) Deposit(amount decimal.Decimal) error {
	if amount.LessThanOrEqual(decimal.Zero) {
		return errors.New("deposit amount must be greater than zero")
	}

	bk.Balance = bk.Balance.Add(amount)
	bk.Transactions = append(bk.Transactions, &Transaction{
		Amount:    amount,
		Kind:      CREDIT,
		Operation: DEPOSIT,
	})

	return nil
}

func (bk *BankAccount) Withdraw(amount decimal.Decimal) error {
	if amount.LessThanOrEqual(decimal.Zero) {
		return errors.New("withdraw amount must be greater than zero")
	}

	if bk.Balance.LessThan(amount) {
		return errors.New("insufficient funds for withdrawal")
	}

	bk.Balance = bk.Balance.Sub(amount)
	bk.Transactions = append(bk.Transactions, &Transaction{
		Amount:    amount,
		Kind:      DEBIT,
		Operation: WITHDRAWAL,
	})

	return nil
}

func (bk *BankAccount) Transfer(to *BankAccount, amount decimal.Decimal) error {
	if amount.LessThanOrEqual(decimal.Zero) {
		return errors.New("transfer amount must be greater than zero")
	}

	if bk.Balance.LessThan(amount) {
		return errors.New("insufficient funds for transfer")
	}

	if to == nil {
		return errors.New("transfer target account cannot be nil")
	}

	bk.Balance = bk.Balance.Sub(amount)
	to.Balance = to.Balance.Add(amount)

	bk.Transactions = append(bk.Transactions, &Transaction{
		Amount:    amount,
		Kind:      DEBIT,
		Operation: TRANSFER,
	})
	to.Transactions = append(to.Transactions, &Transaction{
		Amount:    amount,
		Kind:      CREDIT,
		Operation: TRANSFER,
	})

	return nil
}
