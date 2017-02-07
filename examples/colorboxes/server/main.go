//+build ignore

package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"

	"github.com/influx6/faux/context"
	"github.com/influx6/fractals/fhttp"

	"github.com/gu-io/gu"
	"github.com/gu-io/gu/app"
	_ "github.com/gu-io/gu/examples/colorboxes/resources"
)

func main() {

	base, _ := os.Getwd()
	assets := filepath.Join(base, "../assets")

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
		LocalMW: fhttp.FileServer(filepath.Join(base, "../manifest.json")),
	})

	approuter(fhttp.Endpoint{
		Path:   "/colors",
		Method: "GET",
		Action: func(ctx context.Context, rw *fhttp.Request) error {
			newColorDigit := int64(math.Floor(rand.Float64() * 16777215))
			newColorString := strconv.FormatInt(newColorDigit, 16)
			rw.RespondAny(200, "text/plain", []byte(fmt.Sprintf("#%s", newColorString)))
			return nil
		},
	})

	// approuter(fhttp.Endpoint{
	// 	Path:    "/",
	// 	Method:  "GET",
	// 	Action:  func(ctx context.Context, rw *fhttp.Request) error { return nil },
	// 	LocalMW: fhttp.FileServer(filepath.Join(base, "../index.html")),
	// })

	boxes := app.New("colorboxes.v1", &gu.Options{
		Mode:        gu.ProductionMode,
		ManifestURI: "http://localhost:6040/manifest.json",
	})

	approuter(fhttp.Endpoint{
		Path:   "/",
		Method: "GET",
		Action: func(ctx context.Context, rw *fhttp.Request) error {
			content := boxes.Init(true).RenderWithScript("/#", "assets/box.js")
			rw.RespondAny(200, "text/html", []byte(content.HTML()))
			return nil
		},
	})

	apphttp.Serve(":6040")
}
