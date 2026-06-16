package demo

import (
	"fmt"
	"strings"

	gx "github.com/busdk/bus-gx/pkg/gx"
	ui "github.com/busdk/bus-ui/pkg/ui"
)

// PlaceholderHTML renders one checked static docs placeholder for a Bus UI demo root.
func PlaceholderHTML(id string, label string) (string, error) {
	node, err := PlaceholderNode(id, label)
	if err != nil {
		return "", err
	}
	return gx.RenderHTML(node)
}

// PlaceholderNode renders one static docs placeholder that uses the shared Bus UI loader.
func PlaceholderNode(id string, label string) (gx.Node, error) {
	demoID := strings.TrimSpace(id)
	if demoID == "" {
		return nil, fmt.Errorf("demo placeholder id required")
	}
	loader, err := ui.Loader(ui.LoaderProps{
		Label: label,
		Attrs: map[string]string{
			"class": "bus-ui-demo-placeholder-loader",
		},
	})
	if err != nil {
		return nil, err
	}
	return gx.Element("div", gx.Props{"data-bus-ui-demo": demoID}, loader), nil
}
