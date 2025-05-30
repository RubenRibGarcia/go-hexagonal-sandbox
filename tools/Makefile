GOLANGCI_LINT_VERSION := "v2.1.6"
ARCH_GO_VERSION := "v1.7.0"

.PHONY: install
install: install-golangci-lint install-arch-go

.PHONY: install-golangci-lint
install-golangci-lint:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b $$(go env GOPATH)/bin $(GOLANGCI_LINT_VERSION)

.PHONY: install-arch-go
install-arch-go:
	@go install -v github.com/arch-go/arch-go@$(ARCH_GO_VERSION)