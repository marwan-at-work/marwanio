package components

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

// MainView is the parent level view.
type MainView struct {
	vecty.Core
}

// Render returns a <body> element with the entire app inside of it.
func (pv *MainView) Render() *vecty.HTML {
	return elem.Div(
		prop.ID("home-view"),
		pv.renderMainView(),
		pv.renderFooter(),
	)
}

func (pv *MainView) renderMainView() *vecty.HTML {
	return elem.Div(
		prop.ID("home-top"),
		&NameAndTitleView{},
		&Links{},
	)
}

func (pv *MainView) renderFooter() *vecty.HTML {
	return elem.Footer(
		prop.Class("footer-container"),
		elem.Div(
			vecty.Text("This website is written in "),
			elem.Anchor(
				prop.Href("https://www.github.com/gopherjs/vecty"),
				vecty.Text("Vecty"),
			),
			vecty.Text(" and "),
			elem.Anchor(
				prop.Href("https://www.github.com/gopherjs/gopherjs"),
				vecty.Text("GopherJS"),
			),
			vecty.Text(" - "),
			elem.Anchor(
				prop.Href("https://www.github.com/marwan-at-work/marwanio"),
				vecty.Text("(source)"),
			),
		),
		elem.Div(
			vecty.UnsafeHTML(
				"<span>It was also encrypted encrypt with love, I assume, by <a href=\"https://letsencrypt.org\">Let's Encrypt</a><span>",
			),
		),
	)
}
