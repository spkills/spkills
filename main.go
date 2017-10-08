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

//go:generate ready-router -file=route.conf -templatesPath=templates
//go:generate ready-db -file=db.conf
func main() {
	_, err := toml.DecodeFile("config/config.toml", &conf)
	if err != nil {
		// Error Handling
	}

	model.InitDB(conf)
	model.InitCache()
	//model.InitRedis(conf) // redis起動してないと落ちる

	router := fasthttprouter.New()

	router.GET("/", controller.RootController)
	router.GET("/warming", controller.WarmingController)
	controller.AddRouting(router)

	if conf.Server.SocketFile != "" {
		socketFile := conf.Server.SocketFile
		log.Println("Start listening: " + socketFile)
		log.Fatal(fasthttp.ListenAndServeUNIX(socketFile, os.FileMode(777), router.Handler))
	} else {
		listenAddr := "0.0.0.0:" + conf.Server.Port
		log.Println("Start listening: " + listenAddr)
		log.Fatal(fasthttp.ListenAndServe(listenAddr, router.Handler))
	}

}
