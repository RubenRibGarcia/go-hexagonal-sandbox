package rest

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
)

type HandlersRegister interface {
	MountOn(group *huma.Group)
}

type RestAPI struct {
	Router   *http.ServeMux
	api      huma.API
	apiGroup *huma.Group
}

func NewAPI() RestAPI {
	router := http.NewServeMux()
	api := humago.New(router, huma.DefaultConfig("Go Hexagonal Sandbox API", "v0.0.1"))
	apiGroup := huma.NewGroup(api, "/api")

	return RestAPI{
		Router:   router,
		api:      api,
		apiGroup: apiGroup,
	}
}

func (rapi RestAPI) Mount(handlers ...HandlersRegister) {
	for _, handler := range handlers {
		handler.MountOn(rapi.apiGroup)
	}
}
