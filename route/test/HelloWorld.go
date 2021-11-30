package test

import (
	"GoWild/controller/test"
	"github.com/gin-gonic/gin"
)

func Route(group *gin.RouterGroup) {
	g := group.Group("test")

	g.GET("hello", test.Hello)
}
