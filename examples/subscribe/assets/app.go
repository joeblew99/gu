package main

import "github.com/gu-io/gu/app"
import _ "github.com/gu-io/gu/examples/subscribe/pages"

func main() {
	app := app.New("subscribe.v1")
	app.Init(true)
}
