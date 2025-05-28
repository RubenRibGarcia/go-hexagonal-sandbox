package repositories

import "github.com/RubenRibGarcia/go-hexagonal-sandbox/internal/domain"

type ClientOrderRepository interface {
	Repository[domain.ClientOrder]
}
