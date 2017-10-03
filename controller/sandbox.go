package controller

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

func sandboxGetHandler(ctx *fasthttp.RequestCtx) {

	name := "sandbox"

	ctx.SetContentType("text/html; charset=utf-8")
	ctx.SetStatusCode(fasthttp.StatusOK)
	fmt.Fprintf(ctx, name)

}

func sandboxPostHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("text/html; charset=utf-8")
	ctx.SetStatusCode(fasthttp.StatusNotImplemented)
	fmt.Fprintf(ctx, "NotImplemented")
}

func sandboxPutHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("text/html; charset=utf-8")
	ctx.SetStatusCode(fasthttp.StatusNotImplemented)
	fmt.Fprintf(ctx, "NotImplemented")
}

func sandboxDeleteHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("text/html; charset=utf-8")
	ctx.SetStatusCode(fasthttp.StatusNotImplemented)
	fmt.Fprintf(ctx, "NotImplemented")
}

func sandboxHeadHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("text/html; charset=utf-8")
	ctx.SetStatusCode(fasthttp.StatusOK)
	fmt.Fprintf(ctx, "")
}
