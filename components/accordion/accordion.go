package accordion

import (
	"strconv"

	"github.com/jfbus/templ-components/components/accordion/element"
)

type D struct {
	ID string
	//templplayground:import:github.com/jfbus/templ-components/components/helper
	//templplayground:default:[]element.D{{Title: "Section 1", Content:helpers.S("Content 1")},{Title: "Section 2", Content:helpers.S("Content 2")})
	Children []element.D
}

func (def D) id() string {
	return def.ID
}

func (def D) defaultState() string {
	st := "{ _open:'"
	for i := range def.Children {
		if def.Children[i].Open {
			st += def.ID + strconv.Itoa(i+1)
			break
		}
	}
	st += `',
	isOpen(value) {
		return this._open == value
	},
	toggle(value) {
		if (this._open == value) {
			this._open = ''
		} else {
			this._open = value
		}
	},
}`
	return st
}

func (def D) children() []element.D {
	for i := range def.Children {
		def.Children[i].ID = def.ID + strconv.Itoa(i+1)
	}
	return def.Children
}