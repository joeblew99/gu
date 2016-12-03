# archi


## toc - Future planned aspects for gi in general.
Just putting it here for now, and will move to main GU repo when it makes sense

- mdl components (in seperate repo now)

- Storage client side using PouchDB (https://github.com/flimzy/go-pouchdb)
	- Agnostic
	- Mobile and Desktop.
	- Hooks into Couchbase and others backend that support the PouchDB API

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


