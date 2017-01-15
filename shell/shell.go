package shell

import "github.com/gopherjs/gopherjs/js"

// CacheStrategy defines a type which holds different cache-strategy used for
// requests.
type CacheStrategy string

var (
	// DefaultStrategy defines the default strategy allowed by the request system.
	DefaultStrategy CacheStrategy = "default"

	// UncachedStrategy defines the strateg where resources are never cached and
	// always are retrieved from the network.
	UncachedStrategy CacheStrategy = "uncached"

	// NoCacheStrategy defines the no caching strategy allowed by the request system.
	NoCacheStrategy CacheStrategy = "no-cache"

	// ReloadCacheStrategy defines the reload strategy allowed by the request system.
	// This strategy reloads the data within the cache with the given request response.
	ReloadCacheStrategy CacheStrategy = "reload"

	// LatestSourceCacheStrategy defines the caching strategy which retrieves always
	// from the latest source updating the cache with the request. Goes from server
	// to cache.
	LatestSourceCacheStrategy CacheStrategy = "latest-source"

	// LastCacheStrategy defines the caching strategy which retrieves always
	// from the cache unless no source exists for the request then request from the
	// server then updates the cache with the request's response.
	LastCacheStrategy CacheStrategy = "last-from-cache"
)

// Fetch defines an interface which exposes http capabilities
// to easily make requests with.
type Fetch interface {
	Do(WebRequest) (WebResponse, error)
	DoManifest(ManifestAttr) (WebRequest, WebResponse, error)
	Get(string, CacheStrategy) (WebRequest, WebResponse, error)
}

// Cache defines an interface for systems which store the request and response
// regarding specific request made.
type Cache interface {
	Empty() error
	All() ([]WebPair, error)
	Delete(WebRequest) error
	Put(WebRequest, WebResponse) error
	PutPath(path string, res WebResponse) error
	GetRequest(WebRequest) (WebResponse, error)
	GetPath(path string) (WebRequest, WebResponse, error)
	GetManifest(ManifestAttr) (WebRequest, WebResponse, error)
}

// WebRequest defines the structure which holds request fields for a given
// request.
type WebRequest struct {
	BodyUsed       bool              `json:"bodyUsed"`
	CacheRequest   bool              `json:"cacheRequest"`
	Body           []byte            `json:"body"`
	Method         string            `json:"method"`
	URL            string            `json:"url"`
	Mode           string            `json:"mode"`
	ManifestName   string            `json:"manifest_name"`
	ManifestAddr   string            `json:"manifest_addr"`
	Credentials    string            `json:"credentials,omitempty"`
	Referrer       string            `json:"referrer,oimitempty"`
	ReferrerPolicy string            `json:"referrerPolicy,omitempty"`
	Underline      *js.Object        `json:"underline,omitempty"`
	Cache          CacheStrategy     `json:"cache"`
	Headers        map[string]string `json:"headers"`
}

// WebResponse defines a struct to hold a response for a giving resource.
type WebResponse struct {
	Status       int               `json:"status"`
	Redirected   bool              `json:"redirected"`
	Ok           bool              `json:"ok"`
	Body         []byte            `json:"body"`
	Type         string            `json:"type"`
	StatusText   string            `json:"status_text"`
	FinalURL     string            `json:"use_final_url,omitempty"`
	ManifestName string            `json:"manifest_name,omitempty"`
	ManifestAddr string            `json:"manifest_addr,omitempty"`
	Underline    *js.Object        `json:"underline,omitempty"`
	Headers      map[string]string `json:"headers"`
	Cookies      []string          `json:"cookies"`
}

// ManifestAttr defines a structure which stores a series of
// data pertaining to a specific resource.
type ManifestAttr struct {
	Size     int               `json:"size"`
	Remote   bool              `json:"remote"`
	Localize bool              `json:"localize"`
	ID       string            `json:"appmanifest_id,omitempty"`
	Name     string            `json:"name"`
	Path     string            `json:"path"`
	Content  string            `json:"content"`
	Cache    CacheStrategy     `json:"cache"`
	Meta     map[string]string `json:"meta"`
	HookName string            `json:"hook_name,omitempty"`
}

// WebRequest returns a WebRequest from the manifest.
func (m ManifestAttr) WebRequest() WebRequest {
	return WebRequest{
		Method:       "GET",
		URL:          m.Path,
		ManifestName: m.Name,
		Cache:        m.Cache,
		Credentials:  "same-origin",
	}
}

// AppManifest defines a structure which holds a series of
// manifests data related to specific resources.
type AppManifest struct {
	GlobalScope bool               `json:"global_scope"`
	ID          string             `json:"id"`
	Name        string             `json:"name"`
	Depends     []string           `json:"depends"`
	Manifests   []ManifestAttr     `json:"manifests"`
	Relation    *ComponentRelation `json:"relation"`
}

// WebPair defines a struct which contains the request object and response object
// received for that request.
type WebPair struct {
	Request  WebRequest
	Response WebResponse
}

// WebManifest defines a structure which connects a giving manifest and a series
// of web responses for that giving manifest.
type WebManifest struct {
	Manifest ManifestAttr
	Pair     WebPair
}

// NewAppManifest returns a instance of the AppManifest type.
func NewAppManifest(name string) *AppManifest {
	return &AppManifest{
		Name: name,
		ID:   RandString(15),
	}
}

// ManifestRequests returns a slice of WebRequest related to each ManifestAttr
// within the AppManifest.
func (am *AppManifest) ManifestRequests() []WebManifest {
	var wbs []WebManifest

	for _, mn := range am.Manifests {
		mreq := mn.WebRequest()
		mreq.ManifestAddr = am.ID

		wbs = append(wbs, WebManifest{
			Manifest: mn,
			Pair: WebPair{
				Request: mreq,
			},
		})
	}

	return wbs
}

// ComponentRelation defines a structure which stores specific
// data about the relation of a giving component regarding the
// manifests that correlate to it.
type ComponentRelation struct {
	Name       string   `json:"name"`
	Package    string   `json:"package"`
	Composites []string `json:"composites,omitempty"`
	FieldTypes []string `json:"fieldtypes,omitempty"`
}

// FindByRelation provides a convenient method to search a giving manifests for a specific
// ComponentRelation.
func FindByRelation(apps []*AppManifest, relationName string) *AppManifest {
	for _, app := range apps {
		// fmt.Printf("FindRelation: Wanted{%q} -> Tested{%q, %#v}\n", relationName, app.Relation.Name, app.Relation)

		if app.Relation.Name != relationName {
			continue
		}

		return app
	}

	return nil
}
