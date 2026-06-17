(function () {
  function toArray(list) {
    return Array.prototype.slice.call(list);
  }

  function tabsIn(root) {
    return toArray(root.querySelectorAll('[role="tab"]'));
  }

  function panelsIn(root) {
    return toArray(root.querySelectorAll('[role="tabpanel"]'));
  }

  function activeTab(tabs) {
    for (var i = 0; i < tabs.length; i += 1) {
      if (tabs[i].getAttribute("aria-selected") === "true") {
        return tabs[i];
      }
    }
    return tabs[0] || null;
  }

  function setActive(root, tabs, panels, tab, shouldFocus) {
    if (!tab) {
      return;
    }
    for (var i = 0; i < tabs.length; i += 1) {
      var selected = tabs[i] === tab;
      tabs[i].setAttribute("aria-selected", selected ? "true" : "false");
      tabs[i].tabIndex = selected ? 0 : -1;
    }
    for (var j = 0; j < panels.length; j += 1) {
      var panel = panels[j];
      var isVisible = panel.getAttribute("aria-labelledby") === tab.id;
      panel.hidden = !isVisible;
      panel.setAttribute("data-example-tab-state", isVisible ? "active" : "inactive");
    }
    root.setAttribute("data-example-tabs-state", "ready");
    if (shouldFocus) {
      tab.focus();
    }
  }

  function tabIndexOf(tabs, tab) {
    for (var i = 0; i < tabs.length; i += 1) {
      if (tabs[i] === tab) {
        return i;
      }
    }
    return -1;
  }

  function activateOffset(root, tabs, panels, currentTab, offset) {
    var index = tabIndexOf(tabs, currentTab);
    if (index < 0) {
      return;
    }
    var next = (index + offset + tabs.length) % tabs.length;
    setActive(root, tabs, panels, tabs[next], true);
  }

  function activateEdge(root, tabs, panels, first) {
    if (tabs.length === 0) {
      return;
    }
    setActive(root, tabs, panels, first ? tabs[0] : tabs[tabs.length - 1], true);
  }

  function bindTab(root, tabs, panels, tab) {
    tab.addEventListener("click", function () {
      setActive(root, tabs, panels, tab, false);
    });

    tab.addEventListener("keydown", function (event) {
      switch (event.key) {
        case "ArrowRight":
        case "ArrowDown":
          event.preventDefault();
          activateOffset(root, tabs, panels, tab, 1);
          break;
        case "ArrowLeft":
        case "ArrowUp":
          event.preventDefault();
          activateOffset(root, tabs, panels, tab, -1);
          break;
        case "Home":
          event.preventDefault();
          activateEdge(root, tabs, panels, true);
          break;
        case "End":
          event.preventDefault();
          activateEdge(root, tabs, panels, false);
          break;
      }
    });
  }

  function initRoot(root) {
    var tabs = tabsIn(root);
    var panels = panelsIn(root);
    if (tabs.length === 0 || panels.length === 0) {
      return;
    }
    var selected = activeTab(tabs);
    for (var i = 0; i < tabs.length; i += 1) {
      bindTab(root, tabs, panels, tabs[i]);
    }
    setActive(root, tabs, panels, selected, false);
  }

  function init() {
    var roots = toArray(document.querySelectorAll("[data-example-tabs]"));
    for (var i = 0; i < roots.length; i += 1) {
      initRoot(roots[i]);
    }
  }

  if (document.readyState === "loading") {
    document.addEventListener("DOMContentLoaded", init);
  } else {
    init();
  }
})();
