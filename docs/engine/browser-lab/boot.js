const terminal = document.getElementById("terminal");
const statusEl = document.getElementById("status");
const bootButton = document.getElementById("boot-button");

let manifest;
let booting = false;

function writeLine(line = "") {
  terminal.textContent += `${line}\n`;
  terminal.scrollTop = terminal.scrollHeight;
}

function setStatus(text) {
  statusEl.textContent = text;
}

function artifact(role) {
  const item = manifest.artifacts.find((entry) => entry.role === role);
  if (!item) {
    throw new Error(`manifest is missing ${role}`);
  }
  return item;
}

function artifactUrl(role) {
  return new URL(artifact(role).path, window.location.href).href;
}

async function fetchBytes(role) {
  const item = artifact(role);
  const response = await fetch(new URL(item.path, window.location.href), {
    credentials: "same-origin",
  });
  if (!response.ok) {
    throw new Error(`${role} is unavailable: HTTP ${response.status}`);
  }
  const bytes = new Uint8Array(await response.arrayBuffer());
  if (item.sizeBytes && bytes.byteLength !== item.sizeBytes) {
    writeLine(`warn: ${role} size ${bytes.byteLength} differs from manifest ${item.sizeBytes}`);
  }
  return bytes;
}

async function hasArtifact(role) {
  const item = artifact(role);
  try {
    const response = await fetch(new URL(item.path, window.location.href), {
      credentials: "same-origin",
      method: "HEAD",
    });
    return response.ok;
  } catch {
    return false;
  }
}

function createParentPaths(module, path) {
  const parts = path.split("/").filter(Boolean).slice(0, -1);
  let current = "/";
  for (const part of parts) {
    module.FS_createPath(current, part, true, true);
    current = `${current === "/" ? "" : current}/${part}`;
  }
}

function mountFile(module, path, data) {
  createParentPaths(module, path);
  module.FS.writeFile(path, data);
}

function qemuArgs() {
  return [
    "-M",
    manifest.runtime.machine,
    "-cpu",
    manifest.runtime.cpu,
    "-m",
    manifest.runtime.memory,
    "-accel",
    "tcg,thread=single",
    "-nographic",
    "-no-reboot",
    "-serial",
    "mon:stdio",
    "-monitor",
    "none",
    "-kernel",
    "/kernel",
    "-append",
    manifest.guest.kernelAppend,
    "-drive",
    "file=/rootfs.raw,format=raw,if=virtio",
    "-nic",
    "none",
    "-L",
    "/firmware",
  ];
}

function browserProblem() {
  if (typeof WebAssembly === "undefined") {
    return "WebAssembly is not available in this browser.";
  }
  if (!window.crossOriginIsolated) {
    return "This page needs COOP and COEP headers before the live WebAssembly runtime can start.";
  }
  if (typeof SharedArrayBuffer === "undefined") {
    return "SharedArrayBuffer is unavailable. The browser must allow cross-origin isolated WebAssembly threads.";
  }
  return "";
}

async function checkReady() {
  const problem = browserProblem();
  if (problem) {
    setStatus(`${problem} Showing the captured boot preview.`);
    writeLine("");
    writeLine(`fallback: ${problem}`);
    return;
  }

  const required = [
    "qemu-js",
    "qemu-wasm",
    "kernel",
    "rootfs",
    "firmware-qboot",
    "firmware-linuxboot",
  ];
  const missing = [];
  for (const role of required) {
    if (!(await hasArtifact(role))) {
      missing.push(role);
    }
  }
  if (missing.length > 0) {
    setStatus("Live artifacts are not published in this static build. Showing the captured boot preview.");
    writeLine("");
    writeLine(`fallback: missing artifacts: ${missing.join(", ")}`);
    return;
  }

  bootButton.disabled = false;
  setStatus("Ready to boot Bus Engine OS in this browser.");
}

async function boot() {
  if (booting) {
    return;
  }
  booting = true;
  bootButton.disabled = true;
  terminal.textContent = "";
  setStatus("Loading QEMU/WASM and Bus Engine OS artifacts...");

  writeLine("bus-engine-browser-lab: loading kernel");
  const kernel = await fetchBytes("kernel");
  writeLine("bus-engine-browser-lab: loading root filesystem");
  const rootfs = await fetchBytes("rootfs");
  writeLine("bus-engine-browser-lab: loading firmware");
  const qboot = await fetchBytes("firmware-qboot");
  const linuxboot = await fetchBytes("firmware-linuxboot");
  writeLine("bus-engine-browser-lab: importing QEMU WebAssembly runtime");

  const qemuProgram = artifactUrl("qemu-js");
  const qemuWasm = artifactUrl("qemu-wasm");
  const moduleFactory = (await import(qemuProgram)).default;
  const expected = new Map([
    [manifest.guest.versionMarker, false],
    ...manifest.guest.expectedText.map((text) => [text, false]),
  ]);

  const emit = (line) => {
    writeLine(line);
    for (const text of expected.keys()) {
      if (line.includes(text)) {
        expected.set(text, true);
      }
    }
    if ([...expected.values()].every(Boolean)) {
      setStatus("Bus Engine OS boot evidence reached in browser.");
    }
  };

  setStatus("Starting QEMU/WASM...");
  await moduleFactory({
    arguments: qemuArgs(),
    locateFile(path) {
      if (path === "qemu-system-x86_64.wasm") {
        return qemuWasm;
      }
      return new URL(path, qemuProgram).href;
    },
    mainScriptUrlOrBlob: qemuProgram,
    preRun: [
      (module) => {
        mountFile(module, "/kernel", kernel);
        mountFile(module, "/rootfs.raw", rootfs);
        mountFile(module, "/firmware/qboot.rom", qboot);
        mountFile(module, "/firmware/linuxboot_dma.bin", linuxboot);
      },
    ],
    print: emit,
    printErr: emit,
  });
}

async function init() {
  try {
    const response = await fetch("manifest.json", { credentials: "same-origin" });
    if (!response.ok) {
      throw new Error(`manifest unavailable: HTTP ${response.status}`);
    }
    manifest = await response.json();
    await checkReady();
  } catch (error) {
    setStatus("Live preview setup failed. Showing the captured boot preview.");
    writeLine("");
    writeLine(`fallback: ${error && error.message ? error.message : String(error)}`);
  }
}

bootButton.addEventListener("click", () => {
  boot().catch((error) => {
    setStatus("Live boot failed. Showing terminal output and captured preview.");
    writeLine("");
    writeLine(error && error.stack ? error.stack : String(error));
  });
});

init();
