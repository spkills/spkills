package controller

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

func Test_2GetHandler(ctx *fasthttp.RequestCtx) {

	name := "test_2"

	ctx.SetContentType("text/html; charset=utf-8")
	ctx.SetStatusCode(fasthttp.StatusOK)
	fmt.Fprintf(ctx, name)

}

func Test_2PostHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("text/html; charset=utf-8")
	ctx.SetStatusCode(fasthttp.StatusNotImplemented)
	fmt.Fprintf(ctx, "NotImplemented")
}

func Test_2PutHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("text/html; charset=utf-8")
	ctx.SetStatusCode(fasthttp.StatusNotImplemented)
	fmt.Fprintf(ctx, "NotImplemented")
}

func Test_2DeleteHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("text/html; charset=utf-8")
	ctx.SetStatusCode(fasthttp.StatusNotImplemented)
	fmt.Fprintf(ctx, "NotImplemented")
}

func Test_2HeadHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("text/html; charset=utf-8")
	ctx.SetStatusCode(fasthttp.StatusOK)
	fmt.Fprintf(ctx, "")
}
