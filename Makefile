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
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# args passed to sdk/update
commit=false

# This recipe pulls the latest specs for the given service in cloud-api and re-generates the service's go clients.
.PHONY: sdk/update # service=cloud-foo-service commit=true/false
sdk/update:
	@if [ -z $(GITHUB_TOKEN) ]; then \
		echo "ERROR: GITHUB_TOKEN is not set. Please ensure the token has 'repo' access and is SSO enabled." >&2; \
  		exit 1; \
	fi

	@if [ -z $(service) ]; then \
		echo "ERROR: No service argument provided, please provide in the format 'service=...'" >&2; \
  		exit 1; \
	fi

	bash ./scripts/pull-specs.sh $(service);

	@if [ $(service) = "cloud-shared" ]; then \
		echo "Generating latest SDK for cloud-shared"; \
		bash ./scripts/gen-go-shared-sdk.sh; \
	else \
		echo "Generating latest SDK for $(service)"; \
		bash ./scripts/gen-go-service-sdk.sh $(service); \
	fi

	@if [ $(commit) = true ]; then \
		./scripts/open-pr.sh $(service); \
	fi

# This recipe pulls the specs for the given service from locally cloned cloud-api and re-generates the service's go clients.
.PHONY: sdk/update-local # service=cloud-foo-service
sdk/update-local:
	@if [ -z $(service) ]; then \
		echo "ERROR: No service argument provided, please provide in the format 'service=...'" >&2; \
  		exit 1; \
	fi

	bash ./scripts/pull-specs-local.sh $(service);

	@if [ $(service) = "cloud-shared" ]; then \
		echo "Generating latest SDK for cloud-shared"; \
		bash ./scripts/gen-go-shared-sdk.sh; \
	else \
		echo "Generating latest SDK for $(service)"; \
		bash ./scripts/gen-go-service-sdk.sh $(service); \
	fi