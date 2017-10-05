# SourceMapper

SourceMapper exposes an `http.Handler` that serves any `.go` files that exist in your GOPATH or GOROOT. 
The purpose of this package is to implement `.go` source maps for GopherJS by using your own custom server without having to run `gopherjs serve`. 

This is recommended for development only as your production set up should have its own pipline of minifiers/uglifiers..etc.

## Usage & Docs

[![GoDoc](https://godoc.org/github.com/marwan-at-work/sourcemapper?status.svg)](https://godoc.org/github.com/marwan-at-work/sourcemapper)