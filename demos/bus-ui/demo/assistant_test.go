package demo

import (
	"os"
	"strings"
	"testing"

	gx "github.com/busdk/bus-gx/pkg/gx"
)

type assistantDocsDemoPage struct {
	id    string
	label string
	path  string
}

var assistantDocsDemoPages = []assistantDocsDemoPage{
	{
		id:    "assistant-shell",
		label: "Loading AssistantShell demo...",
		path:  "../../../docs/gx-ui/bus-ui/assistant-shell/index.html",
	},
	{
		id:    "ai-approvals",
		label: "Loading AIApprovals demo...",
		path:  "../../../docs/gx-ui/bus-ui/ai-approvals/index.html",
	},
	{
		id:    "ai-attachment-list",
		label: "Loading AIAttachmentList demo...",
		path:  "../../../docs/gx-ui/bus-ui/ai-attachment-list/index.html",
	},
	{
		id:    "ai-composer",
		label: "Loading AIComposer demo...",
		path:  "../../../docs/gx-ui/bus-ui/ai-composer/index.html",
	},
	{
		id:    "ai-markdown",
		label: "Loading AIMarkdown demo...",
		path:  "../../../docs/gx-ui/bus-ui/ai-markdown/index.html",
	},
	{
		id:    "ai-message",
		label: "Loading AIMessage demo...",
		path:  "../../../docs/gx-ui/bus-ui/ai-message/index.html",
	},
	{
		id:    "ai-model-select",
		label: "Loading AIModelSelect demo...",
		path:  "../../../docs/gx-ui/bus-ui/ai-model-select/index.html",
	},
	{
		id:    "ai-panel",
		label: "Loading AIPanel demo...",
		path:  "../../../docs/gx-ui/bus-ui/ai-panel/index.html",
	},
	{
		id:    "ai-review-status",
		label: "Loading AIReviewStatus demo...",
		path:  "../../../docs/gx-ui/bus-ui/ai-review-status/index.html",
	},
	{
		id:    "ai-thread-isolation",
		label: "Loading AIThreadIsolation demo...",
		path:  "../../../docs/gx-ui/bus-ui/ai-thread-isolation/index.html",
	},
	{
		id:    "ai-thread-list",
		label: "Loading AIThreadList demo...",
		path:  "../../../docs/gx-ui/bus-ui/ai-thread-list/index.html",
	},
}

func TestAssistantDemosRenderRealBusUIComponents(t *testing.T) {
	t.Parallel()

	tests := []struct {
		id   string
		want []string
	}{
		{
			id: "assistant-shell",
			want: []string{
				`data-bus-ui-demo-widget="assistant-shell"`,
				`data-ui-shell="assistant"`,
				`data-ui-callback="onToggle"`,
				"Invoice review",
				"AI Assistant",
			},
		},
		{
			id: "ai-approvals",
			want: []string{
				`data-bus-ui-demo-widget="ai-approvals"`,
				`class="bus-ui-ai-approvals ai-approvals"`,
				`data-approval-id="req-apply"`,
				`data-decision="approve"`,
				"Apply patch",
			},
		},
		{
			id: "ai-attachment-list",
			want: []string{
				`data-bus-ui-demo-widget="ai-attachment-list"`,
				`class="bus-ui-ai-attachments ai-attachments bus-ui-ai-attachments-small"`,
				`data-ai-action="open-attachment"`,
				`data-ai-action="inspect-attachment"`,
				`data-ai-action="remove-attachment"`,
				"notes.md",
			},
		},
		{
			id: "ai-composer",
			want: []string{
				`data-bus-ui-demo-widget="ai-composer"`,
				`class="bus-ui-ai-composer ai-input-row"`,
				`data-ai-action="send"`,
				`data-ai-action="interrupt"`,
				`id="assistant-composer-input"`,
			},
		},
		{
			id: "ai-markdown",
			want: []string{
				`data-bus-ui-demo-widget="ai-markdown"`,
				`href="/docs/gx-ui/bus-ui/assistant-shell/index.html"`,
				`class="ai-code"`,
				"AssistantShellProps",
			},
		},
		{
			id: "ai-message",
			want: []string{
				`data-bus-ui-demo-widget="ai-message"`,
				`data-ai-role="assistant"`,
				`data-ai-sanitizer="bus-markdown/v1"`,
				`class="ai-strong"`,
			},
		},
		{
			id: "ai-model-select",
			want: []string{
				`data-bus-ui-demo-widget="ai-model-select"`,
				`class="bus-ui-ai-model-select"`,
				`data-ai-action="set-model"`,
				`data-thread-id="thread-review"`,
				`<option selected value="gpt-5">GPT-5</option>`,
			},
		},
		{
			id: "ai-panel",
			want: []string{
				`data-bus-ui-demo-widget="ai-panel"`,
				`class="bus-ui-ai-panel ai-panel"`,
				`data-active-thread="thread-review"`,
				`data-ai-action="set-model"`,
				"Apply patch",
				"notes.md",
				"AI working",
			},
		},
		{
			id: "ai-review-status",
			want: []string{
				`data-bus-ui-demo-widget="ai-review-status"`,
				`data-review-state="approved"`,
				"Ready to apply",
				"req-apply",
			},
		},
		{
			id: "ai-thread-isolation",
			want: []string{
				`data-bus-ui-demo-widget="ai-thread-isolation"`,
				`data-isolation-conflict="dirty"`,
				"codex/assistant-shell-demo",
				"working tree has local edits",
			},
		},
		{
			id: "ai-thread-list",
			want: []string{
				`data-bus-ui-demo-widget="ai-thread-list"`,
				`data-ui-component="ai-thread-list"`,
				`data-ai-action="select-thread"`,
				`data-ai-action="archive-thread"`,
				"Assistant shell review",
				"AI working",
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

func TestAssistantDocsPagesUseGeneratedPlaceholderAndSharedScripts(t *testing.T) {
	t.Parallel()

	for _, page := range assistantDocsDemoPages {
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
			if got := strings.Count(text, `data-bus-ui-demo="`); got != 1 {
				t.Fatalf("%s data-bus-ui-demo count = %d, want 1", page.path, got)
			}
			if got := strings.Count(text, "wasm_exec.js"); got != 1 {
				t.Fatalf("%s wasm_exec.js count = %d, want 1", page.path, got)
			}
			if got := strings.Count(text, "loader.js"); got != 1 {
				t.Fatalf("%s loader.js count = %d, want 1", page.path, got)
			}
			if got := strings.Count(text, "data-bus-ui-demo-loader"); got != 1 {
				t.Fatalf("%s data-bus-ui-demo-loader count = %d, want 1", page.path, got)
			}
		})
	}
}
