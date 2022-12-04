package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func main() {
	// 初始化一个http服务对象
	router := gin.Default()

	// 设置一个GET请求的路由，url: '/ping'， 控制器函数： 闭包
	router.GET("/ping", func(c *gin.Context) {
		// 通过请求上下文对象Context，返回json
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.POST("/upload", func(c *gin.Context) {
		// Single file
		file, _ := c.FormFile("file")
		log.Println(file.Filename)

		// Upload the file to specific dst.
		filePath, _ := os.Getwd()
		dst := filePath + "/" + file.Filename
		c.SaveUploadedFile(file, dst)

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
	router.POST("/multi-upload", func(c *gin.Context) {
		// Multipart form
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]

		for _, file := range files {
			log.Println(file.Filename)
			filePath, _ := os.Getwd()
			dst := filePath + "/" + file.Filename

			// Upload the file to specific dst.
			c.SaveUploadedFile(file, dst)
		}
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	})

	// 监听，并在 localhost:8080上启动服务
	router.Run()
	//router.Run(":8080")

}
