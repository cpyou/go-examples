package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
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

	parameterInPath := router.Group("/param-in-path")
	{
		// This handler will match /user/john but will not match /user/ or /user
		parameterInPath.GET("/user/:name", userName)
		// However, this one will match /user/john/ and also /user/john/send
		// If no other routers match /user/john, it will redirect to /user/john/
		parameterInPath.GET("/user/:name/*action", userNameAction)

		// For each matched request Context will hold the route definition
		parameterInPath.POST("/user/:name/*action", func(c *gin.Context) {
			b := c.FullPath() == "/param-in-path/user/:name/*action" // true
			c.String(http.StatusOK, "%t", b)
		})

		// This handler will add a new router for /user/groups.
		// Exact routes are resolved before param routes, regardless of the order they were defined.
		// Routes starting with /user/groups are never interpreted as /user/:name/... routes
		parameterInPath.GET("/user/groups", func(c *gin.Context) {
			c.String(http.StatusOK, "The available groups are [...]")
		})
	}

	// Simple group: file
	file := router.Group("/file")
	{
		// Set a lower memory limit for multipart forms (default is 32 MiB)
		router.MaxMultipartMemory = 8 << 20 // 8 MiB
		file.POST("/upload", upload)
		file.POST("/multi-upload", multiUpload)
	}

	model := router.Group("/model")
	{
		// Example for binding JSON ({"user": "manu", "password": "123"})
		model.POST("/loginJSON", loginJSON)

		// Example for binding XML (
		//  <?xml version="1.0" encoding="UTF-8"?>
		//  <root>
		//    <user>manu</user>
		//    <password>123</password>
		//  </root>)
		router.POST("/loginXML", loginXML)

		// Example for binding a HTML form (user=manu&password=123)
		router.POST("/loginForm", loginForm)
	}

	// 监听，并在 localhost:8080上启动服务
	router.Run()
	//router.Run(":8080")

}
