NAME := cloud-sdk-go

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

.PHONY: test-ci
test-ci: go/lint
	@# TEST_RESULTS is assumed to be set by CircleCI
	@mkdir -p $(TEST_RESULTS)
	$(eval packages = $(shell go list ./... | circleci tests split --split-by=timings --timings-type=classname)) \
	$(GOTESTSUM) \
		--format short-verbose \
		--junitfile $(TEST_RESULTS)/gotestsum-report.xml \
		$(GOTESTSUM_ARGS) \
		-- \
		$(GO_TEST_ARGS) \
		$(packages)