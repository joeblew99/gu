package guviews

import (
	"strings"

	"github.com/influx6/gu/gutrees"
	"github.com/influx6/gu/gutrees/styles"
)

// ViewStates defines the two possible behavioral state of a view's markup
type ViewStates interface {
	Render(gutrees.Markup)
}

// HideView provides a ViewStates for Views inactive state
type HideView struct{}

// Render marks the given markup as display:none
func (v HideView) Render(m gutrees.Markup) {
	if ds, err := gutrees.GetStyle(m, "display"); err == nil {
		name, val := ds.Render()
		if !strings.Contains(val, "none") {
			ds.Reconcile(&gutrees.Style{Name: name, Value: "none"})
			return
		}
	}

	styles.Display("none").Apply(m)
}

// ShowView provides a ViewStates for Views active state
type ShowView struct{}

// Render marks the given markup with a display: block
func (v ShowView) Render(m gutrees.Markup) {
	if ds, err := gutrees.GetStyle(m, "display"); err == nil {
		name, val := ds.Render()
		if !strings.Contains(val, "none") {
			ds.Reconcile(&gutrees.Style{Name: name, Value: "block"})
			return
		}
	}

	styles.Display("block").Apply(m)
}

//==============================================================================
