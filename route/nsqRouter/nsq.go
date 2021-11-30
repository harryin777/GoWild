package nsqRouter

import (
	"GoWild/controller/nsqController"
	"github.com/gin-gonic/gin"
)

func Route(group *gin.RouterGroup) {
	g := group.Group("nsqConsumer")

	g.PUT("genMsg", nsqController.Producer)
	g.POST("stopConsumer", nsqController.StopConsumer)
}
