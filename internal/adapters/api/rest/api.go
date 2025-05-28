package rest

import (
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
	"net/http"
)

type Handlers interface {
	Register(api *huma.API)
}

func NewAPI(handlers ...Handlers) *http.ServeMux {
	router := http.NewServeMux()
	api := humago.New(router, huma.DefaultConfig("Go Hexagonal Sandbox API", "v0.0.1"))
	for _, handler := range handlers {
		handler.Register(&api)
	}

	return router
}
