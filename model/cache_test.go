package model

import (
	"testing"
)

func TestInitCache(t *testing.T) {

	t.Run("InitCache", func(t *testing.T) {
		InitCache()
		if Cache == nil {
			t.Fatalf("model.Cache is nil!!")
		}
	})
}
