package v1

import (
	"context"
	"github.com/RubenRibGarcia/go-hexagonal-sandbox/internal/core/domain"
	"github.com/RubenRibGarcia/go-hexagonal-sandbox/internal/core/service"
	"github.com/danielgtaylor/huma/v2"
	"github.com/gofrs/uuid"
)

const (
	v1_prefix = "/api/v1"
)

type GetBankAccountByIDResponse struct {
	Body struct {
		domain.BankAccount
	}
}

type PostBankAccountResponse struct {
	Body struct {
		domain.BankAccount
	}
}

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
