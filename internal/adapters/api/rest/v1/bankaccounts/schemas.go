package bankaccounts

import "github.com/RubenRibGarcia/go-hexagonal-sandbox/internal/core/domain"

type (
	GetBankAccountByIDResponse struct {
		Body struct {
			domain.BankAccount
		}
	}

	PostBankAccountResponse struct {
		Body struct {
			domain.BankAccount
		}
	}

	PostBankAccountDepositRequest struct {
		ID   string `path:"id" description:"The ID of the bank account to deposit into" example:"123e4567-e89b-12d3-a456-426614174000"`
		Body struct {
			Amount string `json:"amount" description:"The amount to deposit" example:"100.00"`
		}
	}

	PostBankAccountDepositResponse struct {
		Body struct {
			domain.BankAccount
		}
	}

	PostBankAccountWithdrawRequest struct {
		ID   string `path:"id" description:"The ID of the bank account to deposit into" example:"123e4567-e89b-12d3-a456-426614174000"`
		Body struct {
			Amount string `json:"amount" description:"The amount to deposit" example:"100.00"`
		}
	}

	PostBankAccountWithdrawResponse struct {
		Body struct {
			domain.BankAccount
		}
	}

	PostBankAccountTransferRequest struct {
		ID   string `path:"id" description:"The ID of the bank account to deposit into" example:"123e4567-e89b-12d3-a456-426614174000"`
		Body struct {
			To     string `json:"to" description:"The ID of the bank account to transfer to" example:"123e4567-e89b-12d3-a456-426614174002"`
			Amount string `json:"amount" description:"The amount to deposit" example:"100.00"`
		}
	}

	PostBankAccountTransferResponse struct {
		Body struct {
			domain.BankAccount
		}
	}
)
