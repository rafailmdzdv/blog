package shared

import "fmt"
import "github.com/maxence-charriere/go-app/v10/pkg/app"

type Header struct {
	app.Compo

	title string
}

func (h *Header) Render() app.UI {
	return app.Header().Class("border-b").Body(
		app.Div().Class("max-w-7xl mx-auto px-6 py-6 flex justify-between items-center").Body(
			app.H1().Class("text-2xl font-bold tracking-tight text-indigo-600").Text(fmt.Sprintf("%s.blog", h.title)),
			app.Nav().Class("text-sm text-gray-500 space-x-4").Body(
				app.A().Href("/").Class("hover:text-indigo-600 transition").Text("Главная"),
				app.A().Href("/about").Class("hover:text-indigo-600 transition").Text("Обо мне"),
				app.A().Href("/posts").Class("hover:text-indigo-600 transition").Text("Посты"),
			),
		),
	)
}
