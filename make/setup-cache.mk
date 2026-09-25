###############################################################################
###                           Protobuf                                    ###
###############################################################################
ifeq ($(UNAME_OS),Linux)
ifeq ($(UNAME_ARCH),aarch64)
	PROTOC_ZIP ?= protoc-${PROTOC_VERSION}-linux-aarch_64.zip
else
	PROTOC_ZIP ?= protoc-${PROTOC_VERSION}-linux-$(UNAME_ARCH).zip
endif
endif
ifeq ($(UNAME_OS),Darwin)
	PROTOC_ZIP ?= protoc-${PROTOC_VERSION}-osx-universal_binary.zip
endif

SHA256SUM := $(shell command -v sha256sum 2>/dev/null || echo "shasum -a 256")

# verify_sha256 checks a downloaded file against the release's published checksum list
# $(1): checksum list URL, $(2): asset name as listed there, $(3): downloaded file
define verify_sha256
	expected=$$(curl -fsSL "$(1)" | awk '$$2 == "$(2)" { print $$1 }'); \
	actual=$$($(SHA256SUM) "$(3)" | awk '{ print $$1 }'); \
	if [ -z "$$expected" ] || [ "$$expected" != "$$actual" ]; then \
		echo "checksum mismatch for $(2): expected '$$expected', got '$$actual'"; \
		rm -f "$(3)"; \
		exit 1; \
	fi
endef

BUF_RELEASE_URL := https://github.com/bufbuild/buf/releases/download/v$(BUF_VERSION)
BUF_ASSET       := buf-$(UNAME_OS)-$(UNAME_ARCH)

GOLANGCI_LINT_OS          := $(shell echo $(UNAME_OS) | tr '[:upper:]' '[:lower:]')
GOLANGCI_LINT_ARCH        := $(patsubst aarch64,arm64,$(patsubst x86_64,amd64,$(UNAME_ARCH)))
GOLANGCI_LINT_RELEASE_URL := https://github.com/golangci/golangci-lint/releases/download/$(GOLANGCI_LINT_VERSION)
GOLANGCI_LINT_ASSET       := golangci-lint-$(patsubst v%,%,$(GOLANGCI_LINT_VERSION))-$(GOLANGCI_LINT_OS)-$(GOLANGCI_LINT_ARCH)
GOLANGCI_LINT_CHECKSUMS   := golangci-lint-$(patsubst v%,%,$(GOLANGCI_LINT_VERSION))-checksums.txt

$(AKASH_DEVCACHE):
	@echo "creating .cache dir structure..."
	mkdir -p $@
	mkdir -p $(AKASH_DEVCACHE_BIN)
	mkdir -p $(AKASH_DEVCACHE_INCLUDE)
	mkdir -p $(AKASH_DEVCACHE_VERSIONS)
	mkdir -p $(AKASH_DEVCACHE_NODE_MODULES)
	mkdir -p $(AKASH_DEVCACHE)/run

cache: $(AKASH_DEVCACHE)

$(BUF_VERSION_FILE): $(AKASH_DEVCACHE)
	@echo "installing buf v$(BUF_VERSION) ..."
	rm -f $(BUF)
	curl -fsSL -o $(BUF).download "$(BUF_RELEASE_URL)/$(BUF_ASSET)"
	$(call verify_sha256,$(BUF_RELEASE_URL)/sha256.txt,$(BUF_ASSET),$(BUF).download)
	chmod +x $(BUF).download
	mv $(BUF).download $(BUF)
	rm -rf "$(dir $@)"
	mkdir -p "$(dir $@)"
	touch $@
$(BUF): $(BUF_VERSION_FILE)

$(PROTOC_VERSION_FILE): $(AKASH_DEVCACHE)
	@echo "installing protoc compiler v$(PROTOC_VERSION) ..."
	rm -f $(PROTOC)
	(cd /tmp; \
	curl -sOL "https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/${PROTOC_ZIP}"; \
	unzip -oq ${PROTOC_ZIP} -d $(AKASH_DEVCACHE) bin/protoc; \
	rm -f ${PROTOC_ZIP})
	rm -rf "$(dir $@)"
	mkdir -p "$(dir $@)"
	touch $@
$(PROTOC): $(PROTOC_VERSION_FILE)

$(PROTOC_GEN_GOCOSMOS_VERSION_FILE): $(AKASH_DEVCACHE)
	@echo "installing protoc-gen-gocosmos $(PROTOC_GEN_GOCOSMOS_VERSION) ..."
	rm -f $(PROTOC_GEN_GOCOSMOS)
	(cd $(GO_ROOT); GOBIN=$(AKASH_DEVCACHE_BIN) $(GO) install github.com/cosmos/gogoproto/protoc-gen-gocosmos)
	rm -rf "$(dir $@)"
	mkdir -p "$(dir $@)"
	touch $@
$(PROTOC_GEN_GOCOSMOS): $(PROTOC_GEN_GOCOSMOS_VERSION_FILE) #modvendor

$(PROTOC_GEN_GOGO_VERSION_FILE): $(AKASH_DEVCACHE)
	@echo "installing protoc-gen-gogo $(PROTOC_GEN_GOCOSMOS_VERSION) ..."
	rm -f $(PROTOC_GEN_GOGO)
	(cd $(GO_ROOT); GOBIN=$(AKASH_DEVCACHE_BIN) $(GO) install github.com/cosmos/gogoproto/protoc-gen-gogo)
	rm -rf "$(dir $@)"
	mkdir -p "$(dir $@)"
	touch $@
$(PROTOC_GEN_GOGO): $(PROTOC_GEN_GOGO_VERSION_FILE) #modvendor

$(GOGOPROTO_VERSION_FILE): $(AKASH_DEVCACHE)
	@echo "installing gogoproto $(GOGOPROTO_VERSION) ..."
	rm -f $(GOGOPROTO)
	(cd $(GO_ROOT); GOBIN=$(AKASH_DEVCACHE_BIN) $(GO) install github.com/cosmos/gogoproto/gogoproto)
	rm -rf "$(dir $@)"
	mkdir -p "$(dir $@)"
	touch $@
$(GOGOPROTO): $(GOGOPROTO_VERSION_FILE)

$(PROTOC_GEN_GO_PULSAR_VERSION_FILE): $(AKASH_DEVCACHE)
	@echo "installing protoc-gen-go-pulsar $(PROTOC_GEN_GO_PULSAR_VERSION) ..."
	rm -f $(PROTOC_GEN_GO_PULSAR)
	(cd $(GO_ROOT); GOBIN=$(AKASH_DEVCACHE_BIN) $(GO) install github.com/cosmos/cosmos-proto/cmd/protoc-gen-go-pulsar@$(PROTOC_GEN_GO_PULSAR_VERSION))
	rm -rf "$(dir $@)"
	mkdir -p "$(dir $@)"
	touch $@
$(PROTOC_GEN_GO_PULSAR): $(PROTOC_GEN_GO_PULSAR_VERSION_FILE)

$(PROTOC_GEN_GO_VERSION_FILE): $(AKASH_DEVCACHE)
	@echo "installing protoc-gen-go $(PROTOC_GEN_GO_VERSION) ..."
	rm -f $(PROTOC_GEN_GO)
	(cd $(GO_ROOT); GOBIN=$(AKASH_DEVCACHE_BIN) $(GO) install google.golang.org/protobuf/cmd/protoc-gen-go@$(PROTOC_GEN_GO_VERSION))
	rm -rf "$(dir $@)"
	mkdir -p "$(dir $@)"
	touch $@
$(PROTOC_GEN_GO): $(PROTOC_GEN_GO_VERSION_FILE)

$(PROTOC_GEN_PROST_VERSION_FILE): $(AKASH_DEVCACHE)
	@echo "installing protoc-gen-prost $(PROTOC_GEN_PROST_VERSION) ..."
	rm -f $(PROTOC_GEN_PROST)
	cargo install protoc-gen-prost --root $(AKASH_DEVCACHE)
	rm -rf "$(dir $@)"
	mkdir -p "$(dir $@)"
	touch $@
$(PROTOC_GEN_PROST): $(PROTOC_GEN_PROST_VERSION_FILE)

$(PROTOC_GEN_DOC_VERSION_FILE): $(AKASH_DEVCACHE)
	@echo "installing protoc-gen-doc $(PROTOC_GEN_DOC_VERSION) ..."
	rm -f $(PROTOC_GEN_DOC)
	(cd $(GO_ROOT); GOBIN=$(AKASH_DEVCACHE_BIN) $(GO) install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@$(PROTOC_GEN_DOC_VERSION))
	rm -rf "$(dir $@)"
	mkdir -p "$(dir $@)"
	touch $@
$(PROTOC_GEN_DOC): $(PROTOC_GEN_DOC_VERSION_FILE)

$(PROTOC_GEN_GRPC_GATEWAY_VERSION_FILE): $(AKASH_DEVCACHE)
	@echo "Installing protoc-gen-grpc-gateway $(PROTOC_GEN_GRPC_GATEWAY_VERSION) ..."
	rm -f $(PROTOC_GEN_GRPC_GATEWAY)
	(cd $(GO_ROOT); GOBIN=$(AKASH_DEVCACHE_BIN) $(GO) install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway@$(PROTOC_GEN_GRPC_GATEWAY_VERSION))
	rm -rf "$(dir $@)"
	mkdir -p "$(dir $@)"
	touch $@
$(PROTOC_GEN_GRPC_GATEWAY): $(PROTOC_GEN_GRPC_GATEWAY_VERSION_FILE)

$(PROTOC_GEN_SWAGGER_VERSION_FILE): $(AKASH_DEVCACHE)
	@echo "Installing protoc-gen-swagger $(PROTOC_GEN_SWAGGER_VERSION) ..."
	rm -f $(PROTOC_GEN_SWAGGER)
	(cd $(GO_ROOT); GOBIN=$(AKASH_DEVCACHE_BIN) $(GO) install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger@$(PROTOC_GEN_SWAGGER_VERSION))
	rm -rf "$(dir $@)"
	mkdir -p "$(dir $@)"
	touch $@
$(PROTOC_GEN_SWAGGER): $(PROTOC_GEN_SWAGGER_VERSION_FILE)

$(MODVENDOR_VERSION_FILE): $(AKASH_DEVCACHE)
	@echo "installing modvendor $(MODVENDOR_VERSION) ..."
	rm -f $(MODVENDOR)
	(cd $(GO_ROOT); GOBIN=$(AKASH_DEVCACHE_BIN) $(GO) install github.com/goware/modvendor@$(MODVENDOR_VERSION))
	rm -rf "$(dir $@)"
	mkdir -p "$(dir $@)"
	touch $@
$(MODVENDOR): $(MODVENDOR_VERSION_FILE)

$(GIT_CHGLOG_VERSION_FILE): $(AKASH_DEVCACHE)
	@echo "installing git-chglog $(GIT_CHGLOG_VERSION) ..."
	rm -f $(GIT_CHGLOG)
	(cd $(GO_ROOT); GOBIN=$(AKASH_DEVCACHE_BIN) go install github.com/git-chglog/git-chglog/cmd/git-chglog@$(GIT_CHGLOG_VERSION))
	rm -rf "$(dir $@)"
	mkdir -p "$(dir $@)"
	touch $@
$(GIT_CHGLOG): $(GIT_CHGLOG_VERSION_FILE)

$(MOCKERY_VERSION_FILE): $(AKASH_DEVCACHE)
	@echo "installing mockery $(MOCKERY_VERSION) ..."
	rm -f $(MOCKERY)
	(cd $(GO_ROOT); GOBIN=$(AKASH_DEVCACHE_BIN) go install -ldflags '-s -w -X github.com/vektra/mockery/v2/pkg/config.SemVer=$(MOCKERY_VERSION)' github.com/vektra/mockery/v2@v$(MOCKERY_VERSION))
	rm -rf "$(dir $@)"
	mkdir -p "$(dir $@)"
	touch $@
$(MOCKERY): $(MOCKERY_VERSION_FILE)

$(GOLANGCI_LINT_VERSION_FILE): $(AKASH_DEVCACHE)
	@echo "installing golangci-lint $(GOLANGCI_LINT_VERSION) ..."
	rm -f $(GOLANGCI_LINT)
	curl -fsSL -o $(GOLANGCI_LINT).tar.gz "$(GOLANGCI_LINT_RELEASE_URL)/$(GOLANGCI_LINT_ASSET).tar.gz"
	$(call verify_sha256,$(GOLANGCI_LINT_RELEASE_URL)/$(GOLANGCI_LINT_CHECKSUMS),$(GOLANGCI_LINT_ASSET).tar.gz,$(GOLANGCI_LINT).tar.gz)
	tar -xzf $(GOLANGCI_LINT).tar.gz -C $(AKASH_DEVCACHE_BIN) --strip-components=1 $(GOLANGCI_LINT_ASSET)/golangci-lint
	rm -f $(GOLANGCI_LINT).tar.gz
	rm -rf "$(dir $@)"
	mkdir -p "$(dir $@)"
	touch $@
$(GOLANGCI_LINT): $(GOLANGCI_LINT_VERSION_FILE)

$(SEMVER_VERSION_FILE): $(AKASH_DEVCACHE)
	@echo "installing semver $(SEMVER_VERSION) ..."
	rm -f $(SEMVER)
	(cd $(GO_ROOT); GOBIN=$(AKASH_DEVCACHE_BIN) go install github.com/troian/semver/cmd/semver@$(SEMVER_VERSION))
	rm -rf "$(dir $@)"
	mkdir -p "$(dir $@)"
	touch $@
$(SEMVER): $(SEMVER_VERSION_FILE)

semver: $(SEMVER)

$(NPM):
ifeq (, $(shell which $(NPM) 2>/dev/null))
	$(error "npm installation required")
endif

$(SWAGGER_COMBINE): $(AKASH_DEVCACHE) $(NPM)
ifeq (, $(shell which swagger-combine 2>/dev/null))
	@echo "Installing swagger-combine..."
	npm install swagger-combine --prefix $(AKASH_DEVCACHE_NODE_MODULES)
	chmod +x $(SWAGGER_COMBINE)
else
	@echo "swagger-combine already installed; skipping..."
endif

$(AKASH_TS_NODE_MODULES): $(AKASH_TS_PACKAGE_FILE)
	@echo "installing node modules..."
	cd $(AKASH_TS_ROOT) && npm ci
	@echo "node modules installed."
