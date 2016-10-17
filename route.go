package gu

// TreeRouter defines a interface which exposes a router which renders the
// the appropriate content for its element based on the current provided path.
type TreeRouter interface {
	MarkupRenderer
	Navigate(string) error
	Route(string, MarkupRenderer) error
}
