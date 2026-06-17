package main

import (
	"bufio"
	"bytes"
	"context"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

const buttonDocsPreviewRelativePath = "docs/gx-ui/bus-ui/components/action/button/index.html"

func TestButtonDocsPreviewStaticMarkupUsesPublishedDocsTree(t *testing.T) {
	t.Parallel()

	path, err := buttonDocsPreviewPath()
	if err != nil {
		t.Fatal(err)
	}

	wantSuffix := filepath.FromSlash(buttonDocsPreviewRelativePath)
	if !strings.HasSuffix(path, wantSuffix) {
		t.Fatalf("button docs path = %q, want suffix %q", path, wantSuffix)
	}

	body, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("ReadFile(%s) error = %v", path, err)
	}
	text := string(body)
	for _, want := range []string{
		`data-bus-ui-demo="button"`,
		"assets/bus-ui-demo/wasm_exec.js",
		"assets/bus-ui-demo/loader.js",
		"data-bus-ui-demo-loader",
		"Loading Button demo...",
	} {
		if !strings.Contains(text, want) {
			t.Fatalf("%s missing %q", path, want)
		}
	}
}

func TestButtonDocsPreviewBrowserSmoke(t *testing.T) {
	t.Parallel()

	browser, ok := headlessBrowserBinary()
	if !ok {
		t.Skip("skip button browser regression: headless Chrome/Chromium unavailable")
	}

	buttonPath, err := buttonDocsPreviewPath()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("file-preview", func(t *testing.T) {
		output := runHeadlessBrowser(t, browser, fileURL(buttonPath), true)
		for _, want := range []string{
			`data-bus-ui-demo-state="failed"`,
			`data-bus-ui-demo-asset="css"`,
			"Bus UI demos need a local HTTP server to load WebAssembly.",
		} {
			if !strings.Contains(output, want) {
				t.Fatalf("file:// output missing %q in:\n%s", want, output)
			}
		}
		if strings.Contains(output, `data-bus-ui-demo-widget="button"`) {
			t.Fatalf("file:// preview unexpectedly mounted the Button demo:\n%s", output)
		}
	})

	t.Run("http-preview", func(t *testing.T) {
		baseURL, cleanup := serveRepoOnLoopback(t)
		defer cleanup()

		output := runHeadlessBrowser(t, browser, baseURL+buttonDocsPreviewRelativePath, false)
		for _, want := range []string{
			`data-bus-ui-demo-state="mounted"`,
			`data-bus-ui-demo-asset="css"`,
			`data-bus-ui-demo-widget="button"`,
			"bus-ui-btn",
			"bus-ui-btn-primary",
			`data-bus-ui-demo-status="button"`,
		} {
			if !strings.Contains(output, want) {
				t.Fatalf("http output missing %q in:\n%s", want, output)
			}
		}
		if strings.Contains(output, "Bus UI demos need a local HTTP server to load WebAssembly.") {
			t.Fatalf("http preview unexpectedly used the file:// fallback:\n%s", output)
		}
	})
}

func buttonDocsPreviewPath() (string, error) {
	repoRoot, err := repoRootDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(repoRoot, filepath.FromSlash(buttonDocsPreviewRelativePath)), nil
}

func repoRootDir() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return filepath.Clean(filepath.Join(wd, "..", "..")), nil
}

func fileURL(path string) string {
	return (&url.URL{Scheme: "file", Path: path}).String()
}

func headlessBrowserBinary() (string, bool) {
	candidates := []string{
		"/Applications/Google Chrome.app/Contents/MacOS/Google Chrome",
		"/Applications/Chromium.app/Contents/MacOS/Chromium",
		"/Applications/Microsoft Edge.app/Contents/MacOS/Microsoft Edge",
		"google-chrome",
		"google-chrome-stable",
		"chromium",
		"chromium-browser",
		"chrome",
	}
	for _, candidate := range candidates {
		if candidate == "" {
			continue
		}
		if filepath.IsAbs(candidate) {
			info, err := os.Stat(candidate)
			if err == nil && !info.IsDir() {
				return candidate, true
			}
			continue
		}
		if path, err := exec.LookPath(candidate); err == nil {
			return path, true
		}
	}
	return "", false
}

func runHeadlessBrowser(t *testing.T, browser string, pageURL string, allowFileAccess bool) string {
	t.Helper()

	ctx, cancel := context.WithTimeout(context.Background(), 45*time.Second)
	defer cancel()

	userDataDir := t.TempDir()
	args := []string{
		"--headless",
		"--disable-gpu",
		"--disable-background-networking",
		"--disable-dev-shm-usage",
		"--disable-extensions",
		"--disable-sync",
		"--metrics-recording-only",
		"--no-default-browser-check",
		"--no-first-run",
		"--run-all-compositor-stages-before-draw",
		"--remote-debugging-port=0",
		"--user-data-dir=" + userDataDir,
	}
	if allowFileAccess {
		args = append(args, "--allow-file-access-from-files")
	}
	cmd := exec.CommandContext(ctx, browser, args...)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err := cmd.Start(); err != nil {
		t.Fatalf("start headless browser for %s failed: %v\n%s", pageURL, err, stderr.String())
	}
	cleanup := func() {
		_ = cmd.Process.Kill()
		_ = cmd.Wait()
	}
	defer cleanup()

	wsURL, err := waitForPageWebSocketURL(ctx, userDataDir)
	if err != nil {
		t.Fatalf("discover page websocket for %s failed: %v\n%s", pageURL, err, stderr.String())
	}
	client, err := dialChromeDebugger(wsURL)
	if err != nil {
		t.Fatalf("connect page websocket for %s failed: %v\n%s", pageURL, err, stderr.String())
	}
	defer client.Close()

	if _, err := client.call(ctx, "", "Page.enable", nil); err != nil {
		t.Fatalf("enable Page domain for %s failed: %v\n%s", pageURL, err, stderr.String())
	}
	if _, err := client.call(ctx, "", "Page.bringToFront", nil); err != nil {
		t.Fatalf("bring page to front for %s failed: %v\n%s", pageURL, err, stderr.String())
	}
	if _, err := client.call(ctx, "", "Runtime.enable", nil); err != nil {
		t.Fatalf("enable Runtime domain for %s failed: %v\n%s", pageURL, err, stderr.String())
	}
	if _, err := client.call(ctx, "", "Page.navigate", map[string]any{"url": pageURL}); err != nil {
		t.Fatalf("navigate %s failed: %v\n%s", pageURL, err, stderr.String())
	}

	want := []string{
		`data-bus-ui-demo-state="failed"`,
		`data-bus-ui-demo-asset="css"`,
		"Bus UI demos need a local HTTP server to load WebAssembly.",
	}
	if !allowFileAccess {
		want = []string{
			`data-bus-ui-demo-state="mounted"`,
			`data-bus-ui-demo-asset="css"`,
			"Save draft",
			`data-bus-ui-demo-widget="button"`,
		}
	}

	snapshot, err := client.waitForSnapshot(ctx, "", want...)
	if err != nil {
		t.Fatalf("capture browser DOM for %s failed: %v\n%s", pageURL, err, stderr.String())
	}
	if !allowFileAccess {
		if err := client.clickAndWaitForText(ctx, "", `[data-bus-ui-demo-action="button-click"]`, `[data-bus-ui-demo-status="button"]`, "Button clicked"); err != nil {
			t.Fatalf("click button demo for %s failed: %v\n%s", pageURL, err, stderr.String())
		}
		snapshot, err = client.waitForSnapshot(ctx, "", []string{
			`data-bus-ui-demo-state="mounted"`,
			`data-bus-ui-demo-asset="css"`,
			`data-bus-ui-demo-widget="button"`,
			`data-bus-ui-demo-action="button-click"`,
			`data-bus-ui-demo-status="button"`,
			"Button clicked",
		}...)
		if err != nil {
			t.Fatalf("capture clicked browser DOM for %s failed: %v\n%s", pageURL, err, stderr.String())
		}
	}
	return snapshot
}

func serveRepoOnLoopback(t *testing.T) (string, func()) {
	t.Helper()

	repoRoot, err := repoRootDir()
	if err != nil {
		t.Fatal(err)
	}

	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Skipf("skip http browser regression: loopback bind unavailable: %v", err)
	}

	server := &http.Server{Handler: http.FileServer(http.Dir(repoRoot))}
	go func() {
		_ = server.Serve(listener)
	}()

	cleanup := func() {
		_ = server.Close()
		_ = listener.Close()
	}
	return "http://" + listener.Addr().String() + "/", cleanup
}

func waitForPageWebSocketURL(ctx context.Context, userDataDir string) (string, error) {
	portFile := filepath.Join(userDataDir, "DevToolsActivePort")
	for ctx.Err() == nil {
		data, err := os.ReadFile(portFile)
		if err == nil {
			lines := strings.Split(strings.TrimSpace(string(data)), "\n")
			if len(lines) >= 2 {
				port := lines[0]
				if wsURL, err := pageWebSocketURL(ctx, port); err == nil {
					return wsURL, nil
				}
			}
		}
		time.Sleep(100 * time.Millisecond)
	}
	return "", ctx.Err()
}

func pageWebSocketURL(ctx context.Context, port string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://127.0.0.1:"+port+"/json/version", nil)
	if err != nil {
		return "", err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("json/version status %s", resp.Status)
	}
	var payload struct {
		WebSocketDebuggerURL string `json:"webSocketDebuggerUrl"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		return "", err
	}
	if payload.WebSocketDebuggerURL == "" {
		return "", errors.New("browser websocket url missing")
	}
	listReq, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://127.0.0.1:"+port+"/json/list", nil)
	if err != nil {
		return "", err
	}
	listResp, err := http.DefaultClient.Do(listReq)
	if err != nil {
		return "", err
	}
	defer listResp.Body.Close()
	if listResp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("json/list status %s", listResp.Status)
	}
	var targets []struct {
		Type                 string `json:"type"`
		WebSocketDebuggerURL string `json:"webSocketDebuggerUrl"`
	}
	if err := json.NewDecoder(listResp.Body).Decode(&targets); err != nil {
		return "", err
	}
	for _, target := range targets {
		if target.Type == "page" && target.WebSocketDebuggerURL != "" {
			return target.WebSocketDebuggerURL, nil
		}
	}
	return "", errors.New("page websocket url missing")
}

type chromeDebugger struct {
	conn   *websocketConn
	nextID int
}

func dialChromeDebugger(rawURL string) (*chromeDebugger, error) {
	conn, err := dialWebSocket(rawURL)
	if err != nil {
		return nil, err
	}
	return &chromeDebugger{conn: conn}, nil
}

func (c *chromeDebugger) Close() error {
	return c.conn.Close()
}

func (c *chromeDebugger) call(ctx context.Context, sessionID string, method string, params any) (map[string]any, error) {
	c.nextID++
	id := c.nextID
	req := map[string]any{
		"id":     id,
		"method": method,
	}
	if params != nil {
		req["params"] = params
	}
	if sessionID != "" {
		req["sessionId"] = sessionID
	}
	payload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	if err := c.conn.writeText(payload); err != nil {
		return nil, err
	}
	for {
		msg, err := c.conn.readText(ctx)
		if err != nil {
			return nil, err
		}
		var envelope struct {
			ID     int            `json:"id"`
			Result map[string]any `json:"result"`
			Error  *struct {
				Message string `json:"message"`
			} `json:"error"`
		}
		if err := json.Unmarshal(msg, &envelope); err != nil {
			continue
		}
		if envelope.ID != id {
			continue
		}
		if envelope.Error != nil {
			return nil, errors.New(envelope.Error.Message)
		}
		return envelope.Result, nil
	}
}

func (c *chromeDebugger) waitForSnapshot(ctx context.Context, sessionID string, want ...string) (string, error) {
	for i := 0; i < 200; i++ {
		value, err := c.evaluateString(ctx, sessionID, "document.documentElement.outerHTML")
		if err != nil {
			return "", err
		}
		if containsAll(value, want...) {
			return value, nil
		}
		time.Sleep(250 * time.Millisecond)
	}
	return "", context.DeadlineExceeded
}

func (c *chromeDebugger) clickAndWaitForText(ctx context.Context, sessionID string, clickSelector string, statusSelector string, want string) error {
	clickExpr := fmt.Sprintf(`(() => {
		const target = document.querySelector(%q);
		if (!target) {
			return "missing";
		}
		target.click();
		return "clicked";
	})()`, clickSelector)
	clickResult, err := c.evaluateString(ctx, sessionID, clickExpr)
	if err != nil {
		return err
	}
	if clickResult != "clicked" {
		return fmt.Errorf("click target %s not available", clickSelector)
	}
	statusExpr := fmt.Sprintf(`(() => {
		const status = document.querySelector(%q);
		return status ? status.textContent : "";
	})()`, statusSelector)
	for i := 0; i < 200; i++ {
		value, err := c.evaluateString(ctx, sessionID, statusExpr)
		if err != nil {
			return err
		}
		if value == want {
			return nil
		}
		time.Sleep(250 * time.Millisecond)
	}
	return context.DeadlineExceeded
}

func (c *chromeDebugger) evaluateString(ctx context.Context, sessionID string, expression string) (string, error) {
	raw, err := c.call(ctx, sessionID, "Runtime.evaluate", map[string]any{
		"expression":        expression,
		"returnByValue":     true,
		"awaitPromise":      false,
		"replMode":          false,
		"throwOnSideEffect": false,
	})
	if err != nil {
		return "", err
	}
	result, ok := raw["result"].(map[string]any)
	if !ok {
		return "", errors.New("runtime evaluate missing result")
	}
	value, ok := result["value"].(string)
	if !ok {
		return "", errors.New("runtime evaluate missing value")
	}
	return value, nil
}

type websocketConn struct {
	conn net.Conn
}

func dialWebSocket(rawURL string) (*websocketConn, error) {
	parsed, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}
	if parsed.Scheme != "ws" {
		return nil, fmt.Errorf("unexpected websocket scheme %q", parsed.Scheme)
	}
	conn, err := net.Dial("tcp", parsed.Host)
	if err != nil {
		return nil, err
	}

	keyBytes := make([]byte, 16)
	if _, err := rand.Read(keyBytes); err != nil {
		_ = conn.Close()
		return nil, err
	}
	key := base64.StdEncoding.EncodeToString(keyBytes)
	req, err := http.NewRequest(http.MethodGet, rawURL, nil)
	if err != nil {
		_ = conn.Close()
		return nil, err
	}
	req.Header.Set("Host", parsed.Host)
	req.Header.Set("Upgrade", "websocket")
	req.Header.Set("Connection", "Upgrade")
	req.Header.Set("Sec-WebSocket-Key", key)
	req.Header.Set("Sec-WebSocket-Version", "13")
	if err := req.Write(conn); err != nil {
		_ = conn.Close()
		return nil, err
	}
	resp, err := http.ReadResponse(bufio.NewReader(conn), req)
	if err != nil {
		_ = conn.Close()
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusSwitchingProtocols {
		_ = conn.Close()
		return nil, fmt.Errorf("websocket handshake failed: %s", resp.Status)
	}
	accept := resp.Header.Get("Sec-WebSocket-Accept")
	wantAccept := websocketAcceptKey(key)
	if accept != wantAccept {
		_ = conn.Close()
		return nil, fmt.Errorf("websocket accept mismatch: got %q want %q", accept, wantAccept)
	}
	return &websocketConn{conn: conn}, nil
}

func websocketAcceptKey(key string) string {
	sum := sha1.Sum([]byte(key + "258EAFA5-E914-47DA-95CA-C5AB0DC85B11"))
	return base64.StdEncoding.EncodeToString(sum[:])
}

func (w *websocketConn) Close() error {
	return w.conn.Close()
}

func (w *websocketConn) writeText(payload []byte) error {
	var header bytes.Buffer
	header.WriteByte(0x81)
	switch n := len(payload); {
	case n < 126:
		header.WriteByte(byte(n) | 0x80)
	case n <= 65535:
		header.WriteByte(126 | 0x80)
		var ext [2]byte
		binary.BigEndian.PutUint16(ext[:], uint16(n))
		header.Write(ext[:])
	default:
		header.WriteByte(127 | 0x80)
		var ext [8]byte
		binary.BigEndian.PutUint64(ext[:], uint64(n))
		header.Write(ext[:])
	}
	mask := make([]byte, 4)
	if _, err := rand.Read(mask); err != nil {
		return err
	}
	header.Write(mask)
	masked := make([]byte, len(payload))
	for i := range payload {
		masked[i] = payload[i] ^ mask[i%4]
	}
	if _, err := w.conn.Write(header.Bytes()); err != nil {
		return err
	}
	_, err := w.conn.Write(masked)
	return err
}

func (w *websocketConn) readText(ctx context.Context) ([]byte, error) {
	type result struct {
		opcode byte
		data   []byte
		err    error
	}
	ch := make(chan result, 1)
	go func() {
		opcode, data, err := w.readFrame()
		ch <- result{opcode: opcode, data: data, err: err}
	}()
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case r := <-ch:
		if r.err != nil {
			return nil, r.err
		}
		switch r.opcode {
		case 0x1:
			return r.data, nil
		case 0x9:
			_ = w.writeControl(0xA, r.data)
			return w.readText(ctx)
		case 0x8:
			return nil, io.EOF
		default:
			return w.readText(ctx)
		}
	}
}

func (w *websocketConn) writeControl(opcode byte, payload []byte) error {
	var header bytes.Buffer
	header.WriteByte(0x80 | opcode)
	switch n := len(payload); {
	case n < 126:
		header.WriteByte(byte(n) | 0x80)
	case n <= 65535:
		header.WriteByte(126 | 0x80)
		var ext [2]byte
		binary.BigEndian.PutUint16(ext[:], uint16(n))
		header.Write(ext[:])
	default:
		header.WriteByte(127 | 0x80)
		var ext [8]byte
		binary.BigEndian.PutUint64(ext[:], uint64(n))
		header.Write(ext[:])
	}
	mask := make([]byte, 4)
	if _, err := rand.Read(mask); err != nil {
		return err
	}
	header.Write(mask)
	masked := make([]byte, len(payload))
	for i := range payload {
		masked[i] = payload[i] ^ mask[i%4]
	}
	if _, err := w.conn.Write(header.Bytes()); err != nil {
		return err
	}
	_, err := w.conn.Write(masked)
	return err
}

func (w *websocketConn) readFrame() (byte, []byte, error) {
	var hdr [2]byte
	if _, err := io.ReadFull(w.conn, hdr[:]); err != nil {
		return 0, nil, err
	}
	opcode := hdr[0] & 0x0f
	masked := hdr[1]&0x80 != 0
	length := int64(hdr[1] & 0x7f)
	switch length {
	case 126:
		var ext [2]byte
		if _, err := io.ReadFull(w.conn, ext[:]); err != nil {
			return 0, nil, err
		}
		length = int64(binary.BigEndian.Uint16(ext[:]))
	case 127:
		var ext [8]byte
		if _, err := io.ReadFull(w.conn, ext[:]); err != nil {
			return 0, nil, err
		}
		length = int64(binary.BigEndian.Uint64(ext[:]))
	}
	var mask []byte
	if masked {
		mask = make([]byte, 4)
		if _, err := io.ReadFull(w.conn, mask); err != nil {
			return 0, nil, err
		}
	}
	payload := make([]byte, length)
	if _, err := io.ReadFull(w.conn, payload); err != nil {
		return 0, nil, err
	}
	if masked {
		for i := range payload {
			payload[i] ^= mask[i%4]
		}
	}
	return opcode, payload, nil
}

func containsAll(text string, want ...string) bool {
	for _, item := range want {
		if !strings.Contains(text, item) {
			return false
		}
	}
	return true
}
