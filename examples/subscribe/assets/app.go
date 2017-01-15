package main

import (
	"github.com/gu-io/gu"
	"github.com/gu-io/gu/app"
	_ "github.com/gu-io/gu/examples/subscribe/pages"
)

func main() {
	app := app.New("subscribe.v1", &gu.Options{
		Mode:        gu.DevelopmentMode,
		ManifestURI: "manifest.json",
	})
	app.Init(true)
}
