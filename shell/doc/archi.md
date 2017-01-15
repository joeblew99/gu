# archi - External Resources (shell)

- Why ? 
	- Because we need to be able to load external resources (css and js), but gopherjs is monolithic.
	- So we need a way for these external resources to be loaded, and for components to be able to tell the bootloader what to load at Build & Runtime.
- Uses cases
    - This all started with the MDL ( https://github.com/mh-cbon/mdl-go-components/ )
        - We want to port to GU, but requires external CSS and JS.
        - It will allow us to have a stable MDL that lots of people need
        - We could load manually put the CSS and JS in the index.html, but we want a better way.
    - Also pouchDB ( https://github.com/flimzy/go-pouchdb ).
        - needs to load an external JS file that is the standard pouchdb JS. 
        - he has 2 solutions: https://github.com/flimzy/go-pouchdb#deployment
    - Also there may be a need to load many gophers compiled golang code.
        - I doubt this is possible.
    - Other devs will make GU Components in other repos that have external resources
        - Devs want to be able to use "go get".
        - Because we designed this as part of standard go code gen tooling that will naturally work.
        - The global manifest will container all that is needed.
- Actions
	- Design time
		- Each GU component decorates its class with what external resources it needs for itself. These CSS and JS file are either:
			1. External files, with the golang code modeling its file name;
                 - for example MDL use npm to gets into js and css.
            2. DONT WANT TO DO THIS ! - CSS or JS Strings inside the golang code;  
         - Alex Version
             Manual write the manifest into root.
	- Build time.
		- Using golang code gen, the resources are found from the decorations.
        - For each resource we:
            - embedded the JS and CSS. 
                - 1. Into golang code.
                - 2. copy the resource file to a /Assets dir.
            - add to a global  manifest file (resources.json) that lives at "cmd/main.go" the path to the Resource.
        - For global build, we use "bindata" to embed the gopherjs code & the external resources as Assets into the server.
            - Can also use Browserify too maybe, but not important.
        - Alex Version
            -  Nothing to do.
	- Runtime:
		- When GU boots in the browser, the first thing it does is load the external resources.
        - It knows them from the resources manifest (which is server side) still.
            - it ask the server for the manifest.
            - it loads them up into a placeholder in the head of the index.html.
		        - It will need a placeholder in the index.html i expect.
        - Alex Version
            - Looks up manifest (on server)
            - Injects into Head.