/**
*   @Author: yky
*   @File: route
*   @Version: 1.0
*   @Date: 2021-07-14 21:59
 */
package route

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Route() *gin.Engine {
	r := gin.New()

	r.GET("/someJson", func(context *gin.Context) {
		data := map[string]interface{}{
			"lang": "GO",
			"tag":  "<br>",
		}
		context.AsciiJSON(http.StatusOK, data)
	})

	r.GET("/jsonP", func(context *gin.Context) {
		data := map[string]interface{}{
			"data": 123,
		}
		context.JSONP(http.StatusOK, data)
	})

	//上传单文件到服务器
	/*
		curl -X PUT http://192.168.23.1:8080/upload \-F "file=@/opt/test.txt" \-H "Content-Type: multipart/form-data"
	*/
	r.PUT("upload", func(context *gin.Context) {
		file, _ := context.FormFile("file")
		log.Println(file.Filename)
		context.SaveUploadedFile(file, "./"+file.Filename)
		context.JSON(http.StatusOK, "ok")
	})

	return r
}
