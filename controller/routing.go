package controller

import (
	"github.com/buaazp/fasthttprouter"
)

// AddRouteing add routing path from go generate ready
func AddRouting(router *fasthttprouter.Router) {

	router.GET("/sandbox", SandboxGetHandler)
	router.POST("/sandbox", SandboxPostHandler)
	router.PUT("/sandbox", SandboxPutHandler)
	router.DELETE("/sandbox", SandboxDeleteHandler)
	router.HEAD("/sandbox", SandboxHeadHandler)

}
