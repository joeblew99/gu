//+build ignore

package main

import (
	"github.com/gu-io/gu"
	"github.com/gu-io/gu/app"
	_ "github.com/gu-io/gu/examples/hello/apps"
)

func main() {
	app.New("hello.v1", &gu.Options{
		Mode:        gu.DevelopmentMode,
		ManifestURI: "manifest.json",
	}).Init(true)
}
