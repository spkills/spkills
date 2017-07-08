package main

import (
	"fmt"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "瞬殺の美学")
}

func main() {
	listenPort := "8080"
	http.HandleFunc("/", rootHandler)
	http.ListenAndServe(":"+listenPort, nil)
}
