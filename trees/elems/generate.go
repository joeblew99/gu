// +build ignore
// Credit to Richard Musiol (https://github.com/neelance/dom)
// His code was crafted to fit haiku's use

package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var elemNameMap = map[string]string{
	"g":                "Group",
	"font-face-face":   "Fontface",
	"font-face-format": "FontFaceFormat",
	"font-face-name":   "FontfaceName",
	"font-face-src":    "FontFaceSrc",
	"font-face-uri":    "FontfaceURI",
	"missing-glyph":    "MissingGlyph",
	"a":                "Anchor",
	"article":          "Article",
	"aside":            "Aside",
	"area":             "Area",
	"abbr":             "Abbreviation",
	"b":                "Bold",
	"base":             "Base",
	"bdi":              "BidirectionalIsolation",
	"bdo":              "BidirectionalOverride",
	"blockquote":       "BlockQuote",
	"br":               "Break",
	"cite":             "Citation",
	"col":              "Column",
	"colgroup":         "ColumnGroup",
	"datalist":         "DataList",
	"dialog":           "Dialog",
	"details":          "Details",
	"dd":               "Description",
	"del":              "DeletedText",
	"dfn":              "Definition",
	"Def":              "Definition",
	"dl":               "DescriptionList",
	"dt":               "DefinitionTerm",
	"G":                "Group",
	"em":               "Emphasis",
	"embed":            "Embed",
	"footer":           "Footer",
	"figure":           "Figure",
	"figcaption":       "FigureCaption",
	"fieldset":         "FieldSet",
	"h1":               "Header1",
	"h2":               "Header2",
	"h3":               "Header3",
	"h4":               "Header4",
	"h5":               "Header5",
	"h6":               "Header6",
	"hgroup":           "HeadingsGroup",
	"header":           "Header",
	"hr":               "HorizontalRule",
	"i":                "Italic",
	"iframe":           "InlineFrame",
	"img":              "Image",
	"ins":              "InsertedText",
	"kbd":              "KeyboardInput",
	"keygen":           "KeyGen",
	"li":               "ListItem",
	"meta":             "Meta",
	"menuitem":         "MenuItem",
	"nav":              "Navigation",
	"noframes":         "NoFrames",
	"noscript":         "NoScript",
	"ol":               "OrderedList",
	"option":           "Option",
	"optgroup":         "OptionsGroup",
	"p":                "Paragraph",
	"param":            "Parameter",
	"pre":              "Preformatted",
	"q":                "Quote",
	"rp":               "RubyParenthesis",
	"Ref":              "Reference",
	"rt":               "RubyText",
	"s":                "Strikethrough",
	"samp":             "Sample",
	"source":           "Source",
	"section":          "Section",
	"sub":              "Subscript",
	"sup":              "Superscript",
	"tbody":            "TableBody",
	"textarea":         "TextArea",
	"td":               "TableData",
	"tfoot":            "TableFoot",
	"th":               "TableHeader",
	"thead":            "TableHead",
	"tr":               "TableRow",
	"u":                "Underline",
	"ul":               "UnorderedList",
	"var":              "Variable",
	"track":            "Track",
	"wbr":              "WordBreakOpportunity",
}

//list of self closing tags
var autoclosers = map[string]bool{
	"use":     true,
	"area":    true,
	"base":    true,
	"col":     true,
	"command": true,
	"embed":   true,
	"hr":      true,
	"input":   true,
	"keygen":  true,
	"meta":    true,
	"param":   true,
	"source":  true,
	"track":   true,
	"wbr":     true,
	"br":      true,
}

func pullDoc(url string, fx func(doc *goquery.Document)) error {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return err
	}

	fx(doc)
	return nil
}

func main() {
	file, err := os.Create("elems.gen.go")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	fmt.Fprint(file, `// Package elems contains definition for different html element types and some custom node types

//go:generate go run generate.go

// Documentation source: "HTML element reference" by Mozilla Contributors, https://developer.mozilla.org/en-US/docs/Web/HTML/Element, licensed under CC-BY-SA 2.5.

// Documentation for custom types lies within the  "github.com/influx6/faux/domtrees" package

package elems

import (
	"github.com/influx6/gu/trees"
)

// Text provides the concrete implementation for using the domtrees.Text struct
func Text(content string) *trees.Markup {
	return trees.NewText(content)
}

// SVG provides the markup generator for the <svg> xml tag.
func SVG(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("svg",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

`)

	code := regexp.MustCompile("</?code>")

	err = pullDoc("https://developer.mozilla.org/en-US/docs/Web/SVG/Element", func(doc *goquery.Document) {
		doc.Find(".index ul li a").Each(func(i int, s *goquery.Selection) {
			link, _ := s.Attr("href")

			if !strings.HasPrefix(link, "/en-US/docs/Web/SVG/Element/") {
				return
			}

			if s.Parent().Find(".icon-trash, .icon-thumbs-down-alt, .icon-warning-sign").Length() > 0 {
				return
			}

			desc, _ := s.Attr("title")

			text := code.ReplaceAllString(s.Text(), "")

			name := text[1 : len(text)-1]

			// for key, item := range elemNameMap {
			// 	if strings.HasPrefix(name, key) || strings.HasSuffix(name, key) {
			// 		name = strings.Replace(name, key, item, 1)
			// 	}
			// }

			writeSVGElem(file, name, desc, link)
		})
	})

	if err != nil {
		log.Fatalf("Unable to pull SVG ELEMENTS: %s", err)
	}

	err = pullDoc("https://developer.mozilla.org/en-US/docs/Web/HTML/Element", func(doc *goquery.Document) {
		doc.Find(".quick-links a").Each(func(i int, s *goquery.Selection) {
			link, _ := s.Attr("href")
			if !strings.HasPrefix(link, "/en-US/docs/Web/HTML/Element/") {
				return
			}

			if s.Parent().Find(".icon-trash, .icon-thumbs-down-alt, .icon-warning-sign").Length() > 0 {
				return
			}

			desc, _ := s.Attr("title")

			text := s.Text()
			if text == "Heading elements" {
				writeElem(file, "h1", desc, link)
				writeElem(file, "h2", desc, link)
				writeElem(file, "h3", desc, link)
				writeElem(file, "h4", desc, link)
				writeElem(file, "h5", desc, link)
				writeElem(file, "h6", desc, link)
				return
			}

			name := text[1 : len(text)-1]

			if name == "html" || name == "head" || name == "body" {
				return
			}

			writeElem(file, name, desc, link)
		})
	})

	if err != nil {
		log.Fatalf("Unable to pull HTML ELEMENTS: %s", err)
	}

}

var badsymbs = regexp.MustCompile("-(.+)")

func writeSVGElem(w io.Writer, name, desc, link string) {
	var autocloser = autoclosers[name]
	funName := elemNameMap[name]

	if funName == "" {
		funName = name

		for badsymbs.MatchString(funName) {
			if simbs := badsymbs.FindStringSubmatch(funName); len(simbs) > 0 {
				item := capitalize(simbs[1])
				funName = badsymbs.ReplaceAllString(funName, item)
			}
		}

		funName = capitalize(funName)
	}

	fmt.Fprintf(w, `
// SVG%s provides the following for SVG XML elements ->
// %s
// https://developer.mozilla.org%s
func SVG%s(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("%s",%t)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}
`, funName, desc, link, funName, name, autocloser)
}

func writeElem(w io.Writer, name, desc, link string) {
	var autocloser = autoclosers[name]
	funName := elemNameMap[name]

	if funName == "" {
		funName = capitalize(name)
	}

	fmt.Fprintf(w, `
// %s provides the following for html elements ->
// %s
// https://developer.mozilla.org%s
func %s(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("%s",%t)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}
`, funName, desc, link, funName, name, autocloser)
}

// capitalize capitalizes the first character in a string
func capitalize(s string) string {
	return strings.ToUpper(s[:1]) + s[1:]
}
