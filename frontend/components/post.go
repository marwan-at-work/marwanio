package components

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/marwan-at-work/marwanio/frontend/js-wrappers/marked"
	"github.com/marwan-at-work/marwanio/frontend/stores/blogposts"
	"github.com/marwan-at-work/vecty-router"
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

func (pv *PostView) renderErr() *vecty.HTML {
	return elem.Div(
		vecty.Text("not found"),
	)
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
