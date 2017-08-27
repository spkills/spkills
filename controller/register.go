package controller

import (
	"github.com/buaazp/fasthttprouter"
)

func Regist(router *fasthttprouter.Router) {
	router.GET("/sandbox", sandboxHandler)
}
