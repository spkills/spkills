package controller

import (
	"github.com/buaazp/fasthttprouter"
)

// AddRouteing add routing path from go generate ready
func AddRouting(router *fasthttprouter.Router) {

	router.GET("/sandbox", sandboxHandler)
}
