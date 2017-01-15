package shell

import "github.com/gu-io/gu/trees"

// CSSEmbed defines a struct to hold the Fetcher for "css-embed" hook types.
type CSSEmbed struct{}

// Fetch returns the markup for the giving resource and where it should be inserted
// into the dom.
func (CSSEmbed) Fetch(fetch Fetch, attr ManifestAttr) (*trees.Markup, bool, error) {
	if attr.Content != "" {
		style := trees.NewMarkup("style", false)
		trees.NewText(attr.Content).Apply(style)

		return style, true, nil
	}

	res, err := fetch.Do(attr.WebRequest())
	if err != nil {
		return nil, false, err
	}

	style := trees.NewMarkup("style", false)
	trees.NewText(string(res.Body)).Apply(style)

	return style, true, nil
}

// JavascriptEmbed defines a struct to hold the Fetcher for "js-embed" hook types.
type JavascriptEmbed struct{}

// Fetch returns the markup for the giving resource and where it should be inserted
// into the dom.
func (JavascriptEmbed) Fetch(fetch Fetch, attr ManifestAttr) (*trees.Markup, bool, error) {
	if attr.Content != "" {
		script := trees.NewMarkup("script", false)
		trees.NewText(attr.Content).Apply(script)

		return script, true, nil
	}

	res, err := fetch.Do(attr.WebRequest())
	if err != nil {
		return nil, false, err
	}

	script := trees.NewMarkup("script", false)
	trees.NewText(string(res.Body)).Apply(script)

	return script, true, nil
}

func init() {
	Register("css-embed", CSSEmbed{})
	Register("js-embed", JavascriptEmbed{})
}
