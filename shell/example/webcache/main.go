package main

import (
	"honnef.co/go/js/dom"

	"github.com/gu-io/gu/shell/cache/webcache"
)

func main() {
	doc := dom.GetWindow().Document().(dom.HTMLDocument)
	body := doc.QuerySelector("body").(*dom.HTMLBodyElement)

	webCache, err := webcache.New()
	if err != nil {
		body.AppendChild(doc.CreateTextNode(err.Error()))
		body.AppendChild(doc.CreateElement("br"))
	}

	cacheRes := <-webCache.Open("debi.v1")

	if cacheRes.Error != nil {
		body.AppendChild(doc.CreateTextNode(cacheRes.Error.Error() + "\n"))
		body.AppendChild(doc.CreateElement("br"))
		return
	}

	err = <-cacheRes.Cache.Add("http://localhost:8080/github.com/gu-io/gu/shell/")
	if err != nil {
		body.AppendChild(doc.CreateTextNode(err.Error()))
		body.AppendChild(doc.CreateElement("br"))
	}

	resChan := cacheRes.Cache.MatchPath("http://localhost:8080/github.com/gu-io/gu/shell/", webcache.MatchAttr{})

	res := <-resChan
	if res.Error != nil {
		body.AppendChild(doc.CreateTextNode(res.Error.Error()))
		body.AppendChild(doc.CreateElement("br"))
		return
	}

	item := doc.CreateElement("div")
	item.SetInnerHTML(string(res.Response.Body))

	body.AppendChild(item)
	body.AppendChild(doc.CreateElement("br"))
}
