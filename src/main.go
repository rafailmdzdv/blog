package main

import (
	"github.com/maxence-charriere/go-app/v10/pkg/app"
	"github.com/rafailmdzdv/blog/src/components/pages"
	"net/http"
)

func main() {
	title := "rafailmdzdv"
	app.Route("/", func() app.Composer { return &pages.Main{Title: title} })
	app.Route("/posts", func() app.Composer { return &pages.Posts{Title: title} })
	app.RouteWithRegexp("^/posts/.*.org", func() app.Composer { return &pages.Post{Title: title} })
	app.RunWhenOnBrowser()

	http.Handle("/", &app.Handler{
		Name:        "Main",
		Description: "Main page",
		Title:       title,
		Styles: []string{
			"/web/assets/styles.css",
		},
	})
	http.ListenAndServe(":8000", nil)
}
