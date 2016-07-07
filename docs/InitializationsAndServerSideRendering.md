# Initialization And Server-Side Rendering
There is not a lot to say on these two subjects as the style and approach is really a personal choice of the developer. There exists at least special constructs include in the library to aid if needed, but the most important part is really within the views themselves.

  * You can sample code in [Deployed Example](../examples/deployed) *

## Initialization
When writing your component intialization or creation code for components, the old style of writing them with the `main` function of your client code is fine, if that code will only ever be used on the client.
But if deploying both code for client and server rendering, then a more organized initialization approach should be adopted. There exists two choices:

- Rely on the custom initialization system of `guviews`
The `guviews` library comes with a custom registeration and initalization system, which allows you to register functions that return a single or list of properly initalized Renderables, which are then wrapped in a view.

*It's really a personal choice to use this because it just lets you, have this setup within your component library*

*Note: Always seperate the library that contains your component code and your `main` function, to allow you to easily import that in both on client or server, If your code does not contain exclusive imports that can only work on the client*

```go

// some-component-package

// menu defines a map of name and url links.
type menu map[string]string

func (v menu) Render() gutrees.Markup {
	dom := elems.Div(attrs.ID("bat"))

	for name, src := range v {
		gutrees.Augment(dom, elems.Image(
			attrs.Src(src),
			attrs.Rel(name),
		))
	}

	return dom
}

func init(){
	guviews.Register("video-components/vabbs", func(links map[string]string) guviews.Renderable {
		return menu(links)
	})
}

```

```go
// On Server-Sided

import "net/http"
import "some-component-package"

func init(){
}

func IndexRoute(res http.ResponseWriter, req *http.Request){

	guviews.MustCreate(guviews.ViewConfig{
		Name: "video-components/vabbs",
		ID:   "02-vabbs",
		Elem: "body",
		Param: map[string]string{
			"Joyride": "https://pbs.twimg.com/media/CbtrJu9UAAAXhSs.jpg",
			"Bombs":   "https://pbs.twimg.com/media/Cbs8Ia0VAAAYyfR.jpg",
		},
	})

  vabbs, err := guviews.Get("02-vabbs")
  if err != nil{
    res.WriteHeader(http.StatusInternalServerError)
    res.Write([]byte(err.Error()))
    return
  }

  res.WriteHeader(http.StatusOK)
  res.Write([]byte(vabbs.RenderHTML()))
}

```

```go

// main package
import "some-component-package"

func main(){

	guviews.MustCreate(guviews.ViewConfig{
		Name: "video-components/vabbs",
		ID:   "02-vabbs",
		Elem: "body",
		Param: map[string]string{
			"Joyride": "https://pbs.twimg.com/media/CbtrJu9UAAAXhSs.jpg",
			"Bombs":   "https://pbs.twimg.com/media/Cbs8Ia0VAAAYyfR.jpg",
		},
	})

  vabbs, err := guviews.Get("02-vabbs")
}
```

As above, you can pass an argument to the function which will be called, by providing the value to the `Param` attribute. What makes this useful is you can use this on the server or client as the whole `guviews` libraries ensures to encapsulate as much DOM operation only when the DOM exists.


- Write your own initialization function
This would be more preferred for many developers and it really requires not much effort as well. All you need to do is simply write functions that create and return `Renderables` or `Views` with the needed initialization done within, and ensure this can be called either on the client or server without issues.


```go

// some-component-package

func NewPage() guviews.View {
  //....
}
```

```go
// On Client-Sided

import "some-component-package"


func main(){

  page := someComponentPackage.NewPage()
}

```

```go
// On Server-Sided

import "net/http"
import "some-component-package"

func IndexRoute(res http.ResponseWriter, req *http.Request){
  page := someComponentPackage.NewPage()

  res.WriteHeader(http.StatusOK)
  res.Write([]byte(page.RenderHTML()))
}

```


# Server-Sided Rendering
Getting server-side rendering to work is easily but coordinating and ensuring that once on the client, the rendered page and its content continue as intended is usually the hard part. `Gu` was intended to only provide efficient rendering and virtual DOM like updates and its intended design was to allow the flexibility for the developers to synchronize as they intend.

Hence `Gu` takes the simple approach, where once a page has been rendered on the server, then the views initialized on the client will take over and re-render themselves, thereby
taking ownership of the mount points within the DOM.

Of course getting this right is tricky, hence `guviews` provides the `guviews.NewWithID` constructor, which returns a view always with the giving ID, this in turn allows the update mechanism of the internal differ to identify the mount point and the already rendered component DOM tree, and re-render at that point without minimal conflicts by swapping out the content with it's own.

This does require the developer to thoroughly think about data flow and how that transitions between server and client and synchronization mechanisms, which is intended as this removes any locks on a specific pattern and allows `gu` to only care about rendering.

Now, most important, this means the developers must have a means of generating the needed set of consistent IDs for the views whether it be on the server or on the client, because this is the key that makes this whole approach possible.

```go
// some-component-package

var mainPageID = "323343AFGF&"

func NewPage() guviews.View {
  return guviews.NewWithID(mainPageID,/*...Renderable/Renderables*/)
}
```


```go
// On Server-Sided

import "net/http"
import "some-component-package"

func IndexRoute(res http.ResponseWriter, req *http.Request){
  page := someComponentPackage.NewPage()

  res.WriteHeader(http.StatusOK)
  res.Write([]byte(fmt.Sprintf(`
    <html>
      <head>
        ....
      </head>
      <body>%s</body>
    </html>
  `,page.RenderHTML())))
}

```

```go
// On Client-Sided

import "some-component-package"


func main(){
  doc := dom.GetWindow().Document()
  body := doc.QuerySelector("body")

  page := someComponentPackage.NewPage()
  page.Mount(body.Underlying())
}
```

This concludes all their needs to be known with regards to rendering on the server.
