package pages

import (
	"compress/zlib"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"bytes"
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
	contentDIV := app.Div().Class("h-full mt-12 mb-16 flex flex-col items-center space-y-24 text-center")
	if p.PostContent != "" {
		contentDIV = contentDIV.Body(
			app.Raw(p.PostContent),
			app.Div().Class("w-1/2 mb-8").Body(
				app.Script().
					Src("https://giscus.app/client.js").
					DataSets(map[string]any{
						"repo":              "rafailmdzdv/blog",
						"repo-id":           "R_kgDOPD6V_g",
						"category-id":       "DIC_kwDOPD6V_s4CsloY",
						"mapping":           "pathname",
						"strict":            "0",
						"reactions-enabled": "1",
						"emit-metadata":     "0",
						"input-position":    "bottom",
						"theme":             "preferred_color_scheme",
						"lang":              "en",
					}).
					CrossOrigin("anonymous"),
			),
		)
	}
	return &shared.Base{
		Title:    p.Title,
		Children: []app.UI{contentDIV},
	}
}

type PostJSON struct {
	Content string
}

func (p *Post) OnNav(ctx app.Context) {
	ctx.Async(func() {
		splittedPath := strings.Split(ctx.Page().URL().Path, "/")
		postId := splittedPath[len(splittedPath)-1]
		r, _ := http.Get(fmt.Sprintf("%s/api/v1/posts/content/?id=%s", p.Config.API.Url, postId))
		defer r.Body.Close()
		postJSON := PostJSON{}
		body, _ := io.ReadAll(r.Body)
		json.Unmarshal(body, &postJSON)

		// decompressing
		decoded, _ := base64.StdEncoding.DecodeString(postJSON.Content)
		zlibReader, _ := zlib.NewReader(bytes.NewReader(decoded))

		html, _ := org.New().Parse(zlibReader, "").Write(org.NewHTMLWriter())
		html = strings.Split(html, "</nav>")[1]
		ctx.Dispatch(func(ctx app.Context) {
			p.PostContent = html
		})
	})
}
