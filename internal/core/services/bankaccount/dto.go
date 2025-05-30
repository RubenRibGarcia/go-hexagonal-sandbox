package bankaccount

import (
	"github.com/gofrs/uuid"
	"github.com/shopspring/decimal"
)

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
