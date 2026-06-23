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

func TestGXUITopHeaderUsesGXUILinks(t *testing.T) {
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

func TestGXUITopHeaderDoesNotRenderGlobalSiteLinks(t *testing.T) {
	html := renderGXUITopHeader(t, gxUIBaseURL)
	for _, unwanted := range []string{
		`href="https://busdk.com/docs/index.html#products">Products</a>`,
		`href="https://docs.busdk.com/">Documentation</a>`,
		`href="https://busdk.com/docs/blog/index.html">Blog</a>`,
	} {
		if strings.Contains(html, unwanted) {
			t.Fatalf("GXUITopHeader HTML unexpectedly included %q in %s", unwanted, html)
		}
	}
}

func TestBusDKTopHeaderMarksRequestedCurrentLink(t *testing.T) {
	html := renderBusDKTopHeader(t, "site", "https://busdk.com/docs/", "blog")
	want := `aria-current="page" href="https://busdk.com/docs/blog/index.html">Blog</a>`
	if !strings.Contains(html, want) {
		t.Fatalf("BusDKTopHeader HTML missing %q in %s", want, html)
	}
	if got := strings.Count(html, `aria-current="page"`); got != 1 {
		t.Fatalf("aria-current count = %d, want 1 in %s", got, html)
	}
}

func TestBusDKTopHeaderRendersProductLocalLinks(t *testing.T) {
	html := renderBusDKTopHeader(t, "services", "https://busdk.com/docs/services/", "pricing")
	for _, want := range []string{
		`class="brand" href="https://busdk.com/docs/index.html"`,
		`href="https://busdk.com/docs/services/index.html">Overview</a>`,
		`href="https://busdk.com/docs/services/getting-started/index.html">Getting started</a>`,
		`href="https://busdk.com/docs/services/docker/index.html">Docker</a>`,
		`aria-current="page" href="https://busdk.com/docs/services/pricing/index.html">Pricing</a>`,
		`href="https://busdk.com/docs/services/contact/index.html">Contact</a>`,
	} {
		if !strings.Contains(html, want) {
			t.Fatalf("BusDKTopHeader HTML missing %q in %s", want, html)
		}
	}
	for _, unwanted := range []string{
		`href="https://busdk.com/docs/index.html#products">Products</a>`,
		`href="https://busdk.com/docs/blog/index.html">Blog</a>`,
	} {
		if strings.Contains(html, unwanted) {
			t.Fatalf("BusDKTopHeader HTML unexpectedly included %q in %s", unwanted, html)
		}
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

func TestBusDKProductSideNavRendersProductLocalEntries(t *testing.T) {
	html := renderBusDKProductSideNav(t, "services", "docker", "https://busdk.com/docs/services/")
	for _, want := range []string{
		`class="gx-side-nav-title"`,
		`>Services guide</p>`,
		`href="https://busdk.com/docs/services/index.html">Overview</a>`,
		`href="https://busdk.com/docs/services/getting-started/index.html">Getting started</a>`,
		`href="https://busdk.com/docs/services/examples/index.html">Examples</a>`,
		`aria-current="page" href="https://busdk.com/docs/services/docker/index.html">Docker</a>`,
		`href="https://busdk.com/docs/services/modules/index.html">Modules</a>`,
		`href="https://busdk.com/docs/services/pricing/index.html">Pricing</a>`,
		`href="https://busdk.com/docs/services/contact/index.html">Contact</a>`,
	} {
		if !strings.Contains(html, want) {
			t.Fatalf("BusDKProductSideNav HTML missing %q in %s", want, html)
		}
	}
	if got := BusDKProductSideNavCurrentCount("services", "docker"); got != 1 {
		t.Fatalf("BusDKProductSideNavCurrentCount() = %d, want 1", got)
	}
}

func TestBusDKProductNavRendersEngineOSPages(t *testing.T) {
	headerHTML := renderBusDKTopHeader(t, "engine-os", "https://busdk.com/docs/engine-os/", "getting-started")
	for _, want := range []string{
		`class="brand" href="https://busdk.com/docs/index.html"`,
		`href="https://busdk.com/docs/engine-os/index.html">Overview</a>`,
		`aria-current="page" href="https://busdk.com/docs/engine-os/getting-started/index.html">Getting started</a>`,
		`href="https://busdk.com/docs/engine-os/modules/index.html">Modules</a>`,
		`href="https://busdk.com/docs/engine-os/contact/index.html">Contact</a>`,
	} {
		if !strings.Contains(headerHTML, want) {
			t.Fatalf("BusDKTopHeader engine-os HTML missing %q in %s", want, headerHTML)
		}
	}

	sideNavHTML := renderBusDKProductSideNav(t, "engine-os", "modules", "https://busdk.com/docs/engine-os/")
	for _, want := range []string{
		`class="gx-side-nav-title"`,
		`>Engine OS</p>`,
		`href="https://busdk.com/docs/engine-os/index.html">Overview</a>`,
		`href="https://busdk.com/docs/engine-os/getting-started/index.html">Getting started</a>`,
		`aria-current="page" href="https://busdk.com/docs/engine-os/modules/index.html">Modules</a>`,
		`href="https://busdk.com/docs/engine-os/contact/index.html">Contact</a>`,
	} {
		if !strings.Contains(sideNavHTML, want) {
			t.Fatalf("BusDKProductSideNav engine-os HTML missing %q in %s", want, sideNavHTML)
		}
	}
	if got := BusDKProductSideNavCurrentCount("engine-os", "modules"); got != 1 {
		t.Fatalf("BusDKProductSideNavCurrentCount(engine-os) = %d, want 1", got)
	}
}

func TestGXUISideNavIncludesRequiredGroupsAndIDs(t *testing.T) {
	html := renderGXUISideNav(t, "bus-ui/forms")
	for _, want := range []string{
		`class="gx-side-nav-title"`,
		`class="gx-side-nav-heading" href="https://busdk.com/docs/gx-ui/index.html">Overview</a>`,
		`class="gx-side-nav-heading" href="https://busdk.com/docs/gx-ui/gx/index.html">GX Framework</a>`,
		`>Bus UI Library</p>`,
		`class="gx-side-nav-heading" href="https://busdk.com/docs/gx-ui/authoring/index.html">Tutorials</a>`,
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
				`>Overview</p>`,
				`aria-current="page" href="https://busdk.com/docs/gx-ui/index.html">Overview</a>`,
				`class="gx-side-nav-heading" href="https://busdk.com/docs/gx-ui/gx/index.html">GX Framework</a>`,
				`class="gx-side-nav-heading" href="https://busdk.com/docs/gx-ui/bus-ui/index.html">Bus UI Library</a>`,
				`class="gx-side-nav-heading" href="https://busdk.com/docs/gx-ui/authoring/index.html">Tutorials</a>`,
				`href="https://busdk.com/docs/gx-ui/pricing/index.html">Pricing</a>`,
				`href="https://busdk.com/docs/gx-ui/reference/index.html">Reference</a>`,
				`href="https://busdk.com/docs/gx-ui/modules/index.html">Modules</a>`,
			},
			unwanted: []string{
				`href="https://busdk.com/docs/gx-ui/runtime/index.html">Runtime and testing</a>`,
				`href="https://busdk.com/docs/gx-ui/gx/events/index.html">Events</a>`,
				`href="https://busdk.com/docs/gx-ui/bus-ui/forms/index.html">Forms</a>`,
			},
		},
		{
			name:    "gx framework",
			current: "gx/events",
			want: []string{
				`class="gx-side-nav-heading" href="https://busdk.com/docs/gx-ui/index.html">Overview</a>`,
				`>GX Framework</p>`,
				`href="https://busdk.com/docs/gx-ui/gx/index.html">GX Framework</a>`,
				`aria-current="page" href="https://busdk.com/docs/gx-ui/gx/events/index.html">Events</a>`,
				`href="https://busdk.com/docs/gx-ui/gx/nodes/index.html">Nodes and render tree</a>`,
				`class="gx-side-nav-heading" href="https://busdk.com/docs/gx-ui/bus-ui/index.html">Bus UI Library</a>`,
				`class="gx-side-nav-heading" href="https://busdk.com/docs/gx-ui/authoring/index.html">Tutorials</a>`,
			},
			unwanted: []string{
				`href="https://busdk.com/docs/gx-ui/pricing/index.html">Pricing</a>`,
				`href="https://busdk.com/docs/gx-ui/runtime/index.html">Runtime and testing</a>`,
				`href="https://busdk.com/docs/gx-ui/bus-ui/forms/index.html">Forms</a>`,
			},
		},
		{
			name:    "bus ui",
			current: "bus-ui/forms/text-input",
			want: []string{
				`class="gx-side-nav-heading" href="https://busdk.com/docs/gx-ui/index.html">Overview</a>`,
				`class="gx-side-nav-heading" href="https://busdk.com/docs/gx-ui/gx/index.html">GX Framework</a>`,
				`>Bus UI Library</p>`,
				`href="https://busdk.com/docs/gx-ui/bus-ui/index.html">Bus UI Library</a>`,
				`href="https://busdk.com/docs/gx-ui/bus-ui/forms/index.html">Forms</a>`,
				`aria-current="page" class="gx-side-nav-child" href="https://busdk.com/docs/gx-ui/bus-ui/forms/text-input/index.html">TextInput</a>`,
				`class="gx-side-nav-heading" href="https://busdk.com/docs/gx-ui/authoring/index.html">Tutorials</a>`,
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
				`class="gx-side-nav-heading" href="https://busdk.com/docs/gx-ui/index.html">Overview</a>`,
				`class="gx-side-nav-heading" href="https://busdk.com/docs/gx-ui/gx/index.html">GX Framework</a>`,
				`class="gx-side-nav-heading" href="https://busdk.com/docs/gx-ui/bus-ui/index.html">Bus UI Library</a>`,
				`>Tutorials</p>`,
				`href="https://busdk.com/docs/gx-ui/authoring/index.html">Authoring tutorial</a>`,
				`aria-current="page" href="https://busdk.com/docs/gx-ui/runtime/index.html">Runtime and testing</a>`,
				`href="https://busdk.com/docs/gx-ui/components/index.html">Component tutorial</a>`,
			},
			unwanted: []string{
				`href="https://busdk.com/docs/gx-ui/pricing/index.html">Pricing</a>`,
				`href="https://busdk.com/docs/gx-ui/gx/events/index.html">Events</a>`,
				`href="https://busdk.com/docs/gx-ui/bus-ui/forms/index.html">Forms</a>`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			html := renderGXUISideNav(t, tt.current)
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

func renderBusDKTopHeader(t *testing.T, navID string, baseURL string, currentID string) string {
	t.Helper()
	html, err := gx.RenderHTML(BusDKTopHeader(navID, baseURL, currentID))
	if err != nil {
		t.Fatalf("RenderHTML(BusDKTopHeader %q, %q, %q) failed: %v", navID, baseURL, currentID, err)
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

func renderBusDKProductSideNav(t *testing.T, navID string, currentID string, baseURL string) string {
	t.Helper()
	html, err := gx.RenderHTML(BusDKProductSideNav(navID, currentID, baseURL))
	if err != nil {
		t.Fatalf("RenderHTML(BusDKProductSideNav %q, %q, %q) failed: %v", navID, currentID, baseURL, err)
	}
	return html
}
