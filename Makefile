NAME := hcp-sdk-go

GO ?= go
GO_LINT ?= golangci-lint

GO_LINT_CONFIG_PATH ?= ./golangci-config.yml

# args passed to go test
GOTESTSUM ?= gotestsum

GOTESTSUM_ARGS ?=

GO_SCRIPTS_DIR = $(dir $(GO_MKFILE_PATH))scripts

default: test

.PHONY: test
test: go/lint go/tidy go/test

go/test: ## Run go tests
	$(GOTESTSUM) $(GOTESTSUM_ARGS) -- $(GO_TEST_ARGS) ./...

go/tidy:
	$(GO) mod tidy

go/lint:
	$(GO_LINT) run --config $(GO_LINT_CONFIG_PATH) $(GO_LINT_ARGS)

# Run the test and generate an html coverage file.
.PHONY: test-ci
test-ci: go/lint
	go test -coverprofile=coverage.out -short ./...
	go tool cover -html=coverage.out -o coverage.html
