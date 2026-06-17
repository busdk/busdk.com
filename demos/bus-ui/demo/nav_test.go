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

const gxUIBaseURL = "https://busdk.com/docs/gx-ui/"

func TestGXUITopHeaderIncludesBrandLogoAndPricing(t *testing.T) {
	html := renderGXUITopHeader(t, gxUIBaseURL)
	for _, want := range []string{
		`class="site-header-inner"`,
		`class="brand" href="https://busdk.com/docs/index.html"`,
		`class="brand-logo" src="https://busdk.com/docs/busdk-logo.png" alt="BusDK logo"`,
		`href="https://busdk.com/docs/gx-ui/gx/index.html">GX Framework</a>`,
		`href="https://busdk.com/docs/gx-ui/bus-ui/index.html">Bus UI Library</a>`,
		`href="https://busdk.com/docs/gx-ui/reference/index.html">Reference</a>`,
		`href="https://busdk.com/docs/gx-ui/modules/index.html">Modules</a>`,
		`href="https://busdk.com/docs/gx-ui/pricing/index.html">Pricing</a>`,
	} {
		if !strings.Contains(html, want) {
			t.Fatalf("GXUITopHeader HTML missing %q in %s", want, html)
		}
	}
}

func TestGXUITopHeaderResolvesPricingInsideGXUIRoot(t *testing.T) {
	html := renderGXUITopHeader(t, gxUIBaseURL)
	want := `href="https://busdk.com/docs/gx-ui/pricing/index.html">Pricing</a>`
	if !strings.Contains(html, want) {
		t.Fatalf("GXUITopHeader HTML missing %q in %s", want, html)
	}
	unwanted := `href="https://busdk.com/docs/pricing/index.html">Pricing</a>`
	if strings.Contains(html, unwanted) {
		t.Fatalf("GXUITopHeader HTML unexpectedly included %q in %s", unwanted, html)
	}
}

func TestGXUIFooterIncludesCopyrightFallbackContent(t *testing.T) {
	html := renderGXUIFooter(t)
	for _, want := range []string{
		`class="site-footer-inner"`,
		`© `,
		`href="https://hg.fi/">Heusala Group Ltd</a>.`,
	} {
		if !strings.Contains(html, want) {
			t.Fatalf("GXUIFooter HTML missing %q in %s", want, html)
		}
	}
}

func TestGXUISideNavIncludesRequiredGroupsAndIDs(t *testing.T) {
	html := renderGXUISideNav(t, "bus-ui/forms")
	for _, want := range []string{
		`class="gx-side-nav-title"`,
		`>Overview</p>`,
		`>GX Framework</p>`,
		`>Bus UI Library</p>`,
		`>Tutorials</p>`,
		`href="https://busdk.com/docs/gx-ui/bus-ui/index.html"`,
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

func TestGXUISideNavCollapsesTopLevelContext(t *testing.T) {
	tests := []struct {
		name     string
		current  string
		want     []string
		unwanted []string
	}{
		{
			name:    "overview",
			current: "index",
			want: []string{
				`aria-current="page" href="https://busdk.com/docs/gx-ui/index.html">Overview</a>`,
				`href="https://busdk.com/docs/gx-ui/pricing/index.html">Pricing</a>`,
				`href="https://busdk.com/docs/gx-ui/reference/index.html">Reference</a>`,
				`href="https://busdk.com/docs/gx-ui/modules/index.html">Modules</a>`,
			},
			unwanted: []string{
				`href="https://busdk.com/docs/gx-ui/gx/index.html">GX Framework</a>`,
				`href="https://busdk.com/docs/gx-ui/bus-ui/index.html">Bus UI Library</a>`,
				`href="https://busdk.com/docs/gx-ui/runtime/index.html">Runtime and testing</a>`,
			},
		},
		{
			name:    "gx framework",
			current: "gx/events",
			want: []string{
				`href="https://busdk.com/docs/gx-ui/gx/index.html">GX Framework</a>`,
				`aria-current="page" href="https://busdk.com/docs/gx-ui/gx/events/index.html">Events</a>`,
				`href="https://busdk.com/docs/gx-ui/gx/nodes/index.html">Nodes and render tree</a>`,
			},
			unwanted: []string{
				`href="https://busdk.com/docs/gx-ui/pricing/index.html">Pricing</a>`,
				`href="https://busdk.com/docs/gx-ui/bus-ui/index.html">Bus UI Library</a>`,
				`href="https://busdk.com/docs/gx-ui/runtime/index.html">Runtime and testing</a>`,
			},
		},
		{
			name:    "bus ui",
			current: "bus-ui/forms/text-input",
			want: []string{
				`href="https://busdk.com/docs/gx-ui/bus-ui/index.html">Bus UI Library</a>`,
				`href="https://busdk.com/docs/gx-ui/bus-ui/forms/index.html">Forms</a>`,
				`aria-current="page" class="gx-side-nav-child" href="https://busdk.com/docs/gx-ui/bus-ui/forms/text-input/index.html">TextInput</a>`,
			},
			unwanted: []string{
				`href="https://busdk.com/docs/gx-ui/pricing/index.html">Pricing</a>`,
				`href="https://busdk.com/docs/gx-ui/gx/events/index.html">Events</a>`,
				`href="https://busdk.com/docs/gx-ui/runtime/index.html">Runtime and testing</a>`,
			},
		},
		{
			name:    "tutorials",
			current: "runtime",
			want: []string{
				`href="https://busdk.com/docs/gx-ui/authoring/index.html">Authoring tutorial</a>`,
				`aria-current="page" href="https://busdk.com/docs/gx-ui/runtime/index.html">Runtime and testing</a>`,
				`href="https://busdk.com/docs/gx-ui/components/index.html">Component tutorial</a>`,
			},
			unwanted: []string{
				`href="https://busdk.com/docs/gx-ui/index.html">Overview</a>`,
				`href="https://busdk.com/docs/gx-ui/gx/index.html">GX Framework</a>`,
				`href="https://busdk.com/docs/gx-ui/bus-ui/index.html">Bus UI Library</a>`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			html := renderGXUISideNav(t, tt.current)
			for _, heading := range []string{
				`>Overview</p>`,
				`>GX Framework</p>`,
				`>Bus UI Library</p>`,
				`>Tutorials</p>`,
			} {
				if !strings.Contains(html, heading) {
					t.Fatalf("GXUISideNav HTML missing heading %q in %s", heading, html)
				}
			}
			for _, want := range tt.want {
				if !strings.Contains(html, want) {
					t.Fatalf("GXUISideNav HTML missing %q in %s", want, html)
				}
			}
			for _, unwanted := range tt.unwanted {
				if strings.Contains(html, unwanted) {
					t.Fatalf("GXUISideNav HTML unexpectedly included %q in %s", unwanted, html)
				}
			}
			if got := strings.Count(html, `aria-current="page"`); got != 1 {
				t.Fatalf("aria-current count = %d, want 1 in %s", got, html)
			}
		})
	}
}

func TestGXUISideNavFormsLeafShowsSiblingLeavesOnly(t *testing.T) {
	html := renderGXUISideNav(t, "bus-ui/forms/text-input")
	for _, want := range []string{
		`href="https://busdk.com/docs/gx-ui/bus-ui/forms/form/index.html">Form</a>`,
		`href="https://busdk.com/docs/gx-ui/bus-ui/forms/password-input/index.html">PasswordInput</a>`,
		`aria-current="page" class="gx-side-nav-child" href="https://busdk.com/docs/gx-ui/bus-ui/forms/text-input/index.html">TextInput</a>`,
	} {
		if !strings.Contains(html, want) {
			t.Fatalf("GXUISideNav HTML missing %q in %s", want, html)
		}
	}
	for _, unwanted := range []string{
		`href="https://busdk.com/docs/gx-ui/bus-ui/data/dense-table/index.html">DenseTable</a>`,
		`href="https://busdk.com/docs/gx-ui/bus-ui/assistant-shell/index.html">AssistantShell</a>`,
	} {
		if strings.Contains(html, unwanted) {
			t.Fatalf("GXUISideNav HTML unexpectedly included %q in %s", unwanted, html)
		}
	}
	if got := strings.Count(html, `aria-current="page"`); got != 1 {
		t.Fatalf("aria-current count = %d, want 1 in %s", got, html)
	}
}

func TestGXUISideNavUsesIDForDuplicateLabels(t *testing.T) {
	html := renderGXUISideNav(t, "bus-ui/components/navigation/navigation")
	for _, want := range []string{
		`href="https://busdk.com/docs/gx-ui/bus-ui/components/navigation/menu/index.html">Menu</a>`,
		`href="https://busdk.com/docs/gx-ui/bus-ui/components/navigation/tabs/index.html">Tabs</a>`,
		`href="https://busdk.com/docs/gx-ui/bus-ui/components/navigation/navigation/index.html">Navigation</a>`,
	} {
		if !strings.Contains(html, want) {
			t.Fatalf("GXUISideNav HTML missing %q in %s", want, html)
		}
	}
	want := `aria-current="page" class="gx-side-nav-child" href="https://busdk.com/docs/gx-ui/bus-ui/components/navigation/navigation/index.html">Navigation</a>`
	if !strings.Contains(html, want) {
		t.Fatalf("GXUISideNav HTML missing %q in %s", want, html)
	}
	unwanted := `aria-current="page" href="https://busdk.com/docs/gx-ui/bus-ui/components/navigation/index.html">Navigation</a>`
	if strings.Contains(html, unwanted) {
		t.Fatalf("GXUISideNav incorrectly marked family link current in %s", html)
	}
	if got := strings.Count(html, `aria-current="page"`); got != 1 {
		t.Fatalf("aria-current count = %d, want 1 in %s", got, html)
	}
}

func TestGXUISideNavTopLevelPageKeepsChildrenCollapsed(t *testing.T) {
	html := renderGXUISideNav(t, "reference")
	for _, unwanted := range []string{
		`href="https://busdk.com/docs/gx-ui/bus-ui/components/navigation/menu/index.html">Menu</a>`,
		`href="https://busdk.com/docs/gx-ui/bus-ui/forms/form/index.html">Form</a>`,
		`href="https://busdk.com/docs/gx-ui/bus-ui/data/dense-table/index.html">DenseTable</a>`,
		`href="https://busdk.com/docs/gx-ui/bus-ui/assistant-shell/index.html">AssistantShell</a>`,
	} {
		if strings.Contains(html, unwanted) {
			t.Fatalf("GXUISideNav HTML unexpectedly included %q in %s", unwanted, html)
		}
	}
	want := `aria-current="page" href="https://busdk.com/docs/gx-ui/reference/index.html">Reference</a>`
	if !strings.Contains(html, want) {
		t.Fatalf("GXUISideNav HTML missing %q in %s", want, html)
	}
	if got := strings.Count(html, `aria-current="page"`); got != 1 {
		t.Fatalf("aria-current count = %d, want 1 in %s", got, html)
	}
}

func TestGXUISideNavRepresentativeRenderHasExactlyOneCurrent(t *testing.T) {
	html := renderGXUISideNav(t, "bus-ui/data/provider-error")
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

func renderGXUISideNav(t *testing.T, currentID string) string {
	t.Helper()
	html, err := gx.RenderHTML(GXUISideNav(currentID, gxUIBaseURL))
	if err != nil {
		t.Fatalf("RenderHTML(GXUISideNav %q) failed: %v", currentID, err)
	}
	return html
}

func renderGXUITopHeader(t *testing.T, baseURL string) string {
	t.Helper()
	html, err := gx.RenderHTML(GXUITopHeader(baseURL))
	if err != nil {
		t.Fatalf("RenderHTML(GXUITopHeader %q) failed: %v", baseURL, err)
	}
	return html
}

func renderGXUIFooter(t *testing.T) string {
	t.Helper()
	html, err := gx.RenderHTML(GXUIFooter())
	if err != nil {
		t.Fatalf("RenderHTML(GXUIFooter) failed: %v", err)
	}
	return html
}
