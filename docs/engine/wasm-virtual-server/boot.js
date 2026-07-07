const terminal = document.getElementById("terminal");
const statusEl = document.getElementById("status");
const bootButton = document.getElementById("boot-button");
const display = document.getElementById("display");
const profileSelect = document.getElementById("profile-select");
const profileStatus = document.getElementById("profile-status");
const FIRMWARE_MOUNTS = new Map([
  ["firmware-qboot", "/firmware/qboot.rom"],
  ["firmware-linuxboot", "/firmware/linuxboot_dma.bin"],
  ["firmware-bios-256k", "/firmware/bios-256k.bin"],
  ["firmware-kvmvapic", "/firmware/kvmvapic.bin"],
  ["firmware-vgabios", "/firmware/vgabios.bin"],
  ["firmware-vgabios-stdvga", "/firmware/vgabios-stdvga.bin"],
  ["firmware-efi-virtio", "/firmware/efi-virtio.rom"],
]);

let manifest;
let activeProfile;
let booting = false;

function writeLine(line = "") {
  terminal.textContent += `${line}\n`;
  terminal.scrollTop = terminal.scrollHeight;
}

function setStatus(text) {
  statusEl.textContent = text;
}

function setProfileStatus(text) {
  profileStatus.textContent = text;
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

function runtimeValue(name, fallback) {
  if (activeProfile && activeProfile.runtime && activeProfile.runtime[name] !== undefined) {
    return activeProfile.runtime[name];
  }
  if (activeProfile && activeProfile[name] !== undefined) {
    return activeProfile[name];
  }
  return manifest.runtime && manifest.runtime[name] !== undefined
    ? manifest.runtime[name]
    : fallback;
}

function guestValue(name, fallback) {
  if (activeProfile && activeProfile.guest && activeProfile.guest[name] !== undefined) {
    return activeProfile.guest[name];
  }
  return manifest.guest && manifest.guest[name] !== undefined
    ? manifest.guest[name]
    : fallback;
}

function profileId(profile) {
  return profile.id || profile.name;
}

function profileLabel(profile) {
  return profile.label || profile.name || profile.id || "unnamed profile";
}

function profileRuntimeSummary(profile) {
  const runtime = (profile && profile.runtime) || {};
  const parts = [];
  if (runtime.display || runtime.displayDevice) {
    parts.push([runtime.display, runtime.displayDevice].filter(Boolean).join("/"));
  }
  if (
    runtime.resolution &&
    Number.isInteger(runtime.resolution.width) &&
    Number.isInteger(runtime.resolution.height)
  ) {
    parts.push(`${runtime.resolution.width}x${runtime.resolution.height}`);
  }
  if (runtime.serialDiagnostics) {
    parts.push("serial diagnostics");
  }
  return parts.join(" · ");
}

function updateProfileStatus() {
  if (!activeProfile) {
    setProfileStatus("No named profile in manifest.");
    return;
  }
  const summary = profileRuntimeSummary(activeProfile);
  const status = activeProfile.status || "current-artifact";
  setProfileStatus([status, summary].filter(Boolean).join(" · "));
}

function populateProfiles() {
  const profiles = Array.isArray(manifest.profiles) ? manifest.profiles : [];
  profileSelect.textContent = "";
  for (const profile of profiles) {
    const option = document.createElement("option");
    option.value = profileId(profile);
    option.textContent = profileLabel(profile);
    if (activeProfile && profileId(profile) === profileId(activeProfile)) {
      option.selected = true;
    }
    profileSelect.append(option);
  }
  profileSelect.disabled = profiles.length < 2;
  updateProfileStatus();
}

function selectProfile() {
  const profiles = Array.isArray(manifest.profiles) ? manifest.profiles : [];
  const requested = new URLSearchParams(window.location.search).get("profile") || manifest.guest.profile;
  const profile = profiles.find((entry) => profileId(entry) === requested);
  if (!profile) {
    if (profiles.length > 0) {
      throw new Error(`manifest is missing selected profile: ${requested}`);
    }
    activeProfile = null;
    return;
  }
  activeProfile = profile;
  writeLine(`bus-engine-os-preview: selected profile ${profileLabel(profile)}`);
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
  const displayMode = runtimeValue("display", "none");
  const displayDevice = runtimeValue("displayDevice", "default");
  const args = [
    "-M",
    manifest.runtime.machine,
    "-cpu",
    manifest.runtime.cpu,
    "-m",
    manifest.runtime.memory,
    "-accel",
    "tcg,thread=single",
  ];
  if (displayMode === "sdl") {
    args.push("-display", "sdl,gl=off");
    if (displayDevice === "stdvga") {
      args.push("-vga", "std");
    } else if (displayDevice === "virtio-vga") {
      args.push("-vga", "virtio");
    } else if (displayDevice === "virtio-gpu-pci") {
      args.push("-vga", "none", "-device", "virtio-gpu-pci");
    } else if (displayDevice === "none") {
      args.push("-vga", "none");
    }
  } else {
    args.push("-nographic");
  }
  args.push(
    "-no-reboot",
    "-serial",
    "mon:stdio",
    "-monitor",
    "none",
    "-kernel",
    "/kernel",
    "-append",
    guestValue("kernelAppend", "console=ttyS0 root=/dev/vda rw"),
    "-drive",
    "file=/rootfs.raw,format=raw,if=virtio",
    "-nic",
    "none",
    "-L",
    "/firmware",
  );
  return args;
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

  if (activeProfile && activeProfile.status !== "current-artifact") {
    const status = activeProfile.status || "unavailable";
    setStatus(`The selected ${profileLabel(activeProfile)} profile is ${status}. Showing the captured boot preview.`);
    writeLine("");
    writeLine(`fallback: selected profile is ${status}; boot is disabled for this artifact set`);
    return;
  }

  const missing = [];
  for (const item of manifest.artifacts) {
    if (!(await hasArtifact(item.role))) {
      missing.push(item.role);
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
  configureDisplay();
  setStatus("Loading QEMU/WASM and Bus Engine OS artifacts...");

  writeLine("bus-engine-os-preview: loading kernel");
  const kernel = await fetchBytes("kernel");
  writeLine("bus-engine-os-preview: loading root filesystem");
  const rootfs = await fetchBytes("rootfs");
  writeLine("bus-engine-os-preview: loading firmware");
  const firmware = [];
  for (const [role, path] of FIRMWARE_MOUNTS) {
    firmware.push({ data: await fetchBytes(role), path });
  }
  writeLine("bus-engine-os-preview: importing QEMU WebAssembly runtime");

  const qemuProgram = artifactUrl("qemu-js");
  const qemuWasm = artifactUrl("qemu-wasm");
  const moduleFactory = (await import(qemuProgram)).default;
  const expected = new Map([
    [guestValue("versionMarker", manifest.guest.versionMarker), false],
    ...guestValue("expectedText", manifest.guest.expectedText || []).map((text) => [text, false]),
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
    canvas: display,
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
        for (const file of firmware) {
          mountFile(module, file.path, file.data);
        }
      },
    ],
    print: emit,
    printErr: emit,
  });
}

function configureDisplay() {
  const displayMode = runtimeValue("display", "none");
  const resolution = runtimeValue("resolution", null);
  if (resolution && Number.isInteger(resolution.width) && Number.isInteger(resolution.height)) {
    display.width = resolution.width;
    display.height = resolution.height;
  }
  if (displayMode === "sdl") {
    display.hidden = false;
    display.setAttribute("aria-hidden", "false");
    display.focus();
    writeLine("bus-engine-os-preview: graphics enabled; click the display to focus keyboard input");
  } else {
    display.hidden = true;
    display.setAttribute("aria-hidden", "true");
  }
}

display.addEventListener("pointerdown", () => {
  display.focus();
});

profileSelect.addEventListener("change", () => {
  const nextUrl = new URL(window.location.href);
  nextUrl.searchParams.set("profile", profileSelect.value);
  window.location.href = nextUrl.href;
});

async function init() {
  try {
    const response = await fetch("manifest.json", { credentials: "same-origin" });
    if (!response.ok) {
      throw new Error(`manifest unavailable: HTTP ${response.status}`);
    }
    manifest = await response.json();
    selectProfile();
    populateProfiles();
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
