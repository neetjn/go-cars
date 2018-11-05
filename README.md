# go-cars

This project was created to learn Golang basics, as well as common use cases from other languages I use on a daily basis. I'll be recording my experiments and what I've learned while treking through Go.

## Experiments

* [x] Create a local package, module, or library and utilize it in our entry point.
* [x] Leverage go package manager for 3rd party packages.
* [x] Read from a file in a local directory.
* [x] Create a basic DTO (data transfer object).
* [x] Marshal and unmarshal DTO to string and Go structure.
* [x] Expose REST endpoints for the experiment application.
* [x] Create a service description for my REST endpoits, following the HATEOAS model.
* [x] Handle multiple different HTTP verbs.
* [x] Handle and manipulate different media types.
* [x] Experiment with decorators for simplifying common REST features.
* [ ] Use a go routine.
* [x] Dockerize application (include 3rd party packages).
* [ ] Create unit tests and execute on a CI/CD platform.

## What I've Learned

* All .go files in a directory must have the same package name.
  * All .go files in a directory are considered to be part of the "package".
  * All .go files will be expose all exportable methods or constants to the exported package namespace.
  * The entry point of a .go app is the main function in the main package.
* For a function, struct, or constant to be exposed to the package's namespace, it must be pascal cased.
* Go was designed to be semantically hardy. Unlike javascript or Python, where a single operation can be expressed in numerous forms, the philosophy behind Go was to design a programming language where there's a single way to achieve certain operations, inadvertently making the language incredibly verbose and easy to read.
* All go related projects must be properly placed in the `$GOPATH` to be compiled and ran properly. You **can** specify your project directory outside of the `$GOPATH`, but any packages installed with `go` or `dep` are placed directory in the current `$GOPATH` to be discovered.
* Go structures can be marshaled and unmarshaled into and from JSON by design.
* Because the standard library for Go provides so much out of the gate, the need for a framework or any misc. modules for a basic REST api is highly unnecessary.
  * However, when developing more advanced applications a 3rd party router may be useful for structuring and plain convenience ([httprouter](https://github.com/julienschmidt/httprouter), [mux](https://github.com/gorilla/mux)).
* Validation for Go struct fields is unintuitive out of the box. There are however many validation libraries such as [go-validator](https://github.com/go-validator/validator).

## Closing Thoughts

Go lang is a very intuitive, unintuitive programming language. The philosophy behind Go's design is that simpler code is better code. This can be both a blessing when working on larger projects that require structure and a perpetual nightmare when trying to prototype or flesh out an idea. Enough can be accomplished with the standard library for very basic use cases, but due the nature of the language -- it's impossible to avoid writing what feels like redundant code.

...

## Resources

...

---

Copyright (c) 2018 John Nolette Licensed under the MIT License.
