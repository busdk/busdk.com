#!/usr/bin/env sh
set -eu

ROOT=$(CDPATH= cd -- "$(dirname -- "$0")/.." && pwd)
OUT=${BUS_ENGINE_BROWSER_LAB_ARTIFACT_DIR:-"$ROOT/docs/engine/browser-lab/artifacts"}
QEMU_ARTIFACTS=${BUS_ENGINE_BROWSER_LAB_QEMU_ARTIFACTS:-/tmp/qemu-wasm64-tci-artifacts-pipe2-final}
GUEST_ARTIFACTS=${BUS_ENGINE_BROWSER_LAB_GUEST_ARTIFACTS:-/tmp/bus-engine-os-wasm-proof}
QEMU_SOURCE=${BUS_ENGINE_BROWSER_LAB_QEMU_SOURCE:-"$ROOT/../../qemu"}

mkdir -p "$OUT"

copy_checked() {
  src=$1
  dst=$2
  want=$3

  if [ ! -r "$src" ]; then
    echo "stage-engine-browser-lab-artifacts: missing readable input: $src" >&2
    exit 2
  fi

  got=$(sha256sum "$src" | awk '{print $1}')
  if [ "$got" != "$want" ]; then
    echo "stage-engine-browser-lab-artifacts: checksum mismatch: $src" >&2
    echo "  want $want" >&2
    echo "  got  $got" >&2
    exit 2
  fi

  cp "$src" "$dst"
  echo "stage-engine-browser-lab-artifacts: staged $(basename "$dst")"
}

copy_checked \
  "$QEMU_ARTIFACTS/qemu-system-x86_64.js" \
  "$OUT/qemu-system-x86_64.js" \
  "57ea9090acad40b3587c2648df233e1c3560cc20d10d1884b26226c3a23ce3f9"
copy_checked \
  "$QEMU_ARTIFACTS/qemu-system-x86_64.wasm" \
  "$OUT/qemu-system-x86_64.wasm" \
  "10ee623fc6ddb4c49a1b61bb97a0e77d6a06edfccf297d679a8ce1b0ef0287a8"
copy_checked \
  "$GUEST_ARTIFACTS/bzImage" \
  "$OUT/bzImage" \
  "b37cc4f821877ef34d468eeb9504fb6c0738114cff9bb18e030d88a2f4b76943"
copy_checked \
  "$GUEST_ARTIFACTS/rootfs.raw" \
  "$OUT/rootfs.raw" \
  "6d2222e0f5c8a1ff2d40808e682ffa5f0995e53c28a0f0af98ed0a96c9d49eae"
copy_checked \
  "$QEMU_SOURCE/pc-bios/qboot.rom" \
  "$OUT/qboot.rom" \
  "9b9dfc6c25740d6225625570d71cab6805cc9216e68c8932e343266daaeb8c4b"
copy_checked \
  "$QEMU_SOURCE/pc-bios/linuxboot_dma.bin" \
  "$OUT/linuxboot_dma.bin" \
  "9c49e255340c78fc12e54ed043462bca02fb7fca29b7cfab62ff88a5344b6950"
copy_checked \
  "$QEMU_SOURCE/pc-bios/bios-256k.bin" \
  "$OUT/bios-256k.bin" \
  "ae6f6aa973aaccc143f57aa960fb035fd9de4daee4ad0cd713322f8c259e7650"
copy_checked \
  "$QEMU_SOURCE/pc-bios/kvmvapic.bin" \
  "$OUT/kvmvapic.bin" \
  "cdf057a71b07e3b52b19cbe210bdefa59250d01a9810b960f7fe1f98eed95a27"
copy_checked \
  "$QEMU_SOURCE/pc-bios/vgabios.bin" \
  "$OUT/vgabios.bin" \
  "a7ea86bb06a58ff969fd0942e73b8eae00cb58e4c90bccbd900f6a3a01f54fbb"
copy_checked \
  "$QEMU_SOURCE/pc-bios/vgabios-stdvga.bin" \
  "$OUT/vgabios-stdvga.bin" \
  "e8fc9e55790dbe3cb31f019a3deb57206ba6c54f5e581adb2ab2677a9d391472"
copy_checked \
  "$QEMU_SOURCE/pc-bios/efi-virtio.rom" \
  "$OUT/efi-virtio.rom" \
  "26be36901db7f8181c306cc62bd74891d8646528965a78e40cceadba5dd7c8e7"

cat > "$OUT/README.txt" <<'NOTE'
These local artifacts are intentionally ignored by Git.

Publishing this directory distributes QEMU WebAssembly and Bus Engine OS guest
artifacts. The deployment must provide the corresponding source, build scripts,
license texts, and notices required by the included components' licenses.
NOTE

echo "stage-engine-browser-lab-artifacts: ready: $OUT"
