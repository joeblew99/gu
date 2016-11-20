package main


import (
  . "github.com/influx6/gu/design"
  . "github.com/influx6/gu/redom"
  _ "github.com/influx6/gu/examples/subscribe/pages"
  "honnef.co/go/js/dom"
)

func main(){
  New(&DOMRenderer{
    Document: dom.GetWindow().Document(),
  }).Init(true)
}

