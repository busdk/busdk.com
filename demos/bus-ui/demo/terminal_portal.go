package demo

import (
	gx "github.com/busdk/bus-gx/pkg/gx"
	terminalui "github.com/busdk/bus-ui/pkg/terminalui"
	uiportal "github.com/busdk/bus-ui/pkg/uiportal"
)

func TerminalAdapters() gx.Node {
	view, _, err := terminalui.TerminalSessionAdapter(terminalui.TerminalSessionAdapterInput{
		Metadata: terminalui.TerminalSessionMetadata{
			SessionID:        "sess-adapter-demo",
			Command:          "bus lint ./...",
			WorkingDirectory: "/workspace/busdk.com",
			ProcessID:        "4821",
		},
		Events: []terminalui.TerminalEvent{
			{Type: "terminal.opened", Payload: terminalui.TerminalEventPayload{SessionID: "sess-adapter-demo"}},
			{Type: "terminal.output", Payload: terminalui.TerminalEventPayload{Stream: "stdout", Text: "loading workspace"}},
			{Type: "terminal.output", Payload: terminalui.TerminalEventPayload{Stream: "stderr", Text: "warning: config fallback enabled"}},
		},
	})
	if err != nil {
		return demoError(err)
	}
	props := view.TerminalSessionPanelProps()
	props.Title = "Adapted session view"
	props.Summary = "TerminalSessionAdapter normalizes event streams into panel props."
	props.Elapsed = "00:03"
	props.EmptyOutput = "No output yet."
	return demoWidget("terminal-adapters", terminalui.TerminalSessionPanel(props))
}

func TerminalApprovalPrompt() gx.Node {
	node := terminalui.TerminalApprovalPrompt(terminalui.TerminalApprovalPromptProps{
		ID:      "terminal-approval-demo",
		Title:   "Approve container access?",
		Summary: "Allow the task to open a dedicated build container for this session.",
		Decisions: []terminalui.TerminalApprovalDecision{
			{
				Label:     "Approve",
				Decision:  terminalui.TerminalApprovalDecisionApprove,
				RequestID: "req-approve-container",
				OnClick:   noOpTerminalApproval,
			},
			{
				Label:     "Deny",
				Decision:  terminalui.TerminalApprovalDecisionDeny,
				RequestID: "req-approve-container",
				OnClick:   noOpTerminalApproval,
			},
		},
	})
	return demoWidget("terminal-approval-prompt", node)
}

func TerminalInputBox() gx.Node {
	node := terminalui.TerminalInputBox(terminalui.TerminalInputBoxProps{
		ID:           "terminal-input-demo",
		SessionID:    "sess-input-demo",
		Value:        "make quality",
		OnChange:     noOpTerminalInputChange,
		OnSend:       noOpTerminalInput,
		OnExit:       noOpTerminalExit,
		State:        terminalui.TerminalSessionStateRunning,
		Placeholder:  "Type terminal input",
		SubmitLabel:  "Send",
		SubmitAction: "terminal.send",
		ExitLabel:    "Stop",
		ExitAction:   "terminal.stop",
	})
	return demoWidget("terminal-input-box", node)
}

func TerminalOutputView() gx.Node {
	first := 1
	second := 2
	third := 3
	node := terminalui.TerminalOutputView([]terminalui.TerminalChunk{
		{Stream: "stdin", Text: "make quality", Sequence: &first},
		{Stream: "stdout", Text: "running gofmt and go test", Sequence: &second},
		{Stream: "system", Text: "quality checks still in progress", Sequence: &third},
	}, "No output yet.")
	return demoWidget("terminal-output-view", node)
}

func TerminalSessionPanel() gx.Node {
	exitCode := 0
	node := terminalui.TerminalSessionPanel(terminalui.TerminalSessionPanelProps{
		ID:               "terminal-session-demo",
		Title:            "Quality run",
		Summary:          "Go tests and docs checks complete successfully.",
		State:            terminalui.TerminalSessionStateExited,
		Command:          "make quality",
		WorkingDirectory: "/workspace/busdk.com",
		SessionID:        "sess-quality-demo",
		ProcessID:        "4910",
		Elapsed:          "00:11",
		ExitCode:         &exitCode,
		Output: []terminalui.TerminalChunk{
			{Stream: "stdout", Text: "ok   github.com/busdk/busdk.com/demos/bus-ui/demo"},
			{Stream: "system", Text: "All quality gates passed."},
		},
		EmptyOutput: "No output yet.",
	})
	return demoWidget("terminal-session-panel", node)
}

func Portal() gx.Node {
	ctx := uiportal.HostContext{
		ModuleID:       "accounting",
		BasePath:       "/modules/accounting/",
		PortalBasePath: "/",
		Session: &uiportal.SessionProps{
			Authenticated: true,
			IdentityLabel: "finance@example.com",
			Scopes:        []string{"ledger.read", "reports.export"},
		},
		PublicRuntimeConfig: map[string]any{
			"locale": "fi-FI",
		},
	}
	node := uiportal.PortalShell(uiportal.PortalShellProps{
		Title:       "Accounting",
		HostContext: ctx,
		Nav: []uiportal.NavItem{
			{Label: "Overview", Path: "/"},
			{Label: "Approvals", Path: "/approvals", Active: true},
		},
		MainNodes: []gx.Node{
			uiportal.HostSession(ctx),
			gx.Element("p", nil, gx.Text("Portal helpers resolve module-local links and render shared chrome from host context.")),
		},
	})
	return demoWidget("portal", node)
}

func noOpTerminalInput(terminalui.TerminalInputEvent) terminalui.Result {
	return terminalui.Noop()
}

func noOpTerminalInputChange(terminalui.TerminalInputChangeEvent) terminalui.Result {
	return terminalui.Noop()
}

func noOpTerminalExit(terminalui.TerminalExitEvent) terminalui.Result {
	return terminalui.Noop()
}

func noOpTerminalApproval(terminalui.TerminalApprovalEvent) terminalui.Result {
	return terminalui.Noop()
}
