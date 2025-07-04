package pages

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
	"github.com/rafailmdzdv/blog/src/components/shared"
)

type Main struct {
	app.Compo
	Title string
}

func (m *Main) Render() app.UI {
	return &shared.Base{
		Title: m.Title,
		Children: []app.UI{
			app.Section().Class("max-w-3xl mx-auto px-6 py-20 text-center").Body(
				app.H2().Class("text-4xl font-bold mb-4 tracking-tight").Text(fmt.Sprintf("👋 Привет, я %s", m.Title)),
			),
		},
	}
}
