package domain

import (
	"io"
	"net/http"
	"regexp"
	"strings"

	"github.com/niklasfasching/go-org/org"
	"github.com/rafailmdzdv/blog/src/core"
	"github.com/rafailmdzdv/blog/src/utils"
)

type Post struct {
	Title    string
	IconPath string
	Metadata metadata
}

type metadata struct {
	Title       string
	Description string
}

func PostsFromCDN(c config.Map) []Post {
	r, _ := http.Get(c.CDN.Url)
	defer r.Body.Close()
	body, _ := io.ReadAll(r.Body)
	re := regexp.MustCompile(`href=".*"`)
	postsF := re.FindAllString(string(body), -1)
	posts := []Post{}
	for _, f := range postsF {
		fileName := strings.TrimRight(strings.ReplaceAll(f, `href="`, ""), `"`)
		if fileName == "../" || fileName == "ico/" {
			continue
		}
		postTitle := strings.TrimSuffix(fileName, ".org")
		posts = append(
			posts,
			Post{
				Title:    postTitle,
				IconPath: "ico/" + postTitle + ".jpg",
				Metadata: postMetadata(fileName, c),
			},
		)
	}
	return posts
}

func postMetadata(filename string, c config.Map) metadata {
	m := metadata{}
	r, _ := http.Get(c.CDN.Url + filename)
	defer r.Body.Close()
	orgFile := org.New().Parse(r.Body, "")
	if props := utils.ParseProperties(orgFile.Nodes); len(props) != 0 {
		if val, ok := props["TITLE"]; ok {
			m.Title = val
		}
		if val, ok := props["DESCRIPTION"]; ok {
			m.Description = val
		}
	}
	return m
}
