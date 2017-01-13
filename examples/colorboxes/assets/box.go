package main

import (
	"github.com/gu-io/gu/app"
	_ "github.com/gu-io/gu/examples/colorboxes/resources"
)

func main() {

	boxes := app.New("colorboxes.v1")
	boxes.Init(true)

}
