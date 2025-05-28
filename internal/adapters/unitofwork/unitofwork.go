package unitofwork

import (
	"context"

	adapterrepo "github.com/RubenRibGarcia/go-hexagonal-sandbox/internal/adapters/repositories"
	portrepo "github.com/RubenRibGarcia/go-hexagonal-sandbox/internal/ports/repositories"
	portsuow "github.com/RubenRibGarcia/go-hexagonal-sandbox/internal/ports/unitofwork"
	"github.com/jackc/pgx/v5"
)

type PostgresUnitOfWork struct {
	tx                     pgx.Tx
	clientOrdersRepository adapterrepo.ClientOrderRepositoryImpl
}

func WithUnitOfWork[R any](ctx context.Context, fw func(uow portsuow.UnitOfWork) (*R, error)) (*R, error) {

	conn, err := pgx.Connect(ctx, "")

	if err != nil {
		return nil, err
	}

	tx, err := conn.Begin(ctx)

	if err != nil {
		return nil, err
	}

	uow := &PostgresUnitOfWork{
		tx: tx,
		clientOrdersRepository: adapterrepo.NewClientOrderRepository(
			tx,
		),
	}

	defer uow.Commit(ctx)

	rvalue, err := fw(uow)

	if err == nil {
		uow.Commit(ctx)
	} else {
		uow.Rollback(ctx)
	}

	return rvalue, err
}

func (puow *PostgresUnitOfWork) ClientOrders() portrepo.ClientOrderRepository {
	return puow.clientOrdersRepository
}

func (puow *PostgresUnitOfWork) Commit(ctx context.Context) error {
	return puow.tx.Commit(ctx)
}

func (puow *PostgresUnitOfWork) Rollback(ctx context.Context) error {
	return puow.tx.Rollback(ctx)
}
