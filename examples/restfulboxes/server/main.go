package main

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/influx6/faux/context"
	"github.com/influx6/fractals/fhttp"

	"github.com/gu-io/gu/app"
	_ "github.com/gu-io/gu/examples/restfulboxes/resources"
)

func main() {

	boxes := app.New().Init(true)

	base, _ := os.Getwd()
	assets := filepath.Join(base, "../assets")

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
		Path:   "/colors",
		Method: "GET",
		Action: func(ctx context.Context, rw *fhttp.Request) error {
			rw.Respond(200, []byte(""))
			return nil
		},
	})

	app_router(fhttp.Endpoint{
		Path:   "/",
		Method: "GET",
		Action: func(ctx context.Context, rw *fhttp.Request) error {
			content := boxes.RenderWithScript("/#", "assets/restful.js")
			rw.RespondAny(200, "text/html", []byte(content.HTML()))
			return nil
		},
	})

	http.ListenAndServe(":6040", app_http)
}
