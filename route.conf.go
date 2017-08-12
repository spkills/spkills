package main

import (
	"fmt"
	"net/http"
)

func regist(mux *http.ServeMux) {
	mux.HandleFunc("/sandbox", sandboxHandler)
	mux.HandleFunc("/error", errorHandler)
	mux.HandleFunc("/hogehoge", hogehogeHandler)
	mux.HandleFunc("/test", testHandler)
}

func sandboxHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("test")
}
func errorHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("test")
}
func hogehogeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("test")
}
func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("test")
}
