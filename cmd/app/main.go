package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/RubenRibGarcia/go-hexagonal-sandbox/internal/adapters/api/rest"
	v1 "github.com/RubenRibGarcia/go-hexagonal-sandbox/internal/adapters/api/rest/v1"
	"github.com/RubenRibGarcia/go-hexagonal-sandbox/internal/adapters/api/rest/v1/bankaccounts"
	"github.com/RubenRibGarcia/go-hexagonal-sandbox/internal/adapters/db"
	"github.com/RubenRibGarcia/go-hexagonal-sandbox/internal/adapters/unitofwork/postgres"
	"github.com/RubenRibGarcia/go-hexagonal-sandbox/internal/core/services/bankaccount"
	"github.com/danielgtaylor/huma/v2/humacli"
	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/cobra"
)

type Options struct {
	Host string `doc:"Hostname to list on"`
	Port int    `doc:"Port to listen on" short:"p" default:"8000"`
}

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

	cli := humacli.New(func(hooks humacli.Hooks, opts *Options) {
		server := http.Server{
			Addr:    fmt.Sprintf(":%d", opts.Port),
			Handler: api.Router,
		}

		hooks.OnStart(func() {
			fmt.Printf("Starting server at :%d\n", opts.Port)
			log.Fatal(server.ListenAndServe())
		})

		hooks.OnStop(func() {
			fmt.Println("Shutting down server...")
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			server.Shutdown(ctx)
		})
	})

	cli.Root().AddCommand(&cobra.Command{
		Use:   "openapi",
		Short: "Print OpenAPI spec",
		Run: func(cmd *cobra.Command, args []string) {
			b, err := api.API.OpenAPI().YAML()
			if err != nil {
				log.Fatalf("failed to generate OpenAPI spec: %v", err)
			}
			fmt.Println(string(b))
		},
	})

	cli.Run()
}
