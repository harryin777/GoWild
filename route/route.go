/**
*   @Author: yky
*   @File: route
*   @Version: 1.0
*   @Date: 2021-07-14 21:59
 */
package route

import (
	"GoWild/base"
	"GoWild/route/nsqRouter"
	"GoWild/route/test"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func Route() *gin.Engine {
	// 禁止日志的颜色，当用gin.New的时候，下面的语句不管用
	gin.DisableConsoleColor()

	r := gin.New()
	r.Use(base.CORS)
	//自定义路由的格式
	//gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
	//	log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	//}

	r.GET("someJson", func(context *gin.Context) {
		data := map[string]interface{}{
			"lang": "GO",
			"tag":  "<br>",
		}
		context.AsciiJSON(http.StatusOK, data)
	})

	r.GET("jsonP", func(context *gin.Context) {
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
	r.PUT("multiUpload", func(context *gin.Context) {
		form, _ := context.MultipartForm()
		for _, file := range form.File["upload[]"] {
			log.Println(file.Filename)
			context.SaveUploadedFile(file, "./"+file.Filename)
		}
		context.JSON(http.StatusOK, "ok")
	})

	//shouldBindQuery 和 bindQuery
	//注意 binding
	type Person struct {
		Name string `json:"name" form:"name" binding:"required"`
		Age  int    `json:"age" form:"age" binding:"required"`
	}
	var p1 Person
	r.POST("shouldBindJSON", func(context *gin.Context) {
		err := context.ShouldBindJSON(&p1)
		if err == nil {
			log.Println("==== shouldBindQuery ====")
			log.Println(p1.Name)
			log.Println(p1.Age)
		} else {
			log.Println(err)
		}
		context.JSON(http.StatusOK, "ok")
	})

	r.POST("bindJson", func(context *gin.Context) {
		if context.BindJSON(&p1) == nil {
			log.Println("==== bindJson ====")
			log.Println(p1.Name)
			log.Println(p1.Age)
		}
		context.JSON(http.StatusOK, "ok")
	})

	type people struct {
		Name string `form:"name"`
		Age  int    `form:"age"`
	}
	var peo1 people
	r.POST("bindQuery", func(context *gin.Context) {
		if context.BindQuery(&peo1) == nil {
			log.Println("==== bindQuery ====")
			log.Println(peo1.Name)
			log.Println(peo1.Age)
		}
		context.JSON(http.StatusOK, "ok")
	})

	// 在中间件中使用 Goroutine
	// 当在中间件或 handler 中启动新的 Goroutine 时，不能使用原始的上下文
	r.GET("long_async", func(context *gin.Context) {
		cCp := context.Copy()
		go func() {
			time.Sleep(5 * time.Second)
			log.Printf("async after 5 sec %v", cCp.Request.URL.Path)
		}()
		context.JSON(http.StatusOK, "ok")
	})

	r.GET("long_sync", func(context *gin.Context) {
		time.Sleep(5 * time.Second)
		log.Printf("sync after 5 sec %v", context.Request.URL.Path)
		context.JSON(http.StatusOK, "ok")
	})

	//注册 nsqConsumer
	nsqRouter.Route(r.Group(""))
	test.Route(r.Group(""))

	return r
}
