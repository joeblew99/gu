package gutrees

import (
	"errors"
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

var selfClosing = map[string]bool{
  "base":    true,
  "br":      true,
  "col":     true,
  "command": true,
  "embed":   true,
  "hr":      true,
  "img":     true,
  "input":   true,
  "keygen":  true,
  "link":    true,
  "meta":    true,
  "param":   true,
  "source":  true,
  "track":   true,
  "wbr":     true,
}

// ParseTree takes a string markup and returns a Markup which
// contains the full structure transpiled
// into the gutrees markup block structure.
func ParseTree(markup string) ([]Markup, error) {
	tokens, err := html.NewTokenizer(strings.NewReader(markup))
	if err != nil {
		return nil, err
	}

      depth := 0

      for {


        token := tokens.Next()

        switch token {
        case html.ErrorToken:
          return tokens.Err()

        case html.TextToken:

        case html.StartTagToken, html.EndTagToken:

      }

	return , nil
}


