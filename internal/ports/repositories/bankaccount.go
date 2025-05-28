package repositories

import (
	"github.com/RubenRibGarcia/go-hexagonal-sandbox/internal/core/domain"
)

type BankAccountRepository interface {
	Repository[domain.BankAccount]
}
