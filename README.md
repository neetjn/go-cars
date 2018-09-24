# go-cars

This project was created to learn Golang basics, as well as common use cases from other languages I use on a daily basis. I'll be recording my experiments and what I've learned while treking through Go.

## Experiments

* [x] Create a local package, module, or library and utilize it in our entry point.
* [x] Read from a file in a local directory.
* [x] Create a basic DTO (data transfer object).
* [x] Marshal and unmarshal DTO to string and Go structure.
* [x] Expose REST endpoints for the experiment application.
* [ ] Use a go routine.
* [ ] Dockerize application.
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

## Closing Thoughts

...

## Resources

...

---

Copyright (c) 2018 John Nolette Licensed under the MIT License.
