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

catalog_slug() {
  case "$1" in
    data-table) printf "dense-table" ;;
    image-gallery-component) printf "image-gallery" ;;
    provider-error-component) printf "provider-error" ;;
    session-component) printf "session" ;;
    submit-state) printf "submit" ;;
    terminal-session-adapter) printf "terminal-adapters" ;;
    text-area) printf "textarea" ;;
    *) printf "%s" "$1" ;;
  esac
}

missing=0
while IFS=$'\t' read -r id name layer symbols; do
  slug="$(catalog_slug "$id")"
  page_rel="$(awk -F/ -v slug="$slug" '$NF == slug { print; exit }' "$TMP_DIR/pages.txt")"
  if [ -z "$page_rel" ]; then
    if [ "$missing" -eq 0 ]; then
      printf "Missing GX/UI manual component pages:\n" >&2
    fi
    printf "  - %s (%s) expected page slug: %s; symbols: %s\n" "$id" "$name" "$slug" "$symbols" >&2
    missing=1
    continue
  fi
  page_file="$DOC_ROOT/$page_rel/index.html"
  if ! grep -Fq "data-bus-ui-demo=\"$slug\"" "$page_file"; then
    if [ "$missing" -eq 0 ]; then
      printf "Missing GX/UI manual component page demo hooks:\n" >&2
    fi
    printf "  - %s (%s) page %s missing data-bus-ui-demo=\"%s\"\n" "$id" "$name" "$page_rel/index.html" "$slug" >&2
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
    | select(.kind == "component" and .status == "implemented")
    | [.id, .name, .layer, ((.symbols // []) | join(","))]
    | @tsv
  ' "$TMP_DIR/catalog.json"
)

if [ "$missing" -ne 0 ]; then
  exit 1
fi

printf "OK GX/UI manual component pages cover implemented bus-ui component catalog entries\n"
