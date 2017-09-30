package main

import (
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/spkills/spkills/config"
	"github.com/spkills/spkills/model"
)

func BenchmarkFetchRoomNameByDB(b *testing.B) {
	var conf config.Config
	_, err := toml.DecodeFile("../config/config.toml", &conf)
	if err != nil {
		// Error Handling
	}
	model.InitDB(conf)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		model.FetchRoomNameByDB(1)
	}
}

func BenchmarkFetchRoomNameByCache(b *testing.B) {
	var conf config.Config
	_, err := toml.DecodeFile("../config/config.toml", &conf)
	if err != nil {
		// Error Handling
	}
	model.InitDB(conf)
	model.InitCache()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		model.FetchRoomNameByCache(1)
	}
}

func BenchmarkFetchRoomNameByRedis(b *testing.B) {
	var conf config.Config
	_, err := toml.DecodeFile("../config/config.toml", &conf)
	if err != nil {
		// Error Handling
	}
	model.InitDB(conf)
	model.InitRedis(conf)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		model.FetchRoomNameByRedis(1)
	}
}
