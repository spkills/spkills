package main

import (
	"fmt"
	"net/http"
)

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
