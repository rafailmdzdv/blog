package shared

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
	"time"
)

type Base struct {
	app.Compo
	Title    string
	Children []app.UI
}

func (b *Base) Render() app.UI {
	components := []app.UI{&Header{title: b.Title}}
	components = append(components, b.Children...)
	components = append(
		components,
		app.Footer().Class("footer footer-center absolute bottom-0 py-4 mt-8 border-t border-gray-200 select-none").Body(
			app.Aside().Class("grid ml-auto grid-flow-col justify-self-center").Body(
				app.Div().Text(
					fmt.Sprintf(
						"©%[1]s • Created by %[1]s • Powered by Go in %d",
						b.Title,
						time.Now().Year(),
					),
				),
			),
			app.Nav().Class("grid grid-flow-col justify-self-end mr-5 text-3xl").Body(
				app.A().Href("https://github.com/rafailmdzdv").Text(""),
			),
		),
	)
	return app.Div().Class("min-h-full relative flex flex-col").Body(
		app.Range(components).Slice(func(i int) app.UI {
			return components[i]
		}),
	)
}
