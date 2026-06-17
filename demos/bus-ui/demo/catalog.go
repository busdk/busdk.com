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
	"button":      Button,
	"event-bar":   EventBar,
	"icon-button": IconButton,
	"link-button": LinkButton,
	"menu":        Menu,
	"navigation":  Navigation,
	"tabs":        Tabs,
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

func demoWidget(id string, node gx.Node) gx.Node {
	return gx.Element("div", gx.Props{"data-bus-ui-demo-widget": id}, node)
}

func demoError(err error) gx.Node {
	return gx.Element("p", gx.Props{"class": "bus-ui-demo-error"}, gx.Text(err.Error()))
}

func noOpMenuClick(ui.MenuItemEvent) {}
