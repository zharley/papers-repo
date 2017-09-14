package content

import (
	"fmt"

	"github.com/ponzu-cms/ponzu/management/editor"
	"github.com/ponzu-cms/ponzu/system/item"
)

type Paper struct {
	item.Item

	Title     string   `json:"title"`
	Published string   `json:"published"`
	Summary   string   `json:"summary"`
	Input     string   `json:"input"`
	Output    string   `json:"output"`
	Method    string   `json:"method"`
	Results   string   `json:"results"`
	Extra     string   `json:"extra"`
	Images    []string `json:"images"`
}

// MarshalEditor writes a buffer of html to edit a Paper within the CMS
// and implements editor.Editable
func (p *Paper) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(p,
		// Take note that the first argument to these Input-like functions
		// is the string version of each Paper field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:
		editor.Field{
			View: editor.Input("Title", p, map[string]string{
				"label":       "Title",
				"type":        "text",
				"placeholder": "Enter the Title here",
			}),
		},
		editor.Field{
			View: editor.Input("Published", p, map[string]string{
				"label":       "Published",
				"type":        "text",
				"placeholder": "Enter the Published here",
			}),
		},
		editor.Field{
			View: editor.Richtext("Summary", p, map[string]string{
				"label":       "Summary",
				"placeholder": "Enter the Summary here",
			}),
		},
		editor.Field{
			View: editor.Richtext("Input", p, map[string]string{
				"label":       "Input",
				"placeholder": "Enter the Input here",
			}),
		},
		editor.Field{
			View: editor.Richtext("Output", p, map[string]string{
				"label":       "Output",
				"placeholder": "Enter the Output here",
			}),
		},
		editor.Field{
			View: editor.Richtext("Method", p, map[string]string{
				"label":       "Method",
				"placeholder": "Enter the Method here",
			}),
		},
		editor.Field{
			View: editor.Richtext("Results", p, map[string]string{
				"label":       "Results",
				"placeholder": "Enter the Results here",
			}),
		},
		editor.Field{
			View: editor.Richtext("Extra", p, map[string]string{
				"label":       "Extra",
				"placeholder": "Enter the Extra here",
			}),
		},
		editor.Field{
			View: editor.FileRepeater("Images", p, map[string]string{
				"label":       "Images",
				"placeholder": "Upload the Images here",
			}),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render Paper editor view: %s", err.Error())
	}

	return view, nil
}

func init() {
	item.Types["Paper"] = func() interface{} { return new(Paper) }
}

// String defines how a Paper is printed. Update it using more descriptive
// fields from the Paper struct type
func (p *Paper) String() string {
	return fmt.Sprintf("Paper: %s", p.UUID)
}
