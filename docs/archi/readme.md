# archi


## Future planned aspects
- Storage client side using PouchDB (https://github.com/flimzy/go-pouchdb)
	- Agnostic
	- Mobile and Desktop.
	- Hooks into Couchbase and others backend that support the PouchDB API
- Reflux like Data Binding allowing changes to the Storage to be automatically reflected in the GUI. 
	- This is inspired by React & Reflux patterns.
- Internationalization 
	- Using the excellent https://github.com/nicksnyder/go-i18n
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