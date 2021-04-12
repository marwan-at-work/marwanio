package main

import (
	"github.com/hexops/vecty"
	"marwan.io/marwanio/frontend/components"
	"marwan.io/marwanio/frontend/stores/blogposts"
)

func main() {
	must(blogposts.Fetch())
	vecty.SetTitle("Marwan - Software Engineer")
	body := &components.Body{}
	vecty.RenderBody(body)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
