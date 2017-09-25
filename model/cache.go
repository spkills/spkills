package model

import (
	"time"

	cache "github.com/patrickmn/go-cache"
)

var Cache *cache.Cache

func InitCache() {
	Cache = cache.New(5*time.Minute, 10*time.Minute)
}

func FetchByCache(
	key string,
	expire time.Duration,
	fn func(id int) int,
	arg int,
) int {
	val, found := Cache.Get(key)
	if !found {
		val = fn(arg)
		Cache.Set(key, val, expire)
	}
	return val.(int)
}
