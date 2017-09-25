package model

import (
	"testing"
	"time"

	"github.com/spkills/spkills/repository"
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
	closureA := func(args ...interface{}) interface{} {
		return 1
	}
	closureB := func(args ...interface{}) interface{} {
		return "piyo"
	}
	closureC := func(id int64) *repository.Room {
		room := &repository.Room{}
		return room
	}
	t.Run("キャッシュがなければ関数の値", func(t *testing.T) {
		result := FetchByCache("hoge", 10*time.Second, closureA, 1)
		if result.(int) != 1 {
			t.Fatalf("closureA not called")
		}
	})
	t.Run("キャッシュがあればそっち", func(t *testing.T) {
		Cache.Set("hoge", 10, 10*time.Second)
		result := FetchByCache("hoge", 10*time.Second, closureA, 1)
		if result.(int) != 10 {
			t.Fatalf("closureA called")
		}
	})
	t.Run("キャッシュがなければ関数の値(B)", func(t *testing.T) {
		result := FetchByCache("fuga", 10*time.Second, closureB, 1)
		if result.(string) != "piyo" {
			t.Fatalf("closureB not called")
		}
	})
	t.Run("キャッシュがなければ関数の値(B)", func(t *testing.T) {
		Cache.Set("fuga", "bbb", 10*time.Second)
		result := FetchByCache("fuga", 10*time.Second, closureB, 1)
		if result.(string) != "bbb" {
			t.Fatalf("closureB called")
		}
	})
	t.Run("キャッシュがなければ関数の値(C)", func(t *testing.T) {
		result := FetchByCache("fuga", 10*time.Second, closureC, 1)
		room := result.(*repository.Room)
		if room.Name != "" {
			t.Fatalf("closureB not called")
		}
	})
}
