# Gu
  A fast UI component library for building web components

## Install

```bash
go install github.com/influx6/gu/...
```

## Features
  - Dead Simple API.
  - Rendering from both front end and backend.
  - Flux like event dispatch system.

## Goals
  - Quickly craft your UI without touching HTML
  - Rendering from either the backend or frontend
  - Share code between backend and frontend
  - Extensively simple but broad documentation

## Documentation
  - **[Introductions and Design](./docs/IntroductionsAndDesign.md)**
  - **[Components and Views](./docs/ComponentsAndViews.md)**
  - **[Components and Subcomponents](./docs/ComponentsAndSubComponents.md)**
  - **[Components and Styles](./docs/ComponentsAndStyles.md)**
  - **[Reactivity and Notifications](./docs/ReactivityAndNotifications.md)**
  - **[Initializing and Server Side Rendering](./docs/InitializationsAndServerSide.md)**
  - **[Limitations](./docs/Limitations.md)**

## Samples
  You will find samples within the [Examples](./master/examples/) directory of this project source.


## Example: Basic View
  The example below shows how the view API works and how new views can be
  created without relying on the internal guviews initializing strategy.

  - MenuList

  ```go

      package main

      import (
      	"github.com/influx6/gu/gudom"
      	"github.com/influx6/gu/gutrees"
      	"github.com/influx6/gu/gutrees/attrs"
      	"github.com/influx6/gu/gutrees/elems"
      	"github.com/influx6/gu/guviews"
      )

      // menu defines a map of name and url links.
      type menu map[string]string

      func (v menu) Render() gutrees.Markup {
      	dom := elems.Div()

      	for name, src := range v {
      		elems.Image(
      			attrs.Src(src),
      			attrs.Rel(name),
      		).Apply(dom)
      	}

      	return dom
      }

      func main() {

      	view := guviews.New(menu(map[string]string{
      		"Joyride": "https://pbs.twimg.com/media/CbtrJu9UAAAXhSs.jpg",
      		"Bombs":   "https://pbs.twimg.com/media/Cbs8Ia0VAAAYyfR.jpg",
      	}))

      	gudom.RenderAsBody(view)
      }

  ```

## Example: Deployed View
  This example depicts how views should actually be deployed especially when
  moving from backend/server rendered content to client rendered content.
  Because `gu` uses a simple approach in how views are initialized we are able
  to register different view makers that will generate the appropriate views
  while maintaining key information(eg. UUID for a view). Although this approach  requires a more proactive approach in thinking, it resolves the issue of
  reconciling rendering from server to client sided when deploying views from
  the backend with fully rendered pages.

  - MenuList

  ```go

      package main

      import (
      	"github.com/influx6/gu/gudom"
      	"github.com/influx6/gu/gutrees"
      	"github.com/influx6/gu/gutrees/attrs"
      	"github.com/influx6/gu/gutrees/elems"
      	"github.com/influx6/gu/guviews"
      )

      // menu defines a map of name and url links.
      type menu map[string]string

      func (v menu) Render() gutrees.Markup {
      	dom := elems.Div()

      	for name, src := range v {
      		elems.Image(
      			attrs.Src(src),
      			attrs.Rel(name),
      		).Apply(dom)
      	}

      	return dom
      }

      func main() {

        guviews.Register("video-components/vabbs", func(links map[string]string) guviews.Renderable {
            return menu(links)
        })


        // When deployed either on the frontend or backend, the view is able
        // to reconciled the already rendered content from the backend and
        // take over.
      	guviews.MustCreate(guviews.ViewConfig{
          Name: "video-components/vabbs",
          ID: "02-vabbs",
          Elem: "body",
          Param: map[string]string{
        		"Joyride": "https://pbs.twimg.com/media/CbtrJu9UAAAXhSs.jpg",
        		"Bombs":   "https://pbs.twimg.com/media/Cbs8Ia0VAAAYyfR.jpg",
        	},
        })

      }

  ```
