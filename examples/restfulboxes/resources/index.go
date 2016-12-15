package resources

import (
	. "github.com/gu-io/gu/design"
	"github.com/gu-io/gu/examples/restfulboxes/app"
	"github.com/gu-io/gu/examples/restfulboxes/css"
)

var _ = Resource(func() {

	DoTitle("Resful Boxes")

  DoStyle(css.Index, false)

  DoView(app.New("http://localhost:4040/colors"), 2 * time.Second), "", false, false)
  DoView(app.New("http://localhost:4040/colors"), 1 * time.Second), "", false, false)
  DoView(app.New("http://localhost:4040/colors"), 4 * time.Second), "", false, false)
  DoView(app.New("http://localhost:4040/colors"), 3 * time.Second), "", false, false)
})
