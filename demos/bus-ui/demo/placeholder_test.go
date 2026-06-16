package demo

import (
	"os"
	"strings"
	"testing"
)

func TestPlaceholderHTMLUsesSharedLoader(t *testing.T) {
	html, err := PlaceholderHTML("button", "Loading Button demo...")
	if err != nil {
		t.Fatalf("PlaceholderHTML() error = %v", err)
	}
	for _, want := range []string{
		`data-bus-ui-demo="button"`,
		`data-ui-component="Loader"`,
		`class="bus-ui-loader bus-ui-demo-placeholder-loader bus-ui-loader-md"`,
		`<span class="bus-ui-loader-label">Loading Button demo...</span>`,
	} {
		if !strings.Contains(html, want) {
			t.Fatalf("PlaceholderHTML() missing %q in %s", want, html)
		}
	}
}

func TestButtonDocsPageUsesGeneratedPlaceholder(t *testing.T) {
	placeholder, err := PlaceholderHTML("button", "Loading Button demo...")
	if err != nil {
		t.Fatalf("PlaceholderHTML() error = %v", err)
	}
	page, err := os.ReadFile("../../../docs/gx-ui/bus-ui/components/action/button/index.html")
	if err != nil {
		t.Fatalf("ReadFile(button docs page) error = %v", err)
	}
	if !strings.Contains(string(page), placeholder) {
		t.Fatalf("button docs page does not contain generated placeholder %q", placeholder)
	}
}
