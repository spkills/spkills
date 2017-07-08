package main

import (
	"net/http"

	"spkills/controller"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	controller.RootController(w, r)
}

func main() {
	listenPort := "8080"
	http.HandleFunc("/", rootHandler)
	http.ListenAndServe(":"+listenPort, nil)
}
