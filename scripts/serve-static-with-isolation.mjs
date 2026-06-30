#!/usr/bin/env node
import { createServer } from "node:http";
import { readFile } from "node:fs/promises";
import { extname, join, normalize, resolve, sep } from "node:path";

const root = resolve(process.argv[2] || "docs");
const host = process.env.HOST || "127.0.0.1";
const port = Number(process.env.PORT || "8015");

function contentType(path) {
  switch (extname(path)) {
    case ".css":
      return "text/css; charset=utf-8";
    case ".html":
      return "text/html; charset=utf-8";
    case ".js":
    case ".mjs":
      return "text/javascript; charset=utf-8";
    case ".json":
      return "application/json; charset=utf-8";
    case ".png":
      return "image/png";
    case ".jpg":
    case ".jpeg":
      return "image/jpeg";
    case ".wasm":
      return "application/wasm";
    case ".txt":
      return "text/plain; charset=utf-8";
    default:
      return "application/octet-stream";
  }
}

function routePath(pathname) {
  const decoded = decodeURIComponent(pathname);
  const clean = normalize(decoded).replace(/^(\.\.(\/|\\|$))+/, "");
  const route = decoded.endsWith("/") ? join(clean, "index.html") : clean;
  const file = resolve(join(root, route === "/" ? "index.html" : route));
  const rootPrefix = root.endsWith(sep) ? root : `${root}${sep}`;
  if (file !== root && !file.startsWith(rootPrefix)) {
    return null;
  }
  return file;
}

const server = createServer(async (request, response) => {
  const url = new URL(request.url || "/", `http://${host}:${port}`);
  const file = routePath(url.pathname);

  response.setHeader("Cross-Origin-Opener-Policy", "same-origin");
  response.setHeader("Cross-Origin-Embedder-Policy", "require-corp");
  response.setHeader("Cross-Origin-Resource-Policy", "same-origin");

  if (!file) {
    response.writeHead(404, { "Content-Type": "text/plain; charset=utf-8" });
    response.end("not found\n");
    return;
  }

  try {
    const data = await readFile(file);
    response.writeHead(200, { "Content-Type": contentType(file) });
    response.end(request.method === "HEAD" ? undefined : data);
  } catch (error) {
    response.writeHead(404, { "Content-Type": "text/plain; charset=utf-8" });
    response.end(`${error && error.message ? error.message : String(error)}\n`);
  }
});

server.listen(port, host, () => {
  console.log(`serving ${root} with COOP/COEP at http://${host}:${port}/`);
});
