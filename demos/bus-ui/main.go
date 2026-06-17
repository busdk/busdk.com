//go:build js && wasm

package main

import (
	"fmt"
	"syscall/js"

	gx "github.com/busdk/bus-gx/pkg/gx"
	gxwasm "github.com/busdk/bus-gx/pkg/gx/wasm"
	ui "github.com/busdk/bus-ui/pkg/ui"
	"github.com/busdk/busdk.com/demos/bus-ui/demo"
)

var retainedDemoCallbacks []js.Func

func main() {
	mountAll()
	select {}
}

func mountAll() {
	document := js.Global().Get("document")
	if document.IsUndefined() || document.IsNull() {
		return
	}
	mountBusUIDemos(document)
	mountGXUITopHeaders(document)
	mountGXUIFooters(document)
	mountGXUISideNavs(document)
}

func mountBusUIDemos(document js.Value) {
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
		selector := ensureID(el, "bus-ui-demo-root", i)
		if _, err := gxwasm.Mount(selector, ui.GxWASMRoot(ui.GxNodeRoot(root)), gxwasm.Options{
			OnError: func(err error) {
				setFallback(el, "Bus UI demo failed to render.")
			},
		}); err != nil {
			setFallback(el, "Bus UI demo failed to render.")
			continue
		}
		bindDemoInteractions(el)
		el.Call("setAttribute", "data-bus-ui-demo-state", "mounted")
	}
}

func mountGXUISideNavs(document js.Value) {
	nodes := document.Call("querySelectorAll", "aside.gx-side-nav[data-gx-ui-side-nav]")
	baseURL := gxUIDocsBaseURL(document)
	for i := 0; i < nodes.Length(); i++ {
		el := nodes.Index(i)
		currentID := el.Call("getAttribute", "data-gx-ui-current").String()
		el.Call("setAttribute", "data-gx-ui-side-nav-state", "mounting")
		selector := ensureID(el, "gx-ui-side-nav-root", i)
		navCurrentID := currentID
		navBaseURL := baseURL
		root := func() gx.Node {
			return demo.GXUISideNav(navCurrentID, navBaseURL)
		}
		if _, err := gxwasm.Mount(selector, ui.GxWASMRoot(ui.GxNodeRoot(root)), gxwasm.Options{
			OnError: func(err error) {
				setSideNavFallback(el, "GX/UI navigation failed to render.")
			},
		}); err != nil {
			setSideNavFallback(el, "GX/UI navigation failed to render.")
			continue
		}
		if count := demo.GXUISideNavCurrentCount(currentID); count != 1 {
			js.Global().Get("console").Call("warn", "GX/UI side nav expected exactly one current entry", currentID, count)
		}
		el.Call("setAttribute", "data-gx-ui-side-nav-state", "mounted")
	}
}

func mountGXUITopHeaders(document js.Value) {
	nodes := document.Call("querySelectorAll", "header.site-header[data-gx-ui-top-header]")
	baseURL := gxUIDocsBaseURL(document)
	for i := 0; i < nodes.Length(); i++ {
		el := nodes.Index(i)
		fallbackHTML := el.Get("innerHTML").String()
		el.Call("setAttribute", "data-gx-ui-top-header-state", "mounting")
		selector := ensureID(el, "gx-ui-top-header-root", i)
		headerBaseURL := baseURL
		root := func() gx.Node {
			return demo.GXUITopHeader(headerBaseURL)
		}
		if _, err := gxwasm.Mount(selector, ui.GxWASMRoot(ui.GxNodeRoot(root)), gxwasm.Options{
			OnError: func(err error) {
				setTopHeaderFallback(el, fallbackHTML)
			},
		}); err != nil {
			setTopHeaderFallback(el, fallbackHTML)
			continue
		}
		el.Call("setAttribute", "data-gx-ui-top-header-state", "mounted")
	}
}

func mountGXUIFooters(document js.Value) {
	nodes := document.Call("querySelectorAll", "footer.site-footer[data-gx-ui-footer]")
	for i := 0; i < nodes.Length(); i++ {
		el := nodes.Index(i)
		fallbackHTML := el.Get("innerHTML").String()
		el.Call("setAttribute", "data-gx-ui-footer-state", "mounting")
		selector := ensureID(el, "gx-ui-footer-root", i)
		root := func() gx.Node {
			return demo.GXUIFooter()
		}
		if _, err := gxwasm.Mount(selector, ui.GxWASMRoot(ui.GxNodeRoot(root)), gxwasm.Options{
			OnError: func(err error) {
				setFooterFallback(el, fallbackHTML)
			},
		}); err != nil {
			setFooterFallback(el, fallbackHTML)
			continue
		}
		el.Call("setAttribute", "data-gx-ui-footer-state", "mounted")
	}
}

func gxUIDocsBaseURL(document js.Value) string {
	window := js.Global().Get("window")
	if baseURL := window.Get("__gxUIDocsBaseURL"); !baseURL.IsUndefined() && !baseURL.IsNull() && baseURL.String() != "" {
		return baseURL.String()
	}
	if baseURL := window.Get("__gxUISideNavBaseURL"); !baseURL.IsUndefined() && !baseURL.IsNull() && baseURL.String() != "" {
		return baseURL.String()
	}
	script := document.Call("querySelector", "script[src$='side-nav.js']")
	if !script.IsUndefined() && !script.IsNull() {
		src := script.Get("src")
		if !src.IsUndefined() && !src.IsNull() && src.String() != "" {
			return js.Global().Get("URL").New("./", src).Get("href").String()
		}
	}
	return ""
}

func ensureID(el js.Value, prefix string, index int) string {
	id := el.Get("id").String()
	if id == "" {
		id = fmt.Sprintf("%s-%d", prefix, index+1)
		el.Set("id", id)
	}
	return "#" + id
}

func setFallback(el js.Value, text string) {
	el.Set("textContent", text)
	el.Get("classList").Call("add", "bus-ui-demo-fallback")
	el.Call("setAttribute", "data-bus-ui-demo-state", "failed")
}

func setSideNavFallback(el js.Value, text string) {
	el.Set("textContent", text)
	el.Get("classList").Call("add", "bus-ui-demo-fallback")
	el.Call("setAttribute", "data-gx-ui-side-nav-state", "failed")
}

func setTopHeaderFallback(el js.Value, html string) {
	el.Set("innerHTML", html)
	el.Call("setAttribute", "data-gx-ui-top-header-state", "failed")
}

func setFooterFallback(el js.Value, html string) {
	el.Set("innerHTML", html)
	el.Call("setAttribute", "data-gx-ui-footer-state", "failed")
}

func bindDemoInteractions(root js.Value) {
	if !root.Truthy() {
		return
	}
	button := root.Call("querySelector", `[data-bus-ui-demo-action="button-click"]`)
	status := root.Call("querySelector", `[data-bus-ui-demo-status="button"]`)
	if !button.Truthy() || !status.Truthy() {
		return
	}
	cb := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		status.Set("textContent", "Button clicked")
		return nil
	})
	retainedDemoCallbacks = append(retainedDemoCallbacks, cb)
	button.Call("addEventListener", "click", cb)
}
