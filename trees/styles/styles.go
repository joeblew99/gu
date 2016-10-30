package styles

import (
	"strconv"

	"github.com/influx6/gu/trees"
)

// Size presents a basic stringifed unit
type Size string

// Px returns the value in "%px" format
func Px(pixels int) Size {
	return Size(strconv.Itoa(pixels) + "px")
}

// Background provides the color style value
func Background(value string) trees.Property {
	return trees.NewCSSStyle("background", value)
}

// Color provides the color style value
func Color(value string) trees.Property {
	return trees.NewCSSStyle("color", value)
}

// Display provides the style setter that sets the css display value.
func Display(ops string) trees.Property {
	return trees.NewCSSStyle("display", ops)
}

// Height provides the height style value
func Height(size Size) trees.Property {
	return trees.NewCSSStyle("height", string(size))
}

// FontSize provides the margin style value
func FontSize(size Size) trees.Property {
	return trees.NewCSSStyle("font-size", string(size))
}

// Padding provides the margin style value
func Padding(size Size) trees.Property {
	return trees.NewCSSStyle("padding", string(size))
}

// Margin provides the margin style value
func Margin(size Size) trees.Property {
	return trees.NewCSSStyle("margin", string(size))
}

// Width provides the width style value
func Width(size Size) trees.Property {
	return trees.NewCSSStyle("width", string(size))
}
