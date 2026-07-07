#!/usr/bin/env node

import { readFile } from "node:fs/promises";

const ROOT = new URL("../", import.meta.url);
const MANIFEST_URL = new URL("docs/engine/wasm-virtual-server/manifest.json", ROOT);
const BOOT_URL = new URL("docs/engine/wasm-virtual-server/boot.js", ROOT);
const INDEX_URL = new URL("docs/engine/wasm-virtual-server/index.html", ROOT);

function assert(condition, message) {
  if (!condition) {
    throw new Error(message);
  }
}

function profileByName(manifest, name) {
  const profile = manifest.profiles.find((entry) => entry.name === name);
  assert(profile, `manifest is missing ${name} profile`);
  return profile;
}

function checkManifest(manifest) {
  assert(
    manifest.format === "bus-engine-os-qemu-wasm-preview-v1",
    `unexpected manifest format ${JSON.stringify(manifest.format)}`,
  );
  assert(manifest.guest.profile === "virtual-server", "guest profile must remain virtual-server");
  assert(manifest.runtime.display === "sdl", "runtime.display must remain sdl for graphical preview");
  assert(manifest.runtime.displayDevice === "stdvga", "runtime.displayDevice must remain stdvga");
  assert(manifest.runtime.keyboard === true, "runtime.keyboard must be true");
  assert(
    manifest.runtime.resolution &&
      manifest.runtime.resolution.width === 800 &&
      manifest.runtime.resolution.height === 600,
    "runtime.resolution must remain 800x600",
  );
  assert(Array.isArray(manifest.profiles), "manifest must contain profiles[]");

  const server = profileByName(manifest, "virtual-server");
  assert(server.status === "current-artifact", "virtual-server profile must be current-artifact");
  assert(server.guest.profile === "virtual-server", "virtual-server guest profile must be virtual-server");
  assert(server.runtime.display === "sdl", "virtual-server profile must use sdl display");
  assert(server.runtime.displayDevice === "stdvga", "virtual-server profile must use stdvga");
  assert(server.runtime.keyboard === true, "virtual-server profile must enable keyboard");
  assert(server.runtime.serialDiagnostics === true, "virtual-server profile must keep serial diagnostics");
  assert(server.runtime.resolution.width === 800 && server.runtime.resolution.height === 600, "virtual-server resolution must be 800x600");

  const desktop = profileByName(manifest, "virtual-desktop");
  assert(desktop.status === "planned-artifact", "virtual-desktop profile must remain planned-artifact");
  assert(desktop.guest.profile === "virtual-desktop", "virtual-desktop guest profile must be virtual-desktop");
  assert(desktop.runtime.display === "sdl", "virtual-desktop profile must declare sdl display");
  assert(desktop.runtime.displayDevice === "stdvga", "virtual-desktop profile must use stdvga");
  assert(desktop.runtime.keyboard === true, "virtual-desktop profile must enable keyboard");
  assert(desktop.runtime.serialDiagnostics === true, "virtual-desktop profile must keep serial diagnostics");
  assert(desktop.runtime.resolution.width === 1024 && desktop.runtime.resolution.height === 768, "virtual-desktop resolution must be 1024x768");
}

function checkIndex(html) {
  assert(html.includes('id="display"'), "iframe page must contain display canvas");
  assert(html.includes('class="fallback-image"'), "iframe page must keep screenshot fallback image");
  assert(html.includes('id="terminal"'), "iframe page must keep serial diagnostics terminal");
  assert(html.includes('id="profile-select"'), "iframe page must expose a manifest profile selector");
  assert(html.includes('id="profile-status"'), "iframe page must expose selected profile status");
  assert(html.includes('tabindex="0"'), "display canvas must be focusable for keyboard input");
}

function checkBootScript(script) {
  for (const snippet of [
    'runtimeValue("display"',
    'runtimeValue("displayDevice"',
    'runtimeValue("resolution"',
    "activeProfile.runtime",
    "populateProfiles()",
    "profileSelect.addEventListener(\"change\"",
    "activeProfile.status !== \"current-artifact\"",
    "guestValue(\"kernelAppend\"",
    "selectProfile()",
    "new URLSearchParams(window.location.search).get(\"profile\")",
    '"-display"',
    '"sdl,gl=off"',
    '"-vga"',
    '"std"',
    "configureDisplay()",
    "display.focus()",
    "window.crossOriginIsolated",
    "SharedArrayBuffer",
  ]) {
    assert(script.includes(snippet), `boot.js is missing ${snippet}`);
  }
}

async function main() {
  const [manifestText, bootScript, indexHtml] = await Promise.all([
    readFile(MANIFEST_URL, "utf8"),
    readFile(BOOT_URL, "utf8"),
    readFile(INDEX_URL, "utf8"),
  ]);
  checkManifest(JSON.parse(manifestText));
  checkBootScript(bootScript);
  checkIndex(indexHtml);

  console.log(`ok manifest=${MANIFEST_URL.pathname}`);
  console.log(`ok boot=${BOOT_URL.pathname}`);
  console.log(`ok index=${INDEX_URL.pathname}`);
  console.log("ok profiles=virtual-server,virtual-desktop");
}

main().catch((error) => {
  console.error(`check-engine-wasm-bundle-surface: error: ${error.message}`);
  process.exit(1);
});
