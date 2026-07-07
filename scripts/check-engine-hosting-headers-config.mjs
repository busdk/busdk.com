#!/usr/bin/env node

import { readFile } from "node:fs/promises";

const HEADERS_URL = new URL("../docs/_headers", import.meta.url);
const REQUIRED_ROUTES = ["/engine/", "/engine/*"];
const REQUIRED_HEADERS = new Map([
  ["Cross-Origin-Opener-Policy", "same-origin"],
  ["Cross-Origin-Embedder-Policy", "require-corp"],
]);

function assert(condition, message) {
  if (!condition) {
    throw new Error(message);
  }
}

function parseHeadersFile(text) {
  const routes = new Map();
  let currentRoute = "";
  for (const rawLine of text.split(/\r?\n/)) {
    const line = rawLine.trimEnd();
    if (!line.trim() || line.trimStart().startsWith("#")) {
      continue;
    }
    if (!rawLine.startsWith(" ") && !rawLine.startsWith("\t")) {
      currentRoute = line.trim();
      assert(currentRoute.startsWith("/"), `header route must start with /: ${currentRoute}`);
      assert(!routes.has(currentRoute), `duplicate header route: ${currentRoute}`);
      routes.set(currentRoute, new Map());
      continue;
    }
    assert(currentRoute, `header line without route: ${line.trim()}`);
    const trimmed = line.trim();
    const separator = trimmed.indexOf(":");
    assert(separator > 0, `header line must use Name: value syntax: ${trimmed}`);
    const name = trimmed.slice(0, separator).trim();
    const value = trimmed.slice(separator + 1).trim();
    routes.get(currentRoute).set(name, value);
  }
  return routes;
}

async function main() {
  if (process.argv.includes("--help") || process.argv.includes("-h")) {
    console.error("usage: check-engine-hosting-headers-config.mjs");
    return;
  }
  assert(process.argv.length === 2, "too many arguments");

  const text = await readFile(HEADERS_URL, "utf8");
  const routes = parseHeadersFile(text);
  assert(!routes.has("/*"), "Engine iframe headers must not be applied site-wide from this file");

  for (const route of REQUIRED_ROUTES) {
    const headers = routes.get(route);
    assert(headers, `docs/_headers must contain ${route}`);
    for (const [name, expected] of REQUIRED_HEADERS) {
      const actual = headers.get(name);
      assert(actual === expected, `${route} must send ${name}: ${expected}; got ${actual || "(missing)"}`);
    }
  }

  for (const route of routes.keys()) {
    assert(REQUIRED_ROUTES.includes(route), `unexpected Engine iframe header route: ${route}`);
  }

  console.log(`ok headers=${HEADERS_URL.pathname}`);
  console.log(`ok routes=${REQUIRED_ROUTES.join(",")}`);
  console.log("ok coop=same-origin");
  console.log("ok coep=require-corp");
}

main().catch((error) => {
  console.error(`check-engine-hosting-headers-config: error: ${error.message}`);
  process.exit(1);
});
