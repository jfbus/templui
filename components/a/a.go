package a

import (
	"github.com/a-h/templ"
	"github.com/jfbus/templui/components/style"
)

func init() {
	style.SetDefaults(style.Defaults{
		"a": {
			style.Default: {
				style.Set("hover:underline"),
			},
		},
	})
}

type D struct {
	// Href is the target URL.
	Href string
	// Text is the link text.
	Text string
	// CustomStyle defines a custom style.
	// 	style.Custom{
	// 		"a": style.D{style.Add("text-sm")},
	//	}
	CustomStyle style.Custom
	// Attributes defines additional HTML attributes.
	Attributes templ.Attributes
}

func (def D) class() string {
	return style.CSSClass(style.Default, "a", def.CustomStyle)
}
