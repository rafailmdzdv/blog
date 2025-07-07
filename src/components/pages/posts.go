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
			app.Main().Class("w-1/2 h-full mx-auto px-6 py-10 overflow-y-auto").Body(
				app.Div().Class("flex flex-col w-full h-full space-y-3").Body(
					app.Range(posts).Slice(func(i int) app.UI {
						post := posts[i]
						return app.A().Href("/posts/" + post.Title).Class("w-full h-1/8 px-6 rounded-2xl transition hover:shadow-lg bg-gray-50 hover:bg-white").Body(
							app.Div().Class("flex h-full justify-between items-center space-x-5").Body(
								app.Div().Body(
									app.Img().Class("w-16 rounded-full").Src(p.Config.CDN.Url+post.IconPath),
								),
								app.Div().Class("text-end").Body(
									app.H3().Class("text-xl font-semibold text-gray-900 group-hover:text-indigo-600 transition").Text(post.Title),
									app.P().Text(post.Metadata.Description),
								),
							),
						)
					}),
				),
			),
		},
	}
}
