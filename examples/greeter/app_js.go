// +build js

package main

import (
	"github.com/gu-io/gu"
	"github.com/gu-io/gu/drivers/gopherjs"
)

// AppSettings defines the attributes for the giving greeter app for using the GopherJS
// driver.
var AppSettings = gu.AppAttr{
	InterceptRequests: true,
	Name:              "GreeterApp",
	Mode:              gu.DevelopmentMode,
	Title:             "GreeterApp - Demonstrate a greeting web application with Gu",
	Manifests:         "assets/manifests.json",
	Driver:            gopherjs.NewJSDriver(),
}
