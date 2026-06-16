package demo

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	gx "github.com/busdk/bus-gx/pkg/gx"
)

var currentMarkerPattern = regexp.MustCompile(`data-gx-ui-current="([^"]+)"`)

func TestGXUISideNavIncludesRequiredGroupsAndIDs(t *testing.T) {
	html, err := gx.RenderHTML(GXUISideNav("bus-ui/forms", "https://busdk.com/docs/gx-ui/"))
	if err != nil {
		t.Fatalf("RenderHTML(GXUISideNav) failed: %v", err)
	}
	for _, want := range []string{
		`class="gx-side-nav-title"`,
		`>Overview</p>`,
		`>GX Framework</p>`,
		`>Bus UI Library</p>`,
		`>Tutorials</p>`,
		`href="https://busdk.com/docs/gx-ui/bus-ui/forms/index.html"`,
		`href="https://busdk.com/docs/gx-ui/bus-ui/data/index.html"`,
		`href="https://busdk.com/docs/gx-ui/bus-ui/evidence/index.html"`,
		`href="https://busdk.com/docs/gx-ui/bus-ui/assistant/index.html"`,
		`href="https://busdk.com/docs/gx-ui/bus-ui/terminal/index.html"`,
		`href="https://busdk.com/docs/gx-ui/bus-ui/portal/index.html"`,
	} {
		if !strings.Contains(html, want) {
			t.Fatalf("GXUISideNav HTML missing %q in %s", want, html)
		}
	}
}

func TestGXUISideNavUsesIDForDuplicateLabels(t *testing.T) {
	html, err := gx.RenderHTML(GXUISideNav("bus-ui/components/navigation/navigation", "https://busdk.com/docs/gx-ui/"))
	if err != nil {
		t.Fatalf("RenderHTML(GXUISideNav duplicate label case) failed: %v", err)
	}
	want := `aria-current="page" class="gx-side-nav-child" href="https://busdk.com/docs/gx-ui/bus-ui/components/navigation/navigation/index.html">Navigation</a>`
	if !strings.Contains(html, want) {
		t.Fatalf("GXUISideNav HTML missing %q in %s", want, html)
	}
	unwanted := `aria-current="page" href="https://busdk.com/docs/gx-ui/bus-ui/components/navigation/index.html">Navigation</a>`
	if strings.Contains(html, unwanted) {
		t.Fatalf("GXUISideNav incorrectly marked family link current in %s", html)
	}
}

func TestGXUISideNavRepresentativeRenderHasExactlyOneCurrent(t *testing.T) {
	html, err := gx.RenderHTML(GXUISideNav("bus-ui/data/provider-error", "https://busdk.com/docs/gx-ui/"))
	if err != nil {
		t.Fatalf("RenderHTML(GXUISideNav representative case) failed: %v", err)
	}
	if got := strings.Count(html, `aria-current="page"`); got != 1 {
		t.Fatalf("aria-current count = %d, want 1 in %s", got, html)
	}
	if got := GXUISideNavCurrentCount("bus-ui/data/provider-error"); got != 1 {
		t.Fatalf("GXUISideNavCurrentCount() = %d, want 1", got)
	}
}

func TestGXUISideNavCoversPublishedCurrentMarkers(t *testing.T) {
	root := filepath.Join("..", "..", "..", "docs", "gx-ui")
	pages := 0
	err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || filepath.Base(path) != "index.html" {
			return nil
		}
		body, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		matches := currentMarkerPattern.FindAllSubmatch(body, -1)
		if len(matches) == 0 {
			return nil
		}
		pages++
		for _, match := range matches {
			currentID := string(match[1])
			if got := GXUISideNavCurrentCount(currentID); got != 1 {
				t.Fatalf("%s current id %q matched %d nav entries, want 1", path, currentID, got)
			}
		}
		return nil
	})
	if err != nil {
		t.Fatalf("walk docs/gx-ui current markers: %v", err)
	}
	if pages == 0 {
		t.Fatal("no docs/gx-ui pages with data-gx-ui-current found")
	}
}
