package postgres

import "github.com/jackc/pgx/v5"

type TransactionalEventPublisher struct {
	tx pgx.Tx
}

func NewTransactionalEventPublisher(tx pgx.Tx) TransactionalEventPublisher {
	return TransactionalEventPublisher{
		tx: tx,
	}
}

func (tp TransactionalEventPublisher) Publish(v any) error {

}
