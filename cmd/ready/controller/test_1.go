package controller

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

func Test_1GetHandler(ctx *fasthttp.RequestCtx) {

	name := "test_1"

	ctx.SetContentType("text/html; charset=utf-8")
	ctx.SetStatusCode(fasthttp.StatusOK)
	fmt.Fprintf(ctx, name)

}

func Test_1PostHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("text/html; charset=utf-8")
	ctx.SetStatusCode(fasthttp.StatusNotImplemented)
	fmt.Fprintf(ctx, "NotImplemented")
}

func Test_1PutHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("text/html; charset=utf-8")
	ctx.SetStatusCode(fasthttp.StatusNotImplemented)
	fmt.Fprintf(ctx, "NotImplemented")
}

func Test_1DeleteHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("text/html; charset=utf-8")
	ctx.SetStatusCode(fasthttp.StatusNotImplemented)
	fmt.Fprintf(ctx, "NotImplemented")
}

func Test_1HeadHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("text/html; charset=utf-8")
	ctx.SetStatusCode(fasthttp.StatusOK)
	fmt.Fprintf(ctx, "")
}
