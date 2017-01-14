package main

import (
	"github.com/gu-io/gu/app"
	_ "github.com/gu-io/gu/examples/hello/apps"
)

func main() {
	app.New("hello.v1").Init(true)
}
