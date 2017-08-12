package controller

import (
	"fmt"
	"strings"

	"github.com/spkills/spkills/model"
	"github.com/spkills/spkills/view"
	"github.com/valyala/fasthttp"
)

func RootController(ctx *fasthttp.RequestCtx) {
	res := model.RootModel()

	ctx.SetContentType("text/html; charset=utf-8")
	ctx.SetStatusCode(fasthttp.StatusOK)
	view.RootView(ctx, res)
}

func WarmingController(ctx *fasthttp.RequestCtx) {

	ctx.SetContentType("text/html; charset=utf-8")
	ctx.SetStatusCode(fasthttp.StatusOK)
	tables, _ := model.WarmUpInnoDB()
	res := strings.Join(tables, ",")
	fmt.Fprintf(ctx, res)
}
