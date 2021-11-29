package nsq

import (
	"GoWild/service/nsqService"
	"github.com/gin-gonic/gin"
)

func NsqConsumer(c *gin.Context) {
	msg := c.Query("msg")
	nsqService.ProducerMsg(msg)

}
