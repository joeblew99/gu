# Getting Started
Gu is a rendering libary which embodies the very principles of simplicity and least 
complexity both in the design and approach it adopts. It has evolved through out it's 
development and adopted different ideas and patterns which have transformed it to 
what it is today.

To understand Gu, we need not go through exhaustive excercise but follow a more 
approachable method of looking at sample codes and dissecting their very parts,
because in truth they contain all their is to know about Gu from a users stand point.

## Gu:ClientSided
Using Gu on the client side is rather easy and only requires the use of a pattern that 
Goa made popular. By taking this approach it provides meaningful context to what is 
to be displayed but also organizes your code into logical blocks. 

**I like to think of the approach adopted from Goa into Gu as an ochestration of 
a page's layout**

```go

import (
	. "github.com/influx6/gu/design"
	. "github.com/influx6/gu/trees"
	. "github.com/influx6/gu/trees/elems"
)

_ = Resource(func() {

	DoTitle("Dashboard App")

	DoScript("https://code.jquery.com/jquery-3.1.1.js", "text/javascript", false)
	DoScript("https://code.jquery.com/jquery-ui.js", "text/javascript", true)
	DoCSS("https://raw.githubusercontent.com/csswizardry/csswizardry-grids/master/csswizardry-grids.scss", false)

	DoMarkup(Header1(Text("Dashboard App")), "", false)

	DoMarkup(func() *Markup {
		return Div(
			Style(
				Text("* { font-size: 1em; }"),
			),
			Section(
				Label(Text("Hello")),
			),
		)
	}, "", false)

})

func main(){
	New(&DOMRenderer{
		Document: dom.GetWindow().Document(),
	}).Init()
}
```
