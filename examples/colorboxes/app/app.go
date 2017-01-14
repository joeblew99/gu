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
//
// shell:component
//
// Resource {
//  Name: roboto.font.css
//	Path: https://fonts.googleapis.com/css?family=Lato|Open+Sans|Roboto
//	Hook: link-embed
//	Localize: true
// }
//
type RestfulBox struct {
	gu.Reactive

	Color    string
	endpoint string
	ticker   *time.Ticker
	closer   chan struct{}
}

// New returns a new instance of the *RestfulBox pointer.
func New(endpoint string, interval time.Duration) *RestfulBox {
	rb := RestfulBox{
		endpoint: endpoint,
		Reactive: gu.NewReactive(),
		ticker:   time.NewTicker(interval),
		closer:   make(chan struct{}),
	}

	return &rb
}

// OnSubscriptions registers the components actions for the two main steps for
// a component lifecycle.
func (r *RestfulBox) OnSubscriptions(mounted, render, unmounted gu.Subscriptions) {
	unmounted.React(func() {
		r.closer <- struct{}{}
	})

	// Lunch the go routine when the component gets mounted
	// which checks for update signals.
	mounted.React(func() {
		go func() {
			r.update()

		mloop:
			for {
				select {
				case <-r.ticker.C:
					r.update()
				case <-r.closer:
					break mloop
				}
			}
		}()
	})

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
	go r.Publish()
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
        display: inline-block;
      }
    `, r),
	)
}
