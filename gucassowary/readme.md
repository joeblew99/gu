# Gucassowary
 Gucassowary provides a cassowary constraint solver for the gu UI library,
 which is a efficient, but general purpose implementation of the cassowary
 constraint solver for User interfaces.

## Install

  ```bash

    go get -u github.com/influx6/gu/gucassowary/...

  ```

## API

```go


  var x = gucassowary.Const("x",40)
  var y = gucassowary.Var("x",50)

  var mux = gucasswoary.Exprs(gucassowary.Add, x,y)

  var z = gucassowary.Var("z",0)
  mux.Do(z)

  z.Value()// z => 90


  var muxEq = gucassowary.Equation()

```

## Example
