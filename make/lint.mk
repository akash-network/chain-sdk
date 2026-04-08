SUB_LINT ?= go \
proto \
shell \
sdl-schema \
ts

BUF_LINT_PACKAGES ?= provider \
node

.PHONY: lint-go
lint-go: $(GOLANGCI_LINT)
	@$(TOOLS) golint "$(GO_MODULES)" "$(GO_TEST_DIRS)"

.PHONY: lint-proto-%
lint-proto-%:
	$(BUF) lint proto/$*

.PHONY: lint-proto
lint-proto: $(BUF) $(patsubst %, lint-proto-%,$(BUF_LINT_PACKAGES))

.PHONY: lint-shell
lint-shell:
	docker run --rm \
	--volume $(PWD):/shellcheck \
	--entrypoint sh \
	koalaman/shellcheck-alpine:stable \
	-x /shellcheck/script/shellcheck.sh

.PHONY: lint
lint: $(patsubst %, lint-%,$(SUB_LINT))

PROTO_AGAINST_DIR := $(AKASH_DEVCACHE_TMP)/proto-against

.PHONY: proto-check-breaking
proto-check-breaking: $(BUF)
	@git worktree remove --force $(PROTO_AGAINST_DIR) 2>/dev/null || true
	@git worktree prune
	@git worktree add $(PROTO_AGAINST_DIR) main
	@cp -r go/vendor $(PROTO_AGAINST_DIR)/go/
	$(BUF) breaking --against '$(PROTO_AGAINST_DIR)'
	@git worktree remove --force $(PROTO_AGAINST_DIR)

.PHONY: proto-format
proto-format:
	$(DOCKER_CLANG) find ./ ! -path "./go/vendor/*" -name *.proto -exec clang-format -i {} \;

.PHONY: lint-sdl-schema
lint-sdl-schema:
	bash $(GO_ROOT)/sdl/lint-schema.sh

.PHONY: lint-ts
lint-ts: $(AKASH_TS_NODE_MODULES)
	cd $(TS_ROOT) && npm run lint;
