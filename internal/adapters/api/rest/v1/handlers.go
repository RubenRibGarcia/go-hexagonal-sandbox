package v1

import (
	"github.com/RubenRibGarcia/go-hexagonal-sandbox/internal/adapters/api/rest"
	"github.com/danielgtaylor/huma/v2"
)

type APIV1Handlers struct {
	handlers []rest.HandlersRegister
}

func (v1h APIV1Handlers) MountOn(group *huma.Group) {
	apiV1Group := huma.NewGroup(group, "/v1")

	for _, handler := range v1h.handlers {
		handler.MountOn(apiV1Group)
	}
}

func NewAPIV1Handlers(handlers ...rest.HandlersRegister) APIV1Handlers {
	return APIV1Handlers{
		handlers: handlers,
	}
}
