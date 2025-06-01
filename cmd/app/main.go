package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/RubenRibGarcia/go-hexagonal-sandbox/internal/adapters/api/rest"
	v1 "github.com/RubenRibGarcia/go-hexagonal-sandbox/internal/adapters/api/rest/v1"
	"github.com/RubenRibGarcia/go-hexagonal-sandbox/internal/adapters/api/rest/v1/bankaccounts"
	"github.com/RubenRibGarcia/go-hexagonal-sandbox/internal/adapters/db"
	"github.com/RubenRibGarcia/go-hexagonal-sandbox/internal/adapters/unitofwork/postgres"
	"github.com/RubenRibGarcia/go-hexagonal-sandbox/internal/core/services/bankaccount"
	"github.com/kelseyhightower/envconfig"
)

func main() {
	ctx := context.Background()

	var databaseConfig db.DatabaseConfig
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

	bankAccountService := bankaccount.NewBankAccountService(unitOfWorkFactory)
	bkHandlers := bankaccounts.NewBankAccountHandlers(bankAccountService)
	apiV1Handlers := v1.NewAPIV1Handlers(bkHandlers)

	api := rest.NewAPI()
	api.Mount(apiV1Handlers)

	fmt.Println("Starting server at :8888")
	log.Fatal(http.ListenAndServe(":8888", api.Router))
}
