#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "$0")/.." && pwd)"
BUS_UI_DIR="${BUS_UI_DIR:-"$ROOT_DIR/../bus-ui"}"
DOC_ROOT="${GX_UI_DOC_ROOT:-"$ROOT_DIR/docs/gx-ui"}"

if ! command -v jq >/dev/null 2>&1; then
  printf "check-gx-ui-component-pages: jq is required\n" >&2
  exit 2
fi

if [ ! -d "$BUS_UI_DIR" ]; then
  printf "check-gx-ui-component-pages: missing bus-ui dir: %s\n" "$BUS_UI_DIR" >&2
  exit 2
fi

if [ ! -d "$DOC_ROOT" ]; then
  printf "check-gx-ui-component-pages: missing GX/UI docs dir: %s\n" "$DOC_ROOT" >&2
  exit 2
fi

TMP_DIR="$(mktemp -d "${TMPDIR:-/tmp}/gx-ui-component-pages.XXXXXX")"
trap 'rm -rf "$TMP_DIR"' EXIT

(
  cd "$BUS_UI_DIR"
  go run ./cmd/bus-ui catalog --format json
) >"$TMP_DIR/catalog.json"

find "$DOC_ROOT" -mindepth 2 -name index.html -print |
  sed "s#^$DOC_ROOT/##; s#/index.html\$##" |
  sort >"$TMP_DIR/pages.txt"

catalog_page() {
  case "$1" in
    apiurl-resolver) printf "bus-ui/portal/api-url-resolver" ;;
    ai-drop-controller) printf "bus-ui/ai-drop-controller" ;;
    browser-open) printf "bus-ui/tooling/browser-open" ;;
    callback-props) printf "gx/events" ;;
    callback-props-contract) printf "gx/events" ;;
    cli-runtime-flags) printf "bus-ui/tooling/cli-runtime-flags" ;;
    collection) printf "bus-ui/data" ;;
    component) printf "gx/components" ;;
    component-calls) printf "gx/components" ;;
    component-catalog) printf "bus-ui/tooling/component-catalog" ;;
    component-composition) printf "gx/components" ;;
    core-test-helpers) printf "gx/testing" ;;
    css-bundle) printf "bus-ui/tooling/css-bundle" ;;
    data-table) printf "dense-table" ;;
    evidence-url-resolver) printf "bus-ui/evidence/evidence-url-resolver" ;;
    expression-children) printf "gx/props-children" ;;
    go-wasm-frontend-runtime) printf "gx/runtime" ;;
    go-wasm-runtime) printf "gx/runtime" ;;
    image-gallery-component) printf "image-gallery" ;;
    navigation) printf "bus-ui/components/navigation/navigation" ;;
    node) printf "gx/nodes" ;;
    portal-host-context) printf "bus-ui/portal/host-context" ;;
    provider-error-component) printf "provider-error" ;;
    resource) printf "gx/effects" ;;
    runtime-config-component) printf "bus-ui/portal/runtime-config" ;;
    runtime-diagnostics) printf "gx/runtime" ;;
    session-component) printf "session" ;;
    shared-interfaces) printf "gx/nodes" ;;
    state) printf "bus-ui/components/status" ;;
    submit-state) printf "submit" ;;
    terminal-session-adapter) printf "terminal-adapters" ;;
    text-area) printf "textarea" ;;
    ui-artifact-metadata) printf "bus-ui/tooling/ui-artifact-metadata" ;;
    *) printf "%s" "$1" ;;
  esac
}

missing=0
while IFS=$'\t' read -r id name layer symbols; do
  page_id="$(catalog_page "$id")"
  if printf "%s" "$page_id" | grep -Fq "/"; then
    page_rel="$(awk -v page="$page_id" '$0 == page { print; exit }' "$TMP_DIR/pages.txt")"
  else
    page_rel="$(awk -F/ -v slug="$page_id" '$NF == slug { print; exit }' "$TMP_DIR/pages.txt")"
  fi
  if [ -z "$page_rel" ]; then
    if [ "$missing" -eq 0 ]; then
      printf "Missing GX/UI catalog reference pages:\n" >&2
    fi
    printf "  - %s (%s) expected page: %s; symbols: %s\n" "$id" "$name" "$page_id" "$symbols" >&2
    missing=1
    continue
  fi
  kind="$(jq -r --arg id "$id" '.groups[] | .entries[] | select(.id == $id) | .kind' "$TMP_DIR/catalog.json")"
  if [ "$kind" != "component" ]; then
    continue
  fi
  page_file="$DOC_ROOT/$page_rel/index.html"
  demo_id="$(basename "$page_rel")"
  if ! grep -Fq "data-bus-ui-demo=\"$demo_id\"" "$page_file"; then
    if [ "$missing" -eq 0 ]; then
      printf "Missing GX/UI manual component page demo hooks:\n" >&2
    fi
    printf "  - %s (%s) page %s missing data-bus-ui-demo=\"%s\"\n" "$id" "$name" "$page_rel/index.html" "$demo_id" >&2
    missing=1
  fi
  if ! grep -Fq "data-bus-ui-demo-loader" "$page_file"; then
    if [ "$missing" -eq 0 ]; then
      printf "Missing GX/UI manual component page demo hooks:\n" >&2
    fi
    printf "  - %s (%s) page %s missing shared demo loader script\n" "$id" "$name" "$page_rel/index.html" >&2
    missing=1
  fi
done < <(
  jq -r '
    .groups[] as $group
    | $group.entries[]
    | select(.status == "implemented")
    | [.id, .name, .layer, ((.symbols // []) | join(","))]
    | @tsv
  ' "$TMP_DIR/catalog.json"
)

if [ "$missing" -ne 0 ]; then
  exit 1
fi

printf "OK GX/UI reference pages cover implemented bus-ui catalog entries, with demo hooks for components\n"
