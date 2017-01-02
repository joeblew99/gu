# Parse
Parse is a simple library which parsers a set of meta values from targeted go packages provided through a absolute file path. It returns a slice of Resources structures which contains all the details collected for use.

## Install

```bash
go get -u github.com/gu-io/gu/shell/parse
```

## Run Parser with Example

```bash
> cd ./github.com/gu-io/gu/shell/parser/example
> go run main.go
```

## Format
Parse enforces a strict format for the way the meta resource values must be written. It is very basic and uses newlines as separation between fields and their values.
All `Parse` key and value are seperated by newlines and each is defined as demonstrated below.

```coffee
Resource {
    Name: apple.js
    Path: http://cdl.gog.com/res/statics/apple.js
    Hook: javascript
}

Resource {
    Name: jungle.css
    Hook: css.remote
    Path: assets/statics/styles/jungle.css
    Remote: true
}
```


When defining resources with textual data or of larger text of content for specific fields then it's required to begin the enclosure of such text with a (\`\`\`) after the field name and ensure to end such text with a (\`\`\`) indented at both sides. Each must have a newline after its occurrence.

Such as demonstrated below.

```coffee
Resource {
    Name: crock.js
    Content: ```
        <script src="">
					{
						name: "bob",
						age: 12,
					}
					</script>
    ```
    Hook: embed.js
    Size: 440030
}
```

Parse has certain standard names(field keys) that hold special meaning and purpose, which should be used when describing resources fields.

These field keys are as follows:

- Name: This field defines the name of the resource which should be used in identifying the resource. This is a non-optional field and must exists.

- Path: This is an optional field which indicates the external url for the resource.

- Content: This is an optional field which contains the content which must be used for requests for this resource.

- Size: This is an optional field to indicate the size to specify for the resource content.

- Remote: This is an optional field to indicate the resource is remote and should be skipped when attempting
to locate resource data from the local file system. Has usefulness when using relative paths but wish to have the resource served from the server which serves up the manifests.

- Hook: This is an optional field to indicate the specific hook for processing this resource.


## Format in Go Sources
Within the Go source file, Parse requires there exist a marker to set a certain comments set after it as resource declarations.

```coffee
shell:component
```

This precondition which is required should generally exists after the documentation for that type and also should be preceded by a empty line.

Demonstrated as below:

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

## Parse Output
Parse produces a slice of all resource parsed out from running through it's target path and produces structures which contains details of each specific resource meta.

**Generated from the [Example Package](./example)**

```json
[
	{
		"Line": 0,
		"Global": true,
		"File": "/home/influx6/Labs/repos/go/src/github.com/gu-io/gu/shell/parse/example/detos/doc.go",
		"ComponentName": "detos",
		"PackageName": "github.com/gu-io/gu/shell/parse/example/detos",
		"CompositeTypeNames": null,
		"FieldTypeNames": null,
		"Resources": [
			{
				"Name": "sandles.js",
				"Path": "http://cdl.gog.com/res/statics/sandles.js",
				"HookPkg": "",
				"HookName": "javascript",
				"ContentSize": "",
				"Data": "",
				"Meta": {}
			}
		]
	},
	{
		"Line": 0,
		"Global": false,
		"File": "/home/influx6/Labs/repos/go/src/github.com/gu-io/gu/shell/parse/example/detro/detro.go",
		"ComponentName": "DetroComponent",
		"PackageName": "github.com/gu-io/gu/shell/parse/example/detro",
		"CompositeTypeNames": null,
		"FieldTypeNames": null,
		"Resources": [
			{
				"Name": "bendrox.js",
				"Path": "http://cdl.gog.com/res/statics/bendrox.js",
				"HookPkg": "",
				"HookName": "javascript",
				"ContentSize": "",
				"Data": "",
				"Meta": {}
			},
			{
				"Name": "caldox.js",
				"Path": "",
				"HookPkg": "github.com/resk/compo-web/hooks",
				"HookName": "embed.js",
				"ContentSize": "440030",
				"Data": "        \u003cscript src=\"\"\u003e\u003c/script\u003e\n",
				"Meta": {}
			}
		]
	},
	{
		"Line": 0,
		"Global": false,
		"File": "/home/influx6/Labs/repos/go/src/github.com/gu-io/gu/shell/parse/example/detro/detro.go",
		"ComponentName": "Scene",
		"PackageName": "github.com/gu-io/gu/shell/parse/example/detro",
		"CompositeTypeNames": null,
		"FieldTypeNames": null,
		"Resources": null
	},
	{
		"Line": 0,
		"Global": true,
		"File": "/home/influx6/Labs/repos/go/src/github.com/gu-io/gu/shell/parse/example/detos/detos.go",
		"ComponentName": "detos",
		"PackageName": "github.com/gu-io/gu/shell/parse/example/detos",
		"CompositeTypeNames": null,
		"FieldTypeNames": null,
		"Resources": [
			{
				"Name": "globa-dui.js",
				"Path": "http://cdl.gog.com/res/statics/global-dui.js",
				"HookPkg": "",
				"HookName": "javascript",
				"ContentSize": "",
				"Data": "",
				"Meta": {}
			}
		]
	},
	{
		"Line": 0,
		"Global": false,
		"File": "/home/influx6/Labs/repos/go/src/github.com/gu-io/gu/shell/parse/example/detos/detos.go",
		"ComponentName": "HoldComponent",
		"PackageName": "github.com/gu-io/gu/shell/parse/example/detos",
		"CompositeTypeNames": [
			"DetroComponent"
		],
		"FieldTypeNames": [
			"SolidComponent"
		],
		"Resources": [
			{
				"Name": "fantom.js",
				"Path": "http://cdl.gog.com/res/statics/fantom.js",
				"HookPkg": "",
				"HookName": "javascript",
				"ContentSize": "",
				"Data": "",
				"Meta": {}
			},
			{
				"Name": "hul_hub.js",
				"Path": "",
				"HookPkg": "github.com/resk/compo-web/hooks",
				"HookName": "embed.js",
				"ContentSize": "440030",
				"Data": "        \u003cscript src=\"\"\u003e\n\t\t\t\t\tvar mo = func() Resource {\n\t\t\t\t\t\treturn ```block Resource {\n\t\t\t\t\t\t\tname: 'bock',\n\t\t\t\t\t\t}```\n\t\t\t\t\t}\n\t\t\t\t\u003c/script\u003e\n",
				"Meta": {
					"Pkg": "DentusVuz",
					"Version": "1.4.1"
				}
			}
		]
	},
	{
		"Line": 0,
		"Global": false,
		"File": "/home/influx6/Labs/repos/go/src/github.com/gu-io/gu/shell/parse/example/detos/detos.go",
		"ComponentName": "SolidComponent",
		"PackageName": "github.com/gu-io/gu/shell/parse/example/detos",
		"CompositeTypeNames": null,
		"FieldTypeNames": null,
		"Resources": [
			{
				"Name": "shell-su.js",
				"Path": "http://cdl.gog.com/res/statics/shell-su.js",
				"HookPkg": "",
				"HookName": "javascript",
				"ContentSize": "",
				"Data": "",
				"Meta": {}
			},
			{
				"Name": "crock.js",
				"Path": "",
				"HookPkg": "github.com/resk/compo-web/hooks",
				"HookName": "embed.js",
				"ContentSize": "440030",
				"Data": "        \u003cscript src=\"\"\u003e\n\t\t\t\t\t{\n\t\t\t\t\t\tname: \"bob\",\n\t\t\t\t\t\tage: 12,\n\t\t\t\t\t}\n\t\t\t\t\t\u003c/script\u003e\n",
				"Meta": {}
			}
		]
	}
]
```
