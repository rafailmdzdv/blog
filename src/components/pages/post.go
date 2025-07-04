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
	html = fmt.Sprintf(`
		<div class="h-64 bg-gray-500">%s</div>
	`, html)
	fmt.Println(html)
	fmt.Println(p.PostName)
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
