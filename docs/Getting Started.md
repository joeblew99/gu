# Getting Started
Gu has evolved over time as a personal project, evolving in the ideas that make up its internals. I think having a visual overview of what it is in code helps, has the contents will help in explaining all that needs be learned when dealing with gu as a rendering layer for either the backend or frontend.

```go
_ = Resource(func(){

	Route("/home",false)
	
	Script("http://whatever.cdn.org/scripts/script.js")
	CSS("http://whatever.cnd.org/assets/main.css")
	
	Markup(func() Viewable {
	  return Div(Label("Date:"))
	})
	
	Markup(func() Viewable {
	  return Div(Label("Date:"), home.NewFrom("/users"))		
	}, ".buddy")
	
	Markup(func() Viewable {
	return Div(
		Label("Date:"), 
		buller,
		Div(
		  Label("Drecker"),
		),
	  )		
	}, ".buddy", true)
	
	View(func() Viewable {
	  return &Counter{
		Title: "The Wall Counter",
		Duration: 3000,
	  }
	})
	
	RoutedView(func() Viewable{
	   return &Counter{
			Title: "The Ball Counter",
			Duration: 3000,
	   }
	})

	View(func() Viewable {
	  return &Counter{
		Title: "The Ball Counter",
		Duration: 3000,
	  }
	},"#ball-wrapper", true)
	
})
```
