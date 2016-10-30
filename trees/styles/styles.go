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
	return &trees.Style{Name: "background", Value: value}
}

// Color provides the color style value
func Color(value string) trees.Property {
	return &trees.Style{Name: "color", Value: value}
}

// Display provides the style setter that sets the css display value.
func Display(ops string) trees.Property {
	return &trees.Style{Name: "display", Value: ops}
}

// Height provides the height style value
func Height(size Size) trees.Property {
	return &trees.Style{Name: "height", Value: string(size)}
}

// FontSize provides the margin style value
func FontSize(size Size) trees.Property {
	return &trees.Style{Name: "font-size", Value: string(size)}
}

// Padding provides the margin style value
func Padding(size Size) trees.Property {
	return &trees.Style{Name: "padding", Value: string(size)}
}

// Margin provides the margin style value
func Margin(size Size) trees.Property {
	return &trees.Style{Name: "margin", Value: string(size)}
}

// Width provides the width style value
func Width(size Size) trees.Property {
	return &trees.Style{Name: "width", Value: string(size)}
}
