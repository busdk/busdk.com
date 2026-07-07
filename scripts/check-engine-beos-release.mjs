#!/usr/bin/env node

const FALLBACK_BASE_URL = "https://dev.hg.fi/beos/";
const DEFAULT_BASE_URL = process.env.BUS_ENGINE_BEOS_RELEASE_URL || FALLBACK_BASE_URL;
const DEFAULT_PROFILE_PATH = process.env.BUS_ENGINE_BEOS_PROFILE_PATH || "virtual-server/";
const DEFAULT_FETCH_TIMEOUT_MS = Number(process.env.BUS_ENGINE_CHECK_TIMEOUT_MS || "30000");
const REQUIRED_HEADERS = new Map([
  ["cross-origin-opener-policy", "same-origin"],
  ["cross-origin-embedder-policy", "require-corp"],
  ["cross-origin-resource-policy", "same-origin"],
]);
const REQUIRED_FILE_ROLES = [
  "browser-runtime",
  "index",
  "kernel",
  "qemu-javascript",
  "qemu-wasm",
  "rootfs",
];
const REQUIRED_PARAMETERS = [
  "display",
  "displayDevice",
  "kernel",
  "kernelAppend",
  "machine",
  "memory",
  "program",
  "rootfs",
  "targetArch",
  "wasm",
];
const REQUIRED_PROFILE_PAGE_SNIPPETS = [
  'data-default-display="wasm"',
  'id="canvas"',
  'data-qemu-display="true"',
  'id="details-terminal"',
  'id="serial-input"',
];
const REQUIRED_RUNTIME_SNIPPETS = [
  "crossOriginIsolated",
  'SharedArrayBuffer',
  "export function qemuArgs(config)",
  "displayDevice",
  "virtio-gpu-pci",
  "_qemu_wasm_chardev_write_pending",
];
const SMALL_FILE_BODY_SIZE_LIMIT = 1024 * 1024;

function usage() {
  console.error("usage: check-engine-beos-release.mjs [BASE_URL]");
  console.error("");
  console.error(`Default BASE_URL: ${DEFAULT_BASE_URL}`);
  console.error("Environment:");
  console.error("  BUS_ENGINE_BEOS_RELEASE_URL   release base URL");
  console.error("  BUS_ENGINE_BEOS_PROFILE_PATH  profile path under release base URL");
  console.error("  BUS_ENGINE_CHECK_TIMEOUT_MS   per-request timeout in milliseconds");
}

function assert(condition, message) {
  if (!condition) {
    throw new Error(message);
  }
}

assert(
  Number.isInteger(DEFAULT_FETCH_TIMEOUT_MS) && DEFAULT_FETCH_TIMEOUT_MS > 0,
  "BUS_ENGINE_CHECK_TIMEOUT_MS must be a positive integer",
);

function normalizeBaseUrl(raw) {
  const url = new URL(raw || DEFAULT_BASE_URL);
  if (!url.pathname.endsWith("/")) {
    url.pathname += "/";
  }
  return url;
}

function profileNameFromPath(path) {
  return path.split("/").filter(Boolean).at(-1) || "";
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

function checkIsolationHeaders(response, url) {
  for (const [header, expected] of REQUIRED_HEADERS) {
    const actual = response.headers.get(header);
    assert(
      actual === expected,
      `${url.href} must send ${header}: ${expected}; got ${actual || "(missing)"}`,
    );
  }
}

function checkManifest(manifest, manifestUrl) {
  assert(
    manifest.format === "bus-engine-os-browser-hosted-bundle-v1",
    `${manifestUrl.href} has unsupported format ${JSON.stringify(manifest.format)}`,
  );
  assert(Array.isArray(manifest.files), `${manifestUrl.href} must contain files[]`);
  assert(
    manifest.default_parameters && typeof manifest.default_parameters === "object",
    `${manifestUrl.href} must contain default_parameters`,
  );

  const roles = new Set(manifest.files.map((entry) => entry.role));
  for (const role of REQUIRED_FILE_ROLES) {
    assert(roles.has(role), `${manifestUrl.href} is missing files[] role ${role}`);
  }

  for (const item of manifest.files) {
    assert(item.path && typeof item.path === "string", `manifest file entry ${item.role} is missing path`);
    assert(item.sha256 && /^[a-f0-9]{64}$/.test(item.sha256), `manifest file ${item.path} is missing sha256`);
    assert(Number.isInteger(item.size_bytes) && item.size_bytes > 0, `manifest file ${item.path} is missing size_bytes`);
  }

  for (const parameter of REQUIRED_PARAMETERS) {
    assert(
      Object.prototype.hasOwnProperty.call(manifest.default_parameters, parameter),
      `${manifestUrl.href} is missing default_parameters.${parameter}`,
    );
  }

  const params = manifest.default_parameters;
  assert(params.targetArch === manifest.target_arch, "default_parameters.targetArch must match target_arch");
  assert(["none", "sdl", "wasm"].includes(params.display), "default_parameters.display must be none, sdl, or wasm");
  assert(typeof params.kernelAppend === "string" && params.kernelAppend.length > 0, "kernelAppend must be non-empty");
  assert(params.program === manifest.files.find((entry) => entry.role === "qemu-javascript").path, "program must point at qemu-javascript file");
  assert(params.wasm === manifest.files.find((entry) => entry.role === "qemu-wasm").path, "wasm must point at qemu-wasm file");
  assert(params.kernel === manifest.files.find((entry) => entry.role === "kernel").path, "kernel must point at kernel file");
  assert(params.rootfs === manifest.files.find((entry) => entry.role === "rootfs").path, "rootfs must point at rootfs file");
}

function checkProfileShape(manifest, expectedProfileName, manifestUrl) {
  const explicitProfile = manifest.profile || manifest.name || manifest.id;
  if (explicitProfile) {
    assert(
      explicitProfile === expectedProfileName,
      `${manifestUrl.href} declares profile ${explicitProfile}; expected ${expectedProfileName}`,
    );
  }

  if (!Array.isArray(manifest.profiles)) {
    return {
      kind: explicitProfile ? "single-explicit" : "path-implied",
      profiles: explicitProfile ? explicitProfile : "",
    };
  }

  assert(manifest.profiles.length > 0, `${manifestUrl.href} has empty profiles[]`);
  const profileNames = manifest.profiles.map((profile) => profile.id || profile.name || profile.profile).filter(Boolean);
  assert(
    profileNames.includes(expectedProfileName),
    `${manifestUrl.href} profiles[] must include ${expectedProfileName}; got ${profileNames.join(",") || "(none)"}`,
  );
  return { kind: "profiles-array", profiles: profileNames.join(",") };
}

function fileByRole(manifest, role) {
  const file = manifest.files.find((entry) => entry.role === role);
  assert(file, `manifest is missing files[] role ${role}`);
  return file;
}

function checkProfilePage(html, manifest, profileUrl) {
  for (const snippet of REQUIRED_PROFILE_PAGE_SNIPPETS) {
    assert(html.includes(snippet), `${profileUrl.href} is missing ${snippet}`);
  }

  const browserRuntime = fileByRole(manifest, "browser-runtime");
  assert(
    html.includes(`src="./${browserRuntime.path}"`) || html.includes(`src="${browserRuntime.path}"`),
    `${profileUrl.href} must load browser runtime ${browserRuntime.path}`,
  );
  assert(
    html.includes('params.set("display", "wasm")'),
    `${profileUrl.href} must default to wasm display mode`,
  );
  assert(
    html.includes('params.set("displayDevice", "virtio-gpu-pci")'),
    `${profileUrl.href} must default to virtio-gpu-pci display`,
  );
  assert(
    html.includes('params.set("kernelAppend",'),
    `${profileUrl.href} must pass kernelAppend from the generated profile page`,
  );
  const indexFile = fileByRole(manifest, "index");
  const htmlSize = new TextEncoder().encode(html).byteLength;
  assert(
    htmlSize === indexFile.size_bytes,
    `${profileUrl.href} body size must match manifest index size ${indexFile.size_bytes}; got ${htmlSize}`,
  );
}

async function checkManifestFileHeader(manifest, role, profileUrl) {
  const file = fileByRole(manifest, role);
  return checkManifestFile(file, profileUrl);
}

async function checkManifestFile(file, profileUrl) {
  const fileUrl = new URL(file.path, profileUrl);
  const response = await fetchRequired(fileUrl, { method: "HEAD" });
  checkIsolationHeaders(response, fileUrl);
  const rawLength = response.headers.get("content-length");
  if (rawLength === null) {
    assert(
      file.size_bytes <= SMALL_FILE_BODY_SIZE_LIMIT,
      `${fileUrl.href} must send content-length ${file.size_bytes}; got (missing)`,
    );
    const bodyResponse = await fetchRequired(fileUrl);
    checkIsolationHeaders(bodyResponse, fileUrl);
    const bytes = new Uint8Array(await bodyResponse.arrayBuffer());
    assert(
      bytes.byteLength === file.size_bytes,
      `${fileUrl.href} body size must match manifest size ${file.size_bytes}; got ${bytes.byteLength}`,
    );
    return fileUrl;
  }
  const length = Number(rawLength);
  assert(
    Number.isInteger(length) && length === file.size_bytes,
    `${fileUrl.href} must send content-length ${file.size_bytes}; got ${rawLength || "(missing)"}`,
  );
  return fileUrl;
}

function checkRuntimeScript(script, runtimeUrl) {
  for (const snippet of REQUIRED_RUNTIME_SNIPPETS) {
    assert(script.includes(snippet), `${runtimeUrl.href} is missing ${snippet}`);
  }
}

async function main() {
  if (process.argv.includes("--help") || process.argv.includes("-h")) {
    usage();
    return;
  }
  assert(process.argv.length <= 3, "too many arguments");

  const baseUrl = normalizeBaseUrl(process.argv[2]);
  const profileUrl = new URL(DEFAULT_PROFILE_PATH, baseUrl);
  if (!profileUrl.pathname.endsWith("/")) {
    profileUrl.pathname += "/";
  }
  const expectedProfileName = profileNameFromPath(profileUrl.pathname);
  assert(expectedProfileName, "profile path must include a profile name");
  const manifestUrl = new URL("browser-hosted-manifest.json", profileUrl);

  const baseResponse = await fetchRequired(baseUrl, { method: "HEAD" });
  checkIsolationHeaders(baseResponse, baseUrl);

  const profileResponse = await fetchRequired(profileUrl);
  checkIsolationHeaders(profileResponse, profileUrl);
  const profileHtml = await profileResponse.text();

  const manifestResponse = await fetchRequired(manifestUrl);
  checkIsolationHeaders(manifestResponse, manifestUrl);
  const manifest = await manifestResponse.json();
  checkManifest(manifest, manifestUrl);
  const profileShape = checkProfileShape(manifest, expectedProfileName, manifestUrl);
  checkProfilePage(profileHtml, manifest, profileUrl);
  const artifactUrls = new Map();
  for (const role of REQUIRED_FILE_ROLES) {
    artifactUrls.set(role, await checkManifestFileHeader(manifest, role, profileUrl));
  }
  const seenPaths = new Set([...artifactUrls.values()].map((url) => url.href));
  let checkedFiles = artifactUrls.size;
  for (const file of manifest.files) {
    const fileUrl = new URL(file.path, profileUrl);
    if (seenPaths.has(fileUrl.href)) {
      continue;
    }
    await checkManifestFile(file, profileUrl);
    seenPaths.add(fileUrl.href);
    checkedFiles += 1;
  }
  const runtimeUrl = artifactUrls.get("browser-runtime");
  const indexUrl = artifactUrls.get("index");
  const runtimeResponse = await fetchRequired(runtimeUrl);
  checkIsolationHeaders(runtimeResponse, runtimeUrl);
  checkRuntimeScript(await runtimeResponse.text(), runtimeUrl);

  console.log(`ok base=${baseUrl.href}`);
  console.log(`ok profile=${profileUrl.href}`);
  console.log(`ok manifest=${manifestUrl.href}`);
  console.log(`ok profile_index=${indexUrl.href}`);
  console.log(`ok browser_runtime=${runtimeUrl.href}`);
  console.log(`ok target_arch=${manifest.target_arch}`);
  console.log(`ok profile_path=${expectedProfileName}`);
  console.log(`ok display=${manifest.default_parameters.display}`);
  console.log(`ok display_device=${manifest.default_parameters.displayDevice}`);
  if (profileShape.kind === "path-implied") {
    console.log("info manifest_profile_shape=path-implied");
  } else {
    console.log(`ok manifest_profile_shape=${profileShape.kind}`);
    console.log(`ok manifest_profiles=${profileShape.profiles}`);
  }
  console.log(`ok files=${manifest.files.length}`);
  console.log(`ok published_files=${checkedFiles}`);
}

main().catch((error) => {
  console.error(`check-engine-beos-release: error: ${error.message}`);
  process.exit(1);
});
