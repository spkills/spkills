package main

import (
	"log"
	"os"

	"github.com/buaazp/fasthttprouter"
	"github.com/spkills/spkills/controller"
	"github.com/valyala/fasthttp"
)

func main() {
	socketFile := "/spkills/run/spkills.sock"
	router := fasthttprouter.New()

	router.GET("/", controller.RootController)

	log.Fatal(fasthttp.ListenAndServeUNIX(socketFile, os.FileMode(777), router.Handler))

}
