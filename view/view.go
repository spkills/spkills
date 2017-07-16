package view

import (
	"github.com/spkills/spkills/templates"
	"github.com/valyala/fasthttp"
)

func RootView(ctx *fasthttp.RequestCtx, str string) {
	p := &templates.MainPage{ReceiveMsg: str}
	templates.WritePageTemplate(ctx, p)
	//fmt.Fprintf(w, templates.Hello(str))
}
