# Subscribe
Subscribe demonstrates a basic sets of component which do interact with each other in the
way they behave.

![Image of Example](./index.png)

Because GU compiles to a SPA web app, you can just run the index.html in a browser immediatly.

````
file:///Users/apple/workspace/go/src/github.com/gu-io/gu/examples/subscribe/index.html
````


To run a Build just do:

````
gopherjs build -o assets/app.go assets/app.js
````

T run Continuous Build do:

````
http://localhost:8080/github.com/gu-io/gu/examples/subscribe/
````


## Run As Server
Simple navigate into the `./server` directory in your terminal and run as below.

```bash
go run main.go
```

Navigate your browser to `http://localhost:6060`
