package demo

import (
	gx "github.com/busdk/bus-gx/pkg/gx"
	assistantui "github.com/busdk/bus-ui/pkg/assistantui"
)

// AssistantShell renders one split business-plus-assistant workbench demo.
func AssistantShell() gx.Node {
	node := assistantui.AssistantShell(assistantui.AssistantShellProps{
		HeaderNodes: []assistantui.Node{
			gx.Element("p", gx.Props{"class": "bus-ui-demo-assistant-header"}, gx.Text("Operator workspace")),
		},
		BusinessNodes: []assistantui.Node{
			gx.Element("section", gx.Props{"class": "bus-ui-demo-assistant-business"},
				gx.Element("h2", nil, gx.Text("Invoice review")),
				gx.Element("p", nil, gx.Text("Confirm supplier totals before posting the import.")),
			),
		},
		AssistantNodes: []assistantui.Node{
			assistantui.AIPanel(assistantPanelProps()),
		},
		OnToggle: assistantToggleDemo,
		Width:    "30rem",
	})
	return demoWidget("assistant-shell", node)
}

// AIApprovals renders one approval-list demo backed by the public assistant API.
func AIApprovals() gx.Node {
	node := assistantui.AIApprovals(assistantui.AIApprovalsProps{
		ID:        "assistant-approvals",
		ThreadID:  "thread-review",
		Items:     assistantApprovalItems(),
		OnApprove: assistantApprovalDemo,
		OnReject:  assistantApprovalDemo,
	})
	return demoWidget("ai-approvals", node)
}

// AIAttachmentList renders one attachment-chip list demo.
func AIAttachmentList() gx.Node {
	node := assistantui.AIAttachmentList(assistantui.AIAttachmentListProps{
		ID:       "assistant-attachments",
		ThreadID: "thread-review",
		Size:     assistantui.AIAttachmentListSmall,
		Attachments: []assistantui.AIPanelAttachment{
			{ID: "notes-md", Label: "notes.md", SizeLabel: "4 KB"},
			{ID: "diff-json", Label: "diff-summary.json", SizeLabel: "12 KB"},
		},
		OnOpen:    assistantAttachmentDemo,
		OnInspect: assistantAttachmentDemo,
		OnRemove:  assistantAttachmentDemo,
	})
	return demoWidget("ai-attachment-list", node)
}

// AIComposer renders one controlled composer demo.
func AIComposer() gx.Node {
	node := assistantui.AIComposer(assistantui.AIComposerProps{
		ID:          "assistant-composer",
		InputID:     "assistant-composer-input",
		Value:       "Check the assistant shell spacing before applying.",
		Placeholder: "Ask the assistant",
		OnInput:     assistantComposerInputDemo,
		OnSend:      assistantComposerActionDemo,
		OnInterrupt: assistantComposerActionDemo,
	})
	return demoWidget("ai-composer", node)
}

// AIMarkdown renders one trusted assistant markdown demo.
func AIMarkdown() gx.Node {
	node := assistantui.AIMarkdown(assistantui.AIMarkdownProps{
		Text:  "Review [AssistantShell](./assistant-shell.go) and `AssistantShellProps` before applying.",
		Links: assistantui.AIMarkdownLinksWorkspacePaths,
		ResolveWorkspaceHref: func(target string) string {
			if target == "./assistant-shell.go" {
				return "/docs/gx-ui/bus-ui/assistant-shell/index.html"
			}
			return ""
		},
	})
	return demoWidget("ai-markdown", node)
}

// AIMessage renders one trusted transcript bubble demo.
func AIMessage() gx.Node {
	html := assistantui.AIMarkdownHTML(assistantui.AIMarkdownProps{
		Text:  "Review **assistant shell** layout before applying.",
		Links: assistantui.AIMarkdownLinksNone,
	})
	node := assistantui.AIMessage(assistantui.AIMessageProps{
		Role:      "assistant",
		HTML:      html,
		Sanitizer: assistantui.AIMarkdownSanitizer,
	})
	return demoWidget("ai-message", node)
}

// AIModelSelect renders one model-selector demo.
func AIModelSelect() gx.Node {
	node := assistantui.AIModelSelect(assistantui.AIModelSelectProps{
		ID:       "assistant-model",
		Current:  "gpt-5",
		Options:  assistantModelOptions(),
		OnChange: assistantModelDemo,
		ThreadID: "thread-review",
	})
	return demoWidget("ai-model-select", node)
}

// AIPanel renders one full assistant workbench demo.
func AIPanel() gx.Node {
	return demoWidget("ai-panel", assistantui.AIPanel(assistantPanelProps()))
}

// AIReviewStatus renders one review-before-apply status demo.
func AIReviewStatus() gx.Node {
	node := assistantui.AIReviewStatus(assistantui.AIReviewStatusProps{
		State:     assistantui.AIReviewStateApproved,
		Title:     "Ready to apply",
		Summary:   "Diff review passed and approvals are complete.",
		RequestID: "req-apply",
	})
	return demoWidget("ai-review-status", node)
}

// AIThreadIsolation renders one thread-isolation status demo.
func AIThreadIsolation() gx.Node {
	node := assistantui.AIThreadIsolation(assistantui.AIThreadIsolationProps{
		ThreadID: "thread-review",
		Owner:    "bus-ui",
		Branch:   "codex/assistant-shell-demo",
		Worktree: "/workspace/assistant/thread-review",
		Active:   true,
		Conflict: assistantui.AIThreadConflictDirty,
		Detail:   "working tree has local edits",
	})
	return demoWidget("ai-thread-isolation", node)
}

// AIThreadList renders one thread-list demo.
func AIThreadList() gx.Node {
	node := assistantui.AIThreadList(assistantui.AIThreadListProps{
		Threads:      assistantThreadSummaries(),
		ActiveThread: "thread-review",
		OnSelect:     assistantThreadDemo,
		OnArchive:    assistantThreadDemo,
	})
	return demoWidget("ai-thread-list", node)
}

func assistantPanelProps() assistantui.AIPanelProps {
	return assistantui.AIPanelProps{
		Title:         "AI Assistant",
		ActiveThread:  "thread-review",
		Threads:       assistantThreadSummaries(),
		Messages:      assistantPanelMessages(),
		Approvals:     assistantApprovalItems(),
		Model:         "gpt-5",
		ModelOptions:  assistantModelOptions(),
		Attachments:   []assistantui.AIPanelAttachment{{ID: "notes-md", Label: "notes.md", SizeLabel: "4 KB"}},
		Draft:         "Check the assistant shell spacing before applying.",
		OnSend:        assistantSendDemo,
		OnInterrupt:   assistantInterruptDemo,
		OnModelChange: assistantModelDemo,
		OnAttachment:  assistantAttachmentDemo,
		OnApprove:     assistantApprovalDemo,
		OnReject:      assistantApprovalDemo,
	}
}

func assistantThreadSummaries() []assistantui.AIThreadSummary {
	return []assistantui.AIThreadSummary{
		{ID: "thread-review", Title: "Assistant shell review", Working: true},
		{ID: "thread-copy", Title: "Marketing copy update"},
	}
}

func assistantPanelMessages() []assistantui.AIPanelMessage {
	return []assistantui.AIPanelMessage{
		{Role: "user", Text: "Review the assistant shell layout before applying."},
		{
			Role: "assistant",
			BodyNodes: []assistantui.Node{
				assistantui.AIMarkdown(assistantui.AIMarkdownProps{
					Text:  "Spacing is ready. See `AssistantShellProps` for the split-pane contract.",
					Links: assistantui.AIMarkdownLinksNone,
				}),
			},
		},
	}
}

func assistantApprovalItems() []assistantui.AIApprovalRequest {
	return []assistantui.AIApprovalRequest{
		{
			RequestID: "req-apply",
			Title:     "Apply patch",
			Summary:   "Update the assistant shell spacing in docs.",
			State:     assistantui.AIApprovalStatePending,
		},
	}
}

func assistantModelOptions() []assistantui.AIModelOption {
	return []assistantui.AIModelOption{
		{ID: "auto", Label: "Auto"},
		{ID: "gpt-5", Label: "GPT-5"},
		{ID: "gpt-5-mini", Label: "GPT-5 mini", Disabled: true, Reason: "Unavailable in this thread"},
	}
}

func assistantToggleDemo(event assistantui.AssistantToggleEvent) assistantui.Result {
	return assistantui.Success(event.Collapsed)
}

func assistantComposerInputDemo(event assistantui.AIComposeInputEvent) assistantui.Result {
	return assistantui.Success(event.Value)
}

func assistantComposerActionDemo(event assistantui.AIComposeEvent) assistantui.Result {
	return assistantui.Success(event.SourceID)
}

func assistantThreadDemo(event assistantui.AIThreadEvent) assistantui.Result {
	return assistantui.Success(event.ThreadID)
}

func assistantModelDemo(event assistantui.AIModelEvent) assistantui.Result {
	return assistantui.Success(event.ModelID)
}

func assistantApprovalDemo(event assistantui.AIApprovalEvent) assistantui.Result {
	return assistantui.Success(event.Decision)
}

func assistantAttachmentDemo(event assistantui.AIAttachmentEvent) assistantui.Result {
	return assistantui.Success(event.Operation)
}

func assistantSendDemo(event assistantui.AISendEvent) assistantui.Result {
	return assistantui.Success(event.ThreadID)
}

func assistantInterruptDemo(event assistantui.AIInterruptEvent) assistantui.Result {
	return assistantui.Success(event.ThreadID)
}
