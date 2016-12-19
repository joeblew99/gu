package main

import (
	"github.com/gu-io/gu/app"
	_ "github.com/gu-io/gu/examples/restfulboxes/resources"
)

func main() {

	boxes := app.New()
	boxes.Init(true)

}
