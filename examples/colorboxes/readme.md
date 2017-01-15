# Color Boxes
Color Boxes demonstrates trivial component which communicate with a http backend
for their background color and change style accordingly.

![Image of Example](./boxes.gif)

Because GU compiles to a SPA web app, you can just run the index.html in a browser immediatly.

````
file:///Users/apple/workspace/go/src/github.com/gu-io/gu/examples/colorboxes/index.html
````


To run a Build just do:

````
gopherjs build -o assets/box.go assets/box.js
````

To run Continuous Build do:

````
http://localhost:8080/github.com/gu-io/gu/examples/colorboxes/
````


## Run As Server
Simple navigate into the `./server` directory in your terminal and run as below.

```bash
go run main.go
```

Navigate your browser to `http://localhost:6040`
