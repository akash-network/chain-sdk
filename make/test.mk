GO_COVER_PACKAGES  = $(shell cd $(GO_ROOT); go list $(GO_TEST_DIRS) | grep -v mock | paste -sd, -)

GO_CURR_MODULE     = $(shell go list -m 2>/dev/null)

SUB_TESTS         ?= go \
ts

GO_TEST_OPTS     ?=
GO_TEST_TIMEOUT  ?= 300

test_go_flags := -mod=$(GOMOD) -timeout $(GO_TEST_TIMEOUT)s

ifneq (,$(findstring nocache,$(GO_TEST_OPTS)))
test_go_flags += -count=1
endif

ifneq (,$(findstring race,$(GO_TEST_OPTS)))
test_go_flags += -race
endif

ifneq (,$(findstring verbose,$(GO_TEST_OPTS)))
test_go_flags += -v
endif

.PHONY: test
test: $(patsubst %, test-%,$(SUB_TESTS))

.PHONY: test-coverage
test-coverage: $(patsubst %, test-coverage-%,$(SUB_TESTS))

.PHONY: test-ts
test-ts: $(AKASH_TS_NODE_MODULES) $(BUF)
	cd $(TS_ROOT) && (npm run build && npm run test)

.PHONY: test-coverage-ts
test-coverage-ts: $(AKASH_TS_NODE_MODULES) proto-gen-ts
	cd $(TS_ROOT) && (npm run build && npm run test:cov)

.PHONY: test-go
test-go: export GO111MODULE := $(GO111MODULE)
test-go: export GOWORK := $(GOWORK)
test-go:
	@$(TOOLS) gotest "$(GO_MODULES)" "$(test_go_flags)" "$(GO_TEST_DIRS)"

.PHONY: test-coverage-go
test-coverage-go: export GO111MODULE := $(GO111MODULE)
test-coverage-go: export GOWORK := $(GOWORK)
test-coverage-go:
	@$(TOOLS) gocoverage "$(GO_MODULES)" "$(test_go_flags)" "$(GO_TEST_DIRS)"

.PHONY: generate-sdl-fixtures
generate-sdl-fixtures: ## Generate SDL test fixtures (manifest.json and groups.json from input.yaml files)
	@echo "Generating SDL fixtures..."
	cd  go/sdl/tools/generate-sdl-fixtures && go run .

.PHONY: test-sdl-parity
test-sdl-parity: generate-sdl-fixtures $(AKASH_TS_NODE_MODULES) ## Run SDL parity tests for Go and TypeScript
	@echo "Running Go SDL parity and validation tests..."
	@cd go/sdl && go test -v -run "TestParity|TestInvalidSDLsRejected|TestSchemaValidation"
	@echo ""
	@echo "Running TypeScript SDL parity tests..."
	@echo "  (includes: v2.0 fixtures, v2.1 fixtures, and invalid SDL rejection)"
	@cd ts && npm test -- --testPathPattern=parity.spec.ts
