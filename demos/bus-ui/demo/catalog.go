package demo

import (
	"sort"
	"strings"

	gx "github.com/busdk/bus-gx/pkg/gx"
	ui "github.com/busdk/bus-ui/pkg/ui"
)

// Root renders one static Bus UI demo node.
type Root func() gx.Node

var catalog = map[string]Root{
	"app-shell":          AppShell,
	"button":             Button,
	"empty-state":        EmptyState,
	"error-banner":       ErrorBanner,
	"event-bar":          EventBar,
	"icon-button":        IconButton,
	"link-button":        LinkButton,
	"loading-state":      LoadingState,
	"metric-card":        MetricCard,
	"menu":               Menu,
	"navigation":         Navigation,
	"panel":              Panel,
	"page-shell":         PageShell,
	"result-panel":       ResultPanel,
	"shell-action-panel": ShellActionPanel,
	"sidebar-nav":        SidebarNav,
	"sidebar-shell":      SidebarShell,
	"status-pill":        StatusPill,
	"surface-card":       SurfaceCard,
	"tabs":               Tabs,
}

// IDs returns the stable placeholder ids supported by the static demo runtime.
func IDs() []string {
	ids := make([]string, 0, len(catalog))
	for id := range catalog {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	return ids
}

// Lookup resolves one demo id from a data-bus-ui-demo placeholder.
func Lookup(id string) (Root, bool) {
	root, ok := catalog[strings.TrimSpace(id)]
	return root, ok
}

// Button renders the first website-hosted live Bus UI component demo.
func Button() gx.Node {
	node, err := ui.Button(ui.ButtonProps{
		Label:   "Save draft",
		Variant: ui.ButtonPrimary,
		Size:    ui.ButtonMd,
		Attrs: map[string]string{
			"type":                    "button",
			"data-bus-ui-demo-action": "button-click",
		},
	})
	if err != nil {
		return demoError(err)
	}
	return gx.Element("div", gx.Props{"class": "bus-ui-demo-button", "data-bus-ui-demo-widget": "button"},
		node,
		gx.Element("span", gx.Props{
			"class":                   "bus-ui-demo-status",
			"role":                    "status",
			"aria-live":               "polite",
			"data-bus-ui-demo-status": "button",
		}, gx.Text("Ready")),
	)
}

// LinkButton renders one compact button-styled link demo.
func LinkButton() gx.Node {
	node, err := ui.LinkButton(ui.LinkButtonProps{
		Label:   "Open invoices",
		Href:    "/docs/billing/index.html",
		Variant: ui.ButtonSecondary,
		Size:    ui.ButtonMd,
	})
	if err != nil {
		return demoError(err)
	}
	return demoWidget("link-button", node)
}

// IconButton renders one compact icon-leading button demo.
func IconButton() gx.Node {
	node, err := ui.IconButton(ui.IconButtonProps{
		Label:   "Refresh",
		Icon:    ui.IconProps{Name: ui.IconNameRefresh},
		Variant: ui.ButtonGhost,
		Size:    ui.ButtonSm,
		Control: ui.ControlProps{
			Action:   "refresh-events",
			SourceID: "icon-button-demo",
		},
	})
	if err != nil {
		return demoError(err)
	}
	return demoWidget("icon-button", node)
}

// EventBar renders one compact grouped action-row demo.
func EventBar() gx.Node {
	node, err := ui.EventBar(ui.EventBarProps{
		Label: "File actions",
		Actions: []ui.EventBarAction{
			{
				Label:   "Import",
				Variant: ui.ButtonPrimary,
				Control: ui.ControlProps{
					Action:   "import-file",
					SourceID: "event-bar-import",
				},
			},
			{
				Label:   "Open log",
				Icon:    ui.IconProps{Name: ui.IconNameOpen},
				Variant: ui.ButtonGhost,
				Control: ui.ControlProps{
					Action:   "open-log",
					SourceID: "event-bar-open-log",
				},
			},
		},
	})
	if err != nil {
		return demoError(err)
	}
	return demoWidget("event-bar", node)
}

// Menu renders one bounded command-choice menu demo.
func Menu() gx.Node {
	node, err := ui.Menu(ui.MenuProps{
		TriggerLabel: "Actions",
		Selected:     "refresh",
		Items: []ui.MenuItem{
			{Label: "Refresh", Value: "refresh", OnClick: noOpMenuClick},
			{Label: "Archive", Value: "archive", OnClick: noOpMenuClick},
		},
	})
	if err != nil {
		return demoError(err)
	}
	return demoWidget("menu", node)
}

// Navigation renders one compact section-list navigation demo.
func Navigation() gx.Node {
	node, err := ui.Navigation(ui.NavigationProps{
		Label:  "Workspace sections",
		Active: "billing",
		Items: []ui.NavigationItem{
			{ID: "overview", Label: "Overview", Href: "/docs/overview/index.html"},
			{ID: "billing", Label: "Billing", Href: "/docs/billing/index.html"},
			{ID: "exports", Label: "Exports", Href: "/docs/exports/index.html"},
		},
	})
	if err != nil {
		return demoError(err)
	}
	return demoWidget("navigation", node)
}

// Tabs renders one sibling-view tablist demo.
func Tabs() gx.Node {
	node, err := ui.Tabs(ui.TabsProps{
		Label:  "Accounting views",
		Active: "files",
		Items: []ui.TabItem{
			{ID: "overview", Label: "Overview", Href: "/docs/accounting/index.html"},
			{ID: "files", Label: "Files", Href: "/docs/accounting/files/index.html"},
			{ID: "events", Label: "Events", Href: "/docs/accounting/events/index.html"},
		},
	})
	if err != nil {
		return demoError(err)
	}
	return demoWidget("tabs", node)
}

// Panel renders one titled operational surface demo.
func Panel() gx.Node {
	refresh, err := ui.Button(ui.ButtonProps{
		Label:   "Refresh",
		Variant: ui.ButtonSecondary,
		Size:    ui.ButtonSm,
		Control: ui.ControlProps{
			Action:   "panel.refresh",
			SourceID: "panel-refresh",
		},
	})
	if err != nil {
		return demoError(err)
	}
	node := ui.Panel(ui.PanelProps{
		Title:    "Monthly close",
		Subtitle: "3 statements ready for review",
		ActionNodes: []gx.Node{
			refresh,
		},
		BodyNodes: []gx.Node{
			gx.Element("p", nil, gx.Text("Reconciled entries are ready to post.")),
		},
		FooterNodes: []gx.Node{
			gx.Element("p", nil, gx.Text("Last sync: 09:41")),
		},
	})
	if node == nil {
		return demoError(ui.ErrSurfaceBodyRequired)
	}
	return demoWidget("panel", node)
}

// SurfaceCard renders one compact grouped-content card demo.
func SurfaceCard() gx.Node {
	pill, err := ui.StatusPill(ui.StatusPillProps{
		Label:  "ready",
		Status: ui.StatusSuccess,
	})
	if err != nil {
		return demoError(err)
	}
	node := ui.SurfaceCard(ui.SurfaceCardProps{
		HeaderNodes: []gx.Node{
			gx.Element("div", gx.Props{"class": "bus-ui-demo-card-head"},
				gx.Element("span", gx.Props{"class": "bus-ui-demo-card-title"}, gx.Text("Receipt evidence")),
				pill,
			),
		},
		BodyNodes: []gx.Node{
			gx.Element("p", nil, gx.Text("Image, supplier, and VAT fields are matched.")),
		},
		FooterNodes: []gx.Node{
			gx.Element("a", gx.Props{"href": "/docs/evidence/index.html"}, gx.Text("Open evidence queue")),
		},
	})
	if node == nil {
		return demoError(ui.ErrSurfaceBodyRequired)
	}
	return demoWidget("surface-card", node)
}

// MetricCard renders one summary metric card demo.
func MetricCard() gx.Node {
	node := ui.MetricCard(ui.MetricCardProps{
		Title:  "Accepted rows",
		Value:  "128",
		Detail: "5 waiting for manual review",
		Status: ui.MetricCardStatusSuccess,
	})
	if node == nil {
		return demoError(ui.ErrSurfaceValueRequired)
	}
	return demoWidget("metric-card", node)
}

// StatusPill renders one compact semantic status label demo.
func StatusPill() gx.Node {
	node, err := ui.StatusPill(ui.StatusPillProps{
		Label:  "ready",
		Status: ui.StatusSuccess,
	})
	if err != nil {
		return demoError(err)
	}
	return demoWidget("status-pill", node)
}

// EmptyState renders one visible absence-state demo with recovery action.
func EmptyState() gx.Node {
	node, err := ui.EmptyState(ui.EmptyStateProps{
		Title:      "No files yet",
		EventLabel: "Upload file",
		Event: ui.ControlProps{
			Action:   "empty.upload",
			SourceID: "empty-state",
		},
	})
	if err != nil {
		return demoError(err)
	}
	return demoWidget("empty-state", node)
}

// LoadingState renders one busy progress-aware demo.
func LoadingState() gx.Node {
	progress := 67
	node, err := ui.LoadingState(ui.LoadingStateProps{
		Message:         "Importing evidence",
		ProgressLabel:   "Import progress",
		ProgressPercent: &progress,
	})
	if err != nil {
		return demoError(err)
	}
	return demoWidget("loading-state", node)
}

// ResultPanel renders one operation outcome demo with follow-up actions.
func ResultPanel() gx.Node {
	node, err := ui.ResultPanel(ui.ResultPanelProps{
		Status:  ui.StatusSuccess,
		Title:   "Import complete",
		Summary: "12 rows accepted and 2 flagged for review.",
		Events: []ui.EventBarAction{
			{
				Label:   "View rows",
				Variant: ui.ButtonSecondary,
				Control: ui.ControlProps{
					Action:   "result.view",
					SourceID: "result-panel-view",
				},
			},
			{
				Label:   "Download log",
				Icon:    ui.IconProps{Name: ui.IconNameDownload},
				Variant: ui.ButtonGhost,
				Control: ui.ControlProps{
					Action:   "result.download",
					SourceID: "result-panel-download",
				},
			},
		},
	})
	if err != nil {
		return demoError(err)
	}
	return demoWidget("result-panel", node)
}

// ErrorBanner renders one recoverable error surface demo.
func ErrorBanner() gx.Node {
	node, err := ui.ErrorBanner(ui.ErrorBannerProps{
		Message:      "Provider is temporarily unavailable.",
		RetryLabel:   "Retry",
		Retry:        ui.ControlProps{Action: "error.retry", SourceID: "error-banner"},
		DismissLabel: "Dismiss",
		Dismiss:      ui.ControlProps{Action: "error.dismiss", SourceID: "error-banner"},
	})
	if err != nil {
		return demoError(err)
	}
	return demoWidget("error-banner", node)
}

// AppShell renders one full-document shell demo through the public app-shell API.
func AppShell() gx.Node {
	node, err := ui.AppShell(ui.AppShellProps{
		Title:        "Workspace",
		CSSAssetURLs: []string{"/assets/bus-ui-demo/bus-ui.css"},
		NavNodes: []gx.Node{
			gx.Element("a", gx.Props{"href": "/docs/overview/index.html"}, gx.Text("Overview")),
			gx.Element("a", gx.Props{"href": "/docs/billing/index.html"}, gx.Text("Billing")),
		},
		MainNodes: []gx.Node{
			gx.Element("section", nil,
				gx.Element("h1", nil, gx.Text("Workspace")),
				gx.Element("p", nil, gx.Text("Review imports and open billing events.")),
			),
		},
		FooterNodes: []gx.Node{
			gx.Element("p", nil, gx.Text("Bus UI shell preview")),
		},
	})
	if err != nil {
		return demoError(err)
	}
	return demoWidget("app-shell", node)
}

// PageShell renders one in-document page shell demo.
func PageShell() gx.Node {
	node, err := ui.PageShell(ui.PageShellProps{
		HeaderNodes: []gx.Node{
			gx.Element("h1", nil, gx.Text("Files")),
		},
		MainNodes: []gx.Node{
			gx.Element("section", nil, gx.Text("Uploaded accounting files")),
		},
		FooterNodes: []gx.Node{
			gx.Element("p", nil, gx.Text("Last import finished")),
		},
	})
	if err != nil {
		return demoError(err)
	}
	return demoWidget("page-shell", node)
}

// SidebarNav renders one collapsible left-rail navigation demo.
func SidebarNav() gx.Node {
	node, err := ui.SidebarNav(ui.SidebarNavProps{
		AppLabel:    "Accounting",
		AppHref:     "/docs/accounting/index.html",
		ToggleLabel: "Toggle navigation",
		Items: []ui.SidebarNavItemProps{
			{ID: "overview", Label: "Overview", Href: "/docs/accounting/index.html", Active: true},
			{ID: "files", Label: "Files", Href: "/docs/accounting/files/index.html"},
			{
				ID:    "refresh",
				Label: "Refresh",
				Control: ui.ControlProps{
					Action:   "sidebar.refresh",
					SourceID: "sidebar-nav-refresh",
				},
			},
		},
	})
	if err != nil {
		return demoError(err)
	}
	return demoWidget("sidebar-nav", node)
}

// SidebarShell renders one left-rail shell composition demo.
func SidebarShell() gx.Node {
	navNode, err := ui.SidebarNav(ui.SidebarNavProps{
		AppLabel:    "Workspace",
		AppHref:     "/docs/overview/index.html",
		ToggleLabel: "Toggle navigation",
		Items: []ui.SidebarNavItemProps{
			{ID: "imports", Label: "Imports", Href: "/docs/imports/index.html", Active: true},
			{ID: "exports", Label: "Exports", Href: "/docs/exports/index.html"},
		},
	})
	if err != nil {
		return demoError(err)
	}
	node, err := ui.SidebarShell(ui.SidebarShellProps{
		NavNodes: []gx.Node{navNode},
		MainNodes: []gx.Node{
			gx.Element("section", nil,
				gx.Element("h1", nil, gx.Text("Imports")),
				gx.Element("p", nil, gx.Text("Selected view content stays beside the rail.")),
			),
		},
	})
	if err != nil {
		return demoError(err)
	}
	return demoWidget("sidebar-shell", node)
}

// ShellActionPanel renders one bounded shell action panel demo.
func ShellActionPanel() gx.Node {
	node, err := ui.ShellActionPanelNodeChecked(ui.ShellActionPanelCheckedProps{
		Title:    "Review import",
		Subtitle: "Validate the imported evidence before posting.",
		Body: []gx.Node{
			gx.Element("p", nil, gx.Text("2 receipts still need category confirmation.")),
		},
		Actions: []ui.ShellActionPanelControl{
			{
				Label:   "Approve",
				Variant: ui.ButtonSecondary,
				Control: ui.ControlProps{
					Action:   "shell.approve",
					SourceID: "shell-action-panel-approve",
				},
			},
		},
		Footer: []gx.Node{
			gx.Element("p", nil, gx.Text("Posting stays locked until every receipt matches.")),
		},
	})
	if err != nil {
		return demoError(err)
	}
	return demoWidget("shell-action-panel", node)
}

func demoWidget(id string, node gx.Node) gx.Node {
	return gx.Element("div", gx.Props{"data-bus-ui-demo-widget": id}, node)
}

func demoError(err error) gx.Node {
	return gx.Element("p", gx.Props{"class": "bus-ui-demo-error"}, gx.Text(err.Error()))
}

func noOpMenuClick(ui.MenuItemEvent) {}
