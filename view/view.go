package view

import (
	"fmt"
	"net/http"
)

func RootView(w http.ResponseWriter, str string) {
	fmt.Fprintf(w, str)
}
