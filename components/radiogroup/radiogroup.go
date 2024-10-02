package radiogroup

import (
	"github.com/jfbus/templ-components/components/form/validation/message"
	"github.com/jfbus/templ-components/components/radio"
	"github.com/jfbus/templ-components/components/style"
)

const (
	StyleHorizontal        style.Style = 1 << 8
	StyleBordered          style.Style = 1 << 9
	StyleGrouped           style.Style = 1 << 10
	StyleGroupedHorizontal style.Style = 1 << 11
	StyleLabelOnly         style.Style = 1 << 12
)

var Defaults = style.Defaults{
	style.StyleDefault: {
		"RadioContainerClass": {
			style.Class("flex items-center"),
		},
		"RadioLabelClass": {
			style.Add("py-3"),
		},
	},
	StyleHorizontal: {
		"ContainerClass": {
			style.Class("flex flex-col sm:flex-row sm:gap-4"),
		},
	},
	StyleBordered: {
		"RadioContainerClass": {
			style.Class("flex items-center px-4 border rounded w-full"),
			style.Color("border-gray-200 dark:border-gray-700"),
		},
		"ContainerClass": {
			style.Class("flex flex-col sm:flex-row gap-4"),
		},
	},
	StyleGrouped: {
		"ContainerClass": {
			style.Class("border rounded-lg"),
			style.Color("border-gray-200 dark:bg-gray-700 dark:border-gray-600"),
		},
		"RadioContainerClass": {
			style.Class("flex items-center border-b last:border-b-0 px-4"),
		},
	},
	StyleGroupedHorizontal: {
		"ContainerClass": {
			style.Class("sm:flex border rounded-lg"),
			style.Color("border-gray-200 dark:bg-gray-700 dark:border-gray-600"),
		},
		"RadioContainerClass": {
			style.Class("flex items-center border-b sm:border-b-0 sm:border-r last:border-0 px-4 sm:w-full"),
		},
	},
	StyleLabelOnly: {
		"RadioContainerClass": {
			style.Class("inline-flex items-center justify-between"),
		},
		"RadioInputClass": {
			style.Color(""),
			style.Class("hidden peer"),
		},
		"RadioLabelClass": {
			style.Class("border p-2 rounded-lg cursor-pointer"),
			style.Color("text-gray-500 bg-white border-gray-200 dark:hover:text-gray-300 dark:border-gray-700 dark:peer-checked:text-blue-500 peer-checked:border-blue-600 peer-checked:text-blue-600 hover:text-gray-600 hover:bg-gray-100 dark:text-gray-400 dark:bg-gray-800 dark:hover:bg-gray-700"),
		},
	},
}

type D struct {
	// Name is the Name of all inputs.
	Name string
	// Style is the radiogroup style.
	Style style.Style
	// Radios is the list of radios in the group.
	//playground:import:github.com/jfbus/templ-components/components/radio
	//playground:default:[]radio.D{{Name: "foo", Value: "1", Label: "Choice 1"},{Name: "foo", Value: "2", Label:"Choice 2"}}
	Radios []radio.D
	// Message adds a validation message below the field.
	// Just add &message.D{} to add automatic validation.
	//playground:import:github.com/jfbus/templ-components/components/form/validation/message
	//playground:default:&message.D{Message: "Validation message"}
	Message *message.D
	// ContainerClass overrides the class of the div container.
	ContainerClass style.D
	// RadioContainerClass overrides the class of each radio div container.
	RadioContainerClass style.D
	// RadioContainerClass overrides the class of each radio input.
	RadioInputClass style.D
	// RadioContainerClass overrides the class of each radio label.
	RadioLabelClass style.D
}

func (def D) containerClass() string {
	return def.ContainerClass.CSSClass(Defaults, def.Style, "ContainerClass")
}

func (def D) radios() []radio.D {
	for i := range def.Radios {
		def.Radios[i].ID = def.Name + "-" + def.Radios[i].Value
		def.Radios[i].Name = def.Name
		def.Radios[i].InputClass = append(def.RadioInputClass.WithDefault(Defaults, def.Style, "RadioInputClass"), def.Radios[i].InputClass...)
		def.Radios[i].ContainerClass = append(def.RadioContainerClass.WithDefault(Defaults, def.Style, "RadioContainerClass"), def.Radios[i].ContainerClass...)
		def.Radios[i].LabelClass = append(def.RadioLabelClass.WithDefault(Defaults, def.Style, "RadioLabelClass"), def.Radios[i].LabelClass...)
	}
	return def.Radios
}

func (def D) message() message.D {
	m := *def.Message
	m.InputName = def.Name
	m.Style = def.Style
	return m
}
