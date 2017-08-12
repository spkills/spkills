package main

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/buaazp/fasthttprouter"
	"github.com/spkills/spkills/config"
	"github.com/spkills/spkills/controller"
	"github.com/spkills/spkills/model"
	"github.com/valyala/fasthttp"
)

var conf config.Config

//go:generate cmd/ready/ready -file=route.conf
func main() {
	_, err := toml.DecodeFile("config/config.toml", &conf)
	if err != nil {
		// Error Handling
	}

	model.InitDB(conf)

	socketFile := "/spkills/run/spkills.sock"
	router := fasthttprouter.New()

	router.GET("/", controller.RootController)
	router.GET("/warming", controller.WarmingController)
	controller.Regist(router)

	log.Fatal(fasthttp.ListenAndServeUNIX(socketFile, os.FileMode(777), router.Handler))

}
