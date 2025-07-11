package main

import (
	"github.com/maxence-charriere/go-app/v10/pkg/app"
	"github.com/rafailmdzdv/blog/src/components/pages"
	"github.com/rafailmdzdv/blog/src/core"
	"github.com/rafailmdzdv/blog/src/domain"
)

func main() {
	config := config.DefaultConfig()
	title := "rafailmdzdv"
	app.Route("/", func() app.Composer { return &pages.Main{Title: title} })
	app.Route("/posts", func() app.Composer { return &pages.Posts{Title: title, Config: config} })
	app.RouteWithRegexp("^/posts/.+", func() app.Composer { return &pages.Post{Title: title, Config: config} })
	app.RunWhenOnBrowser()

	posts := domain.PostsFromAPI(config)
	pages := []string{"/", "/posts"}
	for _, post := range posts {
		pages = append(pages, "/posts/"+post.Id)
	}

	app.GenerateStaticWebsite(
		"dist",
		&app.Handler{
			Name:        "Main",
			Description: "Main page",
			Title:       title,
			Styles: []string{
				"/web/assets/styles.css",
			},
			HTML: func() app.HTMLHtml { return app.Html().DataSet("theme", "") },
		},
		pages...,
	)
}
