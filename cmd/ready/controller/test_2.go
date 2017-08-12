package controller

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

func test_2Handler(ctx *fasthttp.RequestCtx) {
	fmt.Println("test")
}
