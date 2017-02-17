// Package gopherjs defines helper methods that use the underline gopherjs js package
// to inteface with the browser dom api.
package gopherjs

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/go-humble/detect"
	"github.com/gopherjs/gopherjs/js"
	"github.com/gu-io/gu/router"
)

//==============================================================================

// ListenForHistory sets up the necessary notifications for history.
func ListenForHistory(behaviour func(router.PushEvent)) {
	if BrowserSupportsPushState() {
		js.Global.Set("onpopstate", func() {
			host, path, hash, location := GetLocation()
			behaviour(router.PushEvent{
				Host:   host,
				Path:   path,
				Hash:   hash,
				Rem:    hash,
				From:   location,
				Params: make(map[string]string),
			})
		})
		return
	}

	js.Global.Set("onhashchange", func() {
		host, path, hash, location := GetLocation()
		behaviour(router.PushEvent{
			Host:   host,
			Path:   path,
			Hash:   hash,
			Rem:    hash,
			From:   location,
			Params: make(map[string]string),
		})
	})
}

// BrowserSupportsPushState checks if browser supports pushState
func BrowserSupportsPushState() bool {
	if !detect.IsBrowser() {
		return false
	}

	return (js.Global.Get("onpopstate") != js.Undefined) &&
		(js.Global.Get("history") != js.Undefined) &&
		(js.Global.Get("history").Get("pushState") != js.Undefined)
}

// SetLocationByPushEvent defines a function to set the current location using
// the provided PushDirective.
func SetLocationByPushEvent(ev router.PushDirectiveEvent) {
	_, path, hash, _ := GetLocation()

	if strings.HasPrefix(ev.To, "#") {
		hash = ev.To
	}

	if !strings.HasPrefix(ev.To, "#") && strings.Contains(ev.To, "#") {
		splits := strings.Split(ev.To, "#")
		path = splits[0]
		hash = splits[1]
	}

	if BrowserSupportsPushState() {
		PushDOMState(path, hash)
		return
	}

	SetDOMHash(path, hash)
}

// SetLocation defines a function to set the current location using the provided
// path and hash.
func SetLocation(path string, hash string) {
	if BrowserSupportsPushState() {
		PushDOMState(path, hash)
		return
	}

	SetDOMHash(path, hash)
}

// PushDOMState adds a new state the dom push history.
func PushDOMState(path string, hash string) {
	panicBrowserDetect()

	if path == "" {
		_, path, _, _ = GetLocation()
	}

	if !strings.HasPrefix(hash, "#") {
		hash = "#" + hash
	}

	// Use the advance pushState feature.
	baseURI := fmt.Sprintf("%s%s", path, hash)
	js.Global.Get("history").Call("pushState", nil, baseURI)

	// Set the browsers hash accordinly.
	js.Global.Get("location").Set("hash", hash)
}

// SetDOMHash sets the dom location hash.
func SetDOMHash(path string, hash string) {
	panicBrowserDetect()

	if path != "" {
		js.Global.Get("location").Set("path", path)
	}

	js.Global.Get("location").Set("hash", hash)
}

// GetLocation returns the path and hash of the browsers location api else
// panics if not in a browser.
func GetLocation() (host string, path string, hash string, location string) {
	if !detect.IsBrowser() {
		return
	}

	loc := js.Global.Get("location").String()
	ups, err := url.Parse(loc)
	if err != nil {
		return
	}

	host = ups.Host
	path = ups.Path
	hash = ups.Fragment
	location = loc

	if hash == "" {
		hash = "/#"
	}

	return
}

// DOMObjectToList takes a jsobjects and returns a list of internal objects by calling the item method
func DOMObjectToList(o *js.Object) []*js.Object {
	var out []*js.Object
	length := o.Get("length").Int()
	for i := 0; i < length; i++ {
		out = append(out, o.Call("item", i))
	}
	return out
}

// ChildNodeList returns the nodes list of the childNodes of the js object
func ChildNodeList(o *js.Object) []*js.Object {
	return DOMObjectToList(o.Get("childNodes"))
}

// Attributes takes a js object and extracts the attribute lists from it
func Attributes(co *js.Object) map[string]string {
	o := co.Get("attributes")

	if o == nil || o == js.Undefined {
		return nil
	}

	attrs := map[string]string{}

	length := o.Get("length").Int()

	for i := 0; i < length; i++ {
		item := o.Call("item", i)
		attrs[item.Get("name").String()] = item.Get("value").String()
	}

	return attrs
}

// GetWindow returns the global object which is the window in the browser
func GetWindow() *js.Object {
	return js.Global
}

// GetDocument returns the document js.object from the global window object
func GetDocument() *js.Object {
	win := GetWindow()
	if win == nil || win == js.Undefined {
		return nil
	}
	return win.Get("document")
}

// CreateElement creates a dom element using the document html js.object
func CreateElement(tag string) *js.Object {
	doc := GetDocument()
	if doc == nil || doc == js.Undefined {
		return nil
	}
	return doc.Call("createElement", tag)
}

// CreateDocumentFragment creates a dom documentFragment using the document html js.object
func CreateDocumentFragment() *js.Object {
	doc := GetDocument()
	if doc == nil {
		return nil
	}
	return doc.Call("createDocumentFragment")
}

// var onlySpace = regexp

// EmptyTextNode returns two bool values, the first indicating if its a text node and the second indicating if the text node is empty
func EmptyTextNode(o *js.Object) (bool, bool) {
	if o.Get("nodeType").Int() == 3 {
		textContent := strings.TrimSpace(o.Get("textContent").String())
		if textContent != "" {
			return true, false
		}
		return true, true
	}
	return false, false
}

// CleanAllTextNode removes all texts nodes within the container root
func CleanAllTextNode(o *js.Object) {
	for _, to := range ChildNodeList(o) {
		if istx, isem := EmptyTextNode(to); istx {
			if !isem {
				o.Call("removeChild", to)
			}
		}
	}
}

// UnWrapSpecialTextElements takes a dom object and unwraps all the Text UnknownELement within the lists
func UnWrapSpecialTextElements(o *js.Object) {
	texts := QuerySelectorAll(o, "text")
	for _, to := range texts {
		parent := to.Get("parentNode")
		SpecialAppendChild(parent, to)
		parent.Call("removeChild", to)
	}
}

// SpecialAppendChild takes a list of objects and calls appendNode on the given object, but checks if the objects contain an unknownelement with a text tag then strip the tagname and only apply its content
func SpecialAppendChild(o *js.Object, osets ...*js.Object) {
	for _, onode := range osets {
		if strings.ToLower(onode.Get("tagName").String()) == "text" {
			SpecialAppendChild(o, ChildNodeList(onode)...)
			continue
		}
		o.Call("appendChild", onode)
	}
}

// InsertBefore inserts the insert object before the chil object with the target
func InsertBefore(target, child, insert *js.Object) {
	target.Call("insertBefore", insert, child)
}

// AppendChild takes a list of objects and calls appendNode on the given object
func AppendChild(o *js.Object, osets ...*js.Object) {
	for _, onode := range osets {
		o.Call("appendChild", onode)
	}
}

var headerKids = map[string]bool{
	"meta":  true,
	"link":  true,
	"title": true,
	"base":  true,
}

var bodyKids = map[string]bool{
	"scripts": true,
}

// ContextAppendChild takes a list of objects and calls appendNode on the given object
func ContextAppendChild(o *js.Object, osets ...*js.Object) {
	for _, onode := range osets {

		if doHeadAppend(onode) {
			continue
		}

		if scripts := QuerySelectorAll(onode, "scripts"); len(scripts) != 0 {
			for _, script := range scripts {
				doHeadAppend(script)
			}
		}

		o.Call("appendChild", onode)
	}
}

func doHeadAppend(onode *js.Object) bool {
	tagNameObject := onode.Get("tagName")
	if tagNameObject == nil || tagNameObject == js.Undefined {
		return false
	}

	tagName := strings.ToLower(tagNameObject.String())
	if !headerKids[tagName] {
		return false
	}

	uid := GetAttribute(onode, "uid")
	header := QuerySelector(GetDocument(), "head")
	body := QuerySelector(GetDocument(), "body")

	if headerKids[tagName] && header != nil && header != js.Undefined {
		possibleNode := QuerySelector(header, tagName+"[uid='"+uid+"']")
		if possibleNode != nil && possibleNode != js.Undefined {
			ReplaceNode(header, onode, possibleNode)
			return true
		}

		header.Call("appendChild", onode)
		return true
	}

	if bodyKids[tagName] && body != nil && body != js.Undefined {
		possibleNode := QuerySelector(body, tagName+"[uid='"+uid+"']")
		if possibleNode != nil && possibleNode != js.Undefined {
			ReplaceNode(body, onode, possibleNode)
			return true
		}

		body.Call("appendChild", onode)
		return true
	}

	return false
}

// RemoveChild takes a target and a list of children to remove
func RemoveChild(o *js.Object, co ...*js.Object) {
	for _, onode := range co {
		o.Get("parentNode").Call("removeChild", onode)
	}
}

// IsEqualNode returns a false/true if the nodes are equal in the eyes of the dom
func IsEqualNode(newNode, oldNode *js.Object) bool {
	return oldNode.Call("isEqualNode", newNode).Bool()
}

// ReplaceNode replaces two unequal nodes using their parents
func ReplaceNode(target, newNode, oldNode *js.Object) {
	if newNode == oldNode {
		return
	}
	target.Call("replaceChild", newNode, oldNode)
}

// QuerySelectorAll returns the result of querySelectorAll on an object
func QuerySelectorAll(o *js.Object, sel string) []*js.Object {
	if sad := o.Get("querySelectorAll"); sad == nil || sad == js.Undefined {
		return nil
	}

	return DOMObjectToList(o.Call("querySelectorAll", sel))
}

// QuerySelector returns the result of querySelector on an object
func QuerySelector(o *js.Object, sel string) *js.Object {
	return o.Call("querySelector", sel)
}

// GetTag returns the tag of the js object
func GetTag(o *js.Object) string {
	return o.Get("tagName").String()
}

// GetAttribute returns a string if a key exists using the jsobject
func GetAttribute(o *js.Object, key string) string {
	return o.Call("getAttribute", key).String()
}

// HasAttribute returns true/false if a key exists using the jsobject
func HasAttribute(o *js.Object, key string) bool {
	return o.Call("hasAttribute", key).Bool()
}

// SetAttribute calls setAttribute on the js object with the value and key
func SetAttribute(o *js.Object, key string, value string) {
	o.Call("setAttribute", key, value)
}

// SetInnerHTML calls the innerHTML setter with the given string
func SetInnerHTML(o *js.Object, html string) {
	o.Set("innerHTML", html)
}

// panicBrowserDetect panics if the current gc is not a browser based
// one.
func panicBrowserDetect() {
	if !detect.IsBrowser() {
		panic("expected to be used in a dom/browser env")
	}
}
