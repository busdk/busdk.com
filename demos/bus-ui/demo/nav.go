package demo

import (
	"net/url"
	"strings"

	gx "github.com/busdk/bus-gx/pkg/gx"
)

type navGroup struct {
	Heading string
	Entries []navEntry
}

type navEntry struct {
	ID       string
	Href     string
	Label    string
	Children []navEntry
}

var gxUISideNav = struct {
	Title  string
	Groups []navGroup
}{
	Title: "GX and Bus UI",
	Groups: []navGroup{
		{
			Heading: "Overview",
			Entries: []navEntry{
				{ID: "index", Href: "index.html", Label: "Overview"},
				{ID: "reference", Href: "reference/index.html", Label: "Reference"},
				{ID: "modules", Href: "modules/index.html", Label: "Modules"},
			},
		},
		{
			Heading: "GX Framework",
			Entries: []navEntry{
				{ID: "gx", Href: "gx/index.html", Label: "GX Framework"},
				{ID: "gx/source-files", Href: "gx/source-files/index.html", Label: "Source files"},
				{ID: "gx/components", Href: "gx/components/index.html", Label: "Component functions"},
				{ID: "gx/props-children", Href: "gx/props-children/index.html", Label: "Props and children"},
				{ID: "gx/generated-go", Href: "gx/generated-go/index.html", Label: "Generated Go"},
				{ID: "gx/events", Href: "gx/events/index.html", Label: "Events"},
				{ID: "gx/rendering", Href: "gx/rendering/index.html", Label: "Rendering"},
				{ID: "gx/runtime", Href: "gx/runtime/index.html", Label: "Runtime bridges"},
				{ID: "gx/effects", Href: "gx/effects/index.html", Label: "Effects"},
				{ID: "gx/nodes", Href: "gx/nodes/index.html", Label: "Nodes and render tree"},
				{ID: "gx/testing", Href: "gx/testing/index.html", Label: "Testing"},
			},
		},
		{
			Heading: "Bus UI Library",
			Entries: []navEntry{
				{ID: "bus-ui", Href: "bus-ui/index.html", Label: "Bus UI Library"},
				{ID: "bus-ui/components", Href: "bus-ui/components/index.html", Label: "Components"},
				{
					ID:    "bus-ui/components/shell",
					Href:  "bus-ui/components/shell/index.html",
					Label: "Shells",
					Children: []navEntry{
						{ID: "bus-ui/components/shell/app-shell", Href: "bus-ui/components/shell/app-shell/index.html", Label: "AppShell"},
						{ID: "bus-ui/components/shell/page-shell", Href: "bus-ui/components/shell/page-shell/index.html", Label: "PageShell"},
						{ID: "bus-ui/components/shell/sidebar-shell", Href: "bus-ui/components/shell/sidebar-shell/index.html", Label: "SidebarShell"},
						{ID: "bus-ui/components/shell/sidebar-nav", Href: "bus-ui/components/shell/sidebar-nav/index.html", Label: "SidebarNav"},
						{ID: "bus-ui/components/shell/shell-action-panel", Href: "bus-ui/components/shell/shell-action-panel/index.html", Label: "ShellActionPanel"},
					},
				},
				{
					ID:    "bus-ui/components/navigation",
					Href:  "bus-ui/components/navigation/index.html",
					Label: "Navigation",
					Children: []navEntry{
						{ID: "bus-ui/components/navigation/menu", Href: "bus-ui/components/navigation/menu/index.html", Label: "Menu"},
						{ID: "bus-ui/components/navigation/tabs", Href: "bus-ui/components/navigation/tabs/index.html", Label: "Tabs"},
						{ID: "bus-ui/components/navigation/navigation", Href: "bus-ui/components/navigation/navigation/index.html", Label: "Navigation"},
					},
				},
				{
					ID:    "bus-ui/components/action",
					Href:  "bus-ui/components/action/index.html",
					Label: "Actions",
					Children: []navEntry{
						{ID: "bus-ui/components/action/button", Href: "bus-ui/components/action/button/index.html", Label: "Button"},
						{ID: "bus-ui/components/action/link-button", Href: "bus-ui/components/action/link-button/index.html", Label: "LinkButton"},
						{ID: "bus-ui/components/action/icon-button", Href: "bus-ui/components/action/icon-button/index.html", Label: "IconButton"},
						{ID: "bus-ui/components/action/event-bar", Href: "bus-ui/components/action/event-bar/index.html", Label: "EventBar"},
					},
				},
				{
					ID:    "bus-ui/components/surface",
					Href:  "bus-ui/components/surface/index.html",
					Label: "Surfaces",
					Children: []navEntry{
						{ID: "bus-ui/components/surface/panel", Href: "bus-ui/components/surface/panel/index.html", Label: "Panel"},
						{ID: "bus-ui/components/surface/surface-card", Href: "bus-ui/components/surface/surface-card/index.html", Label: "SurfaceCard"},
						{ID: "bus-ui/components/surface/metric-card", Href: "bus-ui/components/surface/metric-card/index.html", Label: "MetricCard"},
					},
				},
				{
					ID:    "bus-ui/components/status",
					Href:  "bus-ui/components/status/index.html",
					Label: "Status",
					Children: []navEntry{
						{ID: "bus-ui/components/status/status-pill", Href: "bus-ui/components/status/status-pill/index.html", Label: "StatusPill"},
						{ID: "bus-ui/components/status/empty-state", Href: "bus-ui/components/status/empty-state/index.html", Label: "EmptyState"},
						{ID: "bus-ui/components/status/loading-state", Href: "bus-ui/components/status/loading-state/index.html", Label: "LoadingState"},
						{ID: "bus-ui/components/status/result-panel", Href: "bus-ui/components/status/result-panel/index.html", Label: "ResultPanel"},
						{ID: "bus-ui/components/status/error-banner", Href: "bus-ui/components/status/error-banner/index.html", Label: "ErrorBanner"},
					},
				},
				{
					ID:    "bus-ui/forms",
					Href:  "bus-ui/forms/index.html",
					Label: "Forms",
					Children: []navEntry{
						{ID: "bus-ui/forms/form", Href: "bus-ui/forms/form/index.html", Label: "Form"},
						{ID: "bus-ui/forms/field", Href: "bus-ui/forms/field/index.html", Label: "Field"},
						{ID: "bus-ui/forms/input", Href: "bus-ui/forms/input/index.html", Label: "Input"},
						{ID: "bus-ui/forms/text-input", Href: "bus-ui/forms/text-input/index.html", Label: "TextInput"},
						{ID: "bus-ui/forms/password-input", Href: "bus-ui/forms/password-input/index.html", Label: "PasswordInput"},
						{ID: "bus-ui/forms/date-input", Href: "bus-ui/forms/date-input/index.html", Label: "DateInput"},
						{ID: "bus-ui/forms/textarea", Href: "bus-ui/forms/textarea/index.html", Label: "TextArea"},
						{ID: "bus-ui/forms/select", Href: "bus-ui/forms/select/index.html", Label: "Select"},
						{ID: "bus-ui/forms/submit", Href: "bus-ui/forms/submit/index.html", Label: "SubmitControl"},
						{ID: "bus-ui/forms/filter-toolbar", Href: "bus-ui/forms/filter-toolbar/index.html", Label: "FilterToolbar"},
						{ID: "bus-ui/forms/file-input", Href: "bus-ui/forms/file-input/index.html", Label: "FileInput"},
						{ID: "bus-ui/forms/drop-zone", Href: "bus-ui/forms/drop-zone/index.html", Label: "DropZone"},
						{ID: "bus-ui/forms/credential-login-card", Href: "bus-ui/forms/credential-login-card/index.html", Label: "CredentialLoginCard"},
					},
				},
				{
					ID:    "bus-ui/data",
					Href:  "bus-ui/data/index.html",
					Label: "Data display",
					Children: []navEntry{
						{ID: "bus-ui/data/dense-table", Href: "bus-ui/data/dense-table/index.html", Label: "DenseTable"},
						{ID: "bus-ui/data/text-table", Href: "bus-ui/data/text-table/index.html", Label: "TextTable"},
						{ID: "bus-ui/data/record-list", Href: "bus-ui/data/record-list/index.html", Label: "RecordList"},
						{ID: "bus-ui/data/summary-item", Href: "bus-ui/data/summary-item/index.html", Label: "SummaryItem"},
						{ID: "bus-ui/data/projection-detail", Href: "bus-ui/data/projection-detail/index.html", Label: "ProjectionDetail"},
						{ID: "bus-ui/data/provider-error", Href: "bus-ui/data/provider-error/index.html", Label: "ProviderError"},
						{ID: "bus-ui/data/timeline", Href: "bus-ui/data/timeline/index.html", Label: "Timeline"},
					},
				},
				{
					ID:    "bus-ui/evidence",
					Href:  "bus-ui/evidence/index.html",
					Label: "Evidence and files",
					Children: []navEntry{
						{ID: "bus-ui/evidence/evidence-link", Href: "bus-ui/evidence/evidence-link/index.html", Label: "EvidenceLink"},
						{ID: "bus-ui/evidence/evidence-preview", Href: "bus-ui/evidence/evidence-preview/index.html", Label: "EvidencePreview"},
						{ID: "bus-ui/evidence/image-gallery", Href: "bus-ui/evidence/image-gallery/index.html", Label: "ImageGallery"},
					},
				},
				{
					ID:    "bus-ui/assistant",
					Href:  "bus-ui/assistant/index.html",
					Label: "Assistant",
					Children: []navEntry{
						{ID: "bus-ui/assistant-shell", Href: "bus-ui/assistant-shell/index.html", Label: "AssistantShell"},
						{ID: "bus-ui/ai-panel", Href: "bus-ui/ai-panel/index.html", Label: "AIPanel"},
						{ID: "bus-ui/ai-composer", Href: "bus-ui/ai-composer/index.html", Label: "AIComposer"},
						{ID: "bus-ui/ai-model-select", Href: "bus-ui/ai-model-select/index.html", Label: "AIModelSelect"},
						{ID: "bus-ui/ai-approvals", Href: "bus-ui/ai-approvals/index.html", Label: "AIApprovals"},
						{ID: "bus-ui/ai-review-status", Href: "bus-ui/ai-review-status/index.html", Label: "AIReviewStatus"},
						{ID: "bus-ui/ai-thread-isolation", Href: "bus-ui/ai-thread-isolation/index.html", Label: "AIThreadIsolation"},
						{ID: "bus-ui/ai-thread-list", Href: "bus-ui/ai-thread-list/index.html", Label: "AIThreadList"},
						{ID: "bus-ui/ai-message", Href: "bus-ui/ai-message/index.html", Label: "AIMessage"},
						{ID: "bus-ui/ai-markdown", Href: "bus-ui/ai-markdown/index.html", Label: "AIMarkdown"},
						{ID: "bus-ui/ai-attachment-list", Href: "bus-ui/ai-attachment-list/index.html", Label: "AIAttachmentList"},
					},
				},
				{
					ID:    "bus-ui/terminal",
					Href:  "bus-ui/terminal/index.html",
					Label: "Terminal",
					Children: []navEntry{
						{ID: "bus-ui/terminal-session-panel", Href: "bus-ui/terminal-session-panel/index.html", Label: "TerminalSessionPanel"},
						{ID: "bus-ui/terminal-output-view", Href: "bus-ui/terminal-output-view/index.html", Label: "TerminalOutputView"},
						{ID: "bus-ui/terminal-input-box", Href: "bus-ui/terminal-input-box/index.html", Label: "TerminalInputBox"},
						{ID: "bus-ui/terminal-approval-prompt", Href: "bus-ui/terminal-approval-prompt/index.html", Label: "TerminalApprovalPrompt"},
						{ID: "bus-ui/terminal-adapters", Href: "bus-ui/terminal-adapters/index.html", Label: "TerminalAdapters"},
					},
				},
				{ID: "bus-ui/portal", Href: "bus-ui/portal/index.html", Label: "Portal integration"},
				{ID: "bus-ui/assistant-terminal", Href: "bus-ui/assistant-terminal/index.html", Label: "Assistant and Terminal split"},
			},
		},
		{
			Heading: "Tutorials",
			Entries: []navEntry{
				{ID: "authoring", Href: "authoring/index.html", Label: "Authoring tutorial"},
				{ID: "runtime", Href: "runtime/index.html", Label: "Runtime and testing"},
				{ID: "components", Href: "components/index.html", Label: "Component tutorial"},
				{ID: "surfaces", Href: "surfaces/index.html", Label: "Product surfaces"},
			},
		},
	},
}

// GXUISideNav returns the shared GX/UI docs navigation rendered through GX nodes.
func GXUISideNav(currentID string, baseURL string) gx.Node {
	children, _ := gxUISideNavChildren(currentID, baseURL)
	return gx.Fragment(children...)
}

// GXUISideNavCurrentCount reports how many entries match the requested current id.
func GXUISideNavCurrentCount(currentID string) int {
	_, count := gxUISideNavChildren(currentID, "")
	return count
}

func gxUISideNavChildren(currentID string, baseURL string) ([]gx.Node, int) {
	nodes := []gx.Node{
		gx.Element("p", gx.Props{"class": "gx-side-nav-title"}, gx.Text(gxUISideNav.Title)),
	}
	currentCount := 0
	for _, group := range gxUISideNav.Groups {
		groupChildren := []gx.Node{
			gx.Element("p", gx.Props{"class": "gx-side-nav-heading"}, gx.Text(group.Heading)),
		}
		entries, count := renderNavEntries(group.Entries, strings.TrimSpace(currentID), strings.TrimSpace(baseURL), 0)
		currentCount += count
		groupChildren = append(groupChildren, entries...)
		nodes = append(nodes, gx.Element("div", gx.Props{"class": "gx-side-nav-group"}, groupChildren...))
	}
	return nodes, currentCount
}

func renderNavEntries(entries []navEntry, currentID string, baseURL string, depth int) ([]gx.Node, int) {
	nodes := make([]gx.Node, 0, len(entries))
	currentCount := 0
	for _, entry := range entries {
		attrs := gx.Props{
			"href": resolveNavHref(baseURL, entry.Href),
		}
		if className := linkClass(depth); className != "" {
			attrs["class"] = className
		}
		if entry.ID == currentID {
			attrs["aria-current"] = "page"
			currentCount++
		}
		nodes = append(nodes, gx.Element("a", attrs, gx.Text(entry.Label)))
		if len(entry.Children) == 0 {
			continue
		}
		children, count := renderNavEntries(entry.Children, currentID, baseURL, depth+1)
		currentCount += count
		nodes = append(nodes, children...)
	}
	return nodes, currentCount
}

func linkClass(depth int) string {
	if depth == 1 {
		return "gx-side-nav-child"
	}
	if depth >= 2 {
		return "gx-side-nav-grandchild"
	}
	return ""
}

func resolveNavHref(baseURL string, href string) string {
	trimmedHref := strings.TrimSpace(href)
	if trimmedHref == "" {
		return ""
	}
	trimmedBase := strings.TrimSpace(baseURL)
	if trimmedBase == "" {
		return trimmedHref
	}
	base, err := url.Parse(trimmedBase)
	if err != nil {
		return trimmedHref
	}
	resolved, err := base.Parse(trimmedHref)
	if err != nil {
		return trimmedHref
	}
	return resolved.String()
}
