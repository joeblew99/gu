package fetch

import (
	"bytes"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gu-io/shell"
)

var (
	defaultLayout = "Mon, 01/02/06, 03:04PM"
	sameOrigin    = "same-origin"
	getMethod     = "GET"
)

// API implements an http fetching system which services
// http requests for sources using either ajax or the internal go
// network library.
type API struct {
	cache shell.Cache
}

// New returns a new instance of a API.
func New(cache shell.Cache) *API {
	return &API{
		cache: cache,
	}
}

// Get returns the request and response for the giving path using the provided
// shell.CacheStrategy.
func (f *API) Get(path string, strategy shell.CacheStrategy) (shell.WebRequest, shell.WebResponse, error) {
	if strategy == "" {
		strategy = shell.DefaultStrategy
	}

	var req shell.WebRequest
	var res shell.WebResponse

	req.URL = path
	req.Method = getMethod
	req.Credentials = sameOrigin
	req.Cache = strategy

	res, err := f.makeRequest(req)
	if err != nil {
		return req, res, err
	}

	return req, res, nil
}

// Do serves the provided request and returns the response and error
// associated against that request.
func (f *API) Do(req shell.WebRequest) (shell.WebResponse, error) {
	return f.makeRequest(req)
}

// DoManifest retrives the giving manifest from the backend and uses the attached
// cache strategy to decide whether it retrieves from network first or checks cache
// for wanted resource.
func (f *API) DoManifest(manifest shell.ManifestAttr) (shell.WebRequest, shell.WebResponse, error) {
	var res shell.WebResponse

	req := manifest.WebRequest()

	res, err := f.makeRequest(req)
	if err != nil {
		return req, res, err
	}

	return req, res, nil
}

// makeRequest encapsulates the process of making a request with a shell.WebRequest
// and returns the shell.WebResponse and error associated with the request.
func (f *API) makeRequest(req shell.WebRequest) (shell.WebResponse, error) {
	var res shell.WebResponse

	method := strings.ToUpper(req.Method)

	var buff bytes.Buffer
	buff.Write(req.Body)

	hreq, err := http.NewRequest(method, req.URL, &buff)
	if err != nil {
		return res, err
	}

	for key, val := range req.Headers {
		hreq.Header.Set(key, val)
	}

	if method == "GET" {
		switch req.Cache {
		case shell.NoCacheStrategy:
			cres, err := f.cache.GetRequest(req)
			if err != nil {
				rs, herr := http.DefaultClient.Do(hreq)
				if herr != nil {
					return res, herr
				}

				res = f.pullResponse(rs)
				if !f.successResponse(rs) {
					return res, nil
				}

				if cerr := f.cache.Put(req, res); err != nil {
					return res, cerr
				}

				return res, nil
			}

			etag := cres.Headers["E-Tag"]
			date := cres.Headers["Date"]
			lastModified := cres.Headers["Last-Modified"]

			var dateT time.Time
			var lastMod time.Time

			if date != "" {
				if mdate, perr := time.Parse(defaultLayout, date); perr == nil {
					dateT = mdate
				}
			}

			if lastModified != "" {
				if mdate, perr := time.Parse(defaultLayout, lastModified); perr == nil {
					lastMod = mdate
				}
			}

			maxAge := -1
			cacheHeaders := strings.Split(cres.Headers["Cache-Control"], ",")

			// search for the max-age value
			var maxage []string
			for _, item := range cacheHeaders {
				if strings.Contains(item, "max-age") {
					maxage = strings.Split(strings.TrimSpace(item), "=")
					break
				}
			}

			if len(maxage) > 1 {
				if mage, merr := strconv.Atoi(maxage[1]); merr == nil {
					maxAge = mage
				}
			}

			maxAgeTime := time.Unix(int64(maxAge), 0)

			if maxAge != -1 && !lastMod.IsZero() {

				// We have passed the lifetime for this response.
				if lastMod.After(maxAgeTime) {
					rs, merr := http.DefaultClient.Do(hreq)
					if err != nil {
						return res, merr
					}

					res = f.pullResponse(rs)
					if !f.successResponse(rs) {
						return res, nil
					}

					if cerr := f.cache.Put(req, res); err != nil {
						return res, cerr
					}

					return res, nil
				}

				return cres, nil
			}

			if maxAge != -1 && !dateT.IsZero() {

				// We have passed the lifetime for this response.
				if dateT.After(maxAgeTime) {
					rs, merr := http.DefaultClient.Do(hreq)
					if merr != nil {
						return res, merr
					}

					res = f.pullResponse(rs)
					if !f.successResponse(rs) {
						return res, nil
					}

					if cerr := f.cache.Put(req, res); err != nil {
						return res, cerr
					}

					return res, nil
				}

				return cres, nil
			}

			hreq.Header.Add("If-None-Match", etag)
			hreq.Header.Add("If-Modified-Since", date)

			rs, err := http.DefaultClient.Do(hreq)
			if err != nil {
				return res, err
			}

			if rs.StatusCode == 304 {
				return cres, nil
			}

			res = f.pullResponse(rs)
			if !f.successResponse(rs) {
				return res, nil
			}

			if err := f.cache.Put(req, res); err != nil {
				return res, nil
			}

			return res, nil

		// case shell.DefaultStrategy:
		// 	rs, err := http.DefaultClient.Do(hreq)
		// 	if err != nil {
		// 		return res, err
		// 	}
		//
		// 	res = f.pullResponse(rs)
		// 	if !f.successResponse(rs) {
		// 		return res, nil
		// 	}
		//
		// 	if err := f.cache.Put(req, res); err != nil {
		// 		return res, err
		// 	}
		//
		// 	return res, nil

		case shell.LatestSourceCacheStrategy:
			rs, err := http.DefaultClient.Do(hreq)
			if err != nil {
				return res, err
			}

			res = f.pullResponse(rs)

			if !f.successResponse(rs) {
				return res, nil
			}

			if err := f.cache.Put(req, res); err != nil {
				return res, err
			}

			return res, nil

		case shell.LastCacheStrategy, shell.DefaultStrategy:
			cres, err := f.cache.GetRequest(req)
			if err != nil {
				rs, err := http.DefaultClient.Do(hreq)
				if err != nil {
					return res, err
				}

				res = f.pullResponse(rs)

				if !f.successResponse(rs) {
					return res, nil
				}

				if err := f.cache.Put(req, res); err != nil {
					return res, nil
				}

				return res, nil
			}

			return cres, nil
		}
	}

	switch strings.ToUpper(req.Method) {
	case "PUT":
		rs, err := http.DefaultClient.Do(hreq)
		if err != nil {
			return res, err
		}

		res = f.pullResponse(rs)
		return res, nil

	case "POST":
		rs, err := http.DefaultClient.Do(hreq)
		if err != nil {
			return res, err
		}

		res = f.pullResponse(rs)
		return res, nil

	case "PATCH":
		rs, err := http.DefaultClient.Do(hreq)
		if err != nil {
			return res, err
		}

		res = f.pullResponse(rs)
		return res, nil

	case "DELETE":
		rs, err := http.DefaultClient.Do(hreq)
		if err != nil {
			return res, err
		}

		res = f.pullResponse(rs)
		return res, nil

	default:
		rs, err := http.DefaultClient.Do(hreq)
		if err != nil {
			return res, err
		}

		res = f.pullResponse(rs)
		return res, nil
	}
}

// successResponse returns true/false if a giving http.Response was a success.
func (f *API) successResponse(rs *http.Response) bool {
	if rs.StatusCode >= 200 && rs.StatusCode <= 299 {
		return true
	}

	return false
}

// pullResponse retrieves all response details from the giving http.Response struct.
func (f *API) pullResponse(rs *http.Response) shell.WebResponse {
	var res shell.WebResponse

	res.Type = rs.Request.Method
	res.Status = rs.StatusCode
	res.StatusText = rs.Status

	if rs.StatusCode >= 200 && rs.StatusCode <= 299 {
		res.Ok = true
	}

	if url, err := rs.Location(); err == nil {
		res.FinalURL = url.String()
	}

	for _, cookie := range rs.Cookies() {
		res.Cookies = append(res.Cookies, cookie.String())
	}

	for name, val := range rs.Header {
		res.Headers[name] = strings.Join(val, ";")
	}

	var buff bytes.Buffer
	io.Copy(&buff, rs.Body)

	res.Body = buff.Bytes()

	return res
}
