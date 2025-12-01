# Local Akash node setup for functional tests
# This makes chain-sdk self-sufficient for running tests

LOCAL_NODE_DIR          := $(AKASH_ROOT)/.local-node
LOCAL_NODE_HOME         := $(LOCAL_NODE_DIR)/home
LOCAL_NODE_INIT         := $(LOCAL_NODE_DIR)/.init
LOCAL_NODE_PID          := $(LOCAL_NODE_DIR)/node.pid

export AKASH_KEYRING_BACKEND    = test
export AKASH_GAS_ADJUSTMENT     = 2
export AKASH_CHAIN_ID           = local
export AKASH_YES                = true
export AKASH_GAS_PRICES         = 0.025uakt
export AKASH_GAS                = auto
export AKASH_NODE               = http://localhost:26657
export AKASH_HOME               = $(LOCAL_NODE_HOME)

LOCAL_KEY_OPTS          := --keyring-backend=$(AKASH_KEYRING_BACKEND)
LOCAL_GENESIS_PATH       := $(LOCAL_NODE_HOME)/config/genesis.json

LOCAL_CHAIN_MIN_DEPOSIT        := 10000000000000
LOCAL_CHAIN_ACCOUNT_DEPOSIT    := $(shell echo $$(($(LOCAL_CHAIN_MIN_DEPOSIT) * 10)))
LOCAL_CHAIN_VALIDATOR_DELEGATE := $(shell echo $$(($(LOCAL_CHAIN_MIN_DEPOSIT) / 2)))
LOCAL_CHAIN_TOKEN_DENOM        := uakt

LOCAL_KEY_NAMES := main provider validator other

LOCAL_GENESIS_ACCOUNTS := $(LOCAL_KEY_NAMES)

UNAME_OS_LOWER := $(shell uname -s | tr '[:upper:]' '[:lower:]')
UNAME_ARCH := $(shell uname -m | sed "s/x86_64/amd64/g")
AKASH_INSTALL_ARCH := $(UNAME_ARCH)
ifeq ($(UNAME_OS_LOWER), darwin)
	AKASH_INSTALL_ARCH := all
endif

LOCAL_AKASH_BIN := $(AKASH_DEVCACHE_BIN)/akash
LOCAL_AKASH_VERSION ?= $(shell curl -s https://api.github.com/repos/akash-network/node/releases/latest 2>/dev/null | grep -o '"tag_name": "[^"]*"' | head -1 | cut -d'"' -f4 | sed 's/^v//' || echo "0.28.0")
LOCAL_AKASH_VERSION_FILE := $(AKASH_DEVCACHE_VERSIONS)/akash/$(LOCAL_AKASH_VERSION)

$(LOCAL_NODE_DIR):
	mkdir -p $@

$(LOCAL_NODE_HOME): $(LOCAL_NODE_DIR)
	mkdir -p $@

.PHONY: local-node-install-akash
local-node-install-akash: $(AKASH_DEVCACHE)
	@if [ -f "$(LOCAL_AKASH_VERSION_FILE)" ] && [ -x "$(LOCAL_AKASH_BIN)" ]; then \
		echo "akash $(LOCAL_AKASH_VERSION) already installed"; \
		exit 0; \
	fi
	@echo "Installing akash $(LOCAL_AKASH_VERSION)..."
	@rm -f $(LOCAL_AKASH_BIN)
	@mkdir -p $(AKASH_DEVCACHE_VERSIONS)/akash
	@wget -q https://github.com/akash-network/node/releases/download/v$(LOCAL_AKASH_VERSION)/akash_$(UNAME_OS_LOWER)_$(AKASH_INSTALL_ARCH).zip -O $(AKASH_DEVCACHE)/akash.zip 2>/dev/null || \
		(echo "Warning: wget failed, trying curl..." && \
		 curl -sL https://github.com/akash-network/node/releases/download/v$(LOCAL_AKASH_VERSION)/akash_$(UNAME_OS_LOWER)_$(AKASH_INSTALL_ARCH).zip -o $(AKASH_DEVCACHE)/akash.zip || \
		 (echo "Error: Failed to download akash. Please install manually:" && \
		  echo "  curl -sSfL https://raw.githubusercontent.com/akash-network/node/main/install.sh | sh" && \
		  exit 1))
	@if [ -f "$(AKASH_DEVCACHE)/akash.zip" ]; then \
		unzip -p $(AKASH_DEVCACHE)/akash.zip akash > $(LOCAL_AKASH_BIN) && \
		chmod +x $(LOCAL_AKASH_BIN) && \
		rm $(AKASH_DEVCACHE)/akash.zip && \
		mkdir -p "$(dir $(LOCAL_AKASH_VERSION_FILE))" && \
		touch $(LOCAL_AKASH_VERSION_FILE) && \
		echo "akash $(LOCAL_AKASH_VERSION) installed successfully"; \
	fi

.PHONY: local-node-check-akash
local-node-check-akash: local-node-install-akash
	@if command -v akash >/dev/null 2>&1; then \
		echo "Using system akash: $$(command -v akash)"; \
	elif [ -x "$(LOCAL_AKASH_BIN)" ]; then \
		echo "Using local akash: $(LOCAL_AKASH_BIN)"; \
	else \
		echo "Error: Failed to install akash CLI"; \
		echo "Please install manually:"; \
		echo "  curl -sSfL https://raw.githubusercontent.com/akash-network/node/main/install.sh | sh"; \
		exit 1; \
	fi

.PHONY: local-node-init
local-node-init: local-node-check-akash $(LOCAL_NODE_HOME)
	@if [ -f "$(LOCAL_NODE_INIT)" ]; then \
		echo "Local node already initialized. Run 'make local-node-clean' to reset."; \
		exit 0; \
	fi
	@echo "Initializing local Akash node..."
	@AKASH_TO_USE=$$(command -v akash 2>/dev/null || echo "$(LOCAL_AKASH_BIN)"); \
	if [ ! -x "$$AKASH_TO_USE" ]; then \
		echo "Error: akash command not available"; \
		exit 1; \
	fi; \
	echo "Creating keys..."; \
	for keyname in main provider validator other; do \
		if ! $$AKASH_TO_USE $(LOCAL_KEY_OPTS) keys show "$$keyname" -a --home $(LOCAL_NODE_HOME) >/dev/null 2>&1; then \
			echo "  Adding key: $$keyname"; \
			echo "" | $$AKASH_TO_USE keys add "$$keyname" $(LOCAL_KEY_OPTS) --home $(LOCAL_NODE_HOME) >/dev/null 2>&1 || true; \
		else \
			echo "  Key already exists: $$keyname"; \
		fi; \
	done; \
	echo "Initializing genesis..."; \
	$$AKASH_TO_USE genesis init node0 --home $(LOCAL_NODE_HOME) >/dev/null 2>&1; \
	cp "$(LOCAL_GENESIS_PATH)" "$(LOCAL_GENESIS_PATH).orig"; \
	cat "$(LOCAL_GENESIS_PATH).orig" | \
		jq -M '.app_state.gov.params.voting_period = "30s"' | \
		jq -rM '(..|objects|select(has("denom"))).denom |= "$(LOCAL_CHAIN_TOKEN_DENOM)"' | \
		jq -rM '(..|objects|select(has("bond_denom"))).bond_denom |= "$(LOCAL_CHAIN_TOKEN_DENOM)"' | \
		jq -rM '(..|objects|select(has("mint_denom"))).mint_denom |= "$(LOCAL_CHAIN_TOKEN_DENOM)"' > \
		"$(LOCAL_GENESIS_PATH)"; \
	echo "Adding genesis accounts..."; \
	for key in $(LOCAL_GENESIS_ACCOUNTS); do \
		key_addr=$$($$AKASH_TO_USE $(LOCAL_KEY_OPTS) keys show "$$key" -a --home $(LOCAL_NODE_HOME) 2>/dev/null); \
		if [ -n "$$key_addr" ]; then \
			$$AKASH_TO_USE genesis add-account "$$key_addr" "$(LOCAL_CHAIN_MIN_DEPOSIT)$(LOCAL_CHAIN_TOKEN_DENOM)" --home $(LOCAL_NODE_HOME) >/dev/null 2>&1 || true; \
		fi; \
	done; \
	echo "Creating gentx..."; \
	$$AKASH_TO_USE genesis gentx validator "$(LOCAL_CHAIN_VALIDATOR_DELEGATE)$(LOCAL_CHAIN_TOKEN_DENOM)" \
		--min-self-delegation=1 \
		--gas=auto \
		--gas-prices=0.025uakt \
		--home $(LOCAL_NODE_HOME) >/dev/null 2>&1; \
	echo "Collecting genesis transactions..."; \
	$$AKASH_TO_USE genesis collect --home $(LOCAL_NODE_HOME) >/dev/null 2>&1; \
	echo "Validating genesis..."; \
	$$AKASH_TO_USE genesis validate --home $(LOCAL_NODE_HOME) >/dev/null 2>&1; \
	touch $(LOCAL_NODE_INIT); \
	echo "Local node initialized successfully!"

.PHONY: local-node-run
local-node-run: local-node-init
	@if [ -f "$(LOCAL_NODE_PID)" ] && kill -0 $$(cat $(LOCAL_NODE_PID)) 2>/dev/null; then \
		echo "Local node is already running (PID: $$(cat $(LOCAL_NODE_PID)))"; \
		exit 0; \
	fi
	@AKASH_TO_USE=$$(command -v akash 2>/dev/null || echo "$(LOCAL_AKASH_BIN)"); \
	if [ ! -x "$$AKASH_TO_USE" ]; then \
		echo "Error: akash command not available"; \
		exit 1; \
	fi; \
	echo "Starting local Akash node..."; \
	$$AKASH_TO_USE start --minimum-gas-prices=$(AKASH_GAS_PRICES) --home $(LOCAL_NODE_HOME) > $(LOCAL_NODE_DIR)/node.log 2>&1 & \
		echo $$! > $(LOCAL_NODE_PID); \
	echo "Waiting for node to start..."; \
	for i in $$(seq 1 30); do \
		if curl -s http://localhost:26657/status >/dev/null 2>&1; then \
			echo "Node is running (PID: $$(cat $(LOCAL_NODE_PID)))"; \
			exit 0; \
		fi; \
		sleep 1; \
	done; \
	echo "Error: Node failed to start. Check logs: $(LOCAL_NODE_DIR)/node.log"; \
	exit 1

.PHONY: local-node-stop
local-node-stop:
	@if [ -f "$(LOCAL_NODE_PID)" ]; then \
		pid=$$(cat $(LOCAL_NODE_PID)); \
		if kill -0 $$pid 2>/dev/null; then \
			echo "Stopping local node (PID: $$pid)..."; \
			kill $$pid || true; \
			sleep 2; \
			kill -9 $$pid 2>/dev/null || true; \
		fi; \
		rm -f $(LOCAL_NODE_PID); \
	fi

.PHONY: local-node-status
local-node-status:
	@if [ -f "$(LOCAL_NODE_PID)" ] && kill -0 $$(cat $(LOCAL_NODE_PID)) 2>/dev/null; then \
		echo "Node is running (PID: $$(cat $(LOCAL_NODE_PID)))"; \
		AKASH_TO_USE=$$(command -v akash 2>/dev/null || echo "$(LOCAL_AKASH_BIN)"); \
		$$AKASH_TO_USE status --node $(AKASH_NODE) 2>/dev/null || echo "Node is running but not responding"; \
	else \
		echo "Node is not running"; \
	fi

.PHONY: local-node-clean
local-node-clean: local-node-stop
	@echo "Cleaning local node data..."
	@rm -rf $(LOCAL_NODE_DIR)
	@echo "Local node data cleaned"

.PHONY: local-node-ready
local-node-ready: local-node-run
	@echo "Local node is ready for testing"
