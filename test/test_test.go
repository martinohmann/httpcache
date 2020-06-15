package test_test

import (
	"testing"

	"github.com/martinohmann/httpcache"
	"github.com/martinohmann/httpcache/test"
)

func TestMemoryCache(t *testing.T) {
	test.Cache(t, httpcache.NewMemoryCache())
}
