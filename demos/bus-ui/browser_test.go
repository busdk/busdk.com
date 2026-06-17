package main

import (
	"context"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
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
		output := runHeadlessBrowser(t, browser, fileURL(buttonPath))
		for _, want := range []string{
			`data-bus-ui-demo-state="failed"`,
			`data-bus-ui-demo-asset="css"`,
			"Bus UI demos need a local HTTP server to load WebAssembly.",
		} {
			if !strings.Contains(output, want) {
				t.Fatalf("file:// output missing %q in:\n%s", want, output)
			}
		}
		if strings.Contains(output, "Save draft") {
			t.Fatalf("file:// preview unexpectedly mounted the Button demo:\n%s", output)
		}
	})

	t.Run("http-preview", func(t *testing.T) {
		baseURL, cleanup := serveRepoOnLoopback(t)
		defer cleanup()

		output := runHeadlessBrowser(t, browser, baseURL+buttonDocsPreviewRelativePath)
		for _, want := range []string{
			`data-bus-ui-demo-state="mounted"`,
			`data-bus-ui-demo-asset="css"`,
			"Save draft",
			`data-bus-ui-demo-widget="button"`,
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

func runHeadlessBrowser(t *testing.T, browser string, pageURL string) string {
	t.Helper()

	ctx, cancel := context.WithTimeout(context.Background(), 45*time.Second)
	defer cancel()

	userDataDir := t.TempDir()
	cmd := exec.CommandContext(ctx, browser,
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
		"--virtual-time-budget=12000",
		"--user-data-dir="+userDataDir,
		"--dump-dom",
		pageURL,
	)

	output, err := cmd.CombinedOutput()
	if ctx.Err() == context.DeadlineExceeded {
		t.Fatalf("headless browser timed out for %s", pageURL)
	}
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			if status, ok := exitErr.Sys().(syscall.WaitStatus); ok && status.Signaled() && status.Signal() == syscall.SIGABRT {
				t.Skipf("skip button browser regression: headless browser aborts in this environment: %v", err)
			}
		}
		t.Fatalf("headless browser failed for %s: %v\n%s", pageURL, err, output)
	}
	return string(output)
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
