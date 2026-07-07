#!/usr/bin/env node

import { readFile } from "node:fs/promises";

const ROOT = new URL("../", import.meta.url);
const PLAN_URL = new URL("PLAN.md", ROOT);
const GOAL_URL = new URL("GOAL.md", ROOT);
const STATUS_URL = new URL("STATUS_UPDATE.md", ROOT);
const OPEN_PLAN_ITEMS = [
  "Update the Bus Engine OS iframe from terminal-only boot output to an",
  "Add a browser-hosted OS manifest shape that can describe `virtual-server`",
];
const REQUIRED_STATUS_SNIPPETS = [
  "iframe, landing, and manifest",
  "https://dev.hg.fi/beos/",
  "https://dev.hg.fi/beos/virtual-server/browser-hosted-manifest.json",
  "profile_path=virtual-server",
  "manifest_profile_shape=profiles-array",
  "generated_at=2026-07-07T14:29:04Z",
  "profile=virtual-server",
  "profile_id=virtual-server",
  "profile_name=Virtual Server",
  "profiles=[virtual-server]",
  "c9a173fe",
  "73746c8",
  "no-input artifact",
  "virtual-server explicit metadata blocker is closed",
  "display=wasm",
  "display_device=virtio-gpu-pci",
  "files=14",
  "published_files=14",
  "iframe_state=fallback-required",
  "Landing status:",
  "Public landing status:",
  "live QEMU/WASM release CTA",
  "landing=browser-hosted-release",
  "release_embedding=cross-origin-blocked-corp-same-origin",
  "manifest_link=not-yet-deployed",
  "BUS_ENGINE_REQUIRE_IFRAME_ELIGIBLE=1",
  "BUS_ENGINE_REQUIRE_MANIFEST_LINK=1",
  "make engine-beos-release-profile-gate",
  "virtual-server explicit profile metadata and serial controls are live",
  "id=\"copy-serial-log\"",
  "qemuWasmSendSerialText",
  "make engine-beos-public-deploy-gate",
  "Cross-Origin-Resource-Policy: cross-origin",
  "same-origin parent/proxy",
  "../bus-update/go.mod",
  "Website lane:",
  "Engine/BEO lane:",
];
const REQUIRED_GOAL_SNIPPETS = [
  "Supervise the `busdk.com` website iframe, landing, and manifest surface for",
  "Keep the Bus Engine landing copy aligned with the current browser release",
  "https://dev.hg.fi/beos/",
  "docs/_headers",
  "Do not edit Bus Engine OS, QEMU, or hosted `/beos` runtime artifacts here.",
  "make engine-beos-check",
  "make engine-status-update-check",
  "BUS_ENGINE_REQUIRE_IFRAME_ELIGIBLE=1",
  "BUS_ENGINE_REQUIRE_MANIFEST_LINK=1",
  "make engine-beos-public-deploy-gate",
  "Cross-Origin-Resource-Policy: cross-origin",
  "virtual-desktop artifact",
  "The website lane is complete only when the two open `PLAN.md` items are closed",
];

function assert(condition, message) {
  if (!condition) {
    throw new Error(message);
  }
}

function uncheckedItems(plan) {
  return plan
    .split(/\r?\n/)
    .filter((line) => line.startsWith("- [ ] "))
    .map((line) => line.slice("- [ ] ".length));
}

async function main() {
  if (process.argv.includes("--help") || process.argv.includes("-h")) {
    console.error("usage: check-engine-status-update.mjs");
    return;
  }
  assert(process.argv.length === 2, "too many arguments");

  const [plan, goal, status] = await Promise.all([
    readFile(PLAN_URL, "utf8"),
    readFile(GOAL_URL, "utf8"),
    readFile(STATUS_URL, "utf8"),
  ]);

  const openItems = uncheckedItems(plan);
  assert(openItems.length === OPEN_PLAN_ITEMS.length, `PLAN.md must have ${OPEN_PLAN_ITEMS.length} open items; got ${openItems.length}`);
  for (const expected of OPEN_PLAN_ITEMS) {
    assert(openItems.some((item) => item.startsWith(expected)), `PLAN.md is missing open item prefix: ${expected}`);
    assert(status.includes(expected), `STATUS_UPDATE.md is missing open item prefix: ${expected}`);
  }

  for (const snippet of REQUIRED_STATUS_SNIPPETS) {
    assert(status.includes(snippet), `STATUS_UPDATE.md is missing ${snippet}`);
  }
  for (const snippet of REQUIRED_GOAL_SNIPPETS) {
    assert(goal.includes(snippet), `GOAL.md is missing ${snippet}`);
  }

  assert(!goal.includes("BEGIN") && !goal.includes("PRIVATE KEY"), "GOAL.md must not contain key material markers");
  assert(!status.includes("BEGIN") && !status.includes("PRIVATE KEY"), "STATUS_UPDATE.md must not contain key material markers");

  console.log(`ok plan=${PLAN_URL.pathname}`);
  console.log(`ok goal=${GOAL_URL.pathname}`);
  console.log(`ok status=${STATUS_URL.pathname}`);
  console.log(`ok open_items=${openItems.length}`);
  console.log("ok live_status=bus-engine-os-browser-release");
}

main().catch((error) => {
  console.error(`check-engine-status-update: error: ${error.message}`);
  process.exit(1);
});
