package domain

import (
	"github.com/rafailmdzdv/blog/src/core"
	"io"
	"net/http"
	"regexp"
	"strings"
)

type Post struct {
	Title    string
	IconPath string
}

func PostsFromCDN(c config.Map) []Post {
	r, _ := http.Get(c.CDN.Url)
	defer r.Body.Close()
	body, _ := io.ReadAll(r.Body)
	re := regexp.MustCompile(`href=".*"`)
	postsF := re.FindAllString(string(body), -1)
	posts := []Post{}
	for _, f := range postsF {
		fileName := strings.TrimRight(strings.TrimLeft(f, `href="`), `"`)
		if fileName == "../" || fileName == "ico/" {
			continue
		}
		postTitle := strings.TrimRight(fileName, ".org")
		posts = append(
			posts,
			Post{
				Title:    postTitle,
				IconPath: "ico/" + postTitle + ".jpg",
			},
		)
	}
	return posts
}
