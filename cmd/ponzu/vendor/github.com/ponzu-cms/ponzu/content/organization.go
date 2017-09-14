package content

import (
	"fmt"

	"github.com/ponzu-cms/ponzu/management/editor"
	"github.com/ponzu-cms/ponzu/system/item"
)

type Organization struct {
	item.Item

	Name string `json:"name"`
}

// MarshalEditor writes a buffer of html to edit a Organization within the CMS
// and implements editor.Editable
func (o *Organization) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(o,
		// Take note that the first argument to these Input-like functions
		// is the string version of each Organization field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:
		editor.Field{
			View: editor.Input("Name", o, map[string]string{
				"label":       "Name",
				"type":        "text",
				"placeholder": "Enter the Name here",
			}),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render Organization editor view: %s", err.Error())
	}

	return view, nil
}

func init() {
	item.Types["Organization"] = func() interface{} { return new(Organization) }
}

// String defines how a Organization is printed. Update it using more descriptive
// fields from the Organization struct type
func (o *Organization) String() string {
	return fmt.Sprintf("Organization: %s", o.UUID)
}
