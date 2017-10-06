package controller

import (
	"github.com/buaazp/fasthttprouter"
)

// AddRouteing add routing path from go generate ready
func AddRouting(router *fasthttprouter.Router) {

	router.GET("/test", TestGetHandler)
	router.POST("/test", TestPostHandler)
	router.PUT("/test", TestPutHandler)
	router.DELETE("/test", TestDeleteHandler)
	router.HEAD("/test", TestHeadHandler)

	router.GET("/test_1", Test_1GetHandler)
	router.POST("/test_1", Test_1PostHandler)
	router.PUT("/test_1", Test_1PutHandler)
	router.DELETE("/test_1", Test_1DeleteHandler)
	router.HEAD("/test_1", Test_1HeadHandler)

	router.GET("/test_2", Test_2GetHandler)
	router.POST("/test_2", Test_2PostHandler)
	router.PUT("/test_2", Test_2PutHandler)
	router.DELETE("/test_2", Test_2DeleteHandler)
	router.HEAD("/test_2", Test_2HeadHandler)

}
