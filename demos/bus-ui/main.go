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
	mountBusDKTopHeaders(document)
	mountBusDKFooters(document)
	mountBusDKProductSideNavs(document)
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

func mountBusDKTopHeaders(document js.Value) {
	nodes := document.Call("querySelectorAll", "header.site-header[data-busdk-top-header], header.site-header[data-gx-ui-top-header]")
	for i := 0; i < nodes.Length(); i++ {
		el := nodes.Index(i)
		fallbackHTML := el.Get("innerHTML").String()
		setTopHeaderState(el, "mounting")
		selector := ensureID(el, "busdk-top-header-root", i)
		navID := dataAttribute(el, "busdkTopNav")
		if navID == "" && hasAttribute(el, "data-gx-ui-top-header") {
			navID = "gx-ui"
		}
		if navID == "" {
			navID = "site"
		}
		headerBaseURL := busDKTopHeaderBaseURL(document, el, navID)
		currentID := dataAttribute(el, "busdkCurrent")
		if currentID == "" && navID == "gx-ui" {
			currentID = gxUITopHeaderCurrentID(document)
		}
		headerNavID := navID
		headerCurrentID := currentID
		headerBaseHref := headerBaseURL
		root := func() gx.Node {
			return demo.BusDKTopHeader(headerNavID, headerBaseHref, headerCurrentID)
		}
		if _, err := gxwasm.Mount(selector, ui.GxWASMRoot(ui.GxNodeRoot(root)), gxwasm.Options{
			OnError: func(err error) {
				setTopHeaderFallback(el, fallbackHTML)
			},
		}); err != nil {
			setTopHeaderFallback(el, fallbackHTML)
			continue
		}
		setTopHeaderState(el, "mounted")
	}
}

func mountBusDKFooters(document js.Value) {
	nodes := document.Call("querySelectorAll", "footer.site-footer[data-busdk-footer], footer.site-footer[data-gx-ui-footer]")
	for i := 0; i < nodes.Length(); i++ {
		el := nodes.Index(i)
		fallbackHTML := el.Get("innerHTML").String()
		setFooterState(el, "mounting")
		selector := ensureID(el, "busdk-footer-root", i)
		root := func() gx.Node {
			return demo.BusDKFooter()
		}
		if _, err := gxwasm.Mount(selector, ui.GxWASMRoot(ui.GxNodeRoot(root)), gxwasm.Options{
			OnError: func(err error) {
				setFooterFallback(el, fallbackHTML)
			},
		}); err != nil {
			setFooterFallback(el, fallbackHTML)
			continue
		}
		setFooterState(el, "mounted")
	}
}

func mountBusDKProductSideNavs(document js.Value) {
	nodes := document.Call("querySelectorAll", "aside.gx-side-nav[data-busdk-side-nav]")
	for i := 0; i < nodes.Length(); i++ {
		el := nodes.Index(i)
		fallbackHTML := el.Get("innerHTML").String()
		el.Call("setAttribute", "data-busdk-side-nav-state", "mounting")
		selector := ensureID(el, "busdk-side-nav-root", i)
		navID := dataAttribute(el, "busdkSideNav")
		currentID := dataAttribute(el, "busdkCurrent")
		baseURL := resolvedDataBaseURL(document, el, "busdkSideNavBase")
		navCurrentID := currentID
		navBaseURL := baseURL
		navKey := navID
		root := func() gx.Node {
			return demo.BusDKProductSideNav(navKey, navCurrentID, navBaseURL)
		}
		if _, err := gxwasm.Mount(selector, ui.GxWASMRoot(ui.GxNodeRoot(root)), gxwasm.Options{
			OnError: func(err error) {
				setBusDKSideNavFallback(el, fallbackHTML)
			},
		}); err != nil {
			setBusDKSideNavFallback(el, fallbackHTML)
			continue
		}
		if count := demo.BusDKProductSideNavCurrentCount(navID, currentID); count != 1 {
			js.Global().Get("console").Call("warn", "BusDK side nav expected exactly one current entry", navID, currentID, count)
		}
		el.Call("setAttribute", "data-busdk-side-nav-state", "mounted")
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

func busDKSiteBaseURL(document js.Value, el js.Value) string {
	for _, key := range []string{"busdkSiteBase", "gxUiSiteBase"} {
		if value := dataAttribute(el, key); value != "" {
			return value
		}
	}
	window := js.Global().Get("window")
	if baseURL := window.Get("__busDKSiteBaseURL"); !baseURL.IsUndefined() && !baseURL.IsNull() && baseURL.String() != "" {
		return baseURL.String()
	}
	if gxBase := gxUIDocsBaseURL(document); gxBase != "" {
		return js.Global().Get("URL").New("../", gxBase).Get("href").String()
	}
	return ""
}

func busDKTopHeaderBaseURL(document js.Value, el js.Value, navID string) string {
	if baseURL := resolvedDataBaseURL(document, el, "busdkTopNavBase"); baseURL != "" {
		return baseURL
	}
	if navID == "gx-ui" {
		return gxUIDocsBaseURL(document)
	}
	return busDKSiteBaseURL(document, el)
}

func resolvedDataBaseURL(document js.Value, el js.Value, name string) string {
	baseURL := dataAttribute(el, name)
	if baseURL == "" {
		return ""
	}
	location := document.Get("location")
	if location.IsUndefined() || location.IsNull() {
		return baseURL
	}
	href := location.Get("href")
	if href.IsUndefined() || href.IsNull() || href.String() == "" {
		return baseURL
	}
	return js.Global().Get("URL").New(baseURL, href.String()).Get("href").String()
}

func gxUITopHeaderCurrentID(document js.Value) string {
	node := document.Call("querySelector", "aside.gx-side-nav[data-gx-ui-current]")
	if node.IsUndefined() || node.IsNull() {
		return ""
	}
	return dataAttribute(node, "gxUiCurrent")
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
	setTopHeaderState(el, "failed")
}

func setFooterFallback(el js.Value, html string) {
	el.Set("innerHTML", html)
	setFooterState(el, "failed")
}

func setBusDKSideNavFallback(el js.Value, html string) {
	el.Set("innerHTML", html)
	el.Call("setAttribute", "data-busdk-side-nav-state", "failed")
}

func setTopHeaderState(el js.Value, state string) {
	if hasAttribute(el, "data-busdk-top-header") {
		el.Call("setAttribute", "data-busdk-top-header-state", state)
	}
	if hasAttribute(el, "data-gx-ui-top-header") {
		el.Call("setAttribute", "data-gx-ui-top-header-state", state)
	}
}

func setFooterState(el js.Value, state string) {
	if hasAttribute(el, "data-busdk-footer") {
		el.Call("setAttribute", "data-busdk-footer-state", state)
	}
	if hasAttribute(el, "data-gx-ui-footer") {
		el.Call("setAttribute", "data-gx-ui-footer-state", state)
	}
}

func hasAttribute(el js.Value, name string) bool {
	return el.Call("hasAttribute", name).Bool()
}

func dataAttribute(el js.Value, name string) string {
	value := el.Get("dataset").Get(name)
	if value.IsUndefined() || value.IsNull() {
		return ""
	}
	return value.String()
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
