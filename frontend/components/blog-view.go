package components

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
	"marwan.io/vecty-router"
	"marwan.io/marwanio/blog"
	"marwan.io/marwanio/frontend/stores/blogposts"
)

// BlogView renders the blog collection
type BlogView struct {
	vecty.Core
	Posts []blog.Post
}

// Render renders collection of blog posts
func (b *BlogView) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(prop.ID("blog-container")),
		b.renderHeading(),
		b.getTitles(),
	)
}

func (b *BlogView) renderHeading() *vecty.HTML {
	return elem.Heading1(
		vecty.Markup(vecty.Class("blog-header")),
		vecty.Text("Blog-ish"),
	)
}

func (b *BlogView) getTitles() vecty.List {
	var ts vecty.List
	posts := blogposts.GetAll()
	for _, p := range posts {
		ts = append(ts, b.renderPostTitle(p))
	}

	return ts
}

func (b *BlogView) renderPostTitle(p blog.Post) vecty.ComponentOrHTML {
	return elem.Paragraph(
		vecty.Markup(vecty.Class("post-title")),
		elem.Span(
			vecty.Markup(vecty.Class("post-title-text")),
			router.Link(
				p.Link,
				p.Title,
				router.LinkOptions{},
			),
		),
		elem.Span(
			vecty.Markup(vecty.Class("post-title-date")),
			vecty.Text(p.CreatedAt.Format("2006, Jan 02")),
		),
	)
}
