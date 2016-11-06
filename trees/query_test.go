package trees_test

import (
	"fmt"
	"testing"

	"github.com/influx6/gu/tests"
	"github.com/influx6/gu/trees"
)

func TestParseSelector(t *testing.T) {
	sels := trees.Query.ParseSelector("div.shower.ball")

	if sels == nil {
		tests.Failed(t, "Should have returned lists of selectors")
	}
	tests.Passed(t, "Should have returned lists of selectors")

	elem := sels[0]

	if elem.Tag != "div" {
		tests.Failed(t, "Should have div as selector tag")
	}
	tests.Passed(t, "Should have div as selector tag")

	if elem.Attr != "" {
		tests.Failed(t, "Should not have a attribute match required: %q", elem.Attr)
	}
	tests.Passed(t, "Should not have a attribute match required")

	if len(elem.Classes) != 2 {
		tests.Failed(t, "Should have 2 classes as part of element selector")
	}
	tests.Passed(t, "Should have 2 classes as part of element selector")

	if elem.Classes[0] != "shower" {
		tests.Failed(t, "Should contain %q as part of its classes", elem.Classes[0])
	}
	tests.Passed(t, "Should contain %q as part of its class", elem.Classes[0])

	if elem.Classes[1] != "ball" {
		tests.Failed(t, "Should contain %q as part of its classes", elem.Classes[1])
	}
	tests.Passed(t, "Should contain %q as part of its class", elem.Classes[1])

}

func TestParseMultiSelectors(t *testing.T) {
	sels := trees.Query.ParseSelector("div.shower.ball, a.embeded, h1#header.bots.call")
	if sels == nil {
		tests.Failed(t, "Should have returned lists of selectors")
	}
	tests.Passed(t, "Should have returned lists of selectors")

	for _, sel := range sels {
		fmt.Printf("%#v\n", sel)

		if sel.Children != nil {
			for _, sel := range sel.Children {
				fmt.Printf("\t%#v\n", sel)
			}
		}

		fmt.Println(" ")
	}
}
