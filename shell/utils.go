package shell

import (
	"crypto/rand"
	"errors"

	"github.com/gopherjs/gopherjs/js"
)

// RandString generates a set of random numbers of a set length
func RandString(n int) string {
	const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, n)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes)
}

// ObjectToWebRequest converts a object to a WebRequest.
func ObjectToWebRequest(o *js.Object) (WebRequest, error) {
	if o == nil || o == js.Undefined {
		return WebRequest{}, errors.New("Invalid/Undefined Object")
	}

	reqClone := o.Call("clone")

	body := make(chan []byte, 0)

	o.Call("text").Call("then", func(bd string) {
		go func() {
			body <- []byte(bd)
		}()
	})

	return WebRequest{
		Body:           <-body,
		Underline:      reqClone,
		URL:            o.Get("url").String(),
		Mode:           o.Call("mode").String(),
		B64Encode:      o.Get("b64_encode").Bool(),
		Cache:          CacheStrategy(o.Call("cache").String()),
		BodyUsed:       o.Get("bodyUsed").Bool(),
		Method:         o.Get("method").String(),
		Referrer:       o.Call("referrer").String(),
		Headers:        ObjectToMap(o.Get("headers")),
		Credentials:    o.Call("credentials").String(),
		ReferrerPolicy: o.Call("referrerPolicy").String(),
	}, nil
}

// ObjectToWebResponse converts a object to a WebResponse.
func ObjectToWebResponse(o *js.Object) (WebResponse, error) {
	if o == nil || o == js.Undefined {
		return WebResponse{}, errors.New("Invalid/Undefined Object")
	}

	resClone := o.Call("clone")

	body := make(chan []byte, 0)

	o.Call("text").Call("then", func(bd string) {
		go func() {
			body <- []byte(bd)
		}()
	})

	return WebResponse{
		Body:       <-body,
		Underline:  resClone,
		Ok:         o.Get("ok").Bool(),
		B64Encoded: o.Get("b64_encoded").Bool(),
		Status:     o.Get("status").Int(),
		Type:       o.Get("type").String(),
		Redirected: o.Get("redirected").Bool(),
		FinalURL:   o.Get("useFinalURL").String(),
		StatusText: o.Get("statusText").String(),
		Headers:    ObjectToMap(o.Get("headers")),
	}, nil
}

// AllObjectToWebResponse returns a slice of WebResponses from a slice
// of js.Object.
func AllObjectToWebResponse(o []*js.Object) []WebResponse {
	res := make([]WebResponse, 0)

	for _, ro := range o {
		if rq, err := ObjectToWebResponse(ro); err == nil {
			res = append(res, rq)
		}
	}

	return res
}

// AllObjectToWebRequest returns a slice of WebResponses from a slice
// of js.Object.
func AllObjectToWebRequest(o []*js.Object) []WebRequest {
	res := make([]WebRequest, 0)

	for _, ro := range o {
		if rq, err := ObjectToWebRequest(ro); err == nil {
			res = append(res, rq)
		}
	}

	return res
}

// MapToHeaders converts a map into a js.Headers structure.
func MapToHeaders(res map[string]string) *js.Object {
	header := js.Global.Get("Headers").New()

	for name, val := range res {
		header.Call("set", name, val)
	}

	return header
}

// WebResponseToJSResponse converts a object to a js.Response.
func WebResponseToJSResponse(res *WebResponse) *js.Object {
	if res.Underline != nil {
		return res.Underline
	}

	body := js.NewArrayBuffer(res.Body)
	// bodyBlob := js.Global.Get("Blob").New(body)

	res.Underline = js.Global.Get("Response").New(body, map[string]interface{}{
		"status":      res.Status,
		"statusText":  res.StatusText,
		"b64_encoded": res.B64Encoded,
		"headers":     MapToHeaders(res.Headers),
	})

	res.Underline.Set("b64_encoded", res.B64Encoded)

	return res.Underline
}

// WebRequestToJSRequest converts a object to a js.Request.
func WebRequestToJSRequest(res *WebRequest) *js.Object {
	if res.Underline != nil {
		return res.Underline
	}

	res.Underline = js.Global.Get("Request").New(res.URL, map[string]interface{}{
		"body":           res.Body,
		"mode":           res.Mode,
		"b64_encode":     res.B64Encode,
		"cache":          string(res.Cache),
		"method":         res.Method,
		"referrer":       res.Referrer,
		"credentials":    res.Credentials,
		"referrerPolicy": res.ReferrerPolicy,
		"headers":        MapToHeaders(res.Headers),
	})

	res.Underline.Set("b64_encode", res.B64Encode)

	return res.Underline
}

// ObjectToMap returns a map from a giving object.
func ObjectToMap(o *js.Object) map[string]string {
	res := make(map[string]string)

	for i := 0; i < o.Length(); i++ {
		item := o.Index(i)
		itemName := item.String()
		res[itemName] = o.Get(itemName).String()
	}

	return res
}

// ObjectToStringList returns a map from a giving object.
func ObjectToStringList(o *js.Object) []string {
	res := make([]string, 0)

	for i := 0; i < o.Length(); i++ {
		item := o.Index(i)
		res = append(res, item.String())
	}

	return res
}

// ObjectToList returns a map from a giving object.
func ObjectToList(o *js.Object) []*js.Object {
	res := make([]*js.Object, 0)

	for i := 0; i < o.Length(); i++ {
		item := o.Index(i)
		res = append(res, item)
	}

	return res
}
