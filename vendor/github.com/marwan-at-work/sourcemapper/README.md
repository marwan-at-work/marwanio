# SourceMapper

SourceMapper exposes an `http.Handler` that serves any `.go` files that exist in your GOPATH or GOROOT. 
The purpose of this package is to implement `.go` source maps for GopherJS by using your own custom server without having to run `gopherjs serve`. 

## Usage 