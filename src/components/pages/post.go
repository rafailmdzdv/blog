package pages

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
	"github.com/niklasfasching/go-org/org"
	"github.com/rafailmdzdv/blog/src/components/shared"
	"os"
	"strings"
)

type Post struct {
	app.Compo

	Title    string
	PostName string
}

func (p *Post) Render() app.UI {
	f, _ := os.Open(fmt.Sprintf("posts/test.org"))
	html, _ := org.New().Parse(f, "").Write(org.NewHTMLWriter())
	html = strings.Split(html, "</nav>")[1]
	html = fmt.Sprintf(`
		<div class="h-full my-12 text-center">%s</div>
	`, html)
	return &shared.Base{
		Title: p.Title,
		Children: []app.UI{
			app.Raw(html),
		},
	}
}

func (p *Post) OnNav(ctx app.Context) {
	splittedURL := strings.Split(ctx.Page().URL().String(), "/")
	postTitle := splittedURL[len(splittedURL)-1]
	p.PostName = postTitle
}
