// Package detos defines a series of component for the detos project
//
//shell:component:global
//
// Resource {
//     Name: globa-dui.js
//     Path: http://cdl.gog.com/res/statics/global-dui.js
//     Hook: javascript
// }
//
package detos

import (
	"github.com/gu-io/gu/shell/parse/example/detro"
)

// HoldComponent defines a component to provide a hold type.
//
//shell:component
/*
Resource {
    Name: fantom.js
    Path: http://cdl.gog.com/res/statics/fantom.js
    Hook: javascript
}

Resource {
		Version: 1.4.1
		Pkg: DentusVuz
    Name: hul_hub.js
    Content: ```
        <script src="">
					var mo = func() Resource {
						return ```block Resource {
							name: 'bock',
						}```
					}
				</script>
    ```
    Hook: embed.js
    HookRepo: github.com/resk/compo-web/hooks
    Size: 440030
}
*/
type HoldComponent struct {
	detro.DetroComponent
	GulName string         `json:"gul_name"`
	Bul     SolidComponent `json:"bul"`
}

// SolidComponent defines a component to provide a solid box type.
//
// shell:component
// Resource {
//     Name: shell-su.js
//     Path: http://cdl.gog.com/res/statics/shell-su.js
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
//     HookRepo: github.com/resk/compo-web/hooks
//     Size: 440030
// }
//
type SolidComponent struct {
	Name string `json:"name"`
}
