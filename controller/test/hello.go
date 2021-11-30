package test

import (
	"GoWild/service/helloService"
	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	helloService.Hello()
}
