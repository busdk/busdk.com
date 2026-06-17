package demo

import (
	"reflect"
	"strings"
	"testing"

	gx "github.com/busdk/bus-gx/pkg/gx"
)

func TestIDsExposeRegisteredDemos(t *testing.T) {
	if got, want := IDs(), []string{"ai-approvals", "ai-attachment-list", "ai-composer", "ai-markdown", "ai-message", "ai-model-select", "ai-panel", "ai-review-status", "ai-thread-isolation", "ai-thread-list", "app-shell", "assistant-shell", "button", "credential-login-card", "date-input", "dense-table", "drop-zone", "empty-state", "error-banner", "event-bar", "evidence-link", "evidence-preview", "field", "file-input", "filter-toolbar", "form", "icon-button", "image-gallery", "input", "link-button", "loading-state", "menu", "metric-card", "navigation", "page-shell", "panel", "password-input", "portal", "projection-detail", "provider-error", "record-list", "result-panel", "select", "shell-action-panel", "sidebar-nav", "sidebar-shell", "status-pill", "submit", "summary-item", "surface-card", "tabs", "terminal-adapters", "terminal-approval-prompt", "terminal-input-box", "terminal-output-view", "terminal-session-panel", "text-input", "text-table", "textarea", "timeline"}; !reflect.DeepEqual(got, want) {
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

func TestSurfaceFamilyDemosRenderRealBusUIComponents(t *testing.T) {
	t.Parallel()

	tests := []struct {
		id   string
		want []string
	}{
		{
			id: "panel",
			want: []string{
				`data-bus-ui-demo-widget="panel"`,
				"bus-ui-panel",
				"Monthly close",
				"3 statements ready for review",
				`data-ui-action="panel.refresh"`,
				"Last sync: 09:41",
			},
		},
		{
			id: "surface-card",
			want: []string{
				`data-bus-ui-demo-widget="surface-card"`,
				"bus-ui-card",
				"Receipt evidence",
				"bus-ui-status-pill",
				`data-ui-status="success"`,
				"Open evidence queue",
			},
		},
		{
			id: "metric-card",
			want: []string{
				`data-bus-ui-demo-widget="metric-card"`,
				"bus-ui-metric-card",
				`data-ui-status="success"`,
				"Accepted rows",
				"128",
				"5 waiting for manual review",
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

func TestStatusFamilyDemosRenderRealBusUIComponents(t *testing.T) {
	t.Parallel()

	tests := []struct {
		id   string
		want []string
	}{
		{
			id: "status-pill",
			want: []string{
				`data-bus-ui-demo-widget="status-pill"`,
				"bus-ui-status-pill",
				`data-ui-status="success"`,
				"ready",
			},
		},
		{
			id: "empty-state",
			want: []string{
				`data-bus-ui-demo-widget="empty-state"`,
				"bus-ui-empty",
				"No files yet",
				`data-ui-action="empty.upload"`,
				"Upload file",
			},
		},
		{
			id: "loading-state",
			want: []string{
				`data-bus-ui-demo-widget="loading-state"`,
				"bus-ui-loading",
				`aria-busy="true"`,
				"Importing evidence",
				`value="67"`,
			},
		},
		{
			id: "result-panel",
			want: []string{
				`data-bus-ui-demo-widget="result-panel"`,
				"bus-ui-result-panel",
				`data-ui-status="success"`,
				"Import complete",
				`data-ui-action="result.view"`,
				`data-ui-action="result.download"`,
			},
		},
		{
			id: "error-banner",
			want: []string{
				`data-bus-ui-demo-widget="error-banner"`,
				"bus-ui-error-banner",
				`role="alert"`,
				"Provider is temporarily unavailable.",
				`data-ui-action="error.retry"`,
				`data-ui-action="error.dismiss"`,
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

func TestShellFamilyDemosRenderRealBusUIComponents(t *testing.T) {
	t.Parallel()

	tests := []struct {
		id   string
		want []string
	}{
		{
			id: "app-shell",
			want: []string{
				`data-bus-ui-demo-widget="app-shell"`,
				"<html",
				`data-ui-shell="app"`,
				"bus-ui-app-shell-document",
				"Workspace",
				"Bus UI shell preview",
			},
		},
		{
			id: "page-shell",
			want: []string{
				`data-bus-ui-demo-widget="page-shell"`,
				"bus-ui-page-shell",
				`data-ui-shell="page"`,
				"Uploaded accounting files",
				"Last import finished",
			},
		},
		{
			id: "sidebar-nav",
			want: []string{
				`data-bus-ui-demo-widget="sidebar-nav"`,
				"bus-ui-sidebar-nav-rail",
				"Accounting",
				"is-active",
				`data-ui-action="sidebar.refresh"`,
			},
		},
		{
			id: "sidebar-shell",
			want: []string{
				`data-bus-ui-demo-widget="sidebar-shell"`,
				"bus-ui-sidebar-shell",
				`data-ui-shell="sidebar"`,
				"Selected view content stays beside the rail.",
				"bus-ui-sidebar-nav-rail",
			},
		},
		{
			id: "shell-action-panel",
			want: []string{
				`data-bus-ui-demo-widget="shell-action-panel"`,
				"bus-ui-shell-action-panel",
				`data-ui-component="ShellActionPanel"`,
				"Review import",
				`data-ui-action="shell.approve"`,
				"Posting stays locked until every receipt matches.",
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
