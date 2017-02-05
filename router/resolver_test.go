package router_test

import (
	"testing"

	"github.com/gu-io/gu/router"
	"github.com/gu-io/gu/tests"
)

func TestResolver(t *testing.T) {
	rx := router.New("/:id")
	params, _, state := rx.Test("12")

	if !state {
		tests.Failed(t, "Should have matched giving path")
	}
	tests.Passed(t, "Should have matched giving path")

	val, ok := params["id"]
	if !ok {
		tests.Failed(t, "Should have retrieve parameter :id => %s", val)
	}
	tests.Passed(t, "Should have retrieve parameter :id => %s", val)

	rx.Done(func(px router.PushEvent) {
		tests.Passed(t, "Should have notified with PushEvent %#v", px)
	})

	rx.Failed(func(px router.PushEvent) {
		tests.Failed(t, "Should have notified with PushEvent %#v", px)
	})

	rx.Resolve(router.UseLocation("/12"))
}

func TestResolverLevels(t *testing.T) {
	home := router.New("/home/*")
	rx := router.New("/:id")

	home.Register(rx)

	rx.Done(func(px router.PushEvent) {
		tests.Passed(t, "Should have notified with PushEvent %#v", px)
	})

	rx.Failed(func(px router.PushEvent) {
		tests.Failed(t, "Should have notified with PushEvent %#v", px)
	})

	home.Resolve(router.UseLocation("home/12"))
}

func TestResolverFailed(t *testing.T) {
	rx := router.New("/:id")
	rx.Done(func(px router.PushEvent) {
		tests.Failed(t, "Should have notified with failed PushEvent %#v", px)
	})

	rx.Failed(func(px router.PushEvent) {
		tests.Passed(t, "Should have notified with failed PushEvent %#v", px)
	})

	rx.Resolve(router.UseLocation("/home/12"))
}
