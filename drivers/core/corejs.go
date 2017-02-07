// Package core contains javascript sources which are exported into specific
// drivers webview which exposes similar functionality needed to interoperate with
// the platform.

// Document is auto-generate and should not be modified by hand.

//go:generate go run generate.go

package core

// JavascriptDriverCore contains the javascript code to be injected into a webview
// to provide similar functionality desired to have it working with Gu.
var JavascriptDriverCore = `// Package core.js provides javascript functions which provide similar functionalities
// to allow patching provided virtual DOM and query events and dom nodes as needed.

function PatchDOM(){
  
}
`
