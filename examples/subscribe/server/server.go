package main

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/influx6/faux/context"
	"github.com/influx6/fractals/fhttp"

	"github.com/gu-io/gu/examples/subscribe"
	_ "github.com/gu-io/gu/examples/subscribe/pages"
)

func main() {

	base, _ := os.Getwd()
	assets := filepath.Join(base, "../assets")

	// Initialize the app and set it to use hash.
	subscribe.App.Init(true)

	app_http := fhttp.NewHTTP([]fhttp.DriveMiddleware{
		fhttp.WrapMiddleware(fhttp.Logger()),
	}, nil)

	app_router := fhttp.Route(app_http)

	app_router(fhttp.Endpoint{
		Path:    "/assets/*",
		Method:  "GET",
		Action:  func(ctx context.Context, rw *fhttp.Request) error { return nil },
		LocalMW: fhttp.FileServer(assets, "/assets/"),
	})

	app_router(fhttp.Endpoint{
		Path:   "/",
		Method: "GET",
		Action: func(ctx context.Context, rw *fhttp.Request) error {
			content := subscribe.App.RenderWithScript("/#", "assets/app.js")
			rw.RespondAny(200, "text/html", []byte(content.HTML()))
			return nil
		},
	})

	http.ListenAndServe(":6060", app_http)
}
