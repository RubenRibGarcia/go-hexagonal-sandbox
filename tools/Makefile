GOLANGCI_LINT_VERSION := "v2.1.6"

.PHONY: install
install: install-golangci-lint

.PHONY: install-golangci-lint
install-golangci-lint:
	# binary will be $(go env GOPATH)/bin/golangci-lint
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b $(go env GOPATH)/bin $(GOLANGCI_LINT_VERSION)