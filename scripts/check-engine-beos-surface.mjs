#!/usr/bin/env node

import { readFile } from "node:fs/promises";

const FALLBACK_RELEASE_BASE_URL = "https://dev.hg.fi/beos/";
const DEFAULT_RELEASE_BASE_URL = process.env.BUS_ENGINE_BEOS_RELEASE_URL || FALLBACK_RELEASE_BASE_URL;
const DEFAULT_PROFILE_PATH = process.env.BUS_ENGINE_BEOS_PROFILE_PATH || "virtual-server/";
const ENGINE_INDEX = new URL("../docs/engine/index.html", import.meta.url);

function assert(condition, message) {
  if (!condition) {
    throw new Error(message);
  }
}

function countOccurrences(text, needle) {
  return text.split(needle).length - 1;
}

function normalizeBaseUrl(raw) {
  const url = new URL(raw);
  if (!url.pathname.endsWith("/")) {
    url.pathname += "/";
  }
  return url;
}

function profileManifestUrl(releaseBaseUrl) {
  const profileUrl = new URL(DEFAULT_PROFILE_PATH, releaseBaseUrl);
  if (!profileUrl.pathname.endsWith("/")) {
    profileUrl.pathname += "/";
  }
  return new URL("browser-hosted-manifest.json", profileUrl).href;
}

async function main() {
  if (process.argv.includes("--help") || process.argv.includes("-h")) {
    console.error("usage: check-engine-beos-surface.mjs");
    console.error("");
    console.error("Environment:");
    console.error("  BUS_ENGINE_BEOS_RELEASE_URL   expected release base URL");
    console.error("  BUS_ENGINE_BEOS_PROFILE_PATH  expected profile path under release base URL");
    return;
  }
  assert(process.argv.length === 2, "too many arguments");

  const releaseBaseUrl = normalizeBaseUrl(DEFAULT_RELEASE_BASE_URL);
  const releaseManifestUrl = profileManifestUrl(releaseBaseUrl);
  const html = await readFile(ENGINE_INDEX, "utf8");

  assert(html.includes("Published browser-hosted release"), "Engine landing page must title the browser release section");
  assert(
    html.includes("A real Linux server booting in a browser tab"),
    "Engine landing page must describe the browser-hosted Bus Engine OS release",
  );
  assert(
    html.includes("The published Bus Engine OS") && html.includes("profile manifest"),
    "Engine landing page must explain the published release and profile manifest",
  );
  assert(html.includes("data-engine-live-preview"), "Engine page is missing the live preview container");
  assert(
    html.includes(`data-src="${releaseBaseUrl.href}"`),
    `Engine page iframe must target ${releaseBaseUrl.href}`,
  );
  assert(
    html.includes(`href="${releaseBaseUrl.href}"`),
    `Engine page must expose a direct launch link to ${releaseBaseUrl.href}`,
  );
  assert(
    html.includes(`href="${releaseManifestUrl}"`),
    `Engine page must link the live virtual-server manifest at ${releaseManifestUrl}`,
  );
  assert(
    html.includes('allow="cross-origin-isolated; fullscreen"'),
    "Engine page iframe must request cross-origin-isolated and fullscreen permissions",
  );
  assert(
    html.includes("window.crossOriginIsolated") && html.includes('typeof SharedArrayBuffer === "undefined"'),
    "Engine page iframe loader must guard on cross-origin isolation and SharedArrayBuffer",
  );
  assert(
    html.includes("frame.src = frame.dataset.src") && html.includes("fallback.hidden = true"),
    "Engine page iframe loader must only activate the frame after the guard passes",
  );
  assert(
    countOccurrences(html, releaseManifestUrl) === 1,
    "Engine page should link the profile manifest once from the release description",
  );

  console.log(`ok engine_page=${ENGINE_INDEX.pathname}`);
  console.log(`ok release=${releaseBaseUrl.href}`);
  console.log(`ok manifest=${releaseManifestUrl}`);
}

main().catch((error) => {
  console.error(`check-engine-beos-surface: error: ${error.message}`);
  process.exit(1);
});
