package controller

import (
	"github.com/buaazp/fasthttprouter"
)

// AddRouteing add routing path from go generate ready
func AddRouting(router *fasthttprouter.Router) {

	router.GET("/sandbox", sandboxGetHandler)
	router.POST("/sandbox", sandboxPostHandler)
	router.PUT("/sandbox", sandboxPutHandler)
	router.DELETE("/sandbox", sandboxDeleteHandler)
	router.HEAD("/sandbox", sandboxHeadHandler)

}
