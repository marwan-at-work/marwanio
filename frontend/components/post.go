package components

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"marwan.io/marwanio/frontend/js-wrappers/marked"
	"marwan.io/marwanio/frontend/stores/blogposts"
	"marwan.io/vecty-router"
)

// PostView represents a post
type PostView struct {
	vecty.Core
}

// Render returns every title
func (pv *PostView) Render() vecty.ComponentOrHTML {
	// TODO: safely check with ok var
	id := router.GetNamedVar(pv)["id"]
	p, err := blogposts.GetByID(id)
	if err == blogposts.ErrNotFound {
		return pv.renderErr()
	}

	output := marked.Marked(p.Markdown)

	return elem.Div(
		vecty.Markup(
			vecty.Class("blogpost-container"),
			vecty.UnsafeHTML(output),
		),
		pv.renderFooter(),
	)
}

func (pv *PostView) renderErr() vecty.ComponentOrHTML {
	return &notFound{}
}

func (pv *PostView) renderFooter() *vecty.HTML {
	return elem.Div(
		vecty.Markup(
			vecty.Class("twitter-footer"),
			vecty.UnsafeHTML(
				`<div>Follow me on <a href="https://www.twitter.com/MarwanSulaiman">Twitter</a> for updates and stuff.</div>`,
			),
		),
	)
}
