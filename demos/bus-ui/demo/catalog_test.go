package demo

import (
	"reflect"
	"strings"
	"testing"

	gx "github.com/busdk/bus-gx/pkg/gx"
)

func TestIDsExposeRegisteredDemos(t *testing.T) {
	if got, want := IDs(), []string{"button", "event-bar", "icon-button", "link-button", "menu", "navigation", "tabs"}; !reflect.DeepEqual(got, want) {
		t.Fatalf("IDs() = %#v, want %#v", got, want)
	}
}

func TestButtonDemoRendersRealBusUIButton(t *testing.T) {
	root, ok := Lookup("button")
	if !ok {
		t.Fatal("button demo is not registered")
	}
	html, err := gx.RenderHTML(root())
	if err != nil {
		t.Fatalf("RenderHTML(button demo) failed: %v", err)
	}
	for _, want := range []string{
		"bus-ui-btn",
		"bus-ui-btn-primary",
		"Save draft",
		"data-bus-ui-demo-widget",
		`data-bus-ui-demo-action="button-click"`,
		`data-bus-ui-demo-status="button"`,
		`role="status"`,
		`aria-live="polite"`,
		"Ready",
	} {
		if !strings.Contains(html, want) {
			t.Fatalf("button demo HTML %q does not contain %q", html, want)
		}
	}
}

func TestActionFamilyDemosRenderRealBusUIComponents(t *testing.T) {
	t.Parallel()

	tests := []struct {
		id   string
		want []string
	}{
		{
			id: "link-button",
			want: []string{
				`data-bus-ui-demo-widget="link-button"`,
				"bus-ui-btn",
				"bus-ui-btn-secondary",
				`href="/docs/billing/index.html"`,
				"Open invoices",
			},
		},
		{
			id: "icon-button",
			want: []string{
				`data-bus-ui-demo-widget="icon-button"`,
				"bus-ui-btn",
				"bus-ui-btn-ghost",
				"bus-ui-icon",
				`data-ui-action="refresh-events"`,
				"Refresh",
			},
		},
		{
			id: "event-bar",
			want: []string{
				`data-bus-ui-demo-widget="event-bar"`,
				"bus-ui-event-bar",
				`aria-label="File actions"`,
				`data-ui-action="import-file"`,
				`data-ui-action="open-log"`,
				"Open log",
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.id, func(t *testing.T) {
			t.Parallel()

			root, ok := Lookup(tc.id)
			if !ok {
				t.Fatalf("%s demo is not registered", tc.id)
			}
			html, err := gx.RenderHTML(root())
			if err != nil {
				t.Fatalf("RenderHTML(%s demo) failed: %v", tc.id, err)
			}
			for _, want := range tc.want {
				if !strings.Contains(html, want) {
					t.Fatalf("%s demo HTML %q does not contain %q", tc.id, html, want)
				}
			}
		})
	}
}

func TestNavigationFamilyDemosRenderRealBusUIComponents(t *testing.T) {
	t.Parallel()

	tests := []struct {
		id   string
		want []string
	}{
		{
			id: "menu",
			want: []string{
				`data-bus-ui-demo-widget="menu"`,
				"bus-ui-menu",
				"Actions",
				"Refresh",
				"Archive",
			},
		},
		{
			id: "navigation",
			want: []string{
				`data-bus-ui-demo-widget="navigation"`,
				"bus-ui-navigation",
				`aria-label="Workspace sections"`,
				"Billing",
				"is-active",
			},
		},
		{
			id: "tabs",
			want: []string{
				`data-bus-ui-demo-widget="tabs"`,
				"bus-ui-tabs",
				`aria-label="Accounting views"`,
				"Files",
				`role="tablist"`,
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.id, func(t *testing.T) {
			t.Parallel()

			root, ok := Lookup(tc.id)
			if !ok {
				t.Fatalf("%s demo is not registered", tc.id)
			}
			html, err := gx.RenderHTML(root())
			if err != nil {
				t.Fatalf("RenderHTML(%s demo) failed: %v", tc.id, err)
			}
			for _, want := range tc.want {
				if !strings.Contains(html, want) {
					t.Fatalf("%s demo HTML %q does not contain %q", tc.id, html, want)
				}
			}
		})
	}
}
