// Package elems contains definition for different html element types and some custom node types

//go:generate go run generate.go

// Documentation source: "HTML element reference" by Mozilla Contributors, https://developer.mozilla.org/en-US/docs/Web/HTML/Element, licensed under CC-BY-SA 2.5.

// Documentation for custom types lies within the  "github.com/influx6/faux/domtrees" package

package elems

import (
	"github.com/influx6/gu/trees"
	"github.com/influx6/gu/css"
)

// Text provides custom type for defining text nodes with the trees markup.
func Text(content string) *trees.Markup {
	return trees.NewText(content)
}

// Parse returns the giving markup structure generated from the string.
func Parse(markup string) *trees.Markup {
	return trees.ParseAsRoot("section", markup)
}

// ParseIn returns the giving markup structure generated from the string.
func ParseIn(root string,markup string) *trees.Markup {
	return trees.ParseAsRoot(root, markup)
}

// CSS provides a function that takes style rules which returns a stylesheet embeded into 
// the provided element parent and is built on the gu/css package which collects 
// necessary details from its parent to only target where it gets mounted.
func CSS(styles interface{}, bind interface{}) *trees.Markup {
  var rs *css.Rule 

  switch so := styles.(type) {
    case string:
      rs = css.New(so)
    case *css.Rule:
      rs = so
    default:
      panic("Invalid Acceptable type for css: Only string or *css.Rule")
  }

	return trees.CSSStylesheet(rs, bind)
}

// SvgAnchor provides the following for SVG XML elements ->
// The <a> SVG element defines a hyperlink.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/a
func SvgAnchor(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("a",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgAltGlyph provides the following for SVG XML elements ->
// The <altGlyph> SVG element allows sophisticated selection of the glyphs used to render its child character data.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/altGlyph
func SvgAltGlyph(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("altGlyph",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgAltGlyphDef provides the following for SVG XML elements ->
// The <altGlyphDef> SVG element defines a substitution representation for glyphs.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/altGlyphDef
func SvgAltGlyphDef(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("altGlyphDef",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgAltGlyphItem provides the following for SVG XML elements ->
// The <altGlyphItem> element provides a set of candidates for glyph substitution by the <altGlyph> element.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/altGlyphItem
func SvgAltGlyphItem(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("altGlyphItem",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgAnimate provides the following for SVG XML elements ->
// The animate element is put inside a shape element and defines how an attribute of an element changes over the animation. The attribute will change from the initial value to the end value in the duration specified.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/animate
func SvgAnimate(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("animate",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgAnimateColor provides the following for SVG XML elements ->
// The <animateColor> SVG element specifies a color transformation over time.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/animateColor
func SvgAnimateColor(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("animateColor",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgAnimateMotion provides the following for SVG XML elements ->
// The <animateMotion> element causes a referenced element to move along a motion path.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/animateMotion
func SvgAnimateMotion(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("animateMotion",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgAnimateTransform provides the following for SVG XML elements ->
// The animateTransform element animates a transformation attribute on a target element, thereby allowing animations to control translation, scaling, rotation and/or skewing.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/animateTransform
func SvgAnimateTransform(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("animateTransform",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgCircle provides the following for SVG XML elements ->
// The <circle> SVG element is an SVG basic shape, used to create circles based on a center point and a radius.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/circle
func SvgCircle(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("circle",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgClipPath provides the following for SVG XML elements ->
// The <clipPath> SVG element defines a clipping path. A clipping path is used/referenced using the clip-path property.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/clipPath
func SvgClipPath(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("clipPath",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgColorProfile provides the following for SVG XML elements ->
// The <color-profile> element allows describing the color profile used for the image.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/color-profile
func SvgColorProfile(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("color-profile",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgCursor provides the following for SVG XML elements ->
// The <cursor> SVG element can be used to define a platform-independent custom cursor. A recommended approach for defining a platform-independent custom cursor is to create a PNG image and define a cursor element that references the PNG image and identifies the exact position within the image which is the pointer position (i.e., the hot spot).
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/cursor
func SvgCursor(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("cursor",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgDefs provides the following for SVG XML elements ->
// SVG allows graphical objects to be defined for later reuse. It is recommended that, wherever possible, referenced elements be defined inside of a <defs> element. Defining these elements inside of a <defs> element promotes understandability of the SVG content and thus promotes accessibility. Graphical elements defined in a <defs> element will not be directly rendered. You can use a <use> element to render those elements wherever you want on the viewport.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/defs
func SvgDefs(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("defs",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgDesc provides the following for SVG XML elements ->
// Each container element or graphics element in an SVG drawing can supply a description string using the <desc> element where the description is text-only.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/desc
func SvgDesc(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("desc",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgDiscard provides the following for SVG XML elements ->
// The <discard> SVG element allows authors to specify the time at which particular elements are to be discarded, thereby reducing the resources required by an SVG user agent. This is particularly useful to help SVG viewers conserve memory while displaying long-running documents.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/discard
func SvgDiscard(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("discard",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgEllipse provides the following for SVG XML elements ->
// The ellipse element is an SVG basic shape, used to create ellipses based on a center coordinate, and both their x and y radius.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/ellipse
func SvgEllipse(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("ellipse",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgFeBlend provides the following for SVG XML elements ->
// The <feBlend> SVG filter primitive composes two objects together ruled by a certain blending mode. This is similar to what is known from image editing software when blending two layers. The mode is defined by the mode attribute.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feBlend
func SvgFeBlend(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("feBlend",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgFeColorMatrix provides the following for SVG XML elements ->
// The <feColorMatrix> SVG filter element changes colors based on a transformation matrix. Every pixel's color value (represented by an [R,G,B,A] vector) is matrix multiplied to create a new color.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feColorMatrix
func SvgFeColorMatrix(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("feColorMatrix",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgFeComponentTransfer provides the following for SVG XML elements ->
// Th <feComponentTransfer> SVG filter primitive performs color-component-wise remapping of data for each pixel. It allows operations like brightness adjustment, contrast adjustment, color balance or thresholding.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feComponentTransfer
func SvgFeComponentTransfer(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("feComponentTransfer",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgFeComposite provides the following for SVG XML elements ->
// This filter primitive performs the combination of two input images pixel-wise in image space using one of the Porter-Duff compositing operations: over, in, atop, out, xor and lighter. Additionally, a component-wise arithmetic operation (with the result clamped between [0..1]) can be applied.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feComposite
func SvgFeComposite(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("feComposite",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgFeConvolveMatrix provides the following for SVG XML elements ->
// The <feConvolveMatrix> SVG filter primitive applies a matrix convolution filter effect. A convolution combines pixels in the input image with neighboring pixels to produce a resulting image. A wide variety of imaging operations can be achieved through convolutions, including blurring, edge detection, sharpening, embossing and beveling.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feConvolveMatrix
func SvgFeConvolveMatrix(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("feConvolveMatrix",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgFeDiffuseLighting provides the following for SVG XML elements ->
// The <feDiffuseLighting> SVG filter primitive lights an image using the alpha channel as a bump map. The resulting image, which is an RGBA opaque image, depends on the light color, light position and surface geometry of the input bump map.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feDiffuseLighting
func SvgFeDiffuseLighting(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("feDiffuseLighting",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgFeDisplacementMap provides the following for SVG XML elements ->
// The <feDisplacementMap> SVG filter primitive uses the pixel values from the image from in2 to spatially displace the image from in.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feDisplacementMap
func SvgFeDisplacementMap(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("feDisplacementMap",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgFeDistantLight provides the following for SVG XML elements ->
// The <feDistantLight> filter primitive defines a distant light source that can be used within a lighting filter primitive: <feDiffuseLighting> or <feSpecularLighting>.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feDistantLight
func SvgFeDistantLight(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("feDistantLight",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgFeDropShadow provides the following for SVG XML elements ->
// The <feDropShadow> filter primitive creates a drop shadow of the input image. It is a shorthand filter, and is defined in terms of combinations of other filter primitives.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feDropShadow
func SvgFeDropShadow(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("feDropShadow",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgFeFlood provides the following for SVG XML elements ->
// The <feFlood> SVG filter primitive fills the filter subregion with the color and opacity defined by flood-color and flood-opacity.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feFlood
func SvgFeFlood(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("feFlood",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgFeFuncA provides the following for SVG XML elements ->
// The <feFuncA> SVG filter primitive defines the transfer function for the alpha component of the input graphic of its parent <feComponentTransfer> element.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feFuncA
func SvgFeFuncA(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("feFuncA",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgFeFuncB provides the following for SVG XML elements ->
// The <feFuncB> SVG filter primitive defines the transfer function for the blue component of the input graphic of its parent <feComponentTransfer> element.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feFuncB
func SvgFeFuncB(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("feFuncB",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgFeFuncG provides the following for SVG XML elements ->
// The <feFuncG> SVG filter primitive defines the transfer function for the green component of the input graphic of its parent <feComponentTransfer> element.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feFuncG
func SvgFeFuncG(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("feFuncG",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgFeFuncR provides the following for SVG XML elements ->
// The <feFuncR> SVG filter primitive defines the transfer function for the red component of the input graphic of its parent <feComponentTransfer> element.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feFuncR
func SvgFeFuncR(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("feFuncR",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgFeGaussianBlur provides the following for SVG XML elements ->
// The <feGaussianBlur> SVG filter primitive blurs the input image by the amount specified in stdDeviation, which defines the bell-curve.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feGaussianBlur
func SvgFeGaussianBlur(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("feGaussianBlur",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgFeImage provides the following for SVG XML elements ->
// The <feImage> SVG filter primitive fetches image data from an external source and provides the pixel data as output (meaning if the external source is an SVG image, it is rasterized.)
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feImage
func SvgFeImage(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("feImage",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgFeMerge provides the following for SVG XML elements ->
// The <feMerge> SVG element allows filter effects to be applied concurrently instead of sequentially. This is achieved by other filters storing their output via the result attribute and then accessing it in a <feMergeNode> child.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feMerge
func SvgFeMerge(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("feMerge",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgFeMergeNode provides the following for SVG XML elements ->
// The feMergeNode takes the result of another filter to be processed by its parent <feMerge>.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feMergeNode
func SvgFeMergeNode(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("feMergeNode",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgFeMorphology provides the following for SVG XML elements ->
// The <feMorphology> SVG filter primitive is used to erode or dilate the input image. It's usefulness lies especially in fattening or thinning effects.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feMorphology
func SvgFeMorphology(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("feMorphology",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgFeOffset provides the following for SVG XML elements ->
// The <feOffset> SVG filter primitive allows to offset the input image. The input image as a whole is offset by the values specified in the dx and dy attributes.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feOffset
func SvgFeOffset(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("feOffset",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgFePointLight provides the following for SVG XML elements ->
// The  SVG filter primitive allows to create a point light effect.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/fePointLight
func SvgFePointLight(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("fePointLight",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgFeSpecularLighting provides the following for SVG XML elements ->
// The <feSpecularLighting> SVG filter primitive lights a source graphic using the alpha channel as a bump map. The resulting image is an RGBA image based on the light color. The lighting calculation follows the standard specular component of the Phong lighting model. The resulting image depends on the light color, light position and surface geometry of the input bump map. The result of the lighting calculation is added. The filter primitive assumes that the viewer is at infinity in the z direction.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feSpecularLighting
func SvgFeSpecularLighting(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("feSpecularLighting",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgFeSpotLight provides the following for SVG XML elements ->
// The <feSpotLight> SVG filter primitive allows to create a spotlight effect.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feSpotLight
func SvgFeSpotLight(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("feSpotLight",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgFeTile provides the following for SVG XML elements ->
// The <feTile> SVG filter primitive allows to fill a target rectangle with a repeated, tiled pattern of an input image. The effect is similar to the one of a <pattern>.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feTile
func SvgFeTile(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("feTile",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgFeTurbulence provides the following for SVG XML elements ->
// The <feTurbulence> SVG filter primitive creates an image using the Perlin turbulence function. It allows the synthesis of artificial textures like clouds or marble. The resulting image will fill the entire filter primitive subregion.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feTurbulence
func SvgFeTurbulence(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("feTurbulence",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgFilter provides the following for SVG XML elements ->
// The <filter> SVG element serves as container for atomic filter operations. It is never rendered directly. A filter is referenced by using the filter attribute on the target SVG element or via the filter CSS property.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/filter
func SvgFilter(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("filter",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgFont provides the following for SVG XML elements ->
// The <font> SVG element defines a font to be used for text layout.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/font
func SvgFont(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("font",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgFontFace provides the following for SVG XML elements ->
// The <font-face> SVG element corresponds to the CSS @font-face rule. It defines a font's outer properties.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/font-face
func SvgFontFace(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("font-face",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgFontFaceFormat provides the following for SVG XML elements ->
// The <font-face-format> SVG element describes the type of font referenced by its parent <font-face-uri>.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/font-face-format
func SvgFontFaceFormat(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("font-face-format",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgFontfaceName provides the following for SVG XML elements ->
// The <font-face-name> element points to a locally installed copy of this font, identified by its name.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/font-face-name
func SvgFontfaceName(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("font-face-name",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgFontFaceSrc provides the following for SVG XML elements ->
// The <font-face-src> SVG element corresponds to the src descriptor in CSS @font-face rules. It serves as container for <font-face-name>, pointing to locally installed copies of this font, and <font-face-uri>, utilizing remotely defined fonts.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/font-face-src
func SvgFontFaceSrc(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("font-face-src",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgFontfaceURI provides the following for SVG XML elements ->
// The <font-face-uri> SVG element points to a remote definition of the current font.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/font-face-uri
func SvgFontfaceURI(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("font-face-uri",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgForeignObject provides the following for SVG XML elements ->
// The <foreignObject> SVG element allows for inclusion of a foreign XML namespace which has its graphical content drawn by a different user agent. The included foreign graphical content is subject to SVG transformations and compositing.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/foreignObject
func SvgForeignObject(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("foreignObject",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgGroup provides the following for SVG XML elements ->
// The <g> SVG element is a container used to group other SVG elements.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/g
func SvgGroup(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("g",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgGlyph provides the following for SVG XML elements ->
// A <glyph> defines a single glyph in an SVG font.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/glyph
func SvgGlyph(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("glyph",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgGlyphRef provides the following for SVG XML elements ->
// The glyphRef element provides a single possible glyph to the referencing <altGlyph> substitution.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/glyphRef
func SvgGlyphRef(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("glyphRef",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgHatch provides the following for SVG XML elements ->
// The <hatch> SVG element is used to fill or stroke an object using one or more pre-defined paths that are repeated at fixed intervals in a specified direction to cover the areas to be painted.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/hatch
func SvgHatch(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("hatch",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgHatchpath provides the following for SVG XML elements ->
// The <hatchpath> SVG element defines a hatch path used by the <hatch> element.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/hatchpath
func SvgHatchpath(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("hatchpath",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgHkern provides the following for SVG XML elements ->
// The <hkern> SVG element allows to fine-tweak the horizontal distance between two glyphs. This process is known as kerning.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/hkern
func SvgHkern(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("hkern",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgImage provides the following for SVG XML elements ->
// The <image> SVG element allows a raster image to be included in an SVG document.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/image
func SvgImage(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("image",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgLine provides the following for SVG XML elements ->
// The <line> element is an SVG basic shape used to create a line connecting two points.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/line
func SvgLine(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("line",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgLinearGradient provides the following for SVG XML elements ->
// The <linearGradient> SVG element lets authors define linear gradients to fill or stroke graphical elements.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/linearGradient
func SvgLinearGradient(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("linearGradient",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgMarker provides the following for SVG XML elements ->
// The <marker> element defines the graphics that is to be used for drawing arrowheads or polymarkers on a given <path>, <line>, <polyline> or <polygon> element.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/marker
func SvgMarker(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("marker",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgMask provides the following for SVG XML elements ->
// In SVG, you can specify that any other graphics object or <g> element can be used as an alpha mask for compositing the current object into the background. A mask is defined with the <mask> element. A mask is used/referenced using the mask property.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/mask
func SvgMask(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("mask",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgMesh provides the following for SVG XML elements ->
// The documentation about this has not yet been written; please consider contributing!
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/mesh
func SvgMesh(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("mesh",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgMeshgradient provides the following for SVG XML elements ->
// The documentation about this has not yet been written; please consider contributing!
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/meshgradient
func SvgMeshgradient(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("meshgradient",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgMeshpatch provides the following for SVG XML elements ->
// The documentation about this has not yet been written; please consider contributing!
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/meshpatch
func SvgMeshpatch(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("meshpatch",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgMeshrow provides the following for SVG XML elements ->
// The documentation about this has not yet been written; please consider contributing!
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/meshrow
func SvgMeshrow(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("meshrow",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgMetadata provides the following for SVG XML elements ->
// The <metadata> SVG element allows to add metadata to SVG content. Metadata is structured information about data. The contents of <metadata> elements should be elements from other XML namespaces such as RDF, FOAF, etc.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/metadata
func SvgMetadata(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("metadata",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgMissingGlyph provides the following for SVG XML elements ->
// The <missing-glyph> SVG element's content is rendered, if for a given character the font doesn't define an appropriate <glyph>.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/missing-glyph
func SvgMissingGlyph(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("missing-glyph",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgMpath provides the following for SVG XML elements ->
// The <mpath> sub-element for the <animateMotion> element provides the ability to reference an external <path> element as the definition of a motion path.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/mpath
func SvgMpath(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("mpath",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgPath provides the following for SVG XML elements ->
// The <path> SVG element is the generic element to define a shape. All the basic shapes can be created with a path element.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/path
func SvgPath(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("path",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgPattern provides the following for SVG XML elements ->
// The <pattern> element defines a graphics object which can be redrawn at repeated x and y-coordinate intervals ("tiled") to cover an area.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/pattern
func SvgPattern(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("pattern",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgPolygon provides the following for SVG XML elements ->
// The <polygon> element defines a closed shape consisting of a set of connected straight line segments. The last point is connected to the first point. For open shapes see the <polyline> element.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/polygon
func SvgPolygon(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("polygon",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgPolyline provides the following for SVG XML elements ->
// The <polyline> SVG element is an SVG basic shape that creates straight lines connecting several points. Typically a polyline is used to create open shapes as the last point doesn't have to be connected to the first point. For closed shapes see the <polygon> element.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/polyline
func SvgPolyline(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("polyline",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgRadialGradient provides the following for SVG XML elements ->
// The <radialGradient> SVG element lets authors define radial gradients to fill or stroke graphical elements.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/radialGradient
func SvgRadialGradient(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("radialGradient",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgRect provides the following for SVG XML elements ->
// The rect element is an SVG basic shape, used to create rectangles based on the position of a corner and their width and height. It may also be used to create rectangles with rounded corners.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/rect
func SvgRect(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("rect",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgScript provides the following for SVG XML elements ->
// A SVG script element is equivalent to the script element in HTML and thus is the place for scripts (e.g., ECMAScript).
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/script
func SvgScript(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("script",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgSet provides the following for SVG XML elements ->
// The <set> element provides a simple means of just setting the value of an attribute for a specified duration. It supports all attribute types, including those that cannot reasonably be interpolated, such as string and boolean values. The <set> element is non-additive. The additive and accumulate attributes are not allowed, and will be ignored if specified.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/set
func SvgSet(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("set",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgSolidcolor provides the following for SVG XML elements ->
// The documentation about this has not yet been written; please consider contributing!
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/solidcolor
func SvgSolidcolor(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("solidcolor",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgStop provides the following for SVG XML elements ->
// The <stop> SVG element defines the ramp of colors to use on a gradient, which is a child element to either the <linearGradient> or the <radialGradient> element.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/stop
func SvgStop(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("stop",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgStyle provides the following for SVG XML elements ->
// The <style> SVG element allows style sheets to be embedded directly within SVG content. SVG's style element has the same attributes as the corresponding element in HTML (see HTML's <style> element).
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/style
func SvgStyle(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("style",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Svg provides the following for SVG XML elements ->
// The svg element can be used to embed an SVG fragment inside the current document (for example, an HTML document). This SVG fragment has its own viewport and coordinate system.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/svg
func Svg(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("svg",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgSwitch provides the following for SVG XML elements ->
// The <switch> SVG element evaluates the requiredFeatures, requiredExtensions and systemLanguage attributes on its direct child elements in order, and then processes and renders the first child for which these attributes evaluate to true. All others will be bypassed and therefore not rendered. If the child element is a container element such as a <g>, then the entire subtree is either processed/rendered or bypassed/not rendered.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/switch
func SvgSwitch(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("switch",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgSymbol provides the following for SVG XML elements ->
// The <symbol> element is used to define graphical template objects which can be instantiated by a <use> element. The use of symbol elements for graphics that are used multiple times in the same document adds structure and semantics. Documents that are rich in structure may be rendered graphically, as speech, or as Braille, and thus promote accessibility. Note that a symbol element itself is not rendered. Only instances of a symbol element (i.e., a reference to a symbol by a <use> element) are rendered.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/symbol
func SvgSymbol(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("symbol",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgText provides the following for SVG XML elements ->
// The SVG <text> element defines a graphics element consisting of text. It's possible to apply a gradient, pattern, clipping path, mask, or filter to <text>, just like any other SVG graphics element.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/text
func SvgText(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("text",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgTextPath provides the following for SVG XML elements ->
// In addition to text drawn in a straight line, SVG also includes the ability to place text along the shape of a <path> element. To specify that a block of text is to be rendered along the shape of a <path>, include the given text within a <textPath> element which includes an href attribute with a reference to a <path> element.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/textPath
func SvgTextPath(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("textPath",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgTitle provides the following for SVG XML elements ->
// Each container element or graphics element in an SVG drawing can supply a <title> element containing a description string where the description is text-only. When the current SVG document fragment is rendered as SVG on visual media, <title> element is not rendered as part of the graphics. However, some user agents may, for example, display the <title> element as a tooltip. Alternate presentations are possible, both visual and aural, which display the <title> element but do not display path elements or other graphics elements. The <title> element generally improve accessibility of SVG documents
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/title
func SvgTitle(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("title",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgTref provides the following for SVG XML elements ->
// The textual content for a <text> SVG element can be either character data directly embedded within the <text> element or the character data content of a referenced element, where the referencing is specified with a <tref> element.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/tref
func SvgTref(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("tref",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgTspan provides the following for SVG XML elements ->
// Within a <text> element, text and font properties and the current text position can be adjusted with absolute or relative coordinate values by including a <tspan> element.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/tspan
func SvgTspan(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("tspan",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgUnknown provides the following for SVG XML elements ->
// The documentation about this has not yet been written; please consider contributing!
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/unknown
func SvgUnknown(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("unknown",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgUse provides the following for SVG XML elements ->
// The <use> element takes nodes from within the SVG document, and duplicates them somewhere else. The effect is the same as if the nodes were deeply cloned into a non-exposed DOM, and then pasted where the use element is, much like cloned template elements in HTML5. Since the cloned nodes are not exposed, care must be taken when using CSS to style a use element and its hidden descendants. CSS attributes are not guaranteed to be inherited by the hidden, cloned DOM unless you explicitly request it using CSS inheritance.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/use
func SvgUse(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("use",true)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgView provides the following for SVG XML elements ->
// A view is a defined way to view the image, like a zoom level or a detail view.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/view
func SvgView(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("view",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// SvgVkern provides the following for SVG XML elements ->
// The <vkern> SVG element allows to fine-tweak the vertical distance between two glyphs in top-to-bottom fonts. This process is known as kerning.
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/vkern
func SvgVkern(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("vkern",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Anchor provides the following for HTML elements ->
// The HTML Anchor Element (<a>) defines a hyperlink to a location on the same page or any other page on the Web. It can also be used (in an obsolete way) to create an anchor point—a destination for hyperlinks within the content of a page, so that links aren't limited to connecting simply to the top of a page.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/a
func Anchor(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("a",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Abbreviation provides the following for HTML elements ->
// The HTML <abbr> element (or HTML Abbreviation Element) represents an abbreviation and optionally provides a full description for it. If present, the title attribute must contain this full description and nothing else.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/abbr
func Abbreviation(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("abbr",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Address provides the following for HTML elements ->
// The HTML <address> element supplies contact information for its nearest <article> or <body> ancestor; in the latter case, it applies to the whole document.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/address
func Address(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("address",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Area provides the following for HTML elements ->
// The HTML <area> element defines a hot-spot region on an image, and optionally associates it with a hypertext link. This element is used only within a <map> element.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/area
func Area(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("area",true)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Article provides the following for HTML elements ->
// The HTML <article> element represents a self-contained composition in a document, page, application, or site, which is intended to be independently distributable or reusable (e.g., in syndication). This could be a forum post, a magazine or newspaper article, a blog entry, an object, or any other independent item of content. Each <article> should be identified, typically by including a heading (<h1>-<h6> element) as a child of the <article> element.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/article
func Article(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("article",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Aside provides the following for HTML elements ->
// The HTML <aside> element represents a section of the page with content connected tangentially to the rest, which could be considered separate from that content. These sections are often represented as sidebars or inserts. They often contain the definitions on the sidebars, such as definitions from the glossary; there may also be other types of information, such as related advertisements; the biography of the author; web applications; profile information or related links on the blog.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/aside
func Aside(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("aside",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Audio provides the following for HTML elements ->
// The HTML <audio> element is used to embed sound content in documents. It may contain one or more audio sources, represented using the src attribute or the <source> element; the browser will choose the most suitable one.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/audio
func Audio(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("audio",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Bold provides the following for HTML elements ->
// The HTML <b> Element represents a span of text stylistically different from normal text, without conveying any special importance or relevance. It is typically used for keywords in a summary, product names in a review, or other spans of text whose typical presentation would be boldfaced. Another example of its use is to mark the lead sentence of each paragraph of an article.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/b
func Bold(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("b",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Base provides the following for HTML elements ->
// The HTML <base> element specifies the base URL to use for all relative URLs contained within a document. There can be only one <base> element in a document.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/base
func Base(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("base",true)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// BidirectionalIsolation provides the following for HTML elements ->
// The HTML <bdi> Element (or Bi-Directional Isolation Element) isolates a span of text that might be formatted in a different direction from other text outside it.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/bdi
func BidirectionalIsolation(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("bdi",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// BidirectionalOverride provides the following for HTML elements ->
// The HTML <bdo> Element (or HTML bidirectional override element) is used to override the current directionality of text. It causes the directionality of the characters to be ignored in favor of the specified directionality.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/bdo
func BidirectionalOverride(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("bdo",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// BlockQuote provides the following for HTML elements ->
// The HTML <blockquote> Element (or HTML Block Quotation Element) indicates that the enclosed text is an extended quotation. Usually, this is rendered visually by indentation (see Notes for how to change it). A URL for the source of the quotation may be given using the cite attribute, while a text representation of the source can be given using the <cite> element.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/blockquote
func BlockQuote(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("blockquote",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Break provides the following for HTML elements ->
// The HTML element line break <br> produces a line break in text (carriage-return). It is useful for writing a poem or an address, where the division of lines is significant.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/br
func Break(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("br",true)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Button provides the following for HTML elements ->
// The HTML <button> Element represents a clickable button.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/button
func Button(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("button",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Canvas provides the following for HTML elements ->
// The HTML <canvas> Element can be used to draw graphics via scripting (usually JavaScript). For example, it can be used to draw graphs, make photo compositions or even perform animations. You may (and should) provide alternate content inside the <canvas> block. That content will be rendered both on older browsers that don't support canvas and in browsers with JavaScript disabled.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/canvas
func Canvas(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("canvas",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Caption provides the following for HTML elements ->
// The HTML <caption> Element (or HTML Table Caption Element) represents the title of a table. Though it is always the first descendant of a <table>, its styling, using CSS, may place it elsewhere, relative to the table.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/caption
func Caption(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("caption",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Citation provides the following for HTML elements ->
// The HTML Citation Element (<cite>) represents a reference to a creative work. It must include the title of a work or a URL reference, which may be in an abbreviated form according to the conventions used for the addition of citation metadata.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/cite
func Citation(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("cite",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Code provides the following for HTML elements ->
// The HTML Code Element (<code>) represents a fragment of computer code. By default, it is displayed in the browser's default monospace font.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/code
func Code(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("code",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Column provides the following for HTML elements ->
// The HTML Table Column Element (<col>) defines a column within a table and is used for defining common semantics on all common cells. It is generally found within a <colgroup> element.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/col
func Column(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("col",true)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// ColumnGroup provides the following for HTML elements ->
// The HTML Table Column Group Element (<colgroup>) defines a group of columns within a table.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/colgroup
func ColumnGroup(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("colgroup",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Data provides the following for HTML elements ->
// The HTML <data> Element links a given content with a machine-readable translation. If the content is time- or date-related, the <time> must be used.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/data
func Data(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("data",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// DataList provides the following for HTML elements ->
// The HTML Datalist Element (<datalist>) contains a set of <option> elements that represent the values available for other controls.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/datalist
func DataList(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("datalist",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Description provides the following for HTML elements ->
// The HTML <dd> element (HTML Description Element) indicates the description of a term in a description list (<dl>) element. This element can occur only as a child element of a description list and it must follow a <dt> element.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/dd
func Description(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("dd",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// DeletedText provides the following for HTML elements ->
// The HTML Deleted Text Element (<del>) represents a range of text that has been deleted from a document. This element is often (but need not be) rendered with strike-through text.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/del
func DeletedText(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("del",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Details provides the following for HTML elements ->
// The HTML <details> element is used as a disclosure widget from which the user can retrieve additional information.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/details
func Details(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("details",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Definition provides the following for HTML elements ->
// The HTML Definition Element (<dfn>) represents the defining instance of a term.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/dfn
func Definition(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("dfn",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Dialog provides the following for HTML elements ->
// The HTML <dialog> element represents a dialog box or other interactive component, such as an inspector or window. <form> elements can be integrated within a dialog by specifying them with the attribute method="dialog". When such a form is submitted, the dialog is closed with a returnValue attribute set to the value of the submit button used.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/dialog
func Dialog(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("dialog",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Div provides the following for HTML elements ->
// The HTML <div> element (or HTML Document Division Element) is the generic container for flow content, which does not inherently represent anything. It can be used to group elements for styling purposes (using the class or id attributes), or because they share attribute values, such as lang. It should be used only when no other semantic element (such as <article> or <nav>) is appropriate.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/div
func Div(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("div",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// DescriptionList provides the following for HTML elements ->
// The HTML <dl> element (or HTML Description List Element) encloses a list of pairs of terms and descriptions. Common uses for this element are to implement a glossary or to display metadata (a list of key-value pairs).
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/dl
func DescriptionList(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("dl",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// DefinitionTerm provides the following for HTML elements ->
// The HTML <dt> element (or HTML Definition Term Element) identifies a term in a definition list. This element can occur only as a child element of a <dl>. It is usually followed by a <dd> element; however, multiple <dt> elements in a row indicate several terms that are all defined by the immediate next <dd> element.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/dt
func DefinitionTerm(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("dt",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Element provides the following for HTML elements ->
// The HTML <element> element is used to define new custom DOM elements.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/element
func Element(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("element",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Emphasis provides the following for HTML elements ->
// The HTML element emphasis  <em> marks text that has stress emphasis. The <em> element can be nested, with each level of nesting indicating a greater degree of emphasis.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/em
func Emphasis(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("em",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Embed provides the following for HTML elements ->
// The HTML <embed> Element represents an integration point for an external application or interactive content (in other words, a plug-in).
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/embed
func Embed(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("embed",true)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// FieldSet provides the following for HTML elements ->
// The HTML <fieldset> element is used to group several controls as well as labels (<label>) within a web form.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/fieldset
func FieldSet(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("fieldset",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// FigureCaption provides the following for HTML elements ->
// The HTML <figcaption> element represents a caption or a legend associated with a figure or an illustration described by the rest of the data of the <figure> element which is its immediate ancestor which means <figcaption> can be the first or last element inside a <figure> block. Also, the HTML Figcaption Element is optional; if not provided, then the parent figure element will have no caption.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/figcaption
func FigureCaption(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("figcaption",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Figure provides the following for HTML elements ->
// The HTML <figure> element represents self-contained content, frequently with a caption (<figcaption>), and is typically referenced as a single unit. While it is related to the main flow, its position is independent of the main flow. Usually this is an image, an illustration, a diagram, a code snippet, or a schema that is referenced in the main text, but that can be moved to another page or to an appendix without affecting the main flow.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/figure
func Figure(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("figure",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Footer provides the following for HTML elements ->
// The HTML <footer> element represents a footer for its nearest sectioning content or sectioning root element. A footer typically contains information about the author of the section, copyright data or links to related documents.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/footer
func Footer(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("footer",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Form provides the following for HTML elements ->
// The HTML <form> element represents a document section that contains interactive controls to submit information to a web server.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/form
func Form(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("form",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Header provides the following for HTML elements ->
// The HTML <header> element represents a group of introductory or navigational aids. It may contain some heading elements but also other elements like a logo, wrapped section's header, a search form, and so on.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/header
func Header(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("header",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// HeadingsGroup provides the following for HTML elements ->
// The HTML <hgroup> Element (HTML Headings Group Element) represents the heading of a section. It defines a single title that participates in the outline of the document as the heading of the implicit or explicit section that it belongs to.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/hgroup
func HeadingsGroup(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("hgroup",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// HorizontalRule provides the following for HTML elements ->
// The HTML <hr> element represents a thematic break between paragraph-level elements (for example, a change of scene in a story, or a shift of topic with a section). In previous versions of HTML, it represented a horizontal rule. It may still be displayed as a horizontal rule in visual browsers, but is now defined in semantic terms, rather than presentational terms.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/hr
func HorizontalRule(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("hr",true)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Italic provides the following for HTML elements ->
// The HTML <i> Element represents a range of text that is set off from the normal text for some reason, for example, technical terms, foreign language phrases, or fictional character thoughts. It is typically displayed in italic type.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/i
func Italic(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("i",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// InlineFrame provides the following for HTML elements ->
// The HTML Inline Frame Element (<iframe>) represents a nested browsing context, effectively embedding another HTML page into the current page. In HTML 4.01, a document may contain a head and a body or a head and a frameset, but not both a body and a frameset. However, an <iframe> can be used within a normal document body. Each browsing context has its own session history and active document. The browsing context that contains the embedded content is called the parent browsing context. The top-level browsing context (which has no parent) is typically the browser window.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/iframe
func InlineFrame(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("iframe",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Image provides the following for HTML elements ->
// The HTML <img> element represents an image in the document.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/img
func Image(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("img",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Input provides the following for HTML elements ->
// The HTML element <input> is used to create interactive controls for web-based forms in order to accept data from the user. How an <input> works varies considerably depending on the value of its type attribute.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input
func Input(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("input",true)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// InsertedText provides the following for HTML elements ->
// The HTML <ins> Element (or HTML Inserted Text) HTML represents a range of text that has been added to a document.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/ins
func InsertedText(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("ins",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// KeyboardInput provides the following for HTML elements ->
// The HTML Keyboard Input Element (<kbd>) represents user input and produces an inline element displayed in the browser's default monospace font.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/kbd
func KeyboardInput(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("kbd",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Label provides the following for HTML elements ->
// The HTML Label Element (<label>) represents a caption for an item in a user interface. It can be associated with a control either by placing the control element inside the <label> element, or by using the for attribute. Such a control is called the labeled control of the label element. One input can be associated with multiple labels.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/label
func Label(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("label",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Legend provides the following for HTML elements ->
// The HTML <legend> Element (or HTML Legend Field Element) represents a caption for the content of its parent <fieldset>.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/legend
func Legend(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("legend",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// ListItem provides the following for HTML elements ->
// The HTML <li> element (or HTML List Item Element) is used to represent an item in a list. It must be contained in a parent element: an ordered list (<ol>), an unordered list (<ul>), or a menu (<menu>). In menus and unordered lists, list items are usually displayed using bullet points. In ordered lists, they are usually displayed with an ascending counter on the left, such as a number or letter.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/li
func ListItem(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("li",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Link provides the following for HTML elements ->
// The HTML <link> element specifies relationships between the current document and an external resource. Possible uses for this element include defining a relational framework for navigation. This Element is most used to link to style sheets.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/link
func Link(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("link",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Main provides the following for HTML elements ->
// The HTML <main> element represents the main content of  the <body> of a document or application. The main content area consists of content that is directly related to, or expands upon the central topic of a document or the central functionality of an application. This content should be unique to the document, excluding any content that is repeated across a set of documents such as sidebars, navigation links, copyright information, site logos, and search forms (unless the document's main function is as a search form).
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/main
func Main(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("main",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Map provides the following for HTML elements ->
// The HTML <map> element is used with <area> elements to define an image map (a clickable link area).
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/map
func Map(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("map",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Mark provides the following for HTML elements ->
// The HTML Mark Element (<mark>) represents highlighted text, i.e., a run of text marked for reference purpose, due to its relevance in a particular context. For example it can be used in a page showing search results to highlight every instance of the searched-for word.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/mark
func Mark(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("mark",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Menu provides the following for HTML elements ->
// The HTML <menu> element represents a group of commands that a user can perform or activate. This includes both list menus, which might appear across the top of a screen, as well as context menus, such as those that might appear underneath a button after it has been clicked.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/menu
func Menu(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("menu",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// MenuItem provides the following for HTML elements ->
// The HTML <menuitem> element represents a command that a user is able to invoke through a popup menu. This includes context menus, as well as menus that might be attached to a menu button.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/menuitem
func MenuItem(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("menuitem",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Meta provides the following for HTML elements ->
// The HTML <meta> element represents any metadata information that cannot be represented by one of the other HTML meta-related elements (<base>, <link>, <script>, <style> or <title>).
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/meta
func Meta(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("meta",true)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Meter provides the following for HTML elements ->
// The HTML <meter> Element represents either a scalar value within a known range or a fractional value.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/meter
func Meter(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("meter",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Navigation provides the following for HTML elements ->
// The HTML <nav> element (HTML Navigation Element) represents a section of a page that links to other pages or to parts within the page: a section with navigation links.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/nav
func Navigation(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("nav",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// NoFrames provides the following for HTML elements ->
// <noframes> is an HTML element which is used to supporting browsers which are not able to support <frame> elements or configured to do so.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/noframes
func NoFrames(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("noframes",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// NoScript provides the following for HTML elements ->
// The HTML <noscript> Element defines a section of html to be inserted if a script type on the page is unsupported or if scripting is currently turned off in the browser.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/noscript
func NoScript(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("noscript",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Object provides the following for HTML elements ->
// The HTML Embedded Object Element (<object>) represents an external resource, which can be treated as an image, a nested browsing context, or a resource to be handled by a plugin.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/object
func Object(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("object",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// OrderedList provides the following for HTML elements ->
// The HTML <ol> Element (or HTML Ordered List Element) represents an ordered list of items. Typically, ordered-list items are displayed with a preceding numbering, which can be of any form, like numerals, letters or Romans numerals or even simple bullets. This numbered style is not defined in the HTML description of the page, but in its associated CSS, using the list-style-type property.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/ol
func OrderedList(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("ol",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// OptionsGroup provides the following for HTML elements ->
// In a Web form, the HTML <optgroup> element  creates a grouping of options within a <select> element.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/optgroup
func OptionsGroup(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("optgroup",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Option provides the following for HTML elements ->
// In a Web form, the HTML <option> element is used to create a control representing an item within a <select>, an <optgroup> or a <datalist> HTML5 element.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/option
func Option(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("option",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Output provides the following for HTML elements ->
// The HTML <output> element represents the result of a calculation or user action.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/output
func Output(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("output",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Paragraph provides the following for HTML elements ->
// The HTML <p> element (or HTML Paragraph Element) represents a paragraph of text.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/p
func Paragraph(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("p",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Parameter provides the following for HTML elements ->
// The HTML <param> Element (or HTML Parameter Element) defines parameters for <object>.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/param
func Parameter(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("param",true)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Picture provides the following for HTML elements ->
// The HTML <picture> element is a container used to specify multiple <source> elements for a specific <img> contained in it. The browser will choose the most suitable source according to the current layout of the page (the constraints of the box the image will appear in) and the device it will be displayed on (e.g. a normal or hiDPI device.)
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/picture
func Picture(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("picture",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Preformatted provides the following for HTML elements ->
// The HTML <pre> element (or HTML Preformatted Text) represents preformatted text. Text within this element is typically displayed in a non-proportional ("monospace") font exactly as it is laid out in the file. Whitespace inside this element is displayed as typed.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/pre
func Preformatted(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("pre",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Progress provides the following for HTML elements ->
// The HTML <progress> Element is used to view the completion progress of a task. While the specifics of how it's displayed is left up to the browser developer, it's typically displayed as a progress bar. Javascript can be used to manipulate the value of progress bar.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/progress
func Progress(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("progress",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Quote provides the following for HTML elements ->
// The HTML Quote Element (<q>) indicates that the enclosed text is a short inline quotation. This element is intended for short quotations that don't require paragraph breaks; for long quotations use <blockquote> element.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/q
func Quote(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("q",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// RubyParenthesis provides the following for HTML elements ->
// The HTML <rp> element is used to provide fall-back parenthesis for browsers non-supporting ruby annotations. Ruby annotations are for showing pronunciation of East Asian characters, like using Japanese furigana or Taiwainese bopomofo characters. The <rp> element is used in the case of lack of <ruby> element support its content has what should be displayed in order to indicate the presence of a ruby annotation, usually parentheses.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/rp
func RubyParenthesis(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("rp",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// RubyText provides the following for HTML elements ->
// The HTML <rt> Element embraces pronunciation of characters presented in a ruby annotations, which are used to describe the pronunciation of East Asian characters. This element is always used inside a <ruby> element.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/rt
func RubyText(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("rt",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Rtc provides the following for HTML elements ->
// The HTML <rtc> Element embraces semantic annotations of characters presented in a ruby of <rb> elements used inside of <ruby> element. <rb> elements can have both pronunciation (<rt>) and semantic (<rtc>) annotations.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/rtc
func Rtc(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("rtc",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Ruby provides the following for HTML elements ->
// The HTML <ruby> Element represents a ruby annotation. Ruby annotations are for showing pronunciation of East Asian characters.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/ruby
func Ruby(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("ruby",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Strikethrough provides the following for HTML elements ->
// The HTML Strikethrough Element (<s>) renders text with a strikethrough, or a line through it. Use the <s> element to represent things that are no longer relevant or no longer accurate. However, <s> is not appropriate when indicating document edits; for that, use the <del> and <ins> elements, as appropriate.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/s
func Strikethrough(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("s",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Sample provides the following for HTML elements ->
// The HTML <samp> element is an element intended to identify sample output from a computer program. It is usually displayed in the browser's default monotype font (such as Lucida Console).
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/samp
func Sample(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("samp",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Script provides the following for HTML elements ->
// The HTML Script Element (<script>) is used to embed or reference an executable script within an HTML or XHTML document.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/script
func Script(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("script",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Section provides the following for HTML elements ->
// The HTML <section> element represents a generic section of a document, i.e., a thematic grouping of content, typically with a heading. Each <section> should be identified, typically by including a heading (<h1>-<h6> element) as a child of the <section> element.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/section
func Section(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("section",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Select provides the following for HTML elements ->
// The HTML select (<select>) element represents a control that presents a menu of options. The options within the menu are represented by <option> elements, which can be grouped by <optgroup> elements. Options can be pre-selected for the user.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/select
func Select(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("select",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Shadow provides the following for HTML elements ->
// The HTML <shadow> element is used as a shadow DOM insertion point. You might use it if you have created multiple shadow roots under a shadow host. It is not useful in ordinary HTML. It is used with Web Components.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Shadow
func Shadow(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("shadow",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Small provides the following for HTML elements ->
// The HTML Small Element (<small>) makes the text font size one size smaller (for example, from large to medium, or from small to x-small) down to the browser's minimum font size.  In HTML5, this element is repurposed to represent side-comments and small print, including copyright and legal text, independent of its styled presentation.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/small
func Small(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("small",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Source provides the following for HTML elements ->
// The HTML <source> element specifies multiple media resources for either the <picture>, the <audio> or the <video> element. It is an empty element. It is commonly used to serve the same media content in multiple formats supported by different browsers.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/source
func Source(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("source",true)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Span provides the following for HTML elements ->
// The HTML <span> element is a generic inline container for phrasing content, which does not inherently represent anything. It can be used to group elements for styling purposes (using the class or id attributes), or because they share attribute values, such as lang.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/span
func Span(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("span",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Strong provides the following for HTML elements ->
// The HTML Strong Element (<strong>) gives text strong importance, and is typically displayed in bold.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/strong
func Strong(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("strong",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Style provides the following for HTML elements ->
// The HTML <style> element contains style information for a document, or part of a document. By default, the style instructions written inside that element are expected to be CSS.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/style
func Style(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("style",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Subscript provides the following for HTML elements ->
// The HTML Subscript Element (<sub>) defines a span of text that should be displayed, for typographic reasons, lower, and often smaller, than the main span of text.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/sub
func Subscript(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("sub",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Summary provides the following for HTML elements ->
// The HTML <summary> element is used as a summary, caption, or legend for the content of a <details> element.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/summary
func Summary(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("summary",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Superscript provides the following for HTML elements ->
// The HTML Superscript Element (<sup>) defines a span of text that should be displayed, for typographic reasons, higher, and often smaller, than the main span of text.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/sup
func Superscript(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("sup",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Table provides the following for HTML elements ->
// The HTML Table Element (<table>) represents tabular data - i.e., information expressed via a two dimensional data table.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/table
func Table(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("table",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// TableBody provides the following for HTML elements ->
// The HTML Table Body Element (<tbody>) defines one or more <tr> element data-rows to be the body of its parent <table> element (as long as no <tr> elements are immediate children of that table element.)  In conjunction with a preceding <thead> and/or <tfoot> element, <tbody> provides additional semantic information for devices such as printers and displays. Of the parent table's child elements, <tbody> represents the content which, when longer than a page, will most likely differ for each page printed; while the content of <thead> and <tfoot> will be the same or similar for each page printed. For displays, <tbody> will enable separate scrolling of the <thead>, <tfoot>, and <caption> elements of the same parent <table> element.  Note that unlike the <thead>, <tfoot>, and <caption> elements however, multiple <tbody> elements are permitted (if consecutive), allowing the data-rows in long tables to be divided into different sections, each separately formatted as needed.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/tbody
func TableBody(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("tbody",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// TableData provides the following for HTML elements ->
// The Table cell HTML element (<td>) defines a cell of a table that contains data. It participates in the table model.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/td
func TableData(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("td",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Template provides the following for HTML elements ->
// The HTML template element <template> is a mechanism for holding client-side content that is not to be rendered when a page is loaded but may subsequently be instantiated during runtime using JavaScript. 
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/template
func Template(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("template",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// TextArea provides the following for HTML elements ->
// The HTML <textarea> element represents a multi-line plain-text editing control.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/textarea
func TextArea(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("textarea",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// TableFoot provides the following for HTML elements ->
// The HTML Table Foot Element (<tfoot>) defines a set of rows summarizing the columns of the table.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/tfoot
func TableFoot(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("tfoot",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// TableHeader provides the following for HTML elements ->
// The HTML element table header cell <th> defines a cell as header of a group of table cells. The exact nature of this group is defined by the scope and headers attributes.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/th
func TableHeader(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("th",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// TableHead provides the following for HTML elements ->
// The HTML Table Head Element (<thead>) defines a set of rows defining the head of the columns of the table.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/thead
func TableHead(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("thead",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Time provides the following for HTML elements ->
// The HTML <time> element represents either a time on a 24-hour clock or a precise date in the Gregorian calendar (with optional time and timezone information).
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/time
func Time(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("time",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Title provides the following for HTML elements ->
// The HTML <title> element defines the title of the document, shown in a browser's title bar or on the page's tab. It can only contain text, and any contained tags are ignored.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/title
func Title(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("title",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// TableRow provides the following for HTML elements ->
// The HTML element table row <tr> defines a row of cells in a table. Those can be a mix of <td> and <th> elements.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/tr
func TableRow(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("tr",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Track provides the following for HTML elements ->
// The HTML <track> element is used as a child of the media elements—<audio> and <video>. It lets you specify timed text tracks (or time-based data), for example to automatically handle subtitles. The tracks are formatted in WebVTT format (.vtt files) — Web Video Text Tracks.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/track
func Track(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("track",true)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Underline provides the following for HTML elements ->
// The HTML Underline Element (<u>) renders text with an underline, a line under the baseline of its content.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/u
func Underline(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("u",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// UnorderedList provides the following for HTML elements ->
// The HTML <ul> element (or HTML Unordered List Element) represents an unordered list of items, namely a collection of items that do not have a numerical ordering, and their order in the list is meaningless. Typically, unordered-list items are displayed with a bullet, which can be of several forms, like a dot, a circle or a squared. The bullet style is not defined in the HTML description of the page, but in its associated CSS, using the list-style-type property.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/ul
func UnorderedList(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("ul",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Variable provides the following for HTML elements ->
// The HTML Variable Element (<var>) represents a variable in a mathematical expression or a programming context.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/var
func Variable(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("var",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Video provides the following for HTML elements ->
// Use the  HTML <video> element to embed video content in a document. The video element contains one or more video sources. To specify a video source, use either the src attribute or the <source> element; the browser will choose the most suitable one.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/video
func Video(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("video",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// WordBreakOpportunity provides the following for HTML elements ->
// The HTML element word break opportunity <wbr> represents a position within text where the browser may optionally break a line, though its line-breaking rules would not otherwise create a break at that location.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/wbr
func WordBreakOpportunity(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("wbr",true)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Header1 provides the following for HTML elements ->
// Heading elements implement six levels of document headings, <h1> is the most important and <h6> is the least. A heading element briefly describes the topic of the section it introduces. Heading information may be used by user agents, for example, to construct a table of contents for a document automatically.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Heading_Elements
func Header1(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("h1",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Header2 provides the following for HTML elements ->
// Heading elements implement six levels of document headings, <h1> is the most important and <h6> is the least. A heading element briefly describes the topic of the section it introduces. Heading information may be used by user agents, for example, to construct a table of contents for a document automatically.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Heading_Elements
func Header2(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("h2",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Header3 provides the following for HTML elements ->
// Heading elements implement six levels of document headings, <h1> is the most important and <h6> is the least. A heading element briefly describes the topic of the section it introduces. Heading information may be used by user agents, for example, to construct a table of contents for a document automatically.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Heading_Elements
func Header3(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("h3",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Header4 provides the following for HTML elements ->
// Heading elements implement six levels of document headings, <h1> is the most important and <h6> is the least. A heading element briefly describes the topic of the section it introduces. Heading information may be used by user agents, for example, to construct a table of contents for a document automatically.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Heading_Elements
func Header4(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("h4",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Header5 provides the following for HTML elements ->
// Heading elements implement six levels of document headings, <h1> is the most important and <h6> is the least. A heading element briefly describes the topic of the section it introduces. Heading information may be used by user agents, for example, to construct a table of contents for a document automatically.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Heading_Elements
func Header5(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("h5",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}

// Header6 provides the following for HTML elements ->
// Heading elements implement six levels of document headings, <h1> is the most important and <h6> is the least. A heading element briefly describes the topic of the section it introduces. Heading information may be used by user agents, for example, to construct a table of contents for a document automatically.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Heading_Elements
func Header6(markup ...trees.Appliable) *trees.Markup {
	e := trees.NewMarkup("h6",false)
	for _, m := range markup {
		if m == nil { continue }
		m.Apply(e)
	}
	return e
}
