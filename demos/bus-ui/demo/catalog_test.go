package demo

import (
	"reflect"
	"strings"
	"testing"

	gx "github.com/busdk/bus-gx/pkg/gx"
)

func TestIDsExposeButtonDemo(t *testing.T) {
	if got, want := IDs(), []string{"button"}; !reflect.DeepEqual(got, want) {
		t.Fatalf("IDs() = %#v, want %#v", got, want)
	}
}

func TestButtonDemoRendersRealBusUIButton(t *testing.T) {
	root, ok := Lookup("button")
	if !ok {
		t.Fatal("button demo is not registered")
	}
	html, err := gx.RenderHTML(root())
	if err != nil {
		t.Fatalf("RenderHTML(button demo) failed: %v", err)
	}
	for _, want := range []string{"bus-ui-btn", "bus-ui-btn-primary", "Save draft"} {
		if !strings.Contains(html, want) {
			t.Fatalf("button demo HTML %q does not contain %q", html, want)
		}
	}
}
