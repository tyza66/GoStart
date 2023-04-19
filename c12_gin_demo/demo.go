package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
	"net/http"
)

func main() {
	//创建一个gin服务
	ginServer := gin.Default()
	ginServer.Use(favicon.New("c12_gin_demo/favicon.ico"))

	//加载静态界面
	ginServer.LoadHTMLGlob("c12_gin_demo/templates/*")
	//访问地址 处理请求
	ginServer.GET("/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{"msg": "giao"}) //返回json数据
	})

	ginServer.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK,"index.html", gin.H{
			"msg": "Hello world",
		})
	})

	ginServer.Run("localhost:8080")
}

/*
	Restful Api就是做不同事情用不同的请求 gin支持
*/
