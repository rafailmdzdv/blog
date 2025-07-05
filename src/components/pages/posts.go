package pages

import (
	"github.com/maxence-charriere/go-app/v10/pkg/app"
	"github.com/rafailmdzdv/blog/src/components/shared"
	"github.com/rafailmdzdv/blog/src/core"
	"github.com/rafailmdzdv/blog/src/domain"
)

type Posts struct {
	app.Compo
	Config config.Map

	Title string
}

func (p *Posts) Render() app.UI {
	posts := domain.PostsFromCDN(p.Config)
	return &shared.Base{
		Title: p.Title,
		Children: []app.UI{
			app.Main().Class("max-w-6xl h-full mx-auto px-6 py-10").Body(
				app.Div().Class("grid md:grid-cols-2 lg:grid-cols-3 gap-10").Body(
					app.Range(posts).Slice(func(i int) app.UI {
						post := posts[i]
						return app.A().Href("/posts/" + post.Title).Class("group border rounded-2xl px-16 py-16 transition hover:shadow-lg bg-gray-50 hover:bg-white").Body(
							app.Div().Class("flex items-center space-x-5").Body(
								app.Div().Body(
									app.Img().Class("rounded-full").Src(p.Config.CDN.Url+post.IconPath),
								),
								app.Div().Class("flex flex-col").Body(
									app.H3().Class("text-xl font-semibold text-gray-900 group-hover:text-indigo-600 transition").Text(post.Title),
								),
							),
						)
					}),
				),
			),
		},
	}
}
