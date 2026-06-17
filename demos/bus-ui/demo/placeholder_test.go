package demo

import (
	"os"
	"strings"
	"testing"
)

type docsDemoPage struct {
	id    string
	label string
	path  string
}

var docsDemoPages = []docsDemoPage{
	{
		id:    "button",
		label: "Loading Button demo...",
		path:  "../../../docs/gx-ui/bus-ui/components/action/button/index.html",
	},
	{
		id:    "link-button",
		label: "Loading LinkButton demo...",
		path:  "../../../docs/gx-ui/bus-ui/components/action/link-button/index.html",
	},
	{
		id:    "icon-button",
		label: "Loading IconButton demo...",
		path:  "../../../docs/gx-ui/bus-ui/components/action/icon-button/index.html",
	},
	{
		id:    "event-bar",
		label: "Loading EventBar demo...",
		path:  "../../../docs/gx-ui/bus-ui/components/action/event-bar/index.html",
	},
	{
		id:    "menu",
		label: "Loading Menu demo...",
		path:  "../../../docs/gx-ui/bus-ui/components/navigation/menu/index.html",
	},
	{
		id:    "navigation",
		label: "Loading Navigation demo...",
		path:  "../../../docs/gx-ui/bus-ui/components/navigation/navigation/index.html",
	},
	{
		id:    "tabs",
		label: "Loading Tabs demo...",
		path:  "../../../docs/gx-ui/bus-ui/components/navigation/tabs/index.html",
	},
	{
		id:    "panel",
		label: "Loading Panel demo...",
		path:  "../../../docs/gx-ui/bus-ui/components/surface/panel/index.html",
	},
	{
		id:    "surface-card",
		label: "Loading SurfaceCard demo...",
		path:  "../../../docs/gx-ui/bus-ui/components/surface/surface-card/index.html",
	},
	{
		id:    "metric-card",
		label: "Loading MetricCard demo...",
		path:  "../../../docs/gx-ui/bus-ui/components/surface/metric-card/index.html",
	},
	{
		id:    "status-pill",
		label: "Loading StatusPill demo...",
		path:  "../../../docs/gx-ui/bus-ui/components/status/status-pill/index.html",
	},
	{
		id:    "empty-state",
		label: "Loading EmptyState demo...",
		path:  "../../../docs/gx-ui/bus-ui/components/status/empty-state/index.html",
	},
	{
		id:    "loading-state",
		label: "Loading LoadingState demo...",
		path:  "../../../docs/gx-ui/bus-ui/components/status/loading-state/index.html",
	},
	{
		id:    "result-panel",
		label: "Loading ResultPanel demo...",
		path:  "../../../docs/gx-ui/bus-ui/components/status/result-panel/index.html",
	},
	{
		id:    "error-banner",
		label: "Loading ErrorBanner demo...",
		path:  "../../../docs/gx-ui/bus-ui/components/status/error-banner/index.html",
	},
	{
		id:    "app-shell",
		label: "Loading AppShell demo...",
		path:  "../../../docs/gx-ui/bus-ui/components/shell/app-shell/index.html",
	},
	{
		id:    "page-shell",
		label: "Loading PageShell demo...",
		path:  "../../../docs/gx-ui/bus-ui/components/shell/page-shell/index.html",
	},
	{
		id:    "sidebar-shell",
		label: "Loading SidebarShell demo...",
		path:  "../../../docs/gx-ui/bus-ui/components/shell/sidebar-shell/index.html",
	},
	{
		id:    "sidebar-nav",
		label: "Loading SidebarNav demo...",
		path:  "../../../docs/gx-ui/bus-ui/components/shell/sidebar-nav/index.html",
	},
	{
		id:    "shell-action-panel",
		label: "Loading ShellActionPanel demo...",
		path:  "../../../docs/gx-ui/bus-ui/components/shell/shell-action-panel/index.html",
	},
}

func TestPlaceholderHTMLUsesSharedLoader(t *testing.T) {
	t.Parallel()

	for _, page := range docsDemoPages {
		page := page
		t.Run(page.id, func(t *testing.T) {
			t.Parallel()

			html, err := PlaceholderHTML(page.id, page.label)
			if err != nil {
				t.Fatalf("PlaceholderHTML(%q) error = %v", page.id, err)
			}
			for _, want := range []string{
				`data-bus-ui-demo="` + page.id + `"`,
				`data-ui-component="Loader"`,
				`class="bus-ui-loader bus-ui-demo-placeholder-loader bus-ui-loader-md"`,
				`<span class="bus-ui-loader-label">` + page.label + `</span>`,
			} {
				if !strings.Contains(html, want) {
					t.Fatalf("PlaceholderHTML(%q) missing %q in %s", page.id, want, html)
				}
			}
		})
	}
}

func TestDocsPagesUseGeneratedPlaceholderAndSharedScripts(t *testing.T) {
	t.Parallel()

	for _, page := range docsDemoPages {
		page := page
		t.Run(page.id, func(t *testing.T) {
			t.Parallel()

			placeholder, err := PlaceholderHTML(page.id, page.label)
			if err != nil {
				t.Fatalf("PlaceholderHTML(%q) error = %v", page.id, err)
			}
			body, err := os.ReadFile(page.path)
			if err != nil {
				t.Fatalf("ReadFile(%s) error = %v", page.path, err)
			}
			text := string(body)
			if !strings.Contains(text, placeholder) {
				t.Fatalf("%s does not contain generated placeholder %q", page.path, placeholder)
			}
			for _, want := range []string{
				"assets/bus-ui-demo/wasm_exec.js",
				"assets/bus-ui-demo/loader.js",
				"data-bus-ui-demo-loader",
			} {
				if !strings.Contains(text, want) {
					t.Fatalf("%s missing %q", page.path, want)
				}
			}
		})
	}
}
