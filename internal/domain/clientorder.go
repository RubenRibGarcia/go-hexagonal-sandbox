package domain

import (
	"errors"
	"time"

	"github.com/gofrs/uuid"
)

type ClientOrderChanges struct {
	CancellationTime *time.Time
}

type ClientOrder struct {
	Id          *uuid.UUID
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	CancelledAt *time.Time
	Changes     []ClientOrderChanges
}

func (co *ClientOrder) Cancel(cancelAt time.Time) error {
	now := time.Now()

	if cancelAt.Before(now) {
		return errors.New("trying to cancel a order where the cancel at is in the past")
	}

	co.Changes = append(co.Changes, ClientOrderChanges{CancellationTime: &cancelAt})

	return nil
}
