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

	log.Fatal(fasthttp.ListenAndServeUNIX(socketFile, os.FileMode(777), router.Handler))

}
