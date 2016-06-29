# Gu
  A fast UI component library for building web components

## Install

```bash
go install github.com/influx6/gu/...
```

## Features
  - Dead Simple API.
  - Ability to create custom attribute and styles builders.
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
  - **[Views And Routes](./docs/ViewsAndRoutes.md)**
  - **[Initializing and Server Side Rendering](./docs/InitializationsAndServerSideRendering.md)**
  - **[Notifications](./docs/Notifications.md)**
  - **[Limitations](./docs/Limitations.md)**

## Samples
  You will find samples within the [Examples](./examples/) directory of this project source.


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
