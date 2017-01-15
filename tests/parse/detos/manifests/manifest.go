// Package manifests is auto-generated and should not be modified by hand.
// This package contains a virtual file system for generate resources which are not accessed
// through a remote endpoint (i.e those resources generated from the manifests that are local in the
// filesystem and are not marked as remote in access).
package manifests

import (
	"path/filepath"

	"github.com/spf13/afero"
)

// AppFS defines the global handler for which all assets generated from manifests
// files which are not remote resources are provided as binary embedded assets.
var AppFS = afero.NewMemMapFs()

// addFile adds a giving file name into the file system.
func addFile(path string, content []byte){
	dir, _ := filepath.Split(path)
	if dir != "" {
		AppFS.MkdirAll(dir,0755)
	}

	afero.WriteFile(AppFS, path, content, 0644)
}

func init(){

		addFile("sandles.js", []byte("// Package detos defines a series of component for the detos project\n//\n//shell:component:global\n//\n// Resource {\n//     Name: detos.font.js\n//     Path: https://fonts.googleapis.com/css?family=Lato|Open+Sans|Roboto\n//     Hook: embedded-css\n//\t\t Localize: true\n// }\n//\npackage detos\n\nimport (\n\t\"github.com/gu-io/gu/shell/example/parse/detro\"\n)\n\n// HoldComponent defines a component to provide a hold type.\n//\n//shell:component\n/*\nResource {\n    Name: fantom.js\n    Path: http://cdl.gog.com/res/statics/fantom.js\n    Hook: javascript\n}\n\nResource {\n\t\tVersion: 1.4.1\n\t\tPkg: DentusVuz\n    Name: hul_hub.js\n    Content: ```\n        <script src=\"\">\n\t\t\t\t\tvar mo = func() Resource {\n\t\t\t\t\t\treturn ```block Resource {\n\t\t\t\t\t\t\tname: 'bock',\n\t\t\t\t\t\t}```\n\t\t\t\t\t}\n\t\t\t\t</script>\n    ```\n    Hook: embed.js\n    HookRepo: github.com/resk/compo-web/hooks\n    Size: 440030\n}\n*/\ntype HoldComponent struct {\n\tdetro.DetroComponent\n\tGulName string         `json:\"gul_name\"`\n\tBul     SolidComponent `json:\"bul\"`\n}\n\n// SolidComponent defines a component to provide a solid box type.\n//\n// shell:component\n// Resource {\n//     Name: shell-su.js\n//     Path: http://cdl.gog.com/res/statics/shell-su.js\n//     Hook: javascript\n// }\n//\n// Resource {\n//     Name: crock.js\n//     Content: ```\n//         <script src=\"\">\n//\t\t\t\t\t{\n//\t\t\t\t\t\tname: \"bob\",\n//\t\t\t\t\t\tage: 12,\n//\t\t\t\t\t}\n//\t\t\t\t\t</script>\n//     ```\n//     Hook: embed.js\n//     HookRepo: github.com/resk/compo-web/hooks\n//     Size: 440030\n// }\n//\ntype SolidComponent struct {\n\tName string `json:\"name\"`\n}\n"))

		addFile("detos.font.js", []byte("@font-face {\n  font-family: 'Lato';\n  font-style: normal;\n  font-weight: 400;\n  src: local('Lato Regular'), local('Lato-Regular'), url(https://fonts.gstatic.com/s/lato/v11/v0SdcGFAl2aezM9Vq_aFTQ.ttf) format('truetype');\n}\n@font-face {\n  font-family: 'Open Sans';\n  font-style: normal;\n  font-weight: 400;\n  src: local('Open Sans'), local('OpenSans'), url(https://fonts.gstatic.com/s/opensans/v13/cJZKeOuBrn4kERxqtaUH3aCWcynf_cDxXwCLxiixG1c.ttf) format('truetype');\n}\n@font-face {\n  font-family: 'Roboto';\n  font-style: normal;\n  font-weight: 400;\n  src: local('Roboto'), local('Roboto-Regular'), url(https://fonts.gstatic.com/s/roboto/v15/zN7GBFwfMP4uA6AR0HCoLQ.ttf) format('truetype');\n}\n"))

		addFile("hul_hub.js", []byte("        <script src=\"\">\n\t\t\t\t\tvar mo = func() Resource {\n\t\t\t\t\t\treturn ```block Resource {\n\t\t\t\t\t\t\tname: 'bock',\n\t\t\t\t\t\t}```\n\t\t\t\t\t}\n\t\t\t\t</script>\n"))

		addFile("crock.js", []byte("        <script src=\"\">\n\t\t\t\t\t{\n\t\t\t\t\t\tname: \"bob\",\n\t\t\t\t\t\tage: 12,\n\t\t\t\t\t}\n\t\t\t\t\t</script>\n"))

}

