package gudispatch

import (
	"github.com/go-humble/detect"
	"github.com/influx6/faux/loop"
	"github.com/influx6/faux/mque"
	"github.com/influx6/gu/gustates"
)

//==============================================================================

// history provides a central URL coordinator for gu views.
var history *HistoryProvider

// rootEngine defines the root state engine for managing state of
// view elements registered to it.
var rootEngine *gustates.StateEngine

// dispatch provides a default dispatcher for listening to events.
var dispatch = mque.New()

// Subscribe adds a new listener to the dispatcher.
func Subscribe(q interface{}) loop.Looper {
	return dispatch.Q(q)
}

// Dispatch emits a event into the dispatch callback listeners.
func Dispatch(q interface{}) {
	dispatch.Run(q)
}

// RegisterPage registers a giving state to a path on the root path engine.
func RegisterPage(su gustates.States, path string) {
	rootEngine.UseState(path, su)
}

// Navigate changes the URI path of the host browser accordingly.
func Navigate(q string) {
	if history != nil {
		history.Go(q)
	}
}

//==============================================================================

// init intializes the internal state management variables used in dispatch.
func init() {

	rootEngine = gustates.NewStateEngine()

	Subscribe(func(pd *PathDirective) {
		rootEngine.All(pd.Sequence)
	})

	if detect.IsBrowser() {
		history = History(HashSequencer)

		// n, m, hash := GetLocation()
		// Navigate(hash)
	}

}
