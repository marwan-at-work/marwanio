# Wrapping Node Modules With GopherJS
2017-08-05

GopherJS lets you write Go code for the front end, but it also gives you access to the entire NPM registry.

When you call `gopherjs build .`, the compiler will look for files in the same directory that end 
with `.inc.js` and simply wrap that file in a function and call it at runtime after your `func main()`.

Say you have the following two files in a directory:

```go
// main.go
package main

import "fmt"

func main() {
    fmt.Println("hello")
}
```

```js
// bye.inc.js
console.log('good bye');
```

If you run `gopherjs build .` you will get `<dirname>.js` along with its source maps. If you run `node <dirname>.js` then you will see the following output:

```
$ node gophertest.js
hello
good bye
```

Okay, we're half way there! Sorta. 

A node module out of the box will not work this way.
The reason is that a node module has browser-incompatible syntax, specifically imports/exports. 

Therefore, you will need a tool such as Browserify or Webpack to take a Node.js module and turn it into a browser-friendly script. In this case, I will use Webpack. 

### Folder Structure

Let's create our mini app in a directory in your $GOPATH/src called `gophermodules`
This app will install the [isPrimitive](https://www.npmjs.com/package/is-primitive) npm module and run it in the browser.

First, install the npm library in that same directory: 

`npm i is-primitive` 

Now let's create a global link ot the npm module that the browser can access.

```js
// libs.js
window.isPrimitive = require('is-primitive');
```

Notice the name of this file is `libs.js` and not `libs.inc.js` because if this script ran on the browser it wouldn't know how to deal with `require('is-primitive')`. Therefore, we will use Webpack to bundle libs.js into libs.inc.js

Add a webpack.config.js file in the same directory:

```js
module.exports = {
    entry: [
        './libs.js',
    ],
    output: {
        filename: 'libs.inc.js',
    },
};
```

Webpack has a very large ecosystem for optimizations, so feel free to go nuts here. 

Run `webpack` -- of course if you don't have it: `npm i -g webpack`

And voila! Your newly installed npm library is now accessible globally. All you need to do is write a Go wrapper package. Let's create one. 

Create a new directory in `gophermodules` called `is-primitive` and add the following file

```go
package isprimitive

func IsPrimitive(val interface{}) bool {
    return js.Global.Call("isPrimitive", val).Bool()
}
```

Now let's go back to our main.go function and use it! 

```go
package main

import (
    "fmt"

    "gophermodules/is-primitive"
)

func main() {
    fmt.Println(
        isprimitive.IsPrimitive(3),
        isprimitive.IsPrimitive(map[string]string{}),
    )
}
```

Run `gopherjs build .`

Include `gophermodules.js` into your index.html -- I'll assume you know how to do this one. 

Now run the file in your browser and you should see the following output: 

```
true
false
```

And there! We just installed an npm library, wrapped it with Go bindings, bundled it with our Go code, and called it on the browser! 

Here's what your folder strucute should look like:

```
└── gophertest
    ├── gophertest.js
    ├── gophertest.js.map
    ├── index.html
    ├── is-primitive
    │   └── is-primitive.go
    ├── libs.inc.js
    ├── libs.js
    ├── main.go
    ├── node_modules
    │   └── is-primitive
    │       ├── LICENSE
    │       ├── README.md
    │       ├── index.js
    │       └── package.json
    └── webpack.config.js
```

### Things I like: 

1. Access to the entire npm registry.
2. Wrapping npm modules with Go, so you can use them in Go idiomatically and with defined types.

### Things I don't like: 

1. [Access to the entire npm registry.](https://twitter.com/o_cee/status/892306836199800836)
2. Defining your npm modules globally on `window`. 
3. Your build process just went from `gopherjs build .` to...well, JavaScript land.
