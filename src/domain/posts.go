package domain

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/rafailmdzdv/blog/src/core"
)

type Post struct {
	Id          string
	Title       string
	Description string
	IconName    string `json:"icon_name"`
	Created     time.Time
}

func PostsFromAPI(c config.Map) []Post {
	r, _ := http.Get(c.API.Url + "api/v1/posts/")
	defer r.Body.Close()
	body, _ := io.ReadAll(r.Body)
	posts := []Post{}
	json.Unmarshal(body, &posts)
	return posts
}
