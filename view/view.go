package view

import (
	"fmt"
	"net/http"
	"github.com/spkills/spkills/templates"
)

func RootView(w http.ResponseWriter, str string) {
	fmt.Fprintf(w, templates.Hello("Foo"))
}
