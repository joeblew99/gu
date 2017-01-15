package main

import (
	"github.com/gu-io/gu"
	"github.com/gu-io/gu/app"
	_ "github.com/gu-io/gu/examples/colorboxes/resources"
)

func main() {
	boxes := app.New("colorboxes.v1", &gu.Options{
		Mode:        gu.DevelopmentMode,
		ManifestURI: "manifest.json",
	})

	boxes.Init(true)
}
