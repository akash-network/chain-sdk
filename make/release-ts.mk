.PHONY: release-ts
release-ts: $(AKASH_TS_NODE_MODULES) $(AKASH_TS_ROOT)/dist $(BUF) modvendor
	cd $(AKASH_TS_ROOT) && npm run release

$(AKASH_TS_ROOT)/dist: $(shell find $(AKASH_TS_ROOT)/src -type f)
	cd $(AKASH_TS_ROOT) && npm run build
