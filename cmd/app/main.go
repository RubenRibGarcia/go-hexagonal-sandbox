package main

import (
	"context"
	"fmt"
	"github.com/RubenRibGarcia/go-hexagonal-sandbox/internal/adapters/api/rest"
	v1 "github.com/RubenRibGarcia/go-hexagonal-sandbox/internal/adapters/api/rest/v1"
	postgres "github.com/RubenRibGarcia/go-hexagonal-sandbox/internal/adapters/unitofwork"
	"github.com/RubenRibGarcia/go-hexagonal-sandbox/internal/core/service"
	"github.com/kelseyhightower/envconfig"
	"net/http"
)

func main() {
	ctx := context.Background()

	var databaseConfig postgres.DatabaseConfig
	err := envconfig.Process("", &databaseConfig)
	if err != nil {
		panic(err)
	}

	unitOfWorkFactory, err := postgres.NewPostgresUnitOfWorkFactory(
		ctx,
		databaseConfig,
	)
	if err != nil {
		panic(err)
	}

	bankAccountService := service.NewBankAccountService(unitOfWorkFactory)

	handlers := v1.NewBankAccountHandlers(bankAccountService)

	router := rest.NewAPI(handlers)

	fmt.Println("Starting server 127.0.0.1:8888")
	err = http.ListenAndServe("127.0.0.1:8888", router)
	if err != nil {
		panic(err)
	}
}
