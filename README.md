# Gu
  A fast UI component library for building web components

## Install

  ```bash

    go install github.com/influx6/gu/...

  ```

## Features
  - Dead Simple API.
  - Rendering from both frontend and backend.
  - Flux like event dispatch system.

## Goals
  - Quickly craft your UI without touching HTML
  - Rendering from either the backend or frontend
  - Share code between backend and frontend
  - Extensively simple but documentation

## Example

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
      		gutrees.Augment(dom, elems.Image(
      			attrs.Src(src),
      			attrs.Rel(name),
      		))
      	}

      	return dom
      }

      func main() {

      	view := guviews.View("video-vabbs", menu(map[string]string{
      		"Joyride": "https://pbs.twimg.com/media/CbtrJu9UAAAXhSs.jpg",
      		"Bombs":   "https://pbs.twimg.com/media/Cbs8Ia0VAAAYyfR.jpg",
      	}))

      	gudom.RenderAsBody(view)
      }

  ```
