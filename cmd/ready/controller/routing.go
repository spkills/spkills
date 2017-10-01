package controller

import (
	"github.com/buaazp/fasthttprouter"
)

func Routing(router *fasthttprouter.Router) {

	router.GET("/test", testHandler)
	router.GET("/test_1", test_1Handler)
	router.GET("/test_2", test_2Handler)
}
