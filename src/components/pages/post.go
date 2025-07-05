package pages

import (
	"net/http"

	"strings"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
	"github.com/niklasfasching/go-org/org"
	"github.com/rafailmdzdv/blog/src/components/shared"
	"github.com/rafailmdzdv/blog/src/core"
)

type Post struct {
	app.Compo
	Config config.Map

	Title       string
	PostContent string
}

var _ app.Navigator = (*Post)(nil)

func (p *Post) Render() app.UI {
	contentDIV := app.Div().Class("h-full my-12 text-center")
	if p.PostContent != "" {
		contentDIV = contentDIV.Body(app.Raw(p.PostContent))
	}
	return &shared.Base{
		Title:    p.Title,
		Children: []app.UI{contentDIV},
	}
}

func (p *Post) OnNav(ctx app.Context) {
	ctx.Async(func() {
		splittedPath := strings.Split(ctx.Page().URL().Path, "/")
		postTitle := splittedPath[len(splittedPath)-1]
		r, _ := http.Get(p.Config.CDN.Url + postTitle + ".org")
		defer r.Body.Close()
		html, _ := org.New().Parse(r.Body, "").Write(org.NewHTMLWriter())
		html = strings.Split(html, "</nav>")[1]
		ctx.Dispatch(func(ctx app.Context) {
			p.PostContent = html
		})
	})
}
