package pages

import (
	"github.com/maxence-charriere/go-app/v10/pkg/app"
	"github.com/rafailmdzdv/blog/src/components/shared"
	"github.com/rafailmdzdv/blog/src/domain"
)

type Posts struct {
	app.Compo

	Title    string
	posts	 []domain.Post
}

func (p *Posts) Render() app.UI {
	return &shared.Base{
		Title: p.Title,
		Children: []app.UI{
			app.Main().Class("max-w-6xl h-full mx-auto px-6 py-10").Body(
				app.Div().Class("grid md:grid-cols-2 lg:grid-cols-3 gap-10").Body(
					app.Range(p.posts).Slice(func(i int) app.UI {
						return app.A().Href(p.posts[i].FilePath).Class("group border rounded-2xl px-32 py-16 transition hover:shadow-lg bg-gray-50 hover:bg-white").Body(
							app.H3().Class("text-xl font-semibold text-gray-900 group-hover:text-indigo-600 transition").Text(p.posts[i].Title),
							app.P().Class("mt-2 text-sm text-gray-600").Text(p.posts[i].Description),
						)
					}),
				),
			),
		},
	}
}

func (p *Posts) OnMount(ctx app.Context) {
	ctx.Dispatch(func(ctx app.Context) {
		p.posts = domain.PostsFromDir()
	})
}
