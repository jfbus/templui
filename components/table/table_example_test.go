package table_test

import (
	"context"
	"os"

	"github.com/a-h/templ"
	"github.com/jfbus/templui/components/button"
	"github.com/jfbus/templui/components/style"
	"github.com/jfbus/templui/components/table"
	"github.com/jfbus/templui/components/table/cell"
	"github.com/jfbus/templui/components/table/row"
)

func ExampleC() {
	c := table.C(table.D{
		Style: table.StyleStripedRows,
		Header: &row.D{
			Cells: []string{"Email", "Name", "Status", ""},
		},
		Rows: []row.D{{
			Cells: []any{
				"John Doe",
				"john.doe@example.com",
				"active",
				cell.D{
					Content: button.C(button.D{
						Label: "disable",
						Attributes: templ.Attributes{
							"hx-delete": "/users/1",
						},
					}),
				},
			},
		}},
		CustomStyle: style.Custom{
			"table/cell": style.D{
				style.Set("text-center"),
			},
		},
	})
	_ = c.Render(context.TODO(), os.Stdout)
}
