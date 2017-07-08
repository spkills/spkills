package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/spkills/spkills/controller"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	controller.RootController(w, r)
}

func main() {
	socketFile := "/spkills/run/spkills.sock"
	mux := http.NewServeMux()

	// Routers
	mux.HandleFunc("/", rootHandler)

	ln, err := net.Listen("unix", socketFile)
	if err != nil {
		log.Fatalf("error %v", err)
	}

	defer func(ln net.Listener) {
		if err := ln.Close(); err != nil {
			log.Fatalf("error %v", err)
		}
	}(ln)

	c := make(chan os.Signal, 2)
	go func(ln net.Listener, c chan os.Signal) {
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
		if err := ln.Close(); err != nil {
			log.Fatalf("error %v", err)
		}
		os.Exit(1)
	}(ln, c)

	log.Fatalf("error %v", http.Serve(ln, mux))
}
