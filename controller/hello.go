package controller

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

func helloHandler(ctx *fasthttp.RequestCtx) {
	fmt.Println("hello")
}
