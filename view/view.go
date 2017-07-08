package view

import (
	"net/http"
	"github.com/spkills/spkills/templates"
)

func RootView(w http.ResponseWriter, str string) {
	p := &templates.MainPage{ReceiveMsg: str}
	templates.WritePageTemplate(w,p)
	//fmt.Fprintf(w, templates.Hello(str))
}
