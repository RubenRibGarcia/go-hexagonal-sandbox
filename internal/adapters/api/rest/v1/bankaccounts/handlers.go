package bankaccounts

import (
	"context"

	"github.com/RubenRibGarcia/go-hexagonal-sandbox/internal/core/domain"
	"github.com/RubenRibGarcia/go-hexagonal-sandbox/internal/core/services/bankaccount"
	"github.com/danielgtaylor/huma/v2"
	"github.com/gofrs/uuid"
	"github.com/shopspring/decimal"
)

type BankAccountService interface {
	CreateBankAccount(ctx context.Context) (domain.BankAccount, error)
	GetBankAccount(ctx context.Context, id uuid.UUID) (domain.BankAccount, error)
	Deposit(ctx context.Context, request bankaccount.DepositRequest) (domain.BankAccount, error)
	Withdraw(ctx context.Context, request bankaccount.WithdrawRequest) (domain.BankAccount, error)
	Transfer(ctx context.Context, request bankaccount.TransferRequest) (domain.BankAccount, error)
}

type BankAccountHandlers struct {
	bankAccountService BankAccountService
}

func NewBankAccountHandlers(bankAccountService BankAccountService) *BankAccountHandlers {
	return &BankAccountHandlers{
		bankAccountService: bankAccountService,
	}
}

func (h *BankAccountHandlers) MountOn(group *huma.Group) {
	huma.Post(group, "/bank-accounts", h.postBankAccount, func(o *huma.Operation) {
		o.Tags = append(o.Tags, "Bank Accounts")
		o.Summary = "Create new Bank Account"
	})
	huma.Get(group, "/bank-accounts/{id}", h.getBankAccountByID, func(o *huma.Operation) {
		o.Tags = append(o.Tags, "Bank Accounts")
		o.Summary = "Get Bank Account By ID"
	})
	huma.Post(group, "/bank-accounts/{id}/deposit", h.postBankAccountDeposit, func(o *huma.Operation) {
		o.Tags = append(o.Tags, "Bank Accounts")
		o.Summary = "Deposit money into Bank Account"
	})
	huma.Post(group, "/bank-accounts/{id}/withdraw", h.postBankAccountWithdraw, func(o *huma.Operation) {
		o.Tags = append(o.Tags, "Bank Accounts")
		o.Summary = "Withdraw money from Bank Account"
	})
	huma.Post(group, "/bank-accounts/{id}/transfer", h.postBankAccountTransfer, func(o *huma.Operation) {
		o.Tags = append(o.Tags, "Bank Accounts")
		o.Summary = "Transfer money to another Bank Account"
	})
}

func (h *BankAccountHandlers) postBankAccount(ctx context.Context, params *struct{}) (*PostBankAccountResponse, error) {
	entity, err := h.bankAccountService.CreateBankAccount(ctx)
	if err != nil {
		return nil, huma.Error500InternalServerError("Failed to create bank acount", err)
	}

	return &PostBankAccountResponse{
		Body: struct {
			domain.BankAccount
		}{
			BankAccount: entity,
		},
	}, nil
}

func (h *BankAccountHandlers) getBankAccountByID(ctx context.Context, params *struct {
	ID string `path:"id" description:"The ID of the client order to retrieve" example:"123e4567-e89b-12d3-a456-426614174000"`
}) (*GetBankAccountByIDResponse, error) {
	id, err := uuid.FromString(params.ID)

	if err != nil {
		return nil, huma.Error400BadRequest("Invalid UUID format", err)
	}
	entity, err := h.bankAccountService.GetBankAccount(ctx, id)

	if err != nil {
		return nil, huma.Error500InternalServerError("Failed to get bank acount", err)
	}

	return &GetBankAccountByIDResponse{
		Body: struct {
			domain.BankAccount
		}{
			BankAccount: entity,
		},
	}, nil
}

func (h *BankAccountHandlers) postBankAccountDeposit(ctx context.Context, request *PostBankAccountDepositRequest) (*PostBankAccountDepositResponse, error) {
	id, err := uuid.FromString(request.ID)

	if err != nil {
		return nil, huma.Error400BadRequest("Invalid UUID format", err)
	}

	amount, err := decimal.NewFromString(request.Body.Amount)

	if err != nil {
		return nil, huma.Error400BadRequest("Invalid decimal format", err)
	}

	entity, err := h.bankAccountService.Deposit(ctx, bankaccount.DepositRequest{
		BankAccountID: id,
		Amount:        amount,
	})

	if err != nil {
		return nil, huma.Error500InternalServerError("Failed to deposit funds", err)
	}

	return &PostBankAccountDepositResponse{
		Body: struct {
			domain.BankAccount
		}{
			BankAccount: entity,
		},
	}, nil
}

func (h *BankAccountHandlers) postBankAccountWithdraw(ctx context.Context, request *PostBankAccountWithdrawRequest) (*PostBankAccountWithdrawResponse, error) {
	id, err := uuid.FromString(request.ID)

	if err != nil {
		return nil, huma.Error400BadRequest("Invalid UUID format", err)
	}

	amount, err := decimal.NewFromString(request.Body.Amount)

	if err != nil {
		return nil, huma.Error400BadRequest("Invalid decimal format", err)
	}

	entity, err := h.bankAccountService.Withdraw(ctx, bankaccount.WithdrawRequest{
		BankAccountID: id,
		Amount:        amount,
	})

	if err != nil {
		return nil, huma.Error500InternalServerError("Failed to withdraw funds", err)
	}

	return &PostBankAccountWithdrawResponse{
		Body: struct {
			domain.BankAccount
		}{
			BankAccount: entity,
		},
	}, nil
}

func (h *BankAccountHandlers) postBankAccountTransfer(ctx context.Context, request *PostBankAccountTransferRequest) (*PostBankAccountTransferResponse, error) {
	fromID, err := uuid.FromString(request.ID)

	if err != nil {
		return nil, huma.Error400BadRequest("Invalid UUID format", err)
	}

	toID, err := uuid.FromString(request.Body.To)
	if err != nil {
		return nil, huma.Error400BadRequest("Invalid UUID format", err)
	}

	amount, err := decimal.NewFromString(request.Body.Amount)

	if err != nil {
		return nil, huma.Error400BadRequest("Invalid decimal format", err)
	}

	entity, err := h.bankAccountService.Transfer(ctx, bankaccount.TransferRequest{
		FromBankAccountID: fromID,
		ToBankAccountID:   toID,
		Amount:            amount,
	})

	if err != nil {
		return nil, huma.Error500InternalServerError("Failed to transfer funds", err)
	}

	return &PostBankAccountTransferResponse{
		Body: struct {
			domain.BankAccount
		}{
			BankAccount: entity,
		},
	}, nil
}
