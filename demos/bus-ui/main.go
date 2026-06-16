//go:build js && wasm

package main

import (
	"fmt"
	"syscall/js"

	gxwasm "github.com/busdk/bus-gx/pkg/gx/wasm"
	ui "github.com/busdk/bus-ui/pkg/ui"
	"github.com/busdk/busdk.com/demos/bus-ui/demo"
)

func main() {
	mountAll()
	select {}
}

func mountAll() {
	document := js.Global().Get("document")
	if document.IsUndefined() || document.IsNull() {
		return
	}
	nodes := document.Call("querySelectorAll", "[data-bus-ui-demo]")
	for i := 0; i < nodes.Length(); i++ {
		el := nodes.Index(i)
		id := el.Get("dataset").Get("busUiDemo").String()
		el.Call("setAttribute", "data-bus-ui-demo-state", "mounting")
		root, ok := demo.Lookup(id)
		if !ok {
			setFallback(el, "Bus UI demo is unavailable.")
			continue
		}
		selector := ensureID(el, i)
		if _, err := gxwasm.Mount(selector, ui.GxWASMRoot(ui.GxNodeRoot(root)), gxwasm.Options{
			OnError: func(err error) {
				setFallback(el, "Bus UI demo failed to render.")
			},
		}); err != nil {
			setFallback(el, "Bus UI demo failed to render.")
			continue
		}
		el.Call("setAttribute", "data-bus-ui-demo-state", "mounted")
	}
}

func ensureID(el js.Value, index int) string {
	id := el.Get("id").String()
	if id == "" {
		id = fmt.Sprintf("bus-ui-demo-root-%d", index+1)
		el.Set("id", id)
	}
	return "#" + id
}

func setFallback(el js.Value, text string) {
	el.Set("textContent", text)
	el.Get("classList").Call("add", "bus-ui-demo-fallback")
	el.Call("setAttribute", "data-bus-ui-demo-state", "failed")
}
