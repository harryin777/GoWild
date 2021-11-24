package nsq

import "github.com/gin-gonic/gin"

func Route(group *gin.RouterGroup) {
	g := group.Group("nsq")

	g.PUT("genMsg")

}
