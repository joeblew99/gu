package trees

// MarkupQueries defines a package level variable for access the query interface
// which handles running css queries on markup structures.
var MarkupQueries queryCtrl

type queryCtrl struct{}

// Query returns the first element matching the giving selector.
func (queryCtrl) Query(root *Markup, sel string) *Markup {
	var found *Markup

	return found
}

// QueryAll returns all elements matching the giving selector.
func (queryCtrl) QueryAll(root *Markup, sel string) []*Markup {
	var found []*Markup

	return found
}

func (queryCtrl) classFor(root *Markup, class string) *Markup {
	var found *Markup

	return found
}

func (queryCtrl) idFor(root *Markup, id string) *Markup {
	var found *Markup

	return found
}

func (queryCtrl) contentContainsFor(root *Markup, attr string, content string) *Markup {
	var found *Markup

	return found
}

func (queryCtrl) contentMatchFor(root *Markup, attr string, content string) *Markup {
	var found *Markup

	return found
}
