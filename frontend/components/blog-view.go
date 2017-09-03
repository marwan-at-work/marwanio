package components

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
	"github.com/marwan-at-work/marwanio/blog"
	"github.com/marwan-at-work/marwanio/frontend/stores/blogposts"
	"github.com/marwan-at-work/vecty-router"
)

// BlogView renders the blog collection
type BlogView struct {
	vecty.Core
	Posts []blog.Post
}

// Render renders collection of blog posts
func (b *BlogView) Render() *vecty.HTML {
	return elem.Div(
		prop.ID("blog-container"),
		b.renderHeading(),
		b.getTitles(),
	)
}

func (b *BlogView) renderHeading() *vecty.HTML {
	return elem.Heading1(
		prop.Class("blog-header"),
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
		prop.Class("post-title"),
		elem.Span(
			prop.Class("post-title-text"),
			router.Link(
				p.Link,
				p.Title,
				router.LinkOptions{},
			),
		),
		elem.Span(
			prop.Class("post-title-date"),
			vecty.Text(p.CreatedAt.Format("2006, Jan 02")),
		),
	)
}
