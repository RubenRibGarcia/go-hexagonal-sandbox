package domain

import (
	"testing"
	"time"

	"github.com/gofrs/uuid"
)

func TestCancelClientOrder(t *testing.T) {
	id, _ := uuid.NewV4()
	now := time.Now()
	clientOrder := ClientOrder{
		Id:        &id,
		CreatedAt: &now,
		UpdatedAt: &now,
	}

	t.Run("cancel_at_in_the_past", func(t *testing.T) {
		err := clientOrder.Cancel(now.Add(time.Duration(-10) * time.Minute))
	})

}
