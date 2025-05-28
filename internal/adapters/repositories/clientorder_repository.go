package repositories

import (
	"context"

	"github.com/RubenRibGarcia/go-hexagonal-sandbox/internal/domain"
	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5"
)

type ClientOrderRepositoryImpl struct {
	tx pgx.Tx
}

func NewClientOrderRepository(tx pgx.Tx) ClientOrderRepositoryImpl {
	return ClientOrderRepositoryImpl{
		tx: tx,
	}
}

func (cori ClientOrderRepositoryImpl) Get(ctx context.Context, id uuid.UUID) (domain.ClientOrder, error) {
	//cori.connection.QueryRow()

	return domain.ClientOrder{}, nil
}

func (cori ClientOrderRepositoryImpl) Create(ctx context.Context, entity domain.ClientOrder) (domain.ClientOrder, error) {
	return domain.ClientOrder{}, nil
}

func (cori ClientOrderRepositoryImpl) Update(ctx context.Context, entity domain.ClientOrder) (domain.ClientOrder, error) {
	return domain.ClientOrder{}, nil
}
