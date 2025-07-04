package main

import (
	"context"
	"github.com/rafailmdzdv/blog/src/templates/pages"
	"os"
)

func main() {
	component := pages.Main("rafailmdzdv")
	f, _ := os.Create("index.html")
	component.Render(context.Background(), f)
}
