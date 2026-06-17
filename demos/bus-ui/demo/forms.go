package demo

import (
	gx "github.com/busdk/bus-gx/pkg/gx"
	ui "github.com/busdk/bus-ui/pkg/ui"
)

// CredentialLoginCard renders one sign-in card demo with request and submit actions.
func CredentialLoginCard() gx.Node {
	node, err := ui.CredentialLoginCard(ui.CredentialLoginCardProps{
		ID:                   "credential-login-card-demo",
		Title:                "Sign in",
		Copy:                 "Use your work email and one-time code.",
		FormAction:           "/auth/login",
		OnSubmit:             func(ui.CredentialSubmitEvent) {},
		OnRequest:            func(ui.CredentialRequestEvent) {},
		RequestLabel:         "Send one-time code",
		UsernameLabel:        "Email",
		UsernameName:         "email",
		UsernameType:         string(ui.InputTypeEmail),
		UsernamePlaceholder:  "name@example.com",
		UsernameAutocomplete: "email",
		PasswordLabel:        "One-time code",
		PasswordName:         "code",
		PasswordPlaceholder:  "123456",
		PasswordAutocomplete: "one-time-code",
		Submit: ui.ButtonProps{
			Label: "Continue",
		},
	})
	if err != nil {
		return demoError(err)
	}
	return demoWidget("credential-login-card", node)
}

// DateInput renders one shared date input demo.
func DateInput() gx.Node {
	node, err := ui.DateInputNodeChecked("period-end", "2026-06-30", map[string]string{
		"id": "period-end",
	}, nil)
	if err != nil {
		return demoError(err)
	}
	return demoWidget("date-input", node)
}

// DropZone renders one upload drop surface demo composed with public file-input and button APIs.
func DropZone() gx.Node {
	input, err := ui.FileInput(ui.FileInputProps{
		Name:          "evidence",
		ID:            "dropzone-upload",
		AcceptedTypes: []string{"application/pdf", ".csv"},
		Multiple:      true,
		AriaLabel:     "Choose evidence files",
		Attrs: map[string]string{
			"class": "bus-ui-dropzone-input",
		},
	})
	if err != nil {
		return demoError(err)
	}
	choose, err := ui.Button(ui.ButtonProps{
		Label:   "Choose files",
		Variant: ui.ButtonSecondary,
		Size:    ui.ButtonSm,
		Control: ui.ControlProps{
			Action:   "dropzone.choose",
			SourceID: "dropzone-choose",
		},
	})
	if err != nil {
		return demoError(err)
	}
	node, err := ui.DropZone(ui.DropZoneProps{
		ID:            "evidence-drop",
		Title:         "Drop evidence files",
		Copy:          "PDF receipts and CSV exports are accepted.",
		InputNodes:    []gx.Node{input},
		ActionNodes:   []gx.Node{choose},
		ErrorNodes:    []gx.Node{gx.Element("p", nil, gx.Text("Maximum 25 MB per file."))},
		AcceptedTypes: []string{"application/pdf", ".csv"},
		MaxBytes:      25 * 1024 * 1024,
	})
	if err != nil {
		return demoError(err)
	}
	return demoWidget("drop-zone", node)
}

// Field renders one labeled text field demo.
func Field() gx.Node {
	node, err := ui.Field(ui.FieldProps{
		Label:       "Description",
		ControlID:   "field-description",
		ControlName: "description",
		HintText:    "Shown on the evidence card.",
		RenderControlNode: func(attrs map[string]string) (gx.Node, error) {
			return ui.Input(ui.InputProps{
				Type:        string(ui.InputTypeText),
				Name:        "description",
				Placeholder: "Receipt note",
				Attrs:       attrs,
			})
		},
	})
	if err != nil {
		return demoError(err)
	}
	return demoWidget("field", node)
}

// FileInput renders one standalone file input demo.
func FileInput() gx.Node {
	node, err := ui.FileInput(ui.FileInputProps{
		Name:          "upload",
		ID:            "file-upload",
		AcceptedTypes: []string{"application/pdf", ".png"},
		AriaLabel:     "Attach evidence",
	})
	if err != nil {
		return demoError(err)
	}
	return demoWidget("file-input", node)
}

// FilterToolbar renders one compact search and status filter toolbar demo.
func FilterToolbar() gx.Node {
	node, err := ui.FilterToolbar(ui.FilterToolbarProps{
		Label:    "Filters",
		SourceID: "evidence-filters",
		Fields: []ui.FieldProps{
			{
				Label:       "Query",
				ControlID:   "filter-query",
				ControlName: "query",
				RenderControlNode: func(attrs map[string]string) (gx.Node, error) {
					return ui.Input(ui.InputProps{
						Type:        string(ui.InputTypeSearch),
						Name:        "query",
						Placeholder: "Search evidence",
						Attrs:       attrs,
					})
				},
			},
			{
				Label:       "Status",
				ControlID:   "filter-status",
				ControlName: "status",
				RenderControlNode: func(attrs map[string]string) (gx.Node, error) {
					return ui.Select(ui.SelectProps{
						Name:     "status",
						Selected: "open",
						Options: []ui.SelectOption{
							{ID: "open", Label: "Open"},
							{ID: "closed", Label: "Closed"},
						},
						Attrs: attrs,
					})
				},
			},
		},
		SubmitLabel: "Apply",
		ResetLabel:  "Clear",
	})
	if err != nil {
		return demoError(err)
	}
	return demoWidget("filter-toolbar", node)
}

// Form renders one semantic form demo with a field and submit control.
func Form() gx.Node {
	field, err := ui.Field(ui.FieldProps{
		Label:       "Customer",
		ControlID:   "customer-name",
		ControlName: "customer-name",
		HintText:    "Shown on the summary card.",
		RenderControlNode: func(attrs map[string]string) (gx.Node, error) {
			return ui.Input(ui.InputProps{
				Type:        string(ui.InputTypeText),
				Name:        "customer-name",
				Placeholder: "Example Oy",
				Required:    true,
				Attrs:       attrs,
			})
		},
	})
	if err != nil {
		return demoError(err)
	}
	submit, err := ui.SubmitControl(ui.SubmitControlProps{
		Label:    "Save customer",
		Variant:  ui.ButtonPrimary,
		Size:     ui.ButtonMd,
		SourceID: "customer-form",
	})
	if err != nil {
		return demoError(err)
	}
	node, err := ui.Form(ui.FormProps{
		BodyNodes: []gx.Node{
			field,
			gx.Element("div", gx.Props{"class": "bus-ui-form-actions"}, submit),
		},
		Method:    ui.FormMethodPost,
		Action:    "/customers",
		SourceID:  "customer-form",
		AriaLabel: "Customer form",
		Attrs: map[string]string{
			"id": "customer-form",
		},
	})
	if err != nil {
		return demoError(err)
	}
	return demoWidget("form", node)
}

// Input renders one typed base input demo.
func Input() gx.Node {
	node, err := ui.Input(ui.InputProps{
		Type:        string(ui.InputTypeSearch),
		Name:        "query",
		Placeholder: "Search evidence",
		Attrs: map[string]string{
			"id": "query",
		},
	})
	if err != nil {
		return demoError(err)
	}
	return demoWidget("input", node)
}

// PasswordInput renders one password helper demo.
func PasswordInput() gx.Node {
	node, err := ui.PasswordInputNodeChecked("password", "", "Password", "current-password", map[string]string{
		"id": "password",
	}, nil)
	if err != nil {
		return demoError(err)
	}
	return demoWidget("password-input", node)
}

// Select renders one select control demo.
func Select() gx.Node {
	node, err := ui.Select(ui.SelectProps{
		Name:     "status",
		Selected: "open",
		Options: []ui.SelectOption{
			{ID: "open", Label: "Open"},
			{ID: "closed", Label: "Closed"},
			{ID: "archived", Label: "Archived"},
		},
		Attrs: map[string]string{
			"id": "status",
		},
	})
	if err != nil {
		return demoError(err)
	}
	return demoWidget("select", node)
}

// Submit renders one submit-state control demo.
func Submit() gx.Node {
	node, err := ui.SubmitControl(ui.SubmitControlProps{
		State:        ui.SubmitStateWorking,
		Label:        "Save changes",
		WorkingLabel: "Saving",
		Variant:      ui.ButtonPrimary,
		Size:         ui.ButtonMd,
		SourceID:     "submit-demo",
	})
	if err != nil {
		return demoError(err)
	}
	return demoWidget("submit", node)
}

// TextInput renders one text helper demo.
func TextInput() gx.Node {
	node, err := ui.TextInputNodeChecked(ui.TextInputProps{
		Name:        "title",
		Value:       "June VAT report",
		Placeholder: "Report title",
		Attrs: map[string]string{
			"id": "title",
		},
	})
	if err != nil {
		return demoError(err)
	}
	return demoWidget("text-input", node)
}

// Textarea renders one multiline text control demo.
func Textarea() gx.Node {
	node, err := ui.TextArea(ui.TextAreaProps{
		Name:        "notes",
		Value:       "Reviewed by accounting.",
		Placeholder: "Add notes",
		Rows:        4,
		Attrs: map[string]string{
			"id": "notes",
		},
	})
	if err != nil {
		return demoError(err)
	}
	return demoWidget("textarea", node)
}
