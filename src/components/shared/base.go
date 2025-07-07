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
		app.Footer().Class("mt-24 py-8 border-t border-gray-200").Body(
			app.Div().Class("text-center text-sm text-gray-500").Text(
				fmt.Sprintf(
					"©%[1]s • Created by %[1]s • Powered by Go in %d",
					b.Title,
					time.Now().Year(),
				),
			),
		),
	)
	return app.Div().Class("flex flex-col").Body(
		app.Range(components).Slice(func(i int) app.UI {
			return components[i]
		}),
	)
}
