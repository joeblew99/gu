# Shell
Shell is an experiment on an idea to create a resource system which can be used
to generate how contents for offline loading.


## Why
We need a system which allows us to embedded resources easily with components without
the hassle of manually copying the resource files into their separate areas. More
so there will exists components which require specific resources(e.g Javascript, CSS,
Images,...etc), which will be needed for it's full functionality.

This sub-package provides us that capabilities by using meta descriptions of resources
and how those resources should be handle, we are capable of embedding resources needed
for a giving component or series of components with their inter-related dependencies
and distinct resources.

Shell provides a commandline tooling called `gu` in a self same name as the package.
The CLI tool becomes available when when `gu` has been installed with the
appropriate `go get -u` command.


## How
Shell is fully integrated into the Gu's core and need not be required to call the
package directly. To utilize the resource loading capabilities, all that is required
is the lacing of appropriate meta comments which will provide the shell command line
tool, the ability to infer the resources and their relation.

```go
// SolidComponent defines a component to provide a solid box type.
//
// shell:component
// Resource {
//     Name: apple.js
//     Path: http://cdl.gog.com/res/statics/apple.js
//     Hook: javascript
// }
//
// Resource {
//     Name: crock.js
//     Content: ```
//         <script src="">
//					{
//						name: "bob",
//						age: 12,
//					}
//					</script>
//     ```
//     Hook: embed.js
//     Size: 440030
// }
//
type SolidComponent struct {
	Name string `json:"name"`
}
```

Above is a code which demonstrate how components can be marked with giving meta
comments which then when parsed will generate for Gu, the necessary information
to either retrieve, embedded and install the resources when the giving component
is loaded into the DOM.

Once done, all that must be executed is a pre-build call to generate a `manifest.json`,
which will contain all the resources and their relative components that connect them
and for which they must be loaded for use.

```bash
> gu generate --in-dir=$GOPATH/src/github.com/package/components --out-dir=$GOPATH/src/github.com/package/components
```

This commands checks the giving path within the `--in-dir` flag which becomes the
point where the shell parser scans and generates a `manifest.json` file in the path
provided in the `--out-dir` flag.

The generated `manifest.json` deployed with the application and provided for access
on the server deployed with, has it will be the only means by which the resources
declared will be loaded at startup.
