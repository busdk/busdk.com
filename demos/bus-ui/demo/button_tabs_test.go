package demo

import (
	"os"
	"strings"
	"testing"
)

func TestButtonDocsPageUsesReusableExampleTabs(t *testing.T) {
	t.Parallel()

	body, err := os.ReadFile("../../../docs/gx-ui/bus-ui/components/action/button/index.html")
	if err != nil {
		t.Fatalf("ReadFile(button docs page) error = %v", err)
	}
	text := string(body)

	for _, want := range []string{
		`role="tablist"`,
		`aria-orientation="horizontal"`,
		`id="button-example-tab-go-api"`,
		`id="button-example-tab-gx-source"`,
		`id="button-example-panel-go-api"`,
		`id="button-example-panel-gx-source"`,
		`aria-controls="button-example-panel-go-api"`,
		`aria-controls="button-example-panel-gx-source"`,
		`aria-labelledby="button-example-tab-go-api"`,
		`aria-labelledby="button-example-tab-gx-source"`,
		`aria-selected="true"`,
		`aria-selected="false"`,
		`data-bus-ui-demo="button"`,
		`data-bus-ui-demo-loader`,
		`assets/bus-ui-demo/loader.js`,
		`example-tabs.js`,
	} {
		if !strings.Contains(text, want) {
			t.Fatalf("button docs page missing %q", want)
		}
	}

	for _, check := range []struct {
		name string
		got  int
		want int
	}{
		{name: "tablist", got: strings.Count(text, `role="tablist"`), want: 1},
		{name: "tabs", got: strings.Count(text, `id="button-example-tab-`), want: 2},
		{name: "panels", got: strings.Count(text, `id="button-example-panel-`), want: 2},
		{name: "selected tabs", got: strings.Count(text, `aria-selected="true"`), want: 1},
		{name: "unselected tabs", got: strings.Count(text, `aria-selected="false"`), want: 1},
		{name: "live demo placeholder", got: strings.Count(text, `data-bus-ui-demo="button"`), want: 1},
		{name: "demo loader script", got: strings.Count(text, `data-bus-ui-demo-loader`), want: 1},
		{name: "loader asset path", got: strings.Count(text, `assets/bus-ui-demo/loader.js`), want: 1},
		{name: "tabs script", got: strings.Count(text, `example-tabs.js`), want: 1},
	} {
		if check.got != check.want {
			t.Fatalf("%s count = %d, want %d", check.name, check.got, check.want)
		}
	}
}
