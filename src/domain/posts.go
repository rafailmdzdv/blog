package domain

import (
	"fmt"
	"os"
	"strings"
)

type Post struct {
	Title       string
	Description string
	FilePath    string
}

func PostsFromDir() []Post {
	dirName := "posts"
	d, _ := os.ReadDir(dirName)
	posts := []Post{}
	for _, f := range d {
		fileName := f.Name()
		fullPath := fmt.Sprintf("%s/%s", dirName, fileName)
		of, _ := os.ReadFile(fullPath)
		firstLine := ""
		for _, b := range of {
			if strings.ContainsAny(string(b), "*") {
				continue
			}
			firstLine += string(b)
			if b == '\n' {
				break
			}
		}
		posts = append(posts, Post{Title: fileName, Description: firstLine, FilePath: fullPath})
	}
	return posts
}
