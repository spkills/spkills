package controller

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

func testHandler(ctx *fasthttp.RequestCtx) {
	fmt.Println("test")
}
