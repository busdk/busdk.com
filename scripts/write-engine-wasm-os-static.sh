#!/usr/bin/env sh
set -eu

ROOT=$(CDPATH= cd -- "$(dirname -- "$0")/.." && pwd)
QEMU_ARTIFACTS=${BUS_ENGINE_WASM_OS_QEMU_ARTIFACTS:-/tmp/qemu-wasm64-tci-artifacts-pipe2-final}
GUEST_ARTIFACTS=${BUS_ENGINE_WASM_OS_GUEST_ARTIFACTS:-/tmp/bus-engine-os-wasm-proof}
QEMU_SOURCE=${BUS_ENGINE_WASM_OS_QEMU_SOURCE:-"$ROOT/../../qemu"}
ENGINE_OS_DIR=${BUS_ENGINE_WASM_OS_ENGINE_OS_DIR:-"$ROOT/../bus-engine-os"}
ENGINE_OS_PROFILE=${BUS_ENGINE_WASM_OS_PROFILE:-virtual-server}
ENGINE_OS_TARGET_ARCH=${BUS_ENGINE_WASM_OS_TARGET_ARCH:-x86_64}
ENGINE_OS_SOURCES_CACHE=${BUS_ENGINE_WASM_OS_SOURCES_CACHE:-"$ENGINE_OS_DIR/build/sources"}
ALLOW_PENDING_LICENSES=${BUS_ENGINE_WASM_OS_ALLOW_PENDING_LICENSES:-0}
PUBLIC_PATH=${BUS_ENGINE_WASM_OS_PUBLIC_PATH:-/engine/wasm-virtual-server/}
VERBOSE=${BUS_ENGINE_WASM_OS_VERBOSE:-0}

usage() {
  cat >&2 <<'USAGE'
usage: write-engine-wasm-os-static.sh OUTPUT_DIR

Writes a complete static Bus Engine WASM OS bundle to OUTPUT_DIR. Publish that
directory under a dedicated URL path and embed its index.html with an iframe.

Optional inputs:
  BUS_ENGINE_WASM_OS_QEMU_ARTIFACTS   directory containing qemu-system-x86_64.js/.wasm
  BUS_ENGINE_WASM_OS_GUEST_ARTIFACTS  directory containing bzImage and rootfs.raw
  BUS_ENGINE_WASM_OS_QEMU_SOURCE      QEMU source tree with pc-bios firmware files
  BUS_ENGINE_WASM_OS_ENGINE_OS_DIR    bus-engine-os checkout used for package metadata
  BUS_ENGINE_WASM_OS_PROFILE          image profile for shipped package manifests
  BUS_ENGINE_WASM_OS_TARGET_ARCH      target architecture for shipped package manifests
  BUS_ENGINE_WASM_OS_SOURCES_CACHE    verified source cache for source-material files
  BUS_ENGINE_WASM_OS_ALLOW_PENDING_LICENSES=1
                                      allow development-preview output with pending review
  BUS_ENGINE_WASM_OS_PUBLIC_PATH      public URL path used in generated iframe.html
  BUS_ENGINE_WASM_OS_VERBOSE=1        print every copied file
USAGE
}

info() {
  echo "write-engine-wasm-os-static: info: $*"
}

debug() {
  if [ "$VERBOSE" = "1" ]; then
    echo "write-engine-wasm-os-static: debug: $*"
  fi
}

error() {
  echo "write-engine-wasm-os-static: error: $*" >&2
}

engine_os() {
  if [ -n "${BUS_ENGINE_WASM_OS_ENGINE_OS_BIN:-}" ]; then
    "$BUS_ENGINE_WASM_OS_ENGINE_OS_BIN" "$@"
    return
  fi
  if [ -d "$ENGINE_OS_DIR/cmd/bus-engine-os" ]; then
    (
      cd "$ENGINE_OS_DIR"
      go run -mod=readonly -tags netgo,osusergo ./cmd/bus-engine-os "$@"
    )
    return
  fi
  if command -v bus-engine-os >/dev/null 2>&1; then
    bus-engine-os "$@"
    return
  fi
  error "bus-engine-os source checkout or binary is required"
  exit 2
}

if [ "$#" -eq 1 ] && { [ "$1" = "--help" ] || [ "$1" = "-h" ]; }; then
  usage
  exit 0
fi

if [ "$#" -ne 1 ]; then
  usage
  exit 2
fi

OUT=$1
APP_SRC="$ROOT/docs/engine/wasm-virtual-server"
ARTIFACT_OUT="$OUT/artifacts"
PACKAGE_MANIFEST_LIST="$OUT/.bus-engine-os-package-manifests.txt"

case "$OUT" in
  ""|"/"|".")
    error "refusing unsafe output directory: $OUT"
    exit 2
    ;;
esac

mkdir -p "$ARTIFACT_OUT"

copy_file() {
  src=$1
  dst=$2

  if [ ! -r "$src" ]; then
    error "missing readable input: $src"
    exit 2
  fi

  cp "$src" "$dst"
  debug "wrote $dst"
}

copy_checked() {
  src=$1
  dst=$2
  want=$3

  if [ ! -r "$src" ]; then
    error "missing readable input: $src"
    exit 2
  fi

  got=$(sha256sum "$src" | awk '{print $1}')
  if [ "$got" != "$want" ]; then
    error "checksum mismatch: $src"
    echo "  want $want" >&2
    echo "  got  $got" >&2
    exit 2
  fi

  cp "$src" "$dst"
  debug "wrote $dst"
}

sed 's#src="../preview.png"#src="preview.png"#' \
  "$APP_SRC/index.html" > "$OUT/index.html"
debug "wrote $OUT/index.html"
copy_file "$APP_SRC/style.css" "$OUT/style.css"
copy_file "$APP_SRC/boot.js" "$OUT/boot.js"
copy_file "$APP_SRC/manifest.json" "$OUT/manifest.json"
copy_file "$ROOT/docs/engine/preview.png" "$OUT/preview.png"

copy_checked \
  "$QEMU_ARTIFACTS/qemu-system-x86_64.js" \
  "$ARTIFACT_OUT/qemu-system-x86_64.js" \
  "57ea9090acad40b3587c2648df233e1c3560cc20d10d1884b26226c3a23ce3f9"
copy_checked \
  "$QEMU_ARTIFACTS/qemu-system-x86_64.wasm" \
  "$ARTIFACT_OUT/qemu-system-x86_64.wasm" \
  "10ee623fc6ddb4c49a1b61bb97a0e77d6a06edfccf297d679a8ce1b0ef0287a8"
copy_checked \
  "$GUEST_ARTIFACTS/bzImage" \
  "$ARTIFACT_OUT/bzImage" \
  "b37cc4f821877ef34d468eeb9504fb6c0738114cff9bb18e030d88a2f4b76943"
copy_checked \
  "$GUEST_ARTIFACTS/rootfs.raw" \
  "$ARTIFACT_OUT/rootfs.raw" \
  "6d2222e0f5c8a1ff2d40808e682ffa5f0995e53c28a0f0af98ed0a96c9d49eae"
copy_checked \
  "$QEMU_SOURCE/pc-bios/qboot.rom" \
  "$ARTIFACT_OUT/qboot.rom" \
  "9b9dfc6c25740d6225625570d71cab6805cc9216e68c8932e343266daaeb8c4b"
copy_checked \
  "$QEMU_SOURCE/pc-bios/linuxboot_dma.bin" \
  "$ARTIFACT_OUT/linuxboot_dma.bin" \
  "9c49e255340c78fc12e54ed043462bca02fb7fca29b7cfab62ff88a5344b6950"
copy_checked \
  "$QEMU_SOURCE/pc-bios/bios-256k.bin" \
  "$ARTIFACT_OUT/bios-256k.bin" \
  "ae6f6aa973aaccc143f57aa960fb035fd9de4daee4ad0cd713322f8c259e7650"
copy_checked \
  "$QEMU_SOURCE/pc-bios/kvmvapic.bin" \
  "$ARTIFACT_OUT/kvmvapic.bin" \
  "cdf057a71b07e3b52b19cbe210bdefa59250d01a9810b960f7fe1f98eed95a27"
copy_checked \
  "$QEMU_SOURCE/pc-bios/vgabios.bin" \
  "$ARTIFACT_OUT/vgabios.bin" \
  "a7ea86bb06a58ff969fd0942e73b8eae00cb58e4c90bccbd900f6a3a01f54fbb"
copy_checked \
  "$QEMU_SOURCE/pc-bios/vgabios-stdvga.bin" \
  "$ARTIFACT_OUT/vgabios-stdvga.bin" \
  "e8fc9e55790dbe3cb31f019a3deb57206ba6c54f5e581adb2ab2677a9d391472"
copy_checked \
  "$QEMU_SOURCE/pc-bios/efi-virtio.rom" \
  "$ARTIFACT_OUT/efi-virtio.rom" \
  "26be36901db7f8181c306cc62bd74891d8646528965a78e40cceadba5dd7c8e7"

engine_os packages \
  --profile "$ENGINE_OS_PROFILE" \
  --arch "$ENGINE_OS_TARGET_ARCH" \
  --profiles-dir "$ENGINE_OS_DIR/config/profiles" \
  --recipes "$ENGINE_OS_DIR/packages" \
  --format json |
  python3 -c '
import json
import sys

packages = json.load(sys.stdin)
paths = []
for package in packages:
    if not package.get("runtime"):
        continue
    path = package.get("path", "")
    if path:
        paths.append(path)
for path in sorted(set(paths)):
    print(path)
' > "$PACKAGE_MANIFEST_LIST"

set -- artifact license-bundle \
  --out "$OUT" \
  --release-name "Bus Engine OS browser-hosted preview" \
  --sources-cache "$ENGINE_OS_SOURCES_CACHE" \
  --qemu-source-dir "$QEMU_SOURCE"

if [ "$ALLOW_PENDING_LICENSES" = "1" ]; then
  set -- "$@" --allow-pending
fi

while IFS= read -r manifest_path; do
  [ -n "$manifest_path" ] || continue
  case "$manifest_path" in
    /*) set -- "$@" --package-manifest "$manifest_path" ;;
    *) set -- "$@" --package-manifest "$ENGINE_OS_DIR/$manifest_path" ;;
  esac
done < "$PACKAGE_MANIFEST_LIST"

rm -f "$PACKAGE_MANIFEST_LIST"

engine_os "$@"

cat > "$OUT/README.txt" <<'NOTE'
This directory is a complete static Bus Engine WASM OS bundle.

Publish it under a path served with these browser isolation headers:

Cross-Origin-Opener-Policy: same-origin
Cross-Origin-Embedder-Policy: require-corp
Cross-Origin-Resource-Policy: same-origin

Publishing this directory distributes QEMU WebAssembly and Bus Engine OS guest
artifacts. This directory includes generated license indexes, notices, and
source-material payloads for the shipped artifacts. Source materials are copied
only for shipped packages and artifacts whose recorded licenses require source
delivery.
NOTE

case "$PUBLIC_PATH" in
  */) iframe_src="${PUBLIC_PATH}index.html" ;;
  *) iframe_src="${PUBLIC_PATH}/index.html" ;;
esac

cat > "$OUT/iframe.html" <<NOTE
<iframe
  src="$iframe_src"
  title="Live Bus Engine OS QEMU WebAssembly preview"
  loading="lazy"
></iframe>
NOTE

file_count=$(find "$OUT" -type f | wc -l | tr -d ' ')
size_bytes=$(du -sb "$OUT" | awk '{print $1}')

info "wrote static bundle: $OUT"
info "files=$file_count size_bytes=$size_bytes"
info "public_path=$PUBLIC_PATH"
info "publish this directory with COOP/COEP headers and use iframe.html as the embed template"
