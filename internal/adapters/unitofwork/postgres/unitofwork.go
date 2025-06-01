package postgres

import (
	"context"
	"fmt"

	"github.com/RubenRibGarcia/go-hexagonal-sandbox/internal/adapters/db"
	postgreseventbus "github.com/RubenRibGarcia/go-hexagonal-sandbox/internal/adapters/eventbus/postgres"
	postgresrepo "github.com/RubenRibGarcia/go-hexagonal-sandbox/internal/adapters/repositories/postgres"
	"github.com/RubenRibGarcia/go-hexagonal-sandbox/internal/ports/eventbus"
	"github.com/RubenRibGarcia/go-hexagonal-sandbox/internal/ports/unitofwork"

	portrepo "github.com/RubenRibGarcia/go-hexagonal-sandbox/internal/ports/repositories"
	"github.com/jackc/pgx/v5"
)

type PostgresUnitOfWorkFactory struct {
	conn *pgx.Conn
}

type PostgresUnitOfWork struct {
	tx                          pgx.Tx
	bankAccountRepository       postgresrepo.BankAccountRepositoryImpl
	transactionalEventPublisher postgreseventbus.TransactionalEventPublisher
}

func NewPostgresUnitOfWorkFactory(ctx context.Context, databaseConfig db.DatabaseConfig) (unitofwork.UnitOfWorkFactory, error) {
	connectionString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		databaseConfig.Username,
		databaseConfig.Password,
		databaseConfig.Host,
		databaseConfig.Port,
		databaseConfig.Database,
	)

	conn, err := pgx.Connect(ctx, connectionString)
	if err != nil {
		return nil, err
	}

	return &PostgresUnitOfWorkFactory{
		conn: conn,
	}, nil
}

func (puowf *PostgresUnitOfWorkFactory) NewUnitOfWork(ctx context.Context) (unitofwork.UnitOfWork, error) {
	tx, err := puowf.conn.Begin(ctx)

	if err != nil {
		return nil, err
	}

	return &PostgresUnitOfWork{
		tx:                          tx,
		bankAccountRepository:       postgresrepo.NewBankAccountRepository(tx),
		transactionalEventPublisher: postgreseventbus.NewTransactionalEventPublisher(tx),
	}, nil
}

func (puow *PostgresUnitOfWork) BankAccounts() portrepo.BankAccountRepository {
	return puow.bankAccountRepository
}

func (puow *PostgresUnitOfWork) EventPublisher() eventbus.Publisher {
	return puow.transactionalEventPublisher
}

func (puow *PostgresUnitOfWork) Commit(ctx context.Context) error {
	return puow.tx.Commit(ctx)
}

func (puow *PostgresUnitOfWork) Rollback(ctx context.Context) error {
	return puow.tx.Rollback(ctx)
}
