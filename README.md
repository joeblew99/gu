# Gu
  A component rendering library for go able to render standard HTML both on 
  the frontend and backend.

## Install

```bash
go install github.com/influx6/gu/...
```

## Goals
  - Dead Simple API.
  - Simplicity and Flexibility.
  - Able to render on both front and back end.
  - Quickly craft your UI without touching HTML
  - Share code between backend and frontend

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

## Documentation
  
  **[Getting Started](./docs/Getting Started.md)**

  **[Limitations](./docs/Limitations.md)**

## Examples
Basic examples lay within the [Examples](./examples/) directory. 

All examples have `index.html` files within them, these are the points through which 
they can be runned. Simply open up the corresponding `index.html` within the example directory 
to see the end result within your browser.


## Why Gu and What next for Gu.
Gu main purpose is to provide a very simple and elegant approach in rendering front 
end components and projects with Go. It's simplicity will be heavily guarded and
maintained. 

If more complex ideas arise, then those will be built ontop of Gu has seperate 
libraries to offer that niche. But Gu will remain as so, to allow 
anyone to graft and build on top as they see fit.

Only changes that will ever come to the library will be updates for performance, 
fixes or needed changes that simplify and provide better usage for the developer.


## Contributions
 Contributions will be received with much fan fare and gratitude.
