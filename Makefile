.PHONY: browser-lab-artifacts bus-ui-demo-assets clean quality

# Generated artifacts and local caches for the SDD Jekyll site.
CLEAN_PATHS := \
	_site \
	.jekyll-cache \
	.jekyll-metadata \
	.sass-cache \
	.bundle \
	vendor/bundle \
	tmp

BUS_UI_DEMO_ASSET_DIR := docs/assets/bus-ui-demo
BUS_UI_DEMO_GO_CACHE := $(CURDIR)/tmp/go-build-cache

clean:
	rm -rf $(CLEAN_PATHS)

quality:
	./scripts/check-gx-ui-component-pages.sh
	cd demos/bus-ui && go test ./...

browser-lab-artifacts:
	./scripts/stage-engine-browser-lab-artifacts.sh

bus-ui-demo-assets:
	mkdir -p $(BUS_UI_DEMO_ASSET_DIR)
	cp "$$(go env GOROOT)/lib/wasm/wasm_exec.js" $(BUS_UI_DEMO_ASSET_DIR)/wasm_exec.js
	cp ../bus-ui/pkg/ui/assets/uikit.css $(BUS_UI_DEMO_ASSET_DIR)/bus-ui.css
	cd demos/bus-ui && GOCACHE=$(BUS_UI_DEMO_GO_CACHE) GOOS=js GOARCH=wasm go build -trimpath -buildvcs=false -ldflags="-buildid=" -o ../../$(BUS_UI_DEMO_ASSET_DIR)/bus-ui-demo.wasm .
