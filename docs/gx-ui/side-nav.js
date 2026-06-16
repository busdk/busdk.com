(function () {
  const scriptElement = document.currentScript;
  const baseUrl = scriptElement && scriptElement.src
    ? new URL("./", scriptElement.src)
    : new URL("docs/gx-ui/", window.location.href);

  const nav = {
    title: "GX and Bus UI",
    groups: [
      {
        heading: "Overview",
        entries: [
          { id: "index", href: "index.html", label: "Overview" },
          { id: "reference", href: "reference/index.html", label: "Reference" },
          { id: "modules", href: "modules/index.html", label: "Modules" },
        ],
      },
      {
        heading: "GX Framework",
        entries: [
          { id: "gx", href: "gx/index.html", label: "GX Framework" },
          { id: "gx/source-files", href: "gx/source-files/index.html", label: "Source files" },
          { id: "gx/components", href: "gx/components/index.html", label: "Component functions" },
          { id: "gx/props-children", href: "gx/props-children/index.html", label: "Props and children" },
          { id: "gx/generated-go", href: "gx/generated-go/index.html", label: "Generated Go" },
          { id: "gx/events", href: "gx/events/index.html", label: "Events" },
          { id: "gx/rendering", href: "gx/rendering/index.html", label: "Rendering" },
          { id: "gx/runtime", href: "gx/runtime/index.html", label: "Runtime bridges" },
          { id: "gx/effects", href: "gx/effects/index.html", label: "Effects" },
          { id: "gx/nodes", href: "gx/nodes/index.html", label: "Nodes and render tree" },
          { id: "gx/testing", href: "gx/testing/index.html", label: "Testing" },
        ],
      },
      {
        heading: "Bus UI Library",
        entries: [
          { id: "bus-ui", href: "bus-ui/index.html", label: "Bus UI Library" },
          { id: "bus-ui/components", href: "bus-ui/components/index.html", label: "Components" },
          {
            id: "bus-ui/components/shell",
            href: "bus-ui/components/shell/index.html",
            label: "Shells",
            children: [
              { id: "bus-ui/components/shell/app-shell", href: "bus-ui/components/shell/app-shell/index.html", label: "AppShell" },
              { id: "bus-ui/components/shell/page-shell", href: "bus-ui/components/shell/page-shell/index.html", label: "PageShell" },
              { id: "bus-ui/components/shell/sidebar-shell", href: "bus-ui/components/shell/sidebar-shell/index.html", label: "SidebarShell" },
              { id: "bus-ui/components/shell/sidebar-nav", href: "bus-ui/components/shell/sidebar-nav/index.html", label: "SidebarNav" },
              { id: "bus-ui/components/shell/shell-action-panel", href: "bus-ui/components/shell/shell-action-panel/index.html", label: "ShellActionPanel" },
            ],
          },
          {
            id: "bus-ui/components/navigation",
            href: "bus-ui/components/navigation/index.html",
            label: "Navigation",
            children: [
              { id: "bus-ui/components/navigation/menu", href: "bus-ui/components/navigation/menu/index.html", label: "Menu" },
              { id: "bus-ui/components/navigation/tabs", href: "bus-ui/components/navigation/tabs/index.html", label: "Tabs" },
              { id: "bus-ui/components/navigation/navigation", href: "bus-ui/components/navigation/navigation/index.html", label: "Navigation" },
            ],
          },
          {
            id: "bus-ui/components/action",
            href: "bus-ui/components/action/index.html",
            label: "Actions",
            children: [
              { id: "bus-ui/components/action/button", href: "bus-ui/components/action/button/index.html", label: "Button" },
              { id: "bus-ui/components/action/link-button", href: "bus-ui/components/action/link-button/index.html", label: "LinkButton" },
              { id: "bus-ui/components/action/icon-button", href: "bus-ui/components/action/icon-button/index.html", label: "IconButton" },
              { id: "bus-ui/components/action/event-bar", href: "bus-ui/components/action/event-bar/index.html", label: "EventBar" },
            ],
          },
          {
            id: "bus-ui/components/surface",
            href: "bus-ui/components/surface/index.html",
            label: "Surfaces",
            children: [
              { id: "bus-ui/components/surface/panel", href: "bus-ui/components/surface/panel/index.html", label: "Panel" },
              { id: "bus-ui/components/surface/surface-card", href: "bus-ui/components/surface/surface-card/index.html", label: "SurfaceCard" },
              { id: "bus-ui/components/surface/metric-card", href: "bus-ui/components/surface/metric-card/index.html", label: "MetricCard" },
            ],
          },
          {
            id: "bus-ui/components/status",
            href: "bus-ui/components/status/index.html",
            label: "Status",
            children: [
              { id: "bus-ui/components/status/status-pill", href: "bus-ui/components/status/status-pill/index.html", label: "StatusPill" },
              { id: "bus-ui/components/status/empty-state", href: "bus-ui/components/status/empty-state/index.html", label: "EmptyState" },
              { id: "bus-ui/components/status/loading-state", href: "bus-ui/components/status/loading-state/index.html", label: "LoadingState" },
              { id: "bus-ui/components/status/result-panel", href: "bus-ui/components/status/result-panel/index.html", label: "ResultPanel" },
              { id: "bus-ui/components/status/error-banner", href: "bus-ui/components/status/error-banner/index.html", label: "ErrorBanner" },
            ],
          },
          {
            id: "bus-ui/forms",
            href: "bus-ui/forms/index.html",
            label: "Forms",
            children: [
              { id: "bus-ui/forms/form", href: "bus-ui/forms/form/index.html", label: "Form" },
              { id: "bus-ui/forms/field", href: "bus-ui/forms/field/index.html", label: "Field" },
              { id: "bus-ui/forms/input", href: "bus-ui/forms/input/index.html", label: "Input" },
              { id: "bus-ui/forms/text-input", href: "bus-ui/forms/text-input/index.html", label: "TextInput" },
              { id: "bus-ui/forms/password-input", href: "bus-ui/forms/password-input/index.html", label: "PasswordInput" },
              { id: "bus-ui/forms/date-input", href: "bus-ui/forms/date-input/index.html", label: "DateInput" },
              { id: "bus-ui/forms/textarea", href: "bus-ui/forms/textarea/index.html", label: "TextArea" },
              { id: "bus-ui/forms/select", href: "bus-ui/forms/select/index.html", label: "Select" },
              { id: "bus-ui/forms/submit", href: "bus-ui/forms/submit/index.html", label: "SubmitControl" },
              { id: "bus-ui/forms/filter-toolbar", href: "bus-ui/forms/filter-toolbar/index.html", label: "FilterToolbar" },
              { id: "bus-ui/forms/file-input", href: "bus-ui/forms/file-input/index.html", label: "FileInput" },
              { id: "bus-ui/forms/drop-zone", href: "bus-ui/forms/drop-zone/index.html", label: "DropZone" },
              { id: "bus-ui/forms/credential-login-card", href: "bus-ui/forms/credential-login-card/index.html", label: "CredentialLoginCard" },
            ],
          },
          {
            id: "bus-ui/data",
            href: "bus-ui/data/index.html",
            label: "Data display",
            children: [
              { id: "bus-ui/data/dense-table", href: "bus-ui/data/dense-table/index.html", label: "DenseTable" },
              { id: "bus-ui/data/text-table", href: "bus-ui/data/text-table/index.html", label: "TextTable" },
              { id: "bus-ui/data/record-list", href: "bus-ui/data/record-list/index.html", label: "RecordList" },
              { id: "bus-ui/data/summary-item", href: "bus-ui/data/summary-item/index.html", label: "SummaryItem" },
              { id: "bus-ui/data/projection-detail", href: "bus-ui/data/projection-detail/index.html", label: "ProjectionDetail" },
              { id: "bus-ui/data/provider-error", href: "bus-ui/data/provider-error/index.html", label: "ProviderError" },
              { id: "bus-ui/data/timeline", href: "bus-ui/data/timeline/index.html", label: "Timeline" },
            ],
          },
          {
            id: "bus-ui/evidence",
            href: "bus-ui/evidence/index.html",
            label: "Evidence and files",
            children: [
              { id: "bus-ui/evidence/evidence-link", href: "bus-ui/evidence/evidence-link/index.html", label: "EvidenceLink" },
              { id: "bus-ui/evidence/evidence-preview", href: "bus-ui/evidence/evidence-preview/index.html", label: "EvidencePreview" },
              { id: "bus-ui/evidence/image-gallery", href: "bus-ui/evidence/image-gallery/index.html", label: "ImageGallery" },
            ],
          },
          {
            id: "bus-ui/assistant",
            href: "bus-ui/assistant/index.html",
            label: "Assistant",
            children: [
              { id: "bus-ui/assistant-shell", href: "bus-ui/assistant-shell/index.html", label: "AssistantShell" },
              { id: "bus-ui/ai-panel", href: "bus-ui/ai-panel/index.html", label: "AIPanel" },
              { id: "bus-ui/ai-composer", href: "bus-ui/ai-composer/index.html", label: "AIComposer" },
              { id: "bus-ui/ai-model-select", href: "bus-ui/ai-model-select/index.html", label: "AIModelSelect" },
              { id: "bus-ui/ai-approvals", href: "bus-ui/ai-approvals/index.html", label: "AIApprovals" },
              { id: "bus-ui/ai-review-status", href: "bus-ui/ai-review-status/index.html", label: "AIReviewStatus" },
              { id: "bus-ui/ai-thread-isolation", href: "bus-ui/ai-thread-isolation/index.html", label: "AIThreadIsolation" },
              { id: "bus-ui/ai-thread-list", href: "bus-ui/ai-thread-list/index.html", label: "AIThreadList" },
              { id: "bus-ui/ai-message", href: "bus-ui/ai-message/index.html", label: "AIMessage" },
              { id: "bus-ui/ai-markdown", href: "bus-ui/ai-markdown/index.html", label: "AIMarkdown" },
              { id: "bus-ui/ai-attachment-list", href: "bus-ui/ai-attachment-list/index.html", label: "AIAttachmentList" },
            ],
          },
          {
            id: "bus-ui/terminal",
            href: "bus-ui/terminal/index.html",
            label: "Terminal",
            children: [
              { id: "bus-ui/terminal-session-panel", href: "bus-ui/terminal-session-panel/index.html", label: "TerminalSessionPanel" },
              { id: "bus-ui/terminal-output-view", href: "bus-ui/terminal-output-view/index.html", label: "TerminalOutputView" },
              { id: "bus-ui/terminal-input-box", href: "bus-ui/terminal-input-box/index.html", label: "TerminalInputBox" },
              { id: "bus-ui/terminal-approval-prompt", href: "bus-ui/terminal-approval-prompt/index.html", label: "TerminalApprovalPrompt" },
              { id: "bus-ui/terminal-adapters", href: "bus-ui/terminal-adapters/index.html", label: "TerminalAdapters" },
            ],
          },
          { id: "bus-ui/portal", href: "bus-ui/portal/index.html", label: "Portal integration" },
          { id: "bus-ui/assistant-terminal", href: "bus-ui/assistant-terminal/index.html", label: "Assistant and Terminal split" },
        ],
      },
      {
        heading: "Tutorials",
        entries: [
          { id: "authoring", href: "authoring/index.html", label: "Authoring tutorial" },
          { id: "runtime", href: "runtime/index.html", label: "Runtime and testing" },
          { id: "components", href: "components/index.html", label: "Component tutorial" },
          { id: "surfaces", href: "surfaces/index.html", label: "Product surfaces" },
        ],
      },
    ],
  };

  function linkClass(depth) {
    if (depth === 1) {
      return "gx-side-nav-child";
    }
    if (depth >= 2) {
      return "gx-side-nav-grandchild";
    }
    return "";
  }

  function renderEntries(parent, entries, currentId, baseUrl, depth, currentState) {
    entries.forEach(function (entry) {
      const link = document.createElement("a");
      const className = linkClass(depth);
      if (className) {
        link.className = className;
      }
      link.href = new URL(entry.href, baseUrl).href;
      link.textContent = entry.label;
      if (entry.id === currentId) {
        link.setAttribute("aria-current", "page");
        currentState.count += 1;
      }
      parent.appendChild(link);

      if (entry.children) {
        renderEntries(parent, entry.children, currentId, baseUrl, depth + 1, currentState);
      }
    });
  }

  function renderSideNav(aside) {
    const currentId = aside.getAttribute("data-gx-ui-current");
    const currentState = { count: 0 };

    aside.setAttribute("aria-label", "GX/UI documentation");
    aside.replaceChildren();

    const title = document.createElement("p");
    title.className = "gx-side-nav-title";
    title.textContent = nav.title;
    aside.appendChild(title);

    nav.groups.forEach(function (group) {
      const wrapper = document.createElement("div");
      wrapper.className = "gx-side-nav-group";

      const heading = document.createElement("p");
      heading.className = "gx-side-nav-heading";
      heading.textContent = group.heading;
      wrapper.appendChild(heading);

      renderEntries(wrapper, group.entries, currentId, baseUrl, 0, currentState);
      aside.appendChild(wrapper);
    });

    if (currentState.count !== 1) {
      console.warn("GX/UI side nav expected exactly one current entry", currentId, currentState.count);
    }
  }

  document.querySelectorAll("aside.gx-side-nav[data-gx-ui-side-nav]").forEach(renderSideNav);
})();
