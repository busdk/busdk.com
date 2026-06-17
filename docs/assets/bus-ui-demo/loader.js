(function () {
  var loaderPromiseKey = "__busUIDemoLoaderPromise";

  function currentScript() {
    return (
      document.currentScript ||
      document.querySelector("script[data-bus-ui-demo-loader]") ||
      document.querySelector("script[data-bus-ui-wasm]")
    );
  }

  function assetURL(script, name) {
    if (!script || !script.src) {
      return name;
    }
    return new URL(name, script.src).toString();
  }

  function ensureStylesheet(href) {
    if (!href || document.querySelector('link[data-bus-ui-demo-asset="css"]')) {
      return;
    }
    var link = document.createElement("link");
    link.rel = "stylesheet";
    link.href = href;
    link.setAttribute("data-bus-ui-demo-asset", "css");
    document.head.appendChild(link);
  }

  function demoNodes() {
    return Array.prototype.slice.call(document.querySelectorAll("[data-bus-ui-demo]"));
  }

  function sideNavNodes() {
    return Array.prototype.slice.call(document.querySelectorAll("aside.gx-side-nav[data-gx-ui-side-nav]"));
  }

  function topHeaderNodes() {
    return Array.prototype.slice.call(document.querySelectorAll("header.site-header[data-gx-ui-top-header]"));
  }

  function footerNodes() {
    return Array.prototype.slice.call(document.querySelectorAll("footer.site-footer[data-gx-ui-footer]"));
  }

  function mountNodes() {
    return demoNodes().concat(sideNavNodes(), topHeaderNodes(), footerNodes());
  }

  function setMountState(state) {
    demoNodes().forEach(function (node) {
      node.setAttribute("data-bus-ui-demo-state", state);
    });
    sideNavNodes().forEach(function (node) {
      node.setAttribute("data-gx-ui-side-nav-state", state);
    });
    topHeaderNodes().forEach(function (node) {
      node.setAttribute("data-gx-ui-top-header-state", state);
    });
    footerNodes().forEach(function (node) {
      node.setAttribute("data-gx-ui-footer-state", state);
    });
  }

  function setFallback(message) {
    demoNodes().forEach(function (node) {
      node.textContent = message;
      node.classList.add("bus-ui-demo-fallback");
      node.setAttribute("data-bus-ui-demo-state", "failed");
    });
    sideNavNodes().forEach(function (node) {
      node.textContent = message;
      node.classList.add("bus-ui-demo-fallback");
      node.setAttribute("data-gx-ui-side-nav-state", "failed");
    });
    topHeaderNodes().forEach(function (node) {
      node.setAttribute("data-gx-ui-top-header-state", "failed");
    });
    footerNodes().forEach(function (node) {
      node.setAttribute("data-gx-ui-footer-state", "failed");
    });
  }

  function loadDemoWASM(wasmURL, go) {
    return fetch(wasmURL)
      .then(function (response) {
        if (!response.ok) {
          throw new Error("Bus UI demo fetch failed.");
        }
        return response.arrayBuffer();
      })
      .then(function (bytes) {
        return WebAssembly.instantiate(bytes, go.importObject);
      });
  }

  function unsupportedWASMProtocol(url) {
    try {
      var parsed = new URL(url, window.location.href);
      return parsed.protocol === "file:";
    } catch (_) {
      return window.location.protocol === "file:";
    }
  }

  function start() {
    if (window[loaderPromiseKey]) {
      return window[loaderPromiseKey];
    }
    if (mountNodes().length === 0) {
      return;
    }
    var script = currentScript();
    ensureStylesheet(assetURL(script, "bus-ui.css"));
    var wasmURL = script && script.getAttribute("data-bus-ui-wasm");
    if (!wasmURL) {
      wasmURL = assetURL(script, "bus-ui-demo.wasm");
    }
    if (!wasmURL) {
      setFallback("Bus UI demo asset is missing.");
      return;
    }
    if (typeof Go !== "function") {
      setFallback("Bus UI demo runtime is missing.");
      return;
    }
    if (unsupportedWASMProtocol(wasmURL)) {
      setFallback("Bus UI demos need a local HTTP server to load WebAssembly.");
      return;
    }
    setMountState("loading");
    var go = new Go();
    window[loaderPromiseKey] = loadDemoWASM(wasmURL, go)
      .then(function (result) {
        setMountState("starting");
        go.run(result.instance);
      })
      .catch(function () {
        setFallback("Bus UI demo failed to load.");
      });
    return window[loaderPromiseKey];
  }

  if (document.readyState === "loading") {
    document.addEventListener("DOMContentLoaded", start);
  } else {
    start();
  }
})();
