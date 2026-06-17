package demo

import (
	"os"
	"strings"
	"testing"

	gx "github.com/busdk/bus-gx/pkg/gx"
)

var dataEvidenceDemoPages = []docsDemoPage{
	{
		id:    "dense-table",
		label: "Loading DenseTable demo...",
		path:  "../../../docs/gx-ui/bus-ui/data/dense-table/index.html",
	},
	{
		id:    "projection-detail",
		label: "Loading ProjectionDetail demo...",
		path:  "../../../docs/gx-ui/bus-ui/data/projection-detail/index.html",
	},
	{
		id:    "provider-error",
		label: "Loading ProviderError demo...",
		path:  "../../../docs/gx-ui/bus-ui/data/provider-error/index.html",
	},
	{
		id:    "record-list",
		label: "Loading RecordList demo...",
		path:  "../../../docs/gx-ui/bus-ui/data/record-list/index.html",
	},
	{
		id:    "summary-item",
		label: "Loading SummaryItem demo...",
		path:  "../../../docs/gx-ui/bus-ui/data/summary-item/index.html",
	},
	{
		id:    "text-table",
		label: "Loading TextTable demo...",
		path:  "../../../docs/gx-ui/bus-ui/data/text-table/index.html",
	},
	{
		id:    "timeline",
		label: "Loading Timeline demo...",
		path:  "../../../docs/gx-ui/bus-ui/data/timeline/index.html",
	},
	{
		id:    "evidence-link",
		label: "Loading EvidenceLink demo...",
		path:  "../../../docs/gx-ui/bus-ui/evidence/evidence-link/index.html",
	},
	{
		id:    "evidence-preview",
		label: "Loading EvidencePreview demo...",
		path:  "../../../docs/gx-ui/bus-ui/evidence/evidence-preview/index.html",
	},
	{
		id:    "image-gallery",
		label: "Loading ImageGallery demo...",
		path:  "../../../docs/gx-ui/bus-ui/evidence/image-gallery/index.html",
	},
}

func TestDataEvidenceDemoIDsAreRegistered(t *testing.T) {
	t.Parallel()

	for _, page := range dataEvidenceDemoPages {
		page := page
		t.Run(page.id, func(t *testing.T) {
			t.Parallel()

			if _, ok := Lookup(page.id); !ok {
				t.Fatalf("%s demo is not registered", page.id)
			}
		})
	}
}

func TestDataEvidenceFamilyDemosRenderRealBusUIComponents(t *testing.T) {
	t.Parallel()

	tests := []struct {
		id   string
		want []string
	}{
		{
			id: "dense-table",
			want: []string{
				`data-bus-ui-demo-widget="dense-table"`,
				`data-ui-component="DenseTable"`,
				"Document queue",
				"Vendor invoices",
			},
		},
		{
			id: "projection-detail",
			want: []string{
				`data-bus-ui-demo-widget="projection-detail"`,
				`data-ui-component="ProjectionDetail"`,
				`data-ui-component="EvidencePreview"`,
				"Invoice projection",
				"Download invoice",
			},
		},
		{
			id: "provider-error",
			want: []string{
				`data-bus-ui-demo-widget="provider-error"`,
				`data-ui-component="ProviderError"`,
				`data-provider-request-id="req_12345"`,
				`data-ui-action="provider.retry"`,
				"Provider request failed",
			},
		},
		{
			id: "record-list",
			want: []string{
				`data-bus-ui-demo-widget="record-list"`,
				`data-ui-component="RecordList"`,
				"June receipts",
				"Vendor bills",
				"bus-ui-summary-item",
			},
		},
		{
			id: "summary-item",
			want: []string{
				`data-bus-ui-demo-widget="summary-item"`,
				`data-ui-component="SummaryItem"`,
				"June receipts",
				"Prepared for accounting review.",
			},
		},
		{
			id: "text-table",
			want: []string{
				`data-bus-ui-demo-widget="text-table"`,
				`data-ui-component="DenseTable"`,
				"Title",
				"Vendor invoices",
			},
		},
		{
			id: "timeline",
			want: []string{
				`data-bus-ui-demo-widget="timeline"`,
				`data-ui-component="Timeline"`,
				`data-ui-status="warning"`,
				"Invoice workflow",
				"Provider error reviewed",
			},
		},
		{
			id: "evidence-link",
			want: []string{
				`data-bus-ui-demo-widget="evidence-link"`,
				`data-ui-component="EvidenceLink"`,
				`data-evidence-operation="download"`,
				"Download receipt PDF",
			},
		},
		{
			id: "evidence-preview",
			want: []string{
				`data-bus-ui-demo-widget="evidence-preview"`,
				`data-ui-component="EvidencePreview"`,
				`data-evidence-preview-state="inline"`,
				"Bank statement",
			},
		},
		{
			id: "image-gallery",
			want: []string{
				`data-bus-ui-demo-widget="image-gallery"`,
				`data-ui-component="ImageGallery"`,
				`class="bus-ui-gallery"`,
				"Receipt page 1",
				"Page 2",
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.id, func(t *testing.T) {
			t.Parallel()

			root, ok := Lookup(tc.id)
			if !ok {
				t.Fatalf("%s demo is not registered", tc.id)
			}
			html, err := gx.RenderHTML(root())
			if err != nil {
				t.Fatalf("RenderHTML(%s demo) failed: %v", tc.id, err)
			}
			for _, want := range tc.want {
				if !strings.Contains(html, want) {
					t.Fatalf("%s demo HTML %q does not contain %q", tc.id, html, want)
				}
			}
		})
	}
}

func TestDataEvidenceDocsPagesUseGeneratedPlaceholderAndSharedScripts(t *testing.T) {
	t.Parallel()

	for _, page := range dataEvidenceDemoPages {
		page := page
		t.Run(page.id, func(t *testing.T) {
			t.Parallel()

			placeholder, err := PlaceholderHTML(page.id, page.label)
			if err != nil {
				t.Fatalf("PlaceholderHTML(%q) error = %v", page.id, err)
			}
			body, err := os.ReadFile(page.path)
			if err != nil {
				t.Fatalf("ReadFile(%s) error = %v", page.path, err)
			}
			text := string(body)
			if !strings.Contains(text, placeholder) {
				t.Fatalf("%s does not contain generated placeholder %q", page.path, placeholder)
			}
			if got := strings.Count(text, `data-bus-ui-demo="`+page.id+`"`); got != 1 {
				t.Fatalf("%s has %d placeholders for %q, want 1", page.path, got, page.id)
			}
			if got := strings.Count(text, "assets/bus-ui-demo/wasm_exec.js"); got != 1 {
				t.Fatalf("%s has %d wasm_exec.js includes, want 1", page.path, got)
			}
			if got := strings.Count(text, "assets/bus-ui-demo/loader.js"); got != 1 {
				t.Fatalf("%s has %d loader.js includes, want 1", page.path, got)
			}
			if !strings.Contains(text, "data-bus-ui-demo-loader") {
				t.Fatalf("%s missing data-bus-ui-demo-loader", page.path)
			}
		})
	}
}
