package controller

import (
	"github.com/spkills/spkills/model"
	"github.com/spkills/spkills/view"
	"github.com/valyala/fasthttp"
)

func sandboxHandler(ctx *fasthttp.RequestCtx) {
	res := model.RootModel()

	ctx.SetContentType("text/html; charset=utf-8")
	ctx.SetStatusCode(fasthttp.StatusOK)
	view.RootView(ctx, res)
}
