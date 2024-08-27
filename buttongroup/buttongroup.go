package buttongroup

import (
	"github.com/a-h/templ"
	"github.com/jfbus/templ-components/button"
	"github.com/jfbus/templ-components/size"
)

const DefaultColor = "text-gray-900 bg-white border-gray-300 hover:bg-gray-100 focus:ring-gray-100 dark:bg-gray-800 dark:text-white dark:border-gray-600 dark:hover:bg-gray-700 dark:hover:border-gray-600 dark:focus:ring-gray-700"

type D struct {
	// Color overrides the default color CSS classes.
	Color string
	// Buttons is the list of buttons to display.
	Buttons []button.D
	// Size defines the buttons size (shortcut for Buttons.Size).
	Size size.Size
	// Class defines additional CSS class for the button.
	Class string
	// Attributes stores additional attributes (e.g. HTMX attributes).
	Attributes templ.Attributes
}

func (def D) buttons() []button.D {
	bs := make([]button.D, len(def.Buttons))
	for i := range def.Buttons {
		bs[i] = def.Buttons[i]
		bs[i].Style = button.NoBorder
		if def.Size != size.Inherit {
			bs[i].Size = def.Size
		}
		switch {
		case bs[i].Color != "":
		case def.Color != "":
			bs[i].Color = def.Color
		default:
			bs[i].Color = DefaultColor
		}
		switch i {
		case 0:
			bs[i].Class += " border rounded-s-lg"
		case len(bs) - 1:
			bs[i].Class += " border rounded-e-lg"
		default:
			bs[i].Class += " border-t border-b"
		}
	}
	return bs
}

func (def D) groupClass() string {
	class := "inline-flex rounded-md shadow-sm "
	if def.Color != "" {
		class += def.Color
	} else {
		class += DefaultColor
	}
	if def.Class != "" {
		class += " " + def.Class
	}
	return class
}
