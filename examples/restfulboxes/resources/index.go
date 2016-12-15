package resources

import (
	"time"

	. "github.com/gu-io/gu/design"
	"github.com/gu-io/gu/examples/restfulboxes/app"
	"github.com/gu-io/gu/examples/restfulboxes/css"
)

var _ = Resource(func() {

	DoTitle("Resful Boxes")

	DoStyle(css.Index, nil, false)

	DoView(app.New("http://localhost:6040/colors", 5*time.Second), "", false, false)
	DoView(app.New("http://localhost:6040/colors", 7*time.Second), "", false, false)
	DoView(app.New("http://localhost:6040/colors", 2*time.Second), "", false, false)
	DoView(app.New("http://localhost:6040/colors", 6*time.Second), "", false, false)

})
