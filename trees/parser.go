package trees

import (
	"strings"

	"golang.org/x/net/html"
)

// ParseToRoot passes the markup generated from the markup added to the provided
// root.
func ParseToRoot(root *Markup, markup string) {
	trees := ParseTree(markup)
	for _, child := range trees {
		child.Apply(root)
	}
}

// ParseAndFirst expects the markup provided to only have one root element which
// will be returned.
func ParseAndFirst(markup string) *Markup {
	trees := ParseTree(markup)
	if len(trees) != 1 {
		panic("Markup must only returned single item in tree")
	}

	return trees[0]
}

// ParseAsRoot returns the markup generated from the provided markup,
// returning them as children of the provided root.
func ParseAsRoot(root string, markup string) *Markup {
	tokens := html.NewTokenizer(strings.NewReader(markup))

	var sel *Selector
	if sels := Query.ParseSelector(root); sels != nil {
		sel = sels[0]
	} else {
		sel.Tag = root
	}

	rootElem := NewMarkup(sel.Tag, false)

	if sel.ID != "" {
		NewAttr("id", sel.ID).Apply(rootElem)
	}

	if sel.Classes != nil {
		(&ClassList{list: sel.Classes}).Apply(rootElem)
	}

	pullNode(tokens, rootElem)

	return rootElem
}

// ParseTree takes a string markup and returns a *Markup which
// contains the full structure transpiled
// into the gutrees markup block structure.
func ParseTree(markup string) []*Markup {
	tokens := html.NewTokenizer(strings.NewReader(markup))

	rootElem := NewMarkup("div", false)
	pullNode(tokens, rootElem)

	return rootElem.Children()
}

func pullNode(tokens *html.Tokenizer, root *Markup) {
	var node *Markup

	for {
		token := tokens.Next()

		switch token {
		case html.ErrorToken:
			return

		case html.TextToken, html.CommentToken, html.DoctypeToken:
			text := strings.TrimSpace(string(tokens.Text()))

			if text == "" {
				continue
			}

			if token == html.CommentToken {
				text = "<!--" + text + "-->"
			}

			if node != nil {
				NewText(text).Apply(node)
				continue
			}

			NewText(text).Apply(root)
			continue

		case html.StartTagToken, html.EndTagToken, html.SelfClosingTagToken:
			if token == html.EndTagToken {
				node = nil
				return
			}

			tagName, hasAttr := tokens.TagName()

			node = NewMarkup(string(tagName), token == html.SelfClosingTagToken)
			node.Apply(root)

			if hasAttr {
			attrLoop:
				for {
					key, val, more := tokens.TagAttr()

					NewAttr(string(key), string(val)).Apply(node)

					if !more {
						break attrLoop
					}
				}
			}

			if token == html.SelfClosingTagToken {
				continue
			}

			pullNode(tokens, node)
		}
	}
}
