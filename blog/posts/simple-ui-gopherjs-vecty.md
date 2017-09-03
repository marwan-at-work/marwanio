# Writing A Simple UI With GopherJS and Vecty
2017-07-09

I recently started playing with [GopherJS](https://github.com/gopherjs/gopherjs) & [Vecty](https://github.com/gopherjs/vecty) to build a terribly simple portfolio website for myself -- the one you are currently viewing.

[GopherJS](https://github.com/gopherjs/gopherjs) is a Go library that compiles Go code to Javascript. It also gives you access to Javascript's global objects such as `window` and `global`. 

Vecty, on the other hand, is a React-like library that lets you write UI applications in Go. This will be my focus in this post.

The documentation in Vecty, as of writing date, is almost non-existent, but they have built a couple of [example applications](https://github.com/gopherjs/vecty/tree/master/examples) to play with in order for you to learn how to use Vecty.

## Getting Started

To render a component to the DOM you must first render the `<body>` tag.
To do so, run the following line in your main function:

```go
package main

import "github.com/gopherjs/vecty"

func main() {
	vecty.RenderBody(&MyComponent{})
}
```

Note that `MyComponent` must render a Body element. 

Let's get to the React-like `Components` part. You can define a component by embedding `vecty.Core` into a struct. 


```go
    type MyComponent struct {
        vecty.Core
    }
```

The next thing you need to do is define a `Render` method that returns `*vecty.HTML`. In our `MyComponent` case, it must return a Body element. Other child components can return any element they want. 

``` go
import (
    "github.com/gopherjs/vecty/elem"
    "github.com/gopherjs/vecty"
)

func (mc *MyComponent) Render() *vecty.HTML {
    return elem.Body(
        &MyChildComponent{},
        vecty.Text("some footer text"),
    )
}
```

### So how does the render method work? 

A Component's render method must return a `*vecty.HTML`. Think of it as a component that must return a DOM element. Afterall, you're coding a web page. This is where the `elem` package comes in. It gives you access to all DOM elements, such as the `<body>` tag above.

### What's going on inside the `elem.Body` function? 

If you hover over `elem.Body`, or any other element, you'll notice that it takes a variadic argument of `vecty.MarkupOrChild`. This is an empty interface that panics if you pass it the wrong interface, so don't rely on static typing here.

What this type says is that we can safely pass `Markup`, which we will get to in a second, or `Component`, i.e. `MyChildComponent`, or `HTML` which can be other html elements such as `elem.Div()`, `elem.Span()`, or `vecty.Text()`, as arugments. 

Finally let's look at what `MyChildComponent` renders:

``` go
import "github.com/gopherjs/vecty/prop"

func (mc *MyComponent) Render() *vecty.HTML {
    return elem.Div(
        vecty.Markup(prop.Class("my-main-container")),
        vecty.Text("Welcome to my site"),
    )
}    
```

Notice that the child Component is fundamentally the same as the parent Component. I just took the liberty here to introduce you to the `Markup` part aka the `vecty.Markup(prop.Class("my-main-container"))` part. The `prop` package let's you pass element attributes to your HTML such ass Class, ID, href, etc. 

### Running On The Browser

Now that you wrote your first simple Component, you can see it on the browser using [GopherJS](https://www.github.com/gopherjs/gopherjs)

If you know how to use Go tools, you know how to use GopherJS.

`go get github.com/gopherjs/gopherjs` 

And from your main package directory run 

`gopherjs build .`

This will spit out a JavaScript file as well source maps for it. You can then create an `index.html` file and include the JS file in a typical `<script>` tag. 

### Other useful stuff:

Vecty provides many other tools but the above should be enough to get you started. You can look at the examples I mentioned or you check out this website's source code [here](https://www.github.com/marwan-at-work/marwanio)

## Conclusion

When I first decided to play with GopherJS & Vecty, the intention was to just try them out and not care about any limitations. However, I was surprised by how easy and pleasant the transition was from Go on the backend to Go on the frontend. Kudos to the GopherJS & Vecty contributors!

As of this date, I wouldn't recommend this workflow for large & complex UIs just yet. Mainly because Node.js and React and its family of tools and libraries, are quite mature with a large and supportive community.

Furthermore, one complaint I have for languages that compile to JavaScript is SourceMaps instability. This is not a Go-specific issue. But in my experience, I find that SourceMaps, knowing nothing about their internals, tend to fail when your UI gets to a certain level of complexity which makes debugging annoying to the point of regret. 

All of that being said, I do think once you try GopherJS & Vecty, you will get optimistic very quickly.

Go by default solves a lot of JavaScript's limitations: static analysis, concurrency, dependencies, editor-tools and more. What is currently missing is essential UI-scalablity tools similar to Webpack's eco-system: chunking, progressive-loading, CSS processing, cached builds...etc.

In my experience, these tools are vital for building scalable User Interfaces like the ones we build at [work & co](www.work.co/clients). Therefore, once those tools are introduced to the GopherJS ecosystem, the combination of these tools with what Go offers will make a very compelling case for building heavy-duty UIs.


Next up I will write about writing a History API router for Vecty : ) 
