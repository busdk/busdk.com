.PHONY: bus-ui-demo-assets clean engine-beos-check engine-beos-public-deploy-gate engine-beos-public-page-check engine-beos-release-check engine-beos-release-profile-gate engine-hosting-headers-check engine-status-update-check engine-wasm-bundle-surface-check engine-wasm-os-static quality

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
ENGINE_WASM_OS_STATIC_DIR ?= tmp/engine-wasm-os-static
BUS_ENGINE_BEOS_RELEASE_URL ?= https://dev.hg.fi/beos/
BUS_ENGINE_BEOS_PROFILE_PATH ?= virtual-server/
BUS_ENGINE_PUBLIC_PAGE_URL ?= https://busdk.com/engine/
BUS_ENGINE_CHECK_TIMEOUT_MS ?= 30000
BUS_ENGINE_REQUIRE_EXPLICIT_PROFILE_METADATA ?= 0
BUS_ENGINE_REQUIRE_IFRAME_ELIGIBLE ?= 0
BUS_ENGINE_REQUIRE_MANIFEST_LINK ?= 0

clean:
	rm -rf $(CLEAN_PATHS)

quality:
	node scripts/check-engine-wasm-bundle-surface.mjs
	BUS_ENGINE_BEOS_RELEASE_URL="$(BUS_ENGINE_BEOS_RELEASE_URL)" BUS_ENGINE_BEOS_PROFILE_PATH="$(BUS_ENGINE_BEOS_PROFILE_PATH)" node scripts/check-engine-beos-surface.mjs
	node scripts/check-engine-hosting-headers-config.mjs
	node scripts/check-engine-status-update.mjs
	./scripts/check-gx-ui-component-pages.sh
	cd demos/bus-ui && go test ./...

engine-beos-release-check:
	BUS_ENGINE_BEOS_RELEASE_URL="$(BUS_ENGINE_BEOS_RELEASE_URL)" BUS_ENGINE_BEOS_PROFILE_PATH="$(BUS_ENGINE_BEOS_PROFILE_PATH)" BUS_ENGINE_CHECK_TIMEOUT_MS="$(BUS_ENGINE_CHECK_TIMEOUT_MS)" BUS_ENGINE_REQUIRE_EXPLICIT_PROFILE_METADATA="$(BUS_ENGINE_REQUIRE_EXPLICIT_PROFILE_METADATA)" node scripts/check-engine-beos-release.mjs "$(BUS_ENGINE_BEOS_RELEASE_URL)"

engine-beos-release-profile-gate:
	BUS_ENGINE_BEOS_RELEASE_URL="$(BUS_ENGINE_BEOS_RELEASE_URL)" BUS_ENGINE_BEOS_PROFILE_PATH="$(BUS_ENGINE_BEOS_PROFILE_PATH)" BUS_ENGINE_CHECK_TIMEOUT_MS="$(BUS_ENGINE_CHECK_TIMEOUT_MS)" BUS_ENGINE_REQUIRE_EXPLICIT_PROFILE_METADATA="1" node scripts/check-engine-beos-release.mjs "$(BUS_ENGINE_BEOS_RELEASE_URL)"

engine-beos-public-page-check:
	BUS_ENGINE_BEOS_RELEASE_URL="$(BUS_ENGINE_BEOS_RELEASE_URL)" BUS_ENGINE_BEOS_PROFILE_PATH="$(BUS_ENGINE_BEOS_PROFILE_PATH)" BUS_ENGINE_CHECK_TIMEOUT_MS="$(BUS_ENGINE_CHECK_TIMEOUT_MS)" BUS_ENGINE_REQUIRE_IFRAME_ELIGIBLE="$(BUS_ENGINE_REQUIRE_IFRAME_ELIGIBLE)" BUS_ENGINE_REQUIRE_MANIFEST_LINK="$(BUS_ENGINE_REQUIRE_MANIFEST_LINK)" node scripts/check-engine-public-page.mjs "$(BUS_ENGINE_PUBLIC_PAGE_URL)"

engine-beos-public-deploy-gate:
	BUS_ENGINE_BEOS_RELEASE_URL="$(BUS_ENGINE_BEOS_RELEASE_URL)" BUS_ENGINE_BEOS_PROFILE_PATH="$(BUS_ENGINE_BEOS_PROFILE_PATH)" BUS_ENGINE_CHECK_TIMEOUT_MS="$(BUS_ENGINE_CHECK_TIMEOUT_MS)" BUS_ENGINE_REQUIRE_IFRAME_ELIGIBLE="1" BUS_ENGINE_REQUIRE_MANIFEST_LINK="1" node scripts/check-engine-public-page.mjs "$(BUS_ENGINE_PUBLIC_PAGE_URL)"

engine-wasm-bundle-surface-check:
	node scripts/check-engine-wasm-bundle-surface.mjs

engine-hosting-headers-check:
	node scripts/check-engine-hosting-headers-config.mjs

engine-status-update-check:
	node scripts/check-engine-status-update.mjs

engine-beos-check:
	node scripts/check-engine-wasm-bundle-surface.mjs
	BUS_ENGINE_BEOS_RELEASE_URL="$(BUS_ENGINE_BEOS_RELEASE_URL)" BUS_ENGINE_BEOS_PROFILE_PATH="$(BUS_ENGINE_BEOS_PROFILE_PATH)" node scripts/check-engine-beos-surface.mjs
	node scripts/check-engine-hosting-headers-config.mjs
	node scripts/check-engine-status-update.mjs
	BUS_ENGINE_BEOS_RELEASE_URL="$(BUS_ENGINE_BEOS_RELEASE_URL)" BUS_ENGINE_BEOS_PROFILE_PATH="$(BUS_ENGINE_BEOS_PROFILE_PATH)" BUS_ENGINE_CHECK_TIMEOUT_MS="$(BUS_ENGINE_CHECK_TIMEOUT_MS)" BUS_ENGINE_REQUIRE_EXPLICIT_PROFILE_METADATA="$(BUS_ENGINE_REQUIRE_EXPLICIT_PROFILE_METADATA)" node scripts/check-engine-beos-release.mjs "$(BUS_ENGINE_BEOS_RELEASE_URL)"
	BUS_ENGINE_BEOS_RELEASE_URL="$(BUS_ENGINE_BEOS_RELEASE_URL)" BUS_ENGINE_BEOS_PROFILE_PATH="$(BUS_ENGINE_BEOS_PROFILE_PATH)" BUS_ENGINE_CHECK_TIMEOUT_MS="$(BUS_ENGINE_CHECK_TIMEOUT_MS)" BUS_ENGINE_REQUIRE_IFRAME_ELIGIBLE="$(BUS_ENGINE_REQUIRE_IFRAME_ELIGIBLE)" BUS_ENGINE_REQUIRE_MANIFEST_LINK="$(BUS_ENGINE_REQUIRE_MANIFEST_LINK)" node scripts/check-engine-public-page.mjs "$(BUS_ENGINE_PUBLIC_PAGE_URL)"

engine-wasm-os-static:
	./scripts/write-engine-wasm-os-static.sh "$(ENGINE_WASM_OS_STATIC_DIR)"

bus-ui-demo-assets:
	mkdir -p $(BUS_UI_DEMO_ASSET_DIR)
	cp "$$(go env GOROOT)/lib/wasm/wasm_exec.js" $(BUS_UI_DEMO_ASSET_DIR)/wasm_exec.js
	cp ../bus-ui/pkg/ui/assets/uikit.css $(BUS_UI_DEMO_ASSET_DIR)/bus-ui.css
	cd demos/bus-ui && GOCACHE=$(BUS_UI_DEMO_GO_CACHE) GOOS=js GOARCH=wasm go build -trimpath -buildvcs=false -ldflags="-buildid=" -o ../../$(BUS_UI_DEMO_ASSET_DIR)/bus-ui-demo.wasm .
