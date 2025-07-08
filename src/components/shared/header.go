package shared

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type Header struct {
	app.Compo

	title string
}

func (h *Header) Render() app.UI {
	return app.Header().Class("border-b border-gray-200").Body(
		app.Div().Class("navbar max-w-7xl mx-auto px-6 py-6 flex justify-between items-center").Body(
			app.Div().Class("flex-1 text-info text-2xl font-bold").Body(app.A().Href("/").Text(fmt.Sprintf("%s.blog", h.title))),
			app.Div().Body(&ThemeSwap{}),
			app.Ul().Class("menu menu-horizontal flex-none rounded-box").Body(
				app.Li().Body(app.A().Href("/about").Class("hover:text-secondary").Text("About")),
				app.Li().Body(app.A().Href("/posts").Class("hover:text-secondary").Text("Posts")),
			),
		),
	)
}
