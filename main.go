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
//go:generate sqlboiler mysql -o schema -p model  --no-tests -b room_owners,room_watchers
func main() {
	_, err := toml.DecodeFile("config/config.toml", &conf)
	if err != nil {
		// Error Handling
	}

	model.InitDB(conf)

	router := fasthttprouter.New()

	router.GET("/", controller.RootController)
	router.GET("/warming", controller.WarmingController)
	controller.Regist(router)

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
