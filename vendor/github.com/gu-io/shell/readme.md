# Library does not work and has no implementation of any of the concept below yet.

# Shell
Shell is an experiment on an idea to create a resource system which can be used
to generate how contents for offline loading.


## Install

```go
go get -u github.com/gu-io/shell
```

## Why
We need a system which will effeciently handling loading of external resources(javascript, css)
for apps built with Gu. Especially when using external javascript libraries eg. MDL that
provide interesting features and capabilities. Majorly this trends towards a system which
does not require embedding into a Gopherjs generate `.js` which ends up contributing to
size factor but allows us to easily use external resources with out much pain.

Shell is a library to cater to that need. It will be a sort of bootloader which boostraps
all this resources into the DOM ready for usage before initializing a gu based application.

We hope to extend the capabilities of Shell beyond this but this is the primary objective of
it's existence.

## Use Cases
There exists certain use cases that brought about the need for such a system, the major
3 are as below:

- Using an external javascript file
A common need is to use external javascript libraries within a project which would
be built with GopherJs. Although the `inc.js` method approach exists, it would be desirable
to have a flag that allows us expose

## Test cases
1. Loading an internal js resource (mdl)
2. Loading an external js resource (cdn)
3. OUT OF SCOPE: Loading an internal resource in a external golang package (mdl)
  - We can update the system to use codegen later, so this should work.

- Develop a broad system that allows usage of offline strategy like explained on the [Offline Cookbook](https://jakearchibald.com/2014/offline-cookbook/)
