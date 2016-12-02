# Examples
It's important to note that Gu does not provide a platform means to render through 
the OS but uses the same approach where it generates a structure which produces 
the HTML to be rendered on the backend and through specific structures automatically
handles rendering on the frontend in the browser. 

There are not binaries to run, but simple html files to open.

*Please all there is to running the examples is opening the `index.html` file in 
browser except when stated otherwise. More examples showcasing rendering from 
the server will be included later.*

*Also there are no need to build examples with `go build` if you wish to rebuild 
js files then gopjerjs build would be more appropriate but that is neither necessary
as they are already built within the example directories.*

Two examples are provided here, showcasing the possibility of using Gu for creating components 
and how that can be achieved.

## [Hello](./hello/)
This example showcases how the organization structures of `Gu`, also known as `Resources`,
in helping define a page and it's content.

You can find the `index.html` code for the `Hello` app in the [Hello Directory](./hello)

## [Subscribe](./subscribe)
This example showcases how components can be built with their own markup with `Gu`
and layered within `Gu` resources, to build a application with it's corresponding pages. 
It also shows how the app can be rendered on the server with a sample http server.

You can find the `index.html` code for the `Subscribe` app in the [Subscribe Directory](./subscribe)
You can find the server code for the `Subscribe` app in the [Server Directory](./subscribe/server)


