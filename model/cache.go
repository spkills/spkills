package model

import (
	"time"

	cache "github.com/patrickmn/go-cache"
)

var Cache *cache.Cache

func InitCache() {
	Cache = cache.New(5*time.Minute, 10*time.Minute)
}
