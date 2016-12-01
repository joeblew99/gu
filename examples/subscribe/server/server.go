package main

import (
  "net/http"
	"github.com/influx6/gu/design"
  "github.com/influx6/faux/context"
  "github.com/influx6/fractals/fhttp"

  _ "github.com/influx6/gu/examples/subscribe/pages"
)

func main() {

  app_http := fhttp.NewHTTP([]fhttp.DriveMiddleware{
    fhttp.WrapMiddleware(fhttp.Logger()),
  }, nil)

  app_router := fhttp.Route(app_http)

  app_router(fhttp.Endpoint{
    Path: "/assets/*",
    Method: "GET",
    Action: func(ctx context.Context, rw *fhttp.Request) error {return nil },
    LocalMW: fhttp.FileServer("./assets","/assets/"),
  })

  app_router(fhttp.Endpoint{
    Path: "/",
    Method: "GET",
    Action: func(ctx context.Context, rw *fhttp.Request) error {return nil },
    LocalMW: fhttp.IndexServer("./", "index.html",""),
  })


  http.ListenAndServe(":6060", app_http)
}
