package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
	"net/http"
)

//自定义中间件拦截器

func main() {
	//创建一个gin服务
	ginServer := gin.Default()
	ginServer.Use(favicon.New("c12_gin_demo/favicon.ico"))

	//加载静态界面
	ginServer.LoadHTMLGlob("c12_gin_demo/templates/*")
	//加载静态资源
	ginServer.Static("/static", "./c12_gin_demo/static")
	//访问地址 处理请求
	ginServer.GET("/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{"msg": "giao"}) //返回json数据
	})
	//返回页面
	ginServer.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", gin.H{
			"msg": "Hello world",
		})
	})
	//获取传递的参数 get?方式
	ginServer.GET("/get/user", func(context *gin.Context) {
		userid := context.Query("userId")
		username := context.Query("userName")
		context.JSON(http.StatusOK, gin.H{"userid": userid, "username": username}) //返回json数据
	})

	//使用路径的方式传递参数
	ginServer.GET("/get/user/:username/:userid", func(context *gin.Context) {
		userid := context.Param("userid")
		username := context.Param("username")
		context.JSON(http.StatusOK, gin.H{"userid": userid, "username": username}) //返回json数据
	})

	//前端给后端传递json 就是post传对象
	ginServer.POST("/json", func(context *gin.Context) {
		//request.body 这里面返回的是一个切片的
		data, err := context.GetRawData()
		if err != nil {
			return
		}
		var m map[string]interface{}
		//用这个把切片转成map
		err2 := json.Unmarshal(data, &m)
		if err2 != nil {
			return
		}
		context.JSON(http.StatusOK, m)
	})

	//处理表单传来的信息
	ginServer.POST("/user/add", func(context *gin.Context) {
		username := context.PostForm("username")
		password := context.PostForm("password")
		context.JSON(http.StatusOK, gin.H{"username": username, "password": password})
	})

	//路由
	ginServer.GET("/test", func(context *gin.Context) {
		//重定向
		context.Redirect(301, "https://www.baidu.com")
	})

	//没有路由就404
	ginServer.NoRoute(func(context *gin.Context) {
		context.HTML(http.StatusNotFound, "404.html", nil)
	})

	//路由组
	userGroup := ginServer.Group("/user")
	{
		userGroup.GET("/delete")
		userGroup.GET("/update")
	}

	ginServer.Run("localhost:8080")
}

/*
	Restful Api就是做不同事情用不同的请求 gin支持
*/
