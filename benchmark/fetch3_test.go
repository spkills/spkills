package main

import (
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/spkills/spkills/config"
	"github.com/spkills/spkills/model"
)

func BenchmarkUseSqlBoiler(b *testing.B) {
	var conf config.Config
	_, err := toml.DecodeFile("../config/config.toml", &conf)
	if err != nil {
		// Error Handling
	}
	model.InitDB(conf)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		model.FetchRoomByDB(1)
	}
}

func BenchmarkUseSqlBoiler2(b *testing.B) {
	var conf config.Config
	_, err := toml.DecodeFile("../config/config.toml", &conf)
	if err != nil {
		// Error Handling
	}
	model.InitDB(conf)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		model.FetchRoomByDB2(1)
	}
}

func BenchmarkNoUseSqlBoiler(b *testing.B) {
	var conf config.Config
	_, err := toml.DecodeFile("../config/config.toml", &conf)
	if err != nil {
		// Error Handling
	}
	model.InitDB(conf)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		model.FetchRoomByDB3(1)
	}
}
