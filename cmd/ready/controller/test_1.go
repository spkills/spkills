package controller

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

func test_1Handler(ctx *fasthttp.RequestCtx) {
	fmt.Println("test")
}
