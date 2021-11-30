package nsqController

import (
	"GoWild/service/nsqService"
	"github.com/gin-gonic/gin"
)

func Producer(c *gin.Context) {
	msg := c.Query("msg")
	nsqService.ProducerMsg(msg)
}
