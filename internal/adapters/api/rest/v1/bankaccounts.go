package v1

import (
	"context"
	"github.com/RubenRibGarcia/go-hexagonal-sandbox/internal/core/domain"
	"github.com/RubenRibGarcia/go-hexagonal-sandbox/internal/core/service"
	"github.com/danielgtaylor/huma/v2"
	"github.com/gofrs/uuid"
	"github.com/shopspring/decimal"
)

const (
	v1_prefix = "/api/v1"
)

type Handlers struct {
	bankAccountService service.BankAccountService
}

func NewBankAccountHandlers(bankAccountService service.BankAccountService) *Handlers {
	return &Handlers{
		bankAccountService: bankAccountService,
	}
}

func (h *Handlers) Register(api *huma.API) {
	huma.Post(*api, buildPath("/bank-accounts"), h.postBankAccount)
	huma.Get(*api, buildPath("/bank-accounts/{id}"), h.getBankAccountByID)
	huma.Post(*api, buildPath("/bank-accounts/{id}/deposit"), h.postBankAccountDeposit)
	huma.Post(*api, buildPath("/bank-accounts/{id}/withdraw"), h.postBankAccountWithdraw)
	huma.Post(*api, buildPath("/bank-accounts/{id}/transfer"), h.postBankAccountTransfer)
}

func buildPath(path string) string {
	return v1_prefix + path
}

func (h *Handlers) postBankAccount(ctx context.Context, params *struct{}) (*PostBankAccountResponse, error) {
	entity, err := h.bankAccountService.CreateBankAccount(ctx)
	if err != nil {
		return nil, err
	}

	return &PostBankAccountResponse{
		Body: struct {
			domain.BankAccount
		}{
			BankAccount: entity,
		},
	}, nil
}

func (h *Handlers) getBankAccountByID(ctx context.Context, params *struct {
	ID string `path:"id" description:"The ID of the client order to retrieve" example:"123e4567-e89b-12d3-a456-426614174000"`
}) (*GetBankAccountByIDResponse, error) {
	id, err := uuid.FromString(params.ID)

	if err != nil {
		return nil, huma.Error400BadRequest("Invalid UUID format", err)
	}
	entity, err := h.bankAccountService.GetBankAccount(ctx, id)
	return &GetBankAccountByIDResponse{
		Body: struct {
			domain.BankAccount
		}{
			BankAccount: entity,
		},
	}, nil
}

func (h *Handlers) postBankAccountDeposit(ctx context.Context, request *PostBankAccountDepositRequest) (*PostBankAccountDepositResponse, error) {
	id, err := uuid.FromString(request.ID)

	if err != nil {
		return nil, huma.Error400BadRequest("Invalid UUID format", err)
	}

	amount, err := decimal.NewFromString(request.Body.Amount)

	if err != nil {
		return nil, huma.Error400BadRequest("Invalid decimal format", err)
	}

	entity, err := h.bankAccountService.Deposit(ctx, service.DepositRequest{
		BankAccountID: id,
		Amount:        amount,
	})

	return &PostBankAccountDepositResponse{
		Body: struct {
			domain.BankAccount
		}{
			BankAccount: entity,
		},
	}, nil
}

func (h *Handlers) postBankAccountWithdraw(ctx context.Context, request *PostBankAccountWithdrawRequest) (*PostBankAccountWithdrawResponse, error) {
	id, err := uuid.FromString(request.ID)

	if err != nil {
		return nil, huma.Error400BadRequest("Invalid UUID format", err)
	}

	amount, err := decimal.NewFromString(request.Body.Amount)

	if err != nil {
		return nil, huma.Error400BadRequest("Invalid decimal format", err)
	}

	entity, err := h.bankAccountService.Withdraw(ctx, service.WithdrawRequest{
		BankAccountID: id,
		Amount:        amount,
	})

	return &PostBankAccountWithdrawResponse{
		Body: struct {
			domain.BankAccount
		}{
			BankAccount: entity,
		},
	}, nil
}

func (h *Handlers) postBankAccountTransfer(ctx context.Context, request *PostBankAccountTransferRequest) (*PostBankAccountTransferResponse, error) {
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

	entity, err := h.bankAccountService.Transfer(ctx, service.TransferRequest{
		FromBankAccountID: fromID,
		ToBankAccountID:   toID,
		Amount:            amount,
	})

	return &PostBankAccountTransferResponse{
		Body: struct {
			domain.BankAccount
		}{
			BankAccount: entity,
		},
	}, nil
}
