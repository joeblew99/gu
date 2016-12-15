package app

import (
	"bytes"
	"io"
	"net/http"
	"time"

	"github.com/gu-io/gu"
	"github.com/gu-io/gu/trees"
	"github.com/gu-io/gu/trees/elems"
	"github.com/gu-io/gu/trees/property"
)

// RestfulBox defines a box which renders a rectangular box which changes color
// intervally.
type RestfulBox struct {
	gu.Reactive

	Color    string
	endpoint string
	ticker   *time.Ticker
}

func New(endpoint string, interval time.Duration) *RestfulBox {
	rb := RestfulBox{
		Reactive: gu.NewReactive(),
		endpoint: endpoint,
	}

	// Create a ticker.
	rb.ticker = time.NewTicker(interval)

	// Lunch the go routine which checks for update signals.
	go func() {
		for {
			_, closed := <-rb.ticker.C
			if closed {
				break
			}

			rb.update()
		}
	}()

	return &rb
}

// update sends a request to the server expecting a new color to be delivered
// which will be the colored used for updating the view.
func (r *RestfulBox) update() {
	res, err := http.DefaultClient.Get(r.endpoint)
	if err != nil {
		return
	}

	var buf bytes.Buffer

	defer res.Body.Close()
	io.Copy(&buf, res.Body)

	r.Color = string(buf.Bytes())

	// Publish and update the views.
	r.Publish()
}

// Render returns the markup which defines the look for the struct.
func (r *RestfulBox) Render() *trees.Markup {
	return elems.Div(
		property.ClassAttr("boxes"),
		elems.CSS(`
      $ {
        width: 100px;
        height: 100px;
        margin: 20px;
        padding: 10px;
        background: {{ .Color }};
      }
    `, r),
	)
}
