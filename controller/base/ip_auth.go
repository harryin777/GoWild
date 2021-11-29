package base

import (
	"GoWild/common/ip"
	"GoWild/helper/logger"
	"GoWild/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func IpAuth(c *gin.Context) {
	if !ip.LocationInstances().GetLocation(c.ClientIP()) {
		logger.ErrLogger().WithFields(logrus.Fields{
			"service": utils.GetServiceName(),
			"func":    utils.GetFuncName(),
			"ip":      c.ClientIP(),
		}).Error("foreign ip")
		return
	}
}
