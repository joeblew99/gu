package main

import (
	"fmt"

	"honnef.co/go/js/dom"

	"github.com/gu-io/gu/shell/cache/webcache"
)

func main() {
	document := dom.GetWindow().Document().(dom.HTMLDocument)
	switch readyState := document.ReadyState(); readyState {
	case "loading":
		document.AddEventListener("DOMContentLoaded", false, func(_ dom.Event) {
			go setup()
		})
	case "interactive", "complete":
		setup()
	default:
		panic(fmt.Errorf("internal error: unexpected document.ReadyState value: %v", readyState))
	}
}

func setup() {
	doc := dom.GetWindow().Document().(dom.HTMLDocument)
	body := doc.QuerySelector("body").(*dom.HTMLBodyElement)

	body.AppendChild(doc.CreateTextNode("Opening Caches"))

	webCache, err := webcache.New()
	if err != nil {
		body.AppendChild(doc.CreateTextNode(err.Error()))
		body.AppendChild(doc.CreateElement("br"))
	}

	cache, err := webCache.Open("debi.db.v1")
	if err != nil {
		body.AppendChild(doc.CreateTextNode(err.Error() + "\n"))
		body.AppendChild(doc.CreateElement("br"))
		return
	}

	err = cache.Add("http://localhost:8080/github.com/gu-io/gu/shell/")
	if err != nil {
		body.AppendChild(doc.CreateTextNode(err.Error()))
		body.AppendChild(doc.CreateElement("br"))
	}

	res, err := cache.MatchPath("http://localhost:8080/github.com/gu-io/gu/shell/", webcache.MatchAttr{})
	if err != nil {
		body.AppendChild(doc.CreateTextNode(err.Error()))
		body.AppendChild(doc.CreateElement("br"))
		return
	}

	item := doc.CreateElement("div")
	item.SetInnerHTML(string(res.Body))

	body.AppendChild(item)
	body.AppendChild(doc.CreateElement("br"))
}
