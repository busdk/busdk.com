package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"testing"
)

var localDocsAttrPattern = regexp.MustCompile(`(?i)\b(?:href|src)\s*=\s*(?:"([^"]*)"|'([^']*)'|([^\s>]+))`)

type docsDirectoryLinkHit struct {
	File   string
	Attr   string
	Target string
	Raw    string
}

func TestDocsDirectoryLinksUseIndexHTMLForFilePreview(t *testing.T) {
	repoRoot, err := repoRootDir()
	if err != nil {
		t.Fatal(err)
	}

	hits, filesScanned, localAttrsScanned, err := scanDocsDirectoryLinks(repoRoot)
	if err != nil {
		t.Fatal(err)
	}
	if len(hits) > 0 {
		lines := make([]string, 0, len(hits))
		for _, hit := range hits {
			lines = append(lines, fmt.Sprintf("%s\t%s\t%s\t->\t%s", hit.File, hit.Attr, hit.Raw, hit.Target))
		}
		t.Fatalf("LOCAL_DIRECTORY_LINKS_FOUND\n%s", strings.Join(lines, "\n"))
	}

	t.Logf("NO_LOCAL_DIRECTORY_LINKS_FOUND files=%d local_attrs=%d", filesScanned, localAttrsScanned)
}

func TestBusEngineGeneratedNavBasesUsePageRelativePaths(t *testing.T) {
	repoRoot, err := repoRootDir()
	if err != nil {
		t.Fatal(err)
	}

	for _, tt := range []struct {
		rel  string
		base string
	}{
		{rel: "docs/engine/index.html", base: "./"},
		{rel: "docs/engine/modules/index.html", base: "../"},
		{rel: "docs/engine/runtime-images/index.html", base: "../"},
		{rel: "docs/engine/pricing/index.html", base: "../"},
		{rel: "docs/engine/contact/index.html", base: "../"},
	} {
		body, err := os.ReadFile(filepath.Join(repoRoot, tt.rel))
		if err != nil {
			t.Fatal(err)
		}
		text := string(body)
		for _, attr := range []string{
			`data-busdk-top-nav-base="` + tt.base + `"`,
			`data-busdk-side-nav-base="` + tt.base + `"`,
		} {
			if !strings.Contains(text, attr) {
				t.Fatalf("%s missing %s", tt.rel, attr)
			}
		}
	}
}

func scanDocsDirectoryLinks(repoRoot string) ([]docsDirectoryLinkHit, int, int, error) {
	docsRoot := filepath.Join(repoRoot, "docs")
	var hits []docsDirectoryLinkHit
	filesScanned := 0
	localAttrsScanned := 0

	err := filepath.WalkDir(docsRoot, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || !strings.HasSuffix(strings.ToLower(d.Name()), ".html") {
			return nil
		}
		filesScanned++

		body, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		for _, match := range localDocsAttrPattern.FindAllStringSubmatch(string(body), -1) {
			raw := firstNonEmpty(match[1], match[2], match[3])
			if !shouldScanLocalDocsLink(raw) {
				continue
			}
			localAttrsScanned++

			targetPath := filepath.Clean(filepath.Join(filepath.Dir(path), stripDocsLinkSuffix(raw)))
			info, err := os.Stat(targetPath)
			if err != nil || !info.IsDir() {
				continue
			}

			hits = append(hits, docsDirectoryLinkHit{
				File:   relRepoPath(repoRoot, path),
				Attr:   attrNameForMatch(match[0]),
				Raw:    raw,
				Target: relRepoPath(repoRoot, targetPath),
			})
		}
		return nil
	})
	if err != nil {
		return nil, 0, 0, err
	}

	sort.Slice(hits, func(i, j int) bool {
		if hits[i].File != hits[j].File {
			return hits[i].File < hits[j].File
		}
		if hits[i].Attr != hits[j].Attr {
			return hits[i].Attr < hits[j].Attr
		}
		if hits[i].Raw != hits[j].Raw {
			return hits[i].Raw < hits[j].Raw
		}
		return hits[i].Target < hits[j].Target
	})

	return hits, filesScanned, localAttrsScanned, nil
}

func shouldScanLocalDocsLink(raw string) bool {
	value := strings.TrimSpace(raw)
	if value == "" {
		return false
	}
	if strings.HasPrefix(value, "#") || strings.HasPrefix(value, "//") || strings.HasPrefix(value, "/") {
		return false
	}
	if schemeLikeLink.MatchString(value) {
		return false
	}

	clean := stripDocsLinkSuffix(value)
	if clean == "." || clean == ".." {
		return false
	}
	if isOrdinaryAssetFile(clean) {
		return false
	}
	return true
}

func stripDocsLinkSuffix(raw string) string {
	clean := strings.TrimSpace(raw)
	if idx := strings.IndexByte(clean, '#'); idx >= 0 {
		clean = clean[:idx]
	}
	if idx := strings.IndexByte(clean, '?'); idx >= 0 {
		clean = clean[:idx]
	}
	return clean
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if value != "" {
			return value
		}
	}
	return ""
}

func attrNameForMatch(match string) string {
	lower := strings.ToLower(match)
	if strings.Contains(lower, "href") {
		return "href"
	}
	return "src"
}

var schemeLikeLink = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9+.-]*:`)

func isOrdinaryAssetFile(target string) bool {
	switch strings.ToLower(filepath.Ext(target)) {
	case ".css", ".js", ".mjs", ".map", ".png", ".jpg", ".jpeg", ".gif", ".webp", ".svg", ".ico", ".cur", ".avif", ".woff", ".woff2", ".ttf", ".otf", ".eot", ".mp4", ".webm", ".mp3", ".wav", ".pdf", ".json", ".xml", ".txt", ".csv", ".webmanifest":
		return true
	default:
		return false
	}
}

func relRepoPath(repoRoot, path string) string {
	rel, err := filepath.Rel(repoRoot, path)
	if err != nil {
		return filepath.ToSlash(path)
	}
	return filepath.ToSlash(rel)
}
