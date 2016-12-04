# Architecture


## TOC - Future planned aspects for GU

These are ordered in terms of a sensible roadmap.

- Bootloader
	- Why ? 
		- Because we need to be able to load external resources (css and js), but gopherjs is monolithic.
		- So we need a way for these external resources to be loaded, and for components to be able to tell the bootloader what to load at Build & Runtime.
	- Design time.
		- Each GU component decorates its class with what external resources it needs for itself. These CSS and JS file are either:
			1. Strings inside the golang code;  
			2. External files, with the golang code modeling its file name;
	- Build time.
		- Using golang code gen, the resources are embedded & a global manifest is put in root.
	- Runtime:
		- When GU boots in the browser, the first thing it does is load the external resources.
		- It will need a placeholder in the index.html i expect.

- MDL components (in seperate repo now)
	- on hold until bootloader done.
	- see "gu-mdl" repo

- Storage 
	- Will be in "gu-store" repo.  
	- Why ?
		- GU is designed to made SPA apps and they all need storage of some sort, either for reflux patterns or for long term offline storage
		- For advanced apps that need online / offline / sync suppot, you want a client side persistent Event Store that has some commonality between client and server in terms of API and Behaviour, so that you can do advanced patterns like CQRS.
		- There is low hanging fruit here because of the very pure and simple golang databases, so wh not use them on clients and servers
	- Agnostic for maximum reuse:
		- Runs on Client.
			- Runs on Browser, Mobile and Desktop.
		- Runs on Server
			- Just standard go.
		- Same API for Client and Server for maximum cognitive reuse.
			- This is vital because different topologies are a fact of life and we can insulate the topology differences from the developer easily with JSONRPC2.
			- Wrap with GRPC later, once google finish GRPC-web.
	- Core database:
		- Bleve for Full Text Search DB
		- Bolt for NoSQL & Blob DB
	- Topologies:
		- Desktop (Electron style)
			- Client Frontend 
				- Standard GU web site, with no storage.
				- Client code just calls JSON API and its automagically passed to Client Backend
			- Client Backend
				- All storage here using standard GRPC.
		- Browser (web site)
			- Bleve with indexDB adapter cross compiled.
			- Risky, but worth a try.
		- Mobile (Cordova)
			- Client Frontend has zero storage.
				- Standard GU web site, with no storage.
				- Client code just calls JSON API and its automagically passed to Client Backend
			- Client Backend has all storage.
				- Compiled across to NDK with Android & IOS JSONRPC2 pass through code.
	- Refs:
		- side using PouchDB (https://github.com/flimzy/go-pouchdb)
		- Hooks into Couchbase and others backend that support the PouchDB API



- Forms Handling
	- Everyone needs forms eventually
	- use "formulate", which is what react and everyone else uses.
	- make a "gu-forms" and drop in my code



- Reflux like Data Binding allowing changes to the Storage to be automatically reflected in the GUI. 
	- This is inspired by React & Reflux patterns.

- Internationalization
    - Because all our code is golang, it makes just using this can do everything.
	- Design / compile time
        - https://github.com/nicksnyder/go-i18n
    - Run time:
        - https://github.com/siongui/gopherjs-i18n
            - simply grabs the text as needed
	- Extraction of text in your code
	- Generation of Files for Translation.
	- Integration to the Google Translate API for automatic translation as 1st pass, and then human as 2nd pass.
		- stateful
		- transaltion memory.
	- Bunding back into your code seamlessly.
	- Reconciliation of Code to Translation Files (planned), so that when either gets out of sync you are informed
	- Numerics support (planned) for dates, number differences between locals.
		- e.g Currency in Sweden is using ".", and in England is using ","

- Desktop and Mobile Bundling 
	- Desktop using  the new  https://github.com/alexflint/gallium (work in progress)
	- Mobile using gomobile / cordova wrapper (planned)




## Refs




### bundling
- https://github.com/mh-cbon/gump
    - great way to version bump.

- see fitzy repo. he has a great way of bundling third aprty js with go.


### desktop
- https://github.com/alexflint/gallium
    - electron for golang. Really easy.
    - only for OSX, but windows coming.

- https://github.com/mh-cbon/go-github-release#packaging-for-windows
    - pack a msi for Windows
    - very easy with virtualised build too.


### Storage
Some serious options here.

Mobile
    - gomobile bleve, with a JSN pass- through. SHould work.
        - will need small amount of Android and IOS code.
    - can avoid cordova too, by using the Hyan technicque.
Desktop
    - bleve, with exact same JSON pass-through as Mobile uses. Will work.


