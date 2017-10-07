package controller

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

func TestGetHandler(ctx *fasthttp.RequestCtx) {

	name := "test"

	ctx.SetContentType("text/html; charset=utf-8")
	ctx.SetStatusCode(fasthttp.StatusOK)
	fmt.Fprintf(ctx, name)

}

func TestPostHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("text/html; charset=utf-8")
	ctx.SetStatusCode(fasthttp.StatusNotImplemented)
	fmt.Fprintf(ctx, "NotImplemented")
}

func TestPutHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("text/html; charset=utf-8")
	ctx.SetStatusCode(fasthttp.StatusNotImplemented)
	fmt.Fprintf(ctx, "NotImplemented")
}

func TestDeleteHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("text/html; charset=utf-8")
	ctx.SetStatusCode(fasthttp.StatusNotImplemented)
	fmt.Fprintf(ctx, "NotImplemented")
}

func TestHeadHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("text/html; charset=utf-8")
	ctx.SetStatusCode(fasthttp.StatusOK)
	fmt.Fprintf(ctx, "")
}
