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
	// 为 multipart forms 设置较低的内存限制 (默认是 32 MiB)
	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	r.PUT("upload", func(context *gin.Context) {
		file, _ := context.FormFile("file")
		log.Println(file.Filename)
		context.SaveUploadedFile(file, "./"+file.Filename)
		context.JSON(http.StatusOK, "ok")
	})

	//上传多文件到服务器
	/**
	curl -X PUT http://192.168.23.1:8080/multiUpload \-F "upload[]=@/opt/test.txt" \-F "upload[]=@/opt/test2.txt" \-H "Content-Type: multipart/form-data"
	*/
	r.PUT("/multiUpload", func(context *gin.Context) {
		form, _ := context.MultipartForm()
		for _, file := range form.File["upload[]"] {
			log.Println(file.Filename)
			context.SaveUploadedFile(file, "./"+file.Filename)
		}
		context.JSON(http.StatusOK, "ok")
	})

	//从 reader 读取数据

	return r
}
