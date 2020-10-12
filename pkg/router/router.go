package router

import "github.com/gin-gonic/gin"

var r *gin.Engine

func Factory() interface{} {
	r = gin.Default()
	return r
}
