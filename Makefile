
.PHONY: run-environment
run-environment:
	@docker compose -f deployments/dev/docker-compose.yml up -d

.PHONY: stop-environment
stop-environment:
	@docker compose -f deployments/dev/docker-compose.yml down

.PHONY: local-psql
local-psql:
	@docker compose -f deployments/dev/docker-compose.yml exec postgres psql -U postgres -d go_hexagonal_sandbox -W

.PHONY: create-schema-migration
create-schema-migration:
	ifndef MIGRATION
		$(error MIGRATION variable is not set. Use: make create-schema-migration MIGRATION=<migration_name>)
	endif
	@goose -dir db/migrations create $(MIGRATION) sql

.PHONY: run-schema-migration
run-schema-migration:
	@goose -dir db/migrations postgres "postgres://postgres:password@localhost:5432/go_hexagonal_sandbox?sslmode=disable" up

.PHONY: schema-migration-status
schema-migration-status:
	@goose -dir db/migrations postgres "postgres://postgres:password@localhost:5432/go_hexagonal_sandbox?sslmode=disable" status

reset-schema-migration:
	@goose -dir db/migrations postgres "postgres://postgres:password@localhost:5432/go_hexagonal_sandbox?sslmode=disable" reset

.PHONY: test
test: test-arch

.PHONY: test-arch
test-arch:
	@go test ./test/arch

.PHONY: lint
lint:
	@golangci-lint run

fix-lint:
	@golangci-lint run --fix