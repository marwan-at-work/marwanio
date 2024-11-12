# A Short Note On Go2's Error Handling Draft
2018-09-03

I just finished reading the [Go2 error handling draft design](https://go.googlesource.com/proposal/+/master/design/go2draft-error-handling.md) and since it's meant to be in the review phase, I thought I'd give it a brief review.

The TL;DR is that the new error handling hurts readability. But I think it can be made better in at least one way:

The `handle err {}` block should not be at the top, it should be at the bottom.

Let's see why.

When you're reading a function for the first time, you should be able to quickly understand the happy path *then* be able to read some more and look into the edge cases. This idea is better explained by Alan Shreve's [talk](https://www.youtube.com/watch?v=2T6Prj82adg) about Conceptualizing Large Software Systems.

Therefore, if a function is called `ReadFile(s string) ([]byte, error)`, what should be the first thing in your line of site when you look at that function body? Probably what the function's name entails: opening a file and returning its content.

In today's Go, this is how you'd write this function

```go
func ReadFile(s string) ([]byte, error) {
    f, err := os.Open(s)
    if err != nil {
        return nil, fmt.Errorf("could not open %s: %v", s, err)
    }
    defer f.Close()
    bts, err := ioutil.ReadAll(f)
    if err != nil {
        return nil, fmt.Errorf("could not read contents of %s: %v", s, err)
    }

    return bts, nil
}
```

The readability of this function to me is optimal. This is because you can see the *core* operation of the function right at your site, unindented and unclobbered by the edge cases.

In the New Go, it can now be rewritten like this: 

```go
func ReadFile(s string) ([]byte, error) {
    handle err {
        return nil, fmt.Errorf("error reading %s: %v", s, err)
    }
    f := check os.Open(s)
    defer check f.Close()
    bts := check ioutil.ReadAll(f)

    return bts, nil
}
```

If you're encountering this code for the first time, the first line of this function tells you how you should handle an error that happens inside its body.

**But I don't even know what the function body does yet, why should I start reading how it handles its internal errors first?**

### A Thought

I suggest that we invert the proposal and have error handling be bottom-to-top, not top-to-bottom. So this is how it should look like: 

```go
func ReadFile(s string) ([]byte, error) {
    f := check os.Open(s)
    defer check f.Close()
    bts := check ioutil.ReadAll(f)

    handle err {
        return nil, fmt.Errorf("error reading %s: %v", s, err)
    }
    return bts, nil
}
```

This way, you can follow the happy path all the way down to the end, and the two returns are just two lines apart, which means you can see the returns of both happy and unhappy paths right next to each other. 

### Conclusion
I'm still not convinced that the check/handle design is a good one but I'll stop here and write more about it later.

Go, to me, is meant to be more readable than writeable. This is because writing code is much easier than reading code. Much, much easier. 
Writing code can be made even more easy by having tools help you. For example, I never had a problem writing `if err != nil` because I've set up VSCode Snippets to just expand with 3 key strokes and the code writes itself for me.

In 2018, we don't need to compromise readability when we can have smart tools do the writing for us. 

### PS

My number one favorite feature in Go is error handling. I'm quite a fan of Andrew Gerrand and Rob Pike's [experience report](https://commandcenter.blogspot.com/2017/12/error-handling-in-upspin.html) on error handling in Upspin. Not in how they did error handling, but in the underlying idea that every system should have its own *way* in handling errors. I have taken the idea to open source projects like [Athens](https://github.com/gomods/athens/blob/master/pkg/errors/doc.go) and I will be speaking about error handling at the New York Times [later this month](https://2018.gopherconbr.org/en/) 

I am, therefore, really curious to hear their thoughts on the Go2 proposal.