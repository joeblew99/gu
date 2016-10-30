package trees

import (
	"strings"

	"golang.org/x/net/html"
)

// ParseTree takes a string markup and returns a *Markup which
// contains the full structure transpiled
// into the gutrees markup block structure.
func ParseTree(markup string) ([]*Markup, error) {
	tokens := html.NewTokenizer(strings.NewReader(markup))

	rootElem := NewMarkup("div", false)
	pullNode(tokens, rootElem)
	return rootElem.Children(), nil
}

func pullNode(tokens *html.Tokenizer, root *Markup) {
	var node *Markup

	for {
		token := tokens.Next()

		switch token {
		case html.ErrorToken:
			return

		case html.TextToken, html.CommentToken, html.DoctypeToken:
			text := string(tokens.Text())

			if token == html.CommentToken {
				text = "<!--" + text + "-->"
			}

			if node != nil {
				NewText(text).Apply(node)
				continue
			}

			NewText(string(tokens.Text())).Apply(root)
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
