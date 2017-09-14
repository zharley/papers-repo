package content

import (
	"fmt"

	"github.com/bosssauce/reference"

	"github.com/ponzu-cms/ponzu/management/editor"
	"github.com/ponzu-cms/ponzu/system/item"
)

type Author struct {
	item.Item

	Name        string `json:"name"`
	Email       string `json:"email"`
	Affiliation string `json:"affiliation"`
}

// MarshalEditor writes a buffer of html to edit a Author within the CMS
// and implements editor.Editable
func (a *Author) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(a,
		// Take note that the first argument to these Input-like functions
		// is the string version of each Author field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:
		editor.Field{
			View: editor.Input("Name", a, map[string]string{
				"label":       "Name",
				"type":        "text",
				"placeholder": "Enter the Name here",
			}),
		},
		editor.Field{
			View: editor.Input("Email", a, map[string]string{
				"label":       "Email",
				"type":        "text",
				"placeholder": "Enter the Email here",
			}),
		},
		editor.Field{
			View: reference.Select("Affiliation", a, map[string]string{
				"label": "Affiliation",
			},
				"Organization",
				`Organization: {{ .id }}`,
			),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render Author editor view: %s", err.Error())
	}

	return view, nil
}

func init() {
	item.Types["Author"] = func() interface{} { return new(Author) }
}

// String defines how a Author is printed. Update it using more descriptive
// fields from the Author struct type
func (a *Author) String() string {
	return fmt.Sprintf("Author: %s", a.UUID)
}
