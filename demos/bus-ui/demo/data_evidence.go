package demo

import (
	gx "github.com/busdk/bus-gx/pkg/gx"
	ui "github.com/busdk/bus-ui/pkg/ui"
)

// DenseTable renders one validated comparable-row table demo.
func DenseTable() gx.Node {
	node := ui.DenseTable(ui.DenseTableProps{
		Caption: "Document queue",
		Columns: []ui.DenseTableColumn{
			ui.DenseTableColumnText("title", "Title"),
			ui.DenseTableColumnText("files", "Files"),
			ui.DenseTableColumnText("state", "State"),
		},
		Rows: []ui.DenseTableRow{
			ui.DenseTableRowText("June receipts", "3 files", "Ready"),
			ui.DenseTableRowText("Vendor invoices", "8 files", "Review"),
		},
	})
	return demoWidget("dense-table", node)
}

// ProjectionDetail renders one public-safe projected detail with evidence actions.
func ProjectionDetail() gx.Node {
	node, err := ui.ProjectionDetailNodeChecked(ui.ProjectionDetailProps{
		Title:   "Invoice projection",
		Summary: "Validated before the provider response is exposed.",
		Fields: []ui.ProjectionField{
			{Label: "Provider", Value: "Stripe"},
			{Label: "State", Value: "Ready"},
			{Label: "Total", Value: "128.40 EUR"},
		},
		Evidence: []ui.ProjectionEvidenceAction{
			{ID: "invoice-preview", Label: "Preview invoice", Operation: "preview", URL: "/busdk-logo.png", MediaType: "image/png"},
			{ID: "invoice-download", Label: "Download invoice", Operation: "download", URL: "/busdk-logo.png", Filename: "invoice-2026-06.png"},
		},
	})
	if err != nil {
		return demoError(err)
	}
	return demoWidget("projection-detail", node)
}

// ProviderError renders one public-safe provider failure demo.
func ProviderError() gx.Node {
	node := ui.ProviderError(ui.ProviderErrorProps{
		Title:     "Provider request failed",
		Summary:   "The backend returned a validation error.",
		Code:      "validation.failed",
		Status:    422,
		RequestID: "req_12345",
		Fields: map[string]string{
			"invoice_id": "Invoice ID is required.",
		},
		RetryLabel:   "Retry",
		Retry:        ui.ControlProps{Action: "provider.retry", SourceID: "provider-error-demo"},
		DismissLabel: "Dismiss",
		Dismiss:      ui.ControlProps{Action: "provider.dismiss", SourceID: "provider-error-demo"},
	})
	if node == nil {
		return demoError(ui.ErrProviderErrorTitleRequired)
	}
	return demoWidget("provider-error", node)
}

// RecordList renders one ordered record-summary list demo.
func RecordList() gx.Node {
	node := ui.RecordList(ui.RecordListProps{
		Items: []ui.RecordListItem{
			{Title: "June receipts", Meta: "3 files", Detail: "Prepared for accounting review.", Badge: "ready"},
			{Title: "Vendor bills", Meta: "8 files", Detail: "Two rows still need category review.", Badge: "review"},
		},
		EmptyTitle:  "No receipts yet",
		EmptyDetail: "Upload files to build the list.",
	})
	return demoWidget("record-list", node)
}

// SummaryItem renders one compact record summary demo.
func SummaryItem() gx.Node {
	node := ui.SummaryItem(ui.SummaryItemProps{
		Title:  "June receipts",
		Meta:   "3 files",
		Detail: "Prepared for accounting review.",
		Badge:  "ready",
	})
	if node == nil {
		return demoError(ui.ErrSurfaceTitleRequired)
	}
	return demoWidget("summary-item", node)
}

// TextTable renders one static-header trusted-row table demo.
func TextTable() gx.Node {
	rows := []gx.Node{
		gx.Element("tr", nil,
			gx.Element("td", nil, gx.Text("June receipts")),
			gx.Element("td", nil, gx.Text("3 files")),
			gx.Element("td", nil, gx.Text("Ready")),
		),
		gx.Element("tr", nil,
			gx.Element("td", nil, gx.Text("Vendor invoices")),
			gx.Element("td", nil, gx.Text("8 files")),
			gx.Element("td", nil, gx.Text("Review")),
		),
	}
	node := ui.TextTable([]string{"Title", "Files", "State"}, rows)
	return demoWidget("text-table", node)
}

// Timeline renders one ordered evidence-workflow history demo.
func Timeline() gx.Node {
	node := ui.Timeline(ui.TimelineProps{
		Label: "Invoice workflow",
		Items: []ui.TimelineItem{
			{
				Time:     "2026-06-14T09:00:00Z",
				Body:     "Invoice imported",
				Sequence: "1",
				Status:   ui.TimelineStatusSuccess,
				Meta: []ui.TimelineMeta{
					{Label: "Actor", Value: "worker"},
				},
			},
			{
				Time:     "2026-06-14T09:10:00Z",
				Body:     "Provider error reviewed",
				Sequence: "2",
				Status:   ui.TimelineStatusWarning,
				Meta: []ui.TimelineMeta{
					{Label: "Queue", Value: "manual"},
				},
			},
		},
	})
	return demoWidget("timeline", node)
}

// EvidenceLink renders one checked evidence download action demo.
func EvidenceLink() gx.Node {
	node := ui.EvidenceLink(ui.EvidenceLinkProps{
		Label:     "Download receipt PDF",
		Href:      "/busdk-logo.png",
		Operation: ui.EvidenceOperationDownload,
		Download:  true,
	})
	if node == nil {
		return demoError(ui.ErrEvidenceLabelRequired)
	}
	return demoWidget("evidence-link", node)
}

// EvidencePreview renders one inline evidence preview demo.
func EvidencePreview() gx.Node {
	node := ui.EvidencePreview(ui.EvidencePreviewProps{
		Title:              "Bank statement",
		PreviewURL:         "/busdk-logo.png",
		OpenURL:            "/busdk-logo.png",
		DownloadURL:        "/busdk-logo.png",
		ContentType:        "image/png",
		ContentDisposition: "inline",
		Fallback:           "Preview is unavailable.",
	})
	if node == nil {
		return demoError(ui.ErrEvidencePreviewTitleRequired)
	}
	return demoWidget("evidence-preview", node)
}

// ImageGallery renders one allowlisted image-gallery demo.
func ImageGallery() gx.Node {
	node, err := ui.ImageGalleryNodeChecked(ui.ImageGalleryProps{
		Items: []ui.ImageGalleryItem{
			{
				Src:     "/busdk-logo.png",
				Alt:     "Receipt page 1",
				Href:    "/busdk-logo.png",
				Caption: "Page 1",
			},
			{
				Src:     "/busdk-logo.png",
				Alt:     "Receipt page 2",
				Href:    "/busdk-logo.png",
				Caption: "Page 2",
			},
		},
	})
	if err != nil {
		return demoError(err)
	}
	return demoWidget("image-gallery", node)
}
