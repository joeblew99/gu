package webcache

import (
	"errors"
	"fmt"

	"github.com/gopherjs/gopherjs/js"
	"github.com/gu-io/gu/shell"
)

// API creates a structure which exposes the new Cache API for
// browsers found in webkit browsers. https://developer.mozilla.org/en-US/docs/Web/API/Cache.
type API struct {
	*js.Object
}

// ErrInvalidState is returned when the cache api does not exists in the global
// context.
var ErrInvalidState = errors.New("Cache Not Found")

// New returns a new instance of the webcache.
func New() (*API, error) {

	if js.Global == nil || js.Global == js.Undefined {
		return nil, ErrInvalidState
	}

	caches := js.Global.Get("caches")

	if caches == nil || caches == js.Undefined {
		return nil, ErrInvalidState
	}

	var wc API
	wc.Object = caches
	return &wc, nil
}

// NewCacheResponse defines the response returned when a cache is request from the
// API.
type newCacheResponse struct {
	Error error     `json:"error"`
	Cache *CacheAPI `json:"cache"`
}

// ErrCacheNotFound defines the error returned when the cached desired is not found.
var ErrCacheNotFound = errors.New("Cache Not Found")

// Open retrieves a giving cache from the web cache API backend.
func (wc *API) Open(cacheName string) (*CacheAPI, error) {
	if wc.Object == nil || wc.Object == js.Undefined {
		return nil, ErrInvalidState
	}

	openReq := wc.Call("open", cacheName)
	if openReq == js.Undefined || openReq == nil {
		return nil, ErrCacheNotFound
	}

	var opVal newCacheResponse
	res := make(chan newCacheResponse, 0)

	openReq.Call("then", func(o *js.Object) {
		// go func() {
		res <- newCacheResponse{Cache: NewCacheAPI(o), Error: nil}
		// }()
	}, func(o *js.Object) {
		// go func() {
		res <- newCacheResponse{Error: errors.New(o.String())}
		// }()
	})

	opVal = <-res
	return opVal.Cache, opVal.Error
}

// CacheAPI defines individual cache item received from a call to the internal cache API.
type CacheAPI struct {
	*js.Object
}

// NewCacheAPI returns a new instance of the CacheAPI.
func NewCacheAPI(o *js.Object) *CacheAPI {
	var c CacheAPI
	c.Object = o
	return &c
}

// Add calls the internal caches.Cache.Add and caches.Cache.AddAll function matching against the
// request string/strings.
func (c *CacheAPI) Add(request ...string) error {
	resChan := make(chan error, 0)

	var isList bool

	if len(request) > 1 {
		isList = true
	}

	if !isList {
		c.Call("add", request[0]).Call("then", func() {
			// go func() {
			close(resChan)
			// }()
		}).Call("catch", func(err *js.Object) {
			// go func() {
			resChan <- errors.New(err.String())
			// }()
		})

		return <-resChan
	}

	c.Call("addAll", request).Call("then", func() {
		// go func() {
		close(resChan)
		// }()
	}).Call("catch", func(err *js.Object) {
		// go func() {
		resChan <- errors.New(err.String())
		// }()
	})

	return <-resChan
}

// PutPath the giving path and response into the cache.
func (c *CacheAPI) PutPath(request string, res shell.WebResponse) error {
	resChan := make(chan error, 0)

	resObj := shell.WebResponseToJSResponse(&res)

	c.Call("put", request, resObj).Call("catch", func(err *js.Object) {
		// go func() {
		resChan <- errors.New(err.String())
		// }()
	})

	return <-resChan
}

// Put calls the internal caches.Cache.Put function matching against the
// request string.
func (c *CacheAPI) Put(request shell.WebRequest, res shell.WebResponse) error {
	resChan := make(chan error, 0)

	resObj := shell.WebResponseToJSResponse(&res)
	reqObj := shell.WebRequestToJSRequest(&request)

	c.Call("put", reqObj, resObj).Call("then", func(_ *js.Object) {
		// go func(){
		close(resChan)
		// }()
	}).Call("catch", func(err *js.Object) {
		// go func(){
		resChan <- errors.New(err.String())
		// }()
	})

	return <-resChan
}

// GetManifest returns the giving request, response associated with the given
// manifest in the cache.
func (c *CacheAPI) GetManifest(man shell.ManifestAttr) (shell.WebRequest, shell.WebResponse, error) {
	req := man.WebRequest()

	res, err := c.GetRequest(req)
	if err != nil {
		return req, res, err
	}

	return req, res, nil
}

// MatchAttr defines the attributes which are optional passed to the match
// method to refine the cache search term.
type MatchAttr struct {
	IgnoreSearch bool
	IgnoreMethod bool
	IgnoreVary   bool
	CatchName    string
}

// GetRequest calls CacheAPI.Match and passing in a default MatchAttr value.
func (c *CacheAPI) GetRequest(request shell.WebRequest) (shell.WebResponse, error) {
	return c.Match(request, MatchAttr{})
}

// GetPath calls CacheAPI.MatchPath and passing in a default MatchAttr value.
func (c *CacheAPI) GetPath(request string) (shell.WebRequest, shell.WebResponse, error) {
	wr := shell.WebRequest{
		Method:      "GET",
		URL:         request,
		Credentials: "same-origin",
		Cache:       shell.DefaultStrategy,
	}

	res, err := c.MatchPath(request, MatchAttr{})
	if err != nil {
		return wr, res, err
	}

	return wr, res, nil
}

// CacheResponse defines the response returned when a cache request is made.
type cacheResponse struct {
	Error    error             `json:"error"`
	Response shell.WebResponse `json:"response"`
}

// CacheResponseChannel defines a channel type for the response to be received
// over a request.
type cacheResponseChannel chan cacheResponse

// MatchPath calls the internal caches.Cache.Match function matching against the
// request string.
func (c *CacheAPI) MatchPath(request string, attr MatchAttr) (shell.WebResponse, error) {
	resChn := make(cacheResponseChannel, 0)

	c.Call("match", request, attr).Call("then", func(response *js.Object) {
		res, err := shell.ObjectToWebResponse(response)
		if err != nil {
			// go func() {
			resChn <- cacheResponse{Error: err}
			// }()
			return
		}

		// go func() {
		resChn <- cacheResponse{
			Response: res,
		}
		// }()
	}).Call("catch", func(err *js.Object) {
		// go func() {
		resChn <- cacheResponse{Error: errors.New(err.String())}
		// }()
	})

	opVal := <-resChn
	return opVal.Response, opVal.Error
}

// Match calls the internal caches.Cache.Match function matching against the
// request string for a js.Cache.Match.
func (c *CacheAPI) Match(request shell.WebRequest, attr MatchAttr) (shell.WebResponse, error) {
	resChan := make(cacheResponseChannel, 0)

	c.Call("match", shell.WebRequestToJSRequest(&request), attr).Call("then", func(response *js.Object) {
		res, err := shell.ObjectToWebResponse(response)
		if err != nil {
			// go func() {
			resChan <- cacheResponse{Error: err}
			// }()
			return
		}

		// go func() {
		resChan <- cacheResponse{
			Response: res,
			Error:    nil,
		}
		// }()
	}).Call("catch", func(err *js.Object) {
		// go func() {
		resChan <- cacheResponse{Error: errors.New(err.String())}
		// }()
	})

	opVal := <-resChan
	return opVal.Response, opVal.Error
}

// CacheAllResponse defines the slice of responses returned when a cache request
// is made to matchAll.
type cacheAllResponse struct {
	Error    error               `json:"error"`
	Response []shell.WebResponse `json:"response"`
}

// CacheAllResponseChannel defines a channel type for the response to be received
// over a request to js.Cache.MatchAll.
type cacheAllResponseChannel chan cacheAllResponse

// MatchAllPath calls the internal caches.Cache.MatchAll function matching against the
// request string for a js.Cache.Match.
func (c *CacheAPI) MatchAllPath(request string, attr MatchAttr) ([]shell.WebResponse, error) {
	resChan := make(cacheAllResponseChannel, 0)

	c.Call("match", request, attr).Call("then", func(responses *js.Object) {
		// go func() {
		resChan <- cacheAllResponse{
			Response: shell.AllObjectToWebResponse(shell.ObjectToList(responses)),
		}
		// }()
	}).Call("catch", func(err *js.Object) {
		// go func() {
		resChan <- cacheAllResponse{Error: errors.New(err.String())}
		// }()
	})

	opVal := <-resChan
	return opVal.Response, opVal.Error
}

// MatchAll calls the internal caches.Cache.MatchAll function matching against the
// request string for a js.Cache.Match.
func (c *CacheAPI) MatchAll(request shell.WebRequest, attr MatchAttr) ([]shell.WebResponse, error) {
	resChan := make(cacheAllResponseChannel, 0)

	c.Call("match", shell.WebRequestToJSRequest(&request), attr).Call("then", func(responses *js.Object) {
		// go func() {
		resChan <- cacheAllResponse{
			Response: shell.AllObjectToWebResponse(shell.ObjectToList(responses)),
		}
		// }()
	}).Call("catch", func(err *js.Object) {
		// go func() {
		resChan <- cacheAllResponse{Error: errors.New(err.String())}
		// }()
	})

	opVal := <-resChan
	return opVal.Response, opVal.Error
}

// Delete calls the underline CacheAPI.DeleteRequest.
func (c *CacheAPI) Delete(request shell.WebRequest) error {
	return c.DeleteRequest(request, MatchAttr{})
}

// ErrRequestDoesNotExists is returned when the giving request does not exists in
// the catch.
var ErrRequestDoesNotExists = errors.New("Request not existing")

// DeletePath deletes the given path from the cache if found.
func (c *CacheAPI) DeletePath(request string, attr MatchAttr) error {
	resChn := make(chan error, 0)

	c.Call("delete", request, attr).Call("then", func(response *js.Object) {
		if response.Bool() {
			// go func(){
			close(resChn)
			// }()
			return
		}

		// go func(){
		resChn <- ErrRequestDoesNotExists
		// }()

	}).Call("catch", func(err *js.Object) {
		// go func() {
		resChn <- errors.New(err.String())
		// }()
	})

	return <-resChn
}

// DeleteRequest calls the internal caches.Cache.Delete function matching against the
// request for a js.Cache.Delete.
func (c *CacheAPI) DeleteRequest(request shell.WebRequest, attr MatchAttr) error {
	resChn := make(chan error, 0)

	c.Call("delete", request, attr).Call("then", func(response *js.Object) {
		if response.Bool() {
			// go func(){
			close(resChn)
			// }()
			return
		}

		// go func(){
		resChn <- ErrRequestDoesNotExists
		// }()

	}).Call("catch", func(err *js.Object) {
		// go func() {
		resChn <- errors.New(err.String())
		// }()
	})

	return <-resChn
}

// cacheKeys defines the response received when all keys of the cache is retrieved
// or when filtered by a request value.
type cacheKeys struct {
	Error    error    `json:"error"`
	Response []string `json:"response"`
}

// Keys returns a slice of all cache keys for all request added in the order they
// were added.
func (c *CacheAPI) Keys(request interface{}, attr MatchAttr) ([]string, error) {
	resChn := make(chan cacheKeys, 0)

	if request == nil {
		c.Call("keys", nil, attr).Call("then", func(response *js.Object) {
			// go func() {
			resChn <- cacheKeys{
				Response: shell.ObjectToStringList(response),
			}
			// }()
		}).Call("catch", func(err *js.Object) {
			// go func() {
			resChn <- cacheKeys{
				Error: errors.New(err.String()),
			}
			// }()
		})

		opVal := <-resChn
		return opVal.Response, opVal.Error
	}

	switch ro := request.(type) {
	case string:
		c.Call("keys", ro, attr).Call("then", func(response *js.Object) {
			// go func() {
			resChn <- cacheKeys{
				Response: shell.ObjectToStringList(response),
			}
			// }()
		}).Call("catch", func(err *js.Object) {
			// go func() {
			resChn <- cacheKeys{
				Error: errors.New(err.String()),
			}
			// }()
		})
	case shell.WebRequest:
		c.Call("keys", shell.WebRequestToJSRequest(&ro), attr).Call("then", func(response *js.Object) {
			// go func() {
			resChn <- cacheKeys{
				Response: shell.ObjectToStringList(response),
			}
			// }()
		}).Call("catch", func(err *js.Object) {
			// go func() {
			resChn <- cacheKeys{
				Error: errors.New(err.String()),
			}
			// }()
		})
	case *shell.WebRequest:
		c.Call("keys", shell.WebRequestToJSRequest(ro), attr).Call("then", func(response *js.Object) {
			// go func() {
			resChn <- cacheKeys{
				Response: shell.ObjectToStringList(response),
			}
			// }()
		}).Call("catch", func(err *js.Object) {
			// go func() {
			resChn <- cacheKeys{
				Error: errors.New(err.String()),
			}
			// }()
		})
	default:
		// go func() {
		resChn <- cacheKeys{
			Error: errors.New("Request type not supported in cache"),
		}
		// }()
	}

	opVal := <-resChn
	return opVal.Response, opVal.Error
}

// All returns all the pairs of requests which have been added into the cache in
// the order they were added.
func (c *CacheAPI) All() ([]shell.WebPair, error) {
	keys, err := c.Keys(nil, MatchAttr{})
	if err != nil {
		return nil, err
	}

	var pairs []shell.WebPair

	for _, key := range keys {
		var req shell.WebRequest
		req.URL = key
		req.Method = "GET"
		req.Credentials = "same-origin"
		req.Cache = shell.DefaultStrategy

		res, err := c.MatchPath(req.URL, MatchAttr{})
		if err != nil {
			return pairs, err
		}

		pairs = append(pairs, shell.WebPair{
			Request:  req,
			Response: res,
		})
	}

	return pairs, nil
}

// Empty deletes all giving requests from the underline cache.
func (c *CacheAPI) Empty() error {
	keys, err := c.Keys(nil, MatchAttr{})
	fmt.Printf("Retrieved keys: %+q : Error: %s", keys, err)
	if err != nil {
		return err
	}

	for _, key := range keys {
		c.DeletePath(key, MatchAttr{})
	}

	return nil
}
