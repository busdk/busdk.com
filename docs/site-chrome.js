(function () {
  const scriptElement = document.currentScript;
  const topHeaderNodes = Array.prototype.slice.call(
    document.querySelectorAll("header.site-header[data-busdk-top-header]")
  );
  const footerNodes = Array.prototype.slice.call(
    document.querySelectorAll("footer.site-footer[data-busdk-footer]")
  );
  const sideNavNodes = Array.prototype.slice.call(
    document.querySelectorAll("aside.gx-side-nav[data-busdk-side-nav]")
  );
  const baseUrl = scriptElement && scriptElement.src
    ? new URL("./", scriptElement.src)
    : new URL("./", window.location.href);

  window.__busDKSiteBaseURL = baseUrl.href;

  function setFallback(message) {
    sideNavNodes.forEach(function (node) {
      node.textContent = message;
      node.classList.add("bus-ui-demo-fallback");
      node.setAttribute("data-busdk-side-nav-state", "failed");
    });
    topHeaderNodes.forEach(function (node) {
      node.setAttribute("data-busdk-top-header-state", "failed");
    });
    footerNodes.forEach(function (node) {
      node.setAttribute("data-busdk-footer-state", "failed");
    });
  }

  function ensureLoaderScript() {
    if (document.querySelector("script[data-bus-ui-demo-loader]")) {
      return;
    }

    const loaderScript = document.createElement("script");
    loaderScript.src = new URL("assets/bus-ui-demo/loader.js", baseUrl).href;
    loaderScript.defer = true;
    loaderScript.setAttribute("data-bus-ui-demo-loader", "");

    if (typeof Go === "function") {
      document.head.appendChild(loaderScript);
      return;
    }

    const wasmExecScript = document.createElement("script");
    wasmExecScript.src = new URL("assets/bus-ui-demo/wasm_exec.js", baseUrl).href;
    wasmExecScript.defer = true;
    wasmExecScript.onload = function () {
      document.head.appendChild(loaderScript);
    };
    wasmExecScript.onerror = function () {
      setFallback("BusDK navigation failed to load.");
    };
    document.head.appendChild(wasmExecScript);
  }

  if (topHeaderNodes.length > 0 || footerNodes.length > 0 || sideNavNodes.length > 0) {
    ensureLoaderScript();
  }
})();
