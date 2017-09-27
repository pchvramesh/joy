Golly gee, it's another Go to JS compiler.

![img](https://cldup.com/uQb67D_DJT.png)

## Helpful Resources

- https://astexplorer.net/: I've been using this to figure out how to build the JS tree
- http://goast.yuroyoro.net/: Simple Go AST viewer
- https://github.com/estree/estree/blob/master/es5.md: Link to the ES3 AST format, this is implemented in syntax.go
- To run individual tests: `go test -v -run Compiler/3`
- To run output files: `./nodejs/node testfiles/7-classes/output.js.txt`
- `pretty.Println(ast)` will pretty print the JS AST

## Open Questions

- [x] How will we deal with existing Go standard libraries?
  - Standard libraries are opt-in and won't work out of the box (these + DOM are the only libraries internal to Golly)
- [x] How will the different base Go types map to javascript types? (like maps)
  - This is where Golly may have to be more strict than Go (e.g. map[boolean]string)
- [x] How will we make DOM typesafe, but not include it in the build?
  - DOM will be by default be an server-side API (like jsdom), but include interfaces to JS AST Nodes that Golly will call into over transpiling the Go code into JS
- [ ] How will React/Preact look and work inside Go?
- [x] How can we implement the DOM's massive API
  - We'll try to generate this using Typescript's definitions
- [ ] How can we turn goroutines and channels into async/await?
- [ ] How does the JS variable scope differ from the Go's scope and what do we need to account for?
- [ ] Should Go's interfaces show up in the emitted Javascript in any form?
- [ ] Is it possible to "upgrade" an existing Go package to use Golly when you don't control the package's source, e.g. "fmt"?
- [ ] What is a []byte? Can it just be a string? How does that map?
- [ ] Should we rename this project to Jolly? :-D

## High-level TODO:

In no particular order.

### File-Level:

- [x] implement if statements
- [x] implement loops
- [x] implement nested structs
- [x] implement classes
- [x] implement arrays
- [x] implement maps
- [x] implement functions with multiple return values
- [x] implement anonymous functions
- [ ] implement slices
- [ ] implement built-in functions (append, copy, len, make)
- [ ] implement spreads (users...)
- [ ] implement switch statements
- [ ] handle User{"a"}
- [ ] breaks, continues, panics
- [ ] custom types (type Blah = string)
- [ ] global variables, constants, etc.

### Package-Level:

- [x] support multiple files
  - files are wrapped in closures and the package wraps the file closures in another closure calling main()
- [ ] see if we can mimic rollup to decrease filesize further

### Hard but probably necessary:

- [ ] support goroutines and channels(using async/await)
  - we need some async functionality for things like AJAX
    - otherwise callback hell in Go? 🤷‍♂️
  - probably makes sense to prototype this as a babel transform first 
  - if too hard or too costly, maybe just mutexes? gopherjs does seem to have a working implementation of goroutines and channels
  - for cross-browser we can run it through this: https://github.com/facebook/regenerator
    - this process relies on a Promise polyfill, but if written
    to use generators, we no longer need promises.
- [ ] handle differences in variable scope between JS <-> Golang

### Before you can use it everyday:

- [ ] write basic DOM package using Typescript's libdefs
  - Source: https://github.com/Microsoft/TypeScript/blob/master/lib/lib.dom.d.ts
  - Typescript has a script to generate these files, so maybe we can use that 
- [ ] write preact package (~1000LOC)
- [ ] write fetch package (https://github.com/developit/unfetch ~50LOC)
- [ ] basic development server (livereload is fine for now)

### Later:

- [ ] add spacing (space, tab, newline) tokens into the JS AST to improve our JS output format
- [ ] implement defer function () { ... }
- [ ] error out on maps without key strings
- [ ] recover statement
- [ ] goto + label statements
- [ ] support go's select statement
- [ ] track functions declarations in use to do topological sort and pruning

## Design

- Small single-file builds that run in all browsers
- Compiles to ES3, the most popular JS dialect
- The "bucklescript" for Go
- Fast builds

## Wait, why?

- Javascript has become increasingly difficult to package up and distribute.
- Even with WebAssembly, Javascript will be supported for years to come
- Go has a minimal syntax and clear future priorities
- Go is both easy to read and easy to maintain

## Tradeoffs

**Go runs Golly code, but Golly may not run Go code**

Just because your code is running in Go does not mean it will run on Golly without modification. For everything that browserify & webpack brought to the community,they also brought with it mysteriously large builds.

Golly aims to generate clean, minimal & performant ES3 code that works in every browser. It will not jump hoops to make your Go code work inside the browser. GopherJS is much better suited for that task.

**Golly doesn't include Go's wonderful standard library**

All 3rd party packages are opt-in. This is to make sure that our builds remain purposeful and small. Over time, Golly will add compatibility with certain standard library packages where it makes sense.