package demo

import (
	"strings"
	"testing"

	gx "github.com/busdk/bus-gx/pkg/gx"
)

func TestTerminalAndPortalDemosRenderRealPublicMarkup(t *testing.T) {
	t.Parallel()

	tests := []struct {
		id   string
		want []string
	}{
		{
			id: "terminal-adapters",
			want: []string{
				`data-bus-ui-demo-widget="terminal-adapters"`,
				`data-ui-component="terminal-session-panel"`,
				"Adapted session view",
				"bus lint ./...",
				"loading workspace",
				"warning: config fallback enabled",
			},
		},
		{
			id: "terminal-approval-prompt",
			want: []string{
				`data-bus-ui-demo-widget="terminal-approval-prompt"`,
				`data-ui-component="terminal-approval-prompt"`,
				"Approve container access?",
				`data-terminal-approval-request-id="req-approve-container"`,
				"Approve",
				"Deny",
			},
		},
		{
			id: "terminal-input-box",
			want: []string{
				`data-bus-ui-demo-widget="terminal-input-box"`,
				`data-ui-component="terminal-input-box"`,
				`data-terminal-session-id="sess-input-demo"`,
				"make quality",
				"Send",
				"Stop",
			},
		},
		{
			id: "terminal-output-view",
			want: []string{
				`data-bus-ui-demo-widget="terminal-output-view"`,
				`class="terminal-output"`,
				`data-terminal-stream="stdin"`,
				`data-terminal-stream="stdout"`,
				`data-terminal-stream="system"`,
				"quality checks still in progress",
			},
		},
		{
			id: "terminal-session-panel",
			want: []string{
				`data-bus-ui-demo-widget="terminal-session-panel"`,
				`data-ui-component="terminal-session-panel"`,
				"Quality run",
				"make quality",
				"All quality gates passed.",
				"Exit code",
			},
		},
		{
			id: "portal",
			want: []string{
				`data-bus-ui-demo-widget="portal"`,
				`data-ui-component="PortalShell"`,
				`data-portal-module="accounting"`,
				`href="/modules/accounting/approvals"`,
				`class="bus-ui-session`,
				"finance@example.com",
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
