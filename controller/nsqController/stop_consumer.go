package nsqController

import (
	"GoWild/service/nsqService"
	"github.com/gin-gonic/gin"
)

func StopConsumer(c *gin.Context) {
	nsqService.StopConsumer()
}
