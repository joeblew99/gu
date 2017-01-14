package main

import (
	"os"
	"path/filepath"

	"github.com/influx6/faux/context"
	"github.com/influx6/fractals/fhttp"

	"github.com/gu-io/gu/app"
	_ "github.com/gu-io/gu/examples/subscribe/pages"
)

func main() {
	base, _ := os.Getwd()
	assets := filepath.Join(base, "../assets")

	// Initialize the app and set it to use hash.
	app := app.New("subscribe.v1")

	apphttp := fhttp.Drive(fhttp.MW(fhttp.RequestLogger(os.Stdout)))(fhttp.MW(fhttp.ResponseLogger(os.Stdout)))

	approuter := fhttp.Route(apphttp)

	approuter(fhttp.Endpoint{
		Path:    "/assets/*",
		Method:  "GET",
		Action:  func(ctx context.Context, rw *fhttp.Request) error { return nil },
		LocalMW: fhttp.DirFileServer(assets, "/assets/"),
	})

	approuter(fhttp.Endpoint{
		Path:    "/manifest.json",
		Method:  "GET",
		Action:  func(ctx context.Context, rw *fhttp.Request) error { return nil },
		LocalMW: fhttp.FileServer(filepath.Join(base, "assets/manifest.json")),
	})

	approuter(fhttp.Endpoint{
		Path:   "/",
		Method: "GET",
		Action: func(ctx context.Context, rw *fhttp.Request) error {
			content := app.RenderWithScript("/#", "assets/app.js")
			rw.RespondAny(200, "text/html", []byte(content.HTML()))
			return nil
		},
	})

	apphttp.Serve(":6060")
}
