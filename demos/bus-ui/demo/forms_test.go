package demo

import (
	"os"
	"strings"
	"testing"

	gx "github.com/busdk/bus-gx/pkg/gx"
)

type formsDocsDemoPage struct {
	id    string
	label string
	path  string
}

var formsDocsDemoPages = []formsDocsDemoPage{
	{
		id:    "credential-login-card",
		label: "Loading CredentialLoginCard demo...",
		path:  "../../../docs/gx-ui/bus-ui/forms/credential-login-card/index.html",
	},
	{
		id:    "date-input",
		label: "Loading DateInput demo...",
		path:  "../../../docs/gx-ui/bus-ui/forms/date-input/index.html",
	},
	{
		id:    "drop-zone",
		label: "Loading DropZone demo...",
		path:  "../../../docs/gx-ui/bus-ui/forms/drop-zone/index.html",
	},
	{
		id:    "field",
		label: "Loading Field demo...",
		path:  "../../../docs/gx-ui/bus-ui/forms/field/index.html",
	},
	{
		id:    "file-input",
		label: "Loading FileInput demo...",
		path:  "../../../docs/gx-ui/bus-ui/forms/file-input/index.html",
	},
	{
		id:    "filter-toolbar",
		label: "Loading FilterToolbar demo...",
		path:  "../../../docs/gx-ui/bus-ui/forms/filter-toolbar/index.html",
	},
	{
		id:    "form",
		label: "Loading Form demo...",
		path:  "../../../docs/gx-ui/bus-ui/forms/form/index.html",
	},
	{
		id:    "input",
		label: "Loading Input demo...",
		path:  "../../../docs/gx-ui/bus-ui/forms/input/index.html",
	},
	{
		id:    "password-input",
		label: "Loading PasswordInput demo...",
		path:  "../../../docs/gx-ui/bus-ui/forms/password-input/index.html",
	},
	{
		id:    "select",
		label: "Loading Select demo...",
		path:  "../../../docs/gx-ui/bus-ui/forms/select/index.html",
	},
	{
		id:    "submit",
		label: "Loading SubmitControl demo...",
		path:  "../../../docs/gx-ui/bus-ui/forms/submit/index.html",
	},
	{
		id:    "text-input",
		label: "Loading TextInput demo...",
		path:  "../../../docs/gx-ui/bus-ui/forms/text-input/index.html",
	},
	{
		id:    "textarea",
		label: "Loading TextArea demo...",
		path:  "../../../docs/gx-ui/bus-ui/forms/textarea/index.html",
	},
}

func TestFormsFamilyDemosRenderRealBusUIComponents(t *testing.T) {
	t.Parallel()

	tests := []struct {
		id   string
		want []string
	}{
		{
			id: "credential-login-card",
			want: []string{
				`data-bus-ui-demo-widget="credential-login-card"`,
				"bus-ui-auth-card",
				`action="/auth/login"`,
				"Send one-time code",
				"Continue",
			},
		},
		{
			id: "date-input",
			want: []string{
				`data-bus-ui-demo-widget="date-input"`,
				"bus-ui-input",
				`type="date"`,
				`value="2026-06-30"`,
			},
		},
		{
			id: "drop-zone",
			want: []string{
				`data-bus-ui-demo-widget="drop-zone"`,
				"bus-ui-dropzone",
				"Drop evidence files",
				"Choose files",
				"Maximum 25 MB per file.",
			},
		},
		{
			id: "field",
			want: []string{
				`data-bus-ui-demo-widget="field"`,
				"bus-ui-field",
				"Description",
				"Shown on the evidence card.",
				`type="text"`,
			},
		},
		{
			id: "file-input",
			want: []string{
				`data-bus-ui-demo-widget="file-input"`,
				`type="file"`,
				`accept="application/pdf,.png"`,
				"Attach evidence",
			},
		},
		{
			id: "filter-toolbar",
			want: []string{
				`data-bus-ui-demo-widget="filter-toolbar"`,
				"bus-ui-filter-toolbar-fields",
				"Filters",
				"Apply",
				"Clear",
			},
		},
		{
			id: "form",
			want: []string{
				`data-bus-ui-demo-widget="form"`,
				`action="/customers"`,
				"Customer form",
				"Save customer",
			},
		},
		{
			id: "input",
			want: []string{
				`data-bus-ui-demo-widget="input"`,
				"bus-ui-input",
				`type="search"`,
				`placeholder="Search evidence"`,
			},
		},
		{
			id: "password-input",
			want: []string{
				`data-bus-ui-demo-widget="password-input"`,
				"bus-ui-input",
				`type="password"`,
				`autocomplete="current-password"`,
			},
		},
		{
			id: "select",
			want: []string{
				`data-bus-ui-demo-widget="select"`,
				"bus-ui-select",
				"Open",
				"Closed",
				"Archived",
			},
		},
		{
			id: "submit",
			want: []string{
				`data-bus-ui-demo-widget="submit"`,
				`data-ui-submit-state="working"`,
				`aria-busy="true"`,
				"Saving",
			},
		},
		{
			id: "text-input",
			want: []string{
				`data-bus-ui-demo-widget="text-input"`,
				"bus-ui-input",
				`value="June VAT report"`,
				`placeholder="Report title"`,
			},
		},
		{
			id: "textarea",
			want: []string{
				`data-bus-ui-demo-widget="textarea"`,
				"bus-ui-textarea",
				`rows="4"`,
				"Reviewed by accounting.",
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

func TestFormsDocsPagesUseGeneratedPlaceholderAndSharedScripts(t *testing.T) {
	t.Parallel()

	for _, page := range formsDocsDemoPages {
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
				t.Fatalf("%s data-bus-ui-demo count = %d, want 1", page.path, got)
			}
			if got := strings.Count(text, "assets/bus-ui-demo/wasm_exec.js"); got != 1 {
				t.Fatalf("%s wasm_exec.js count = %d, want 1", page.path, got)
			}
			if got := strings.Count(text, "assets/bus-ui-demo/loader.js"); got != 1 {
				t.Fatalf("%s loader.js count = %d, want 1", page.path, got)
			}
			if got := strings.Count(text, "data-bus-ui-demo-loader"); got != 1 {
				t.Fatalf("%s data-bus-ui-demo-loader count = %d, want 1", page.path, got)
			}
		})
	}
}
