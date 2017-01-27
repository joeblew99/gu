package shell

import (
	"encoding/base64"
	"path/filepath"
	"strings"

	"github.com/gu-io/gu/trees"
	"github.com/gu-io/gu/trees/elems"
)

var imageStyle = `
	-%s-img{
		background-image: url("data:image/%s;base64,%s")
	}
`

// ImageCSSEmbed defines a struct to hold the Fetcher for "img-css-embed" hook types.
// This hook retrieves the giving image data, transformed into
type ImageCSSEmbed struct{}

// Fetch returns the markup for the giving resource and where it should be inserted
// into the dom.
func (m ImageCSSEmbed) Fetch(fetch Fetch, attr ManifestAttr) (*trees.Markup, bool, error) {
	if attr.Content != "" {
		style := trees.NewMarkup("style", false)

		ext := strings.TrimPrefix(filepath.Ext(attr.Name), ".")
		encoded := base64.StdEncoding.EncodeToString([]byte(attr.Content))
		elems.Text(imageStyle, m.getNewName(attr.Name), ext, encoded).Apply(style)

		return style, true, nil
	}

	res, err := fetch.Do(attr.WebRequest())
	if err != nil {
		return nil, false, err
	}

	style := trees.NewMarkup("style", false)
	ext := strings.TrimPrefix(filepath.Ext(attr.Name), ".")
	encoded := base64.StdEncoding.EncodeToString(res.Body)
	elems.Text(imageStyle, m.getNewName(attr.Name), ext, encoded).Apply(style)

	return style, true, nil
}

// getNewName returns an appropriate css name from the provided string.
func (ImageCSSEmbed) getNewName(val string) string {
	newVal := strings.Replace(val, filepath.Ext(val), "", 1)
	return strings.Replace(newVal, ".", "-", -1)
}

// CSSLink defines a struct to hold the Fetcher for "css-link" hook types.
// This hook embeds the css link as a link tag into the dom.
type CSSLink struct{}

// Fetch returns the markup for the giving resource and where it should be inserted
// into the dom.
func (CSSLink) Fetch(fetch Fetch, attr ManifestAttr) (*trees.Markup, bool, error) {
	style := trees.NewMarkup("link", false)
	trees.NewAttr("href", attr.Path).Apply(style)
	trees.NewAttr("ref", "stylesheet").Apply(style)
	trees.NewAttr("type", "text/css").Apply(style)

	return style, true, nil
}

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

// JSLink defines a struct to hold the Fetcher for "js-link" hook types.
// This hook embeds the css link as a link tag into the dom.
type JSLink struct{}

// Fetch returns the markup for the giving resource and where it should be inserted
// into the dom.
func (JSLink) Fetch(fetch Fetch, attr ManifestAttr) (*trees.Markup, bool, error) {
	script := trees.NewMarkup("script", false)
	trees.NewAttr("src", attr.Path).Apply(script)
	trees.NewAttr("type", "text/javascript").Apply(script)

	return script, true, nil
}

func init() {
	// Register CSS hooks.
	Register("css-link", CSSLink{})
	Register("css-embed", CSSEmbed{})

	// Register Javascript hooks.
	Register("js-link", JSLink{})
	Register("js-embed", JavascriptEmbed{})

	// Register Images hooks.
	Register("img-css-embed", ImageCSSEmbed{})
}
