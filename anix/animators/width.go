package animators

import (
	"fmt"

	"github.com/gopherjs/gopherjs/js"
	"github.com/gu-io/gu/anix"
	"honnef.co/go/js/dom"
)

// WidthAnimator defines a animator for the width property.
type WidthAnimator struct {
	To    int
	cache []cachable
}

type cachable struct {
	value int
	item  dom.HTMLElement
}

func (w *WidthAnimator) Apply(item interface{}) {
	obj, ok := item.(*js.Object)
	if !ok {
		return
	}

	domNode := dom.WrapHTMLElement(obj)
	styles := domNode.Style()

	w.cache = append(cachable{
		item:  domNode,
		value: styles.GetPropertyValue("width").Int(),
	})
}

// Update updates all values with the delta values ensuring all target items
// are updated accordingly.
func (w *WidthAnimator) Update(step int64, delta float64) {
	for _, item := range w.cache {
		item.value = item.value * delta
	}
}

// Render updates all values with the interpolated values and syncing those values
// to the views.
func (w *WidthAnimator) Render(interpolate float64) {
	for _, item := range w.cache {
		item.value = (item.value * interpolate) + item.value
		item.item.Style().SetProperty("width", fmt.Sprintf("%spx", item.value), "")
	}
}

// Width returns a giving rotation animator to handle rendering rotation for
// web animations.
func Width(to int) anix.FunnelRenderer {
	return func() anix.Renderable {
		return &WidthAnimator{To: to}
	}
}
