/**
*   @Author: yky
*   @File: cors
*   @Version: 1.0
*   @Date: 2021-07-27 21:41
 */
package base

import "github.com/gin-gonic/gin"

func CORS(c *gin.Context) {
	origin := c.Request.Header.Get("Origin")
	if origin != "" {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "HEAD, POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
	}

	c.Next()
}
