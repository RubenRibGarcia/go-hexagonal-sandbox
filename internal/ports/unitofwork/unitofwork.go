package unitofwork

import (
	"context"

	"github.com/RubenRibGarcia/go-hexagonal-sandbox/internal/ports/repositories"
)

type UnitOfWork interface {
	ClientOrders() repositories.ClientOrderRepository
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}