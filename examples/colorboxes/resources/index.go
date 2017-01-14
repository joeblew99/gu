package resources

import (
	"time"

	"github.com/gu-io/gu"
	"github.com/gu-io/gu/examples/colorboxes/app"
	"github.com/gu-io/gu/examples/colorboxes/css"
)

var _ = gu.Resource(func() {
	gu.Title("Resful Boxes")
	gu.Style(css.Index, nil, false)

	gu.View(app.New("http://localhost:6040/colors", 5*time.Second), "", false, false)
	gu.View(app.New("http://localhost:6040/colors", 7*time.Second), "", false, false)
	gu.View(app.New("http://localhost:6040/colors", 2*time.Second), "", false, false)
	gu.View(app.New("http://localhost:6040/colors", 6*time.Second), "", false, false)
})
