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
