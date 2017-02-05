package notifications

import "github.com/gu-io/gu/notifications/mque"

// dispatch provides a default dispatcher for listening to events.
var dispatch = mque.New()

// Subscribe adds a new listener to the dispatcher.
func Subscribe(q interface{}, end ...func()) mque.End {
	return dispatch.Q(q, end...)
}

// Dispatch emits a event into the dispatch callback listeners.
func Dispatch(q interface{}) {
	dispatch.Run(q)
}
