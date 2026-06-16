(function () {
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

  function setFallback(message) {
    document.querySelectorAll("[data-bus-ui-demo]").forEach(function (node) {
      if (node.childElementCount === 0 && node.textContent.trim() === "") {
        node.textContent = message;
        node.classList.add("bus-ui-demo-fallback");
      }
    });
  }

  function start() {
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
    var go = new Go();
    WebAssembly.instantiateStreaming(fetch(wasmURL), go.importObject)
      .then(function (result) {
        go.run(result.instance);
      })
      .catch(function () {
        return fetch(wasmURL)
          .then(function (response) {
            return response.arrayBuffer();
          })
          .then(function (bytes) {
            return WebAssembly.instantiate(bytes, go.importObject);
          })
          .then(function (result) {
            go.run(result.instance);
          });
      })
      .catch(function () {
        setFallback("Bus UI demo failed to load.");
      });
  }

  if (document.readyState === "loading") {
    document.addEventListener("DOMContentLoaded", start);
  } else {
    start();
  }
})();
