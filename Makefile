
.PHONY: test
test: test-arch

.PHONY: test-arch
test-arch:
	@go test ./test/arch