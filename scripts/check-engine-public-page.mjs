#!/usr/bin/env node

const DEFAULT_ENGINE_URL = "https://busdk.com/engine/";
const FALLBACK_RELEASE_BASE_URL = "https://dev.hg.fi/beos/";
const DEFAULT_RELEASE_BASE_URL = process.env.BUS_ENGINE_BEOS_RELEASE_URL || FALLBACK_RELEASE_BASE_URL;
const DEFAULT_PROFILE_PATH = process.env.BUS_ENGINE_BEOS_PROFILE_PATH || "virtual-server/";
const DEFAULT_FETCH_TIMEOUT_MS = Number(process.env.BUS_ENGINE_CHECK_TIMEOUT_MS || "30000");
const REQUIRE_IFRAME_ELIGIBLE = process.env.BUS_ENGINE_REQUIRE_IFRAME_ELIGIBLE === "1";
const REQUIRE_MANIFEST_LINK = process.env.BUS_ENGINE_REQUIRE_MANIFEST_LINK === "1";

function assert(condition, message) {
  if (!condition) {
    throw new Error(message);
  }
}

assert(
  Number.isInteger(DEFAULT_FETCH_TIMEOUT_MS) && DEFAULT_FETCH_TIMEOUT_MS > 0,
  "BUS_ENGINE_CHECK_TIMEOUT_MS must be a positive integer",
);

function hasIsolationHeaders(response) {
  const coop = response.headers.get("cross-origin-opener-policy");
  const coep = response.headers.get("cross-origin-embedder-policy");
  return coop === "same-origin" && (coep === "require-corp" || coep === "credentialless");
}

function releaseEmbeddingProblem(parentUrl, releaseUrl, releaseResponse) {
  if (parentUrl.origin === releaseUrl.origin) {
    return "";
  }
  const corp = releaseResponse.headers.get("cross-origin-resource-policy");
  if (corp === "cross-origin") {
    return "";
  }
  return `${releaseUrl.href} sends cross-origin-resource-policy: ${corp || "(missing)"}; cross-origin iframe activation from ${parentUrl.origin} needs release CORP cross-origin or a same-origin parent/proxy`;
}

function releaseEmbeddingState(parentUrl, releaseUrl, releaseResponse) {
  if (parentUrl.origin === releaseUrl.origin) {
    return "same-origin";
  }
  const corp = releaseResponse.headers.get("cross-origin-resource-policy");
  return corp === "cross-origin" ? "cross-origin-allowed" : `cross-origin-blocked-corp-${corp || "missing"}`;
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

function usage() {
  console.error("usage: check-engine-public-page.mjs [ENGINE_URL]");
  console.error("");
  console.error(`Default ENGINE_URL: ${DEFAULT_ENGINE_URL}`);
  console.error("Environment:");
  console.error("  BUS_ENGINE_BEOS_RELEASE_URL   expected release base URL");
  console.error("  BUS_ENGINE_BEOS_PROFILE_PATH  expected profile path under release base URL");
  console.error("  BUS_ENGINE_CHECK_TIMEOUT_MS   per-request timeout in milliseconds");
  console.error("  BUS_ENGINE_REQUIRE_IFRAME_ELIGIBLE=1  fail unless ENGINE_URL sends COOP/COEP");
  console.error("  BUS_ENGINE_REQUIRE_MANIFEST_LINK=1    fail unless ENGINE_URL links profile manifest");
}

async function fetchRequired(url, options = {}) {
  const controller = new AbortController();
  const timeout = setTimeout(() => controller.abort(), DEFAULT_FETCH_TIMEOUT_MS);
  try {
    const response = await fetch(url, { ...options, signal: controller.signal });
    assert(response.ok, `${url.href} returned HTTP ${response.status}`);
    return response;
  } catch (error) {
    if (error && error.name === "AbortError") {
      throw new Error(`${url.href} timed out after ${DEFAULT_FETCH_TIMEOUT_MS}ms`);
    }
    throw error;
  } finally {
    clearTimeout(timeout);
  }
}

async function main() {
  if (process.argv.includes("--help") || process.argv.includes("-h")) {
    usage();
    return;
  }
  assert(process.argv.length <= 3, "too many arguments");

  const engineUrl = new URL(process.argv[2] || DEFAULT_ENGINE_URL);
  const releaseBaseUrl = normalizeBaseUrl(DEFAULT_RELEASE_BASE_URL);
  const releaseManifestUrl = profileManifestUrl(releaseBaseUrl);
  const response = await fetchRequired(engineUrl);

  const html = await response.text();
  assert(
    html.includes("Published browser-hosted release"),
    "public Engine landing page must title the browser-hosted release section",
  );
  assert(
    html.includes("The published Bus Engine OS"),
    "public Engine landing page must explain the published Bus Engine OS release",
  );
  assert(
    html.includes("Open live QEMU/WASM release"),
    "public Engine landing page must expose the live release CTA",
  );
  assert(html.includes("data-engine-live-preview"), "public Engine page is missing the live preview container");
  assert(
    html.includes(`data-src="${releaseBaseUrl.href}"`),
    `public Engine page iframe must target ${releaseBaseUrl.href}`,
  );
  assert(
    html.includes(`href="${releaseBaseUrl.href}"`),
    `public Engine page must expose a direct launch link to ${releaseBaseUrl.href}`,
  );
  assert(
    html.includes("window.crossOriginIsolated") && html.includes('typeof SharedArrayBuffer === "undefined"'),
    "public Engine page must guard iframe activation on cross-origin isolation and SharedArrayBuffer",
  );
  assert(
    html.includes("frame.src = frame.dataset.src"),
    "public Engine page must activate the iframe from data-src after the guard passes",
  );

  const iframeState = hasIsolationHeaders(response) ? "iframe-eligible" : "fallback-required";
  assert(
    !REQUIRE_IFRAME_ELIGIBLE || iframeState === "iframe-eligible",
    `${engineUrl.href} is not iframe-eligible: missing COOP/COEP parent-page headers`,
  );
  const releaseResponse = await fetchRequired(releaseBaseUrl, { method: "HEAD" });
  const embeddingProblem = releaseEmbeddingProblem(engineUrl, releaseBaseUrl, releaseResponse);
  assert(!REQUIRE_IFRAME_ELIGIBLE || !embeddingProblem, embeddingProblem);
  const embeddingState = releaseEmbeddingState(engineUrl, releaseBaseUrl, releaseResponse);
  const hasManifestLink = html.includes(releaseManifestUrl);
  assert(
    !REQUIRE_MANIFEST_LINK || hasManifestLink,
    `${engineUrl.href} does not link expected profile manifest ${releaseManifestUrl}`,
  );

  console.log(`ok engine_page=${engineUrl.href}`);
  console.log(`ok release=${releaseBaseUrl.href}`);
  console.log("ok landing=browser-hosted-release");
  console.log(`ok iframe_state=${iframeState}`);
  if (embeddingProblem) {
    console.log(`info release_embedding=${embeddingState}`);
  } else {
    console.log(`ok release_embedding=${embeddingState}`);
  }
  if (hasManifestLink) {
    console.log(`ok manifest_link=${releaseManifestUrl}`);
  } else {
    console.log(`info manifest_link=not-yet-deployed`);
  }
}

main().catch((error) => {
  console.error(`check-engine-public-page: error: ${error.message}`);
  process.exit(1);
});
