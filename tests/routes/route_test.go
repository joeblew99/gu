package routes

import (
	"testing"

	"github.com/influx6/gu"
	"github.com/influx6/gu/tests"
)

func TestRoute(t *testing.T) {
	rm := gu.NewRouteManager()

	home := rm.L("/home/*")
	if _, _, pass := home.Test("/home"); !pass {
		tests.Failed(t, "Should have validated path")
	}
	tests.Passed(t, "Should have validated path")

	index := rm.L("/index/*")
	if _, _, pass := index.Test("/index"); !pass {
		tests.Failed(t, "Should have validated path /route")
	}
	tests.Passed(t, "Should have validated path /route")

	getModel := home.N("/models/*").N(":id")
  getModel.
}
