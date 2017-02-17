package gtkwebview

// WindowAttr defines the attribute structure which defines the options for
// a giving GTK window.
type WindowAttr struct {
	Width  int
	Height int
}

// Window defines the construct that creates a GTK Window and renders an internal
// webview into it for usage.
type Window struct {
	attr WindowAttr
}

// Create creates the new window, locking the thread and initializing gtk.
func (w *Window) Create() {

}
