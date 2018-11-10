package components

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
	"marwan.io/vecty-router"
)

// Body renders the <body> tag
type Body struct {
	vecty.Core
}

// Render renders the <body> tag with the App as its children
func (b *Body) Render() vecty.ComponentOrHTML {
	return elem.Body(
		router.NewRoute("/", &MainView{}, router.NewRouteOpts{ExactMatch: true}),
		router.NewRoute("/blog", &BlogView{}, router.NewRouteOpts{ExactMatch: true}),
		router.NewRoute("/blog/{id}", &PostView{}, router.NewRouteOpts{ExactMatch: true}),
		router.NewRoute("/talks", &Talks{}, router.NewRouteOpts{ExactMatch: true}),
		router.NotFoundHandler(&notFound{}),
	)
}

type notFound struct {
	vecty.Core
}

func (nf *notFound) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(prop.ID("home-view")),
		elem.Div(
			vecty.Markup(prop.ID("home-top")),
			elem.Heading1(
				vecty.Text("page not found 🤦🏻‍♂️"),
			),
		),
	)
}
