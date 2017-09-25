package model

import (
	"testing"
	"time"
)

func TestInitCache(t *testing.T) {

	t.Run("InitCache", func(t *testing.T) {
		InitCache()
		if Cache == nil {
			t.Fatalf("model.Cache is nil!!")
		}
	})
}

func TestFetchByCache(t *testing.T) {
	InitCache()
	closureA := func(id int) int {
		return 0
	}
	t.Run("First fetch", func(t *testing.T) {
		result := FetchByCache("hoge", 10*time.Second, closureA, 1)
		if result != 0 {
			t.Fatalf("closureA not called")
		}
	})
	t.Run("Second fetch", func(t *testing.T) {
		Cache.Set("hoge", 10, 10*time.Second)
		result := FetchByCache("hoge", 10*time.Second, closureA, 1)
		if result != 10 {
			t.Fatalf("closureA called")
		}
	})
}
