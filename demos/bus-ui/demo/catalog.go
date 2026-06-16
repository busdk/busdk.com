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
	"button": Button,
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
		return gx.Element("p", gx.Props{"class": "bus-ui-demo-error"}, gx.Text(err.Error()))
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
