package repositories

import (
	"context"

	"github.com/gofrs/uuid"
)

type Repository[T any] interface{
	Get(ctx context.Context, id uuid.UUID) (T, error)
	Create(ctx context.Context, entity T) (T, error)
	Update(ctx context.Context, entity T) (T, error)
}