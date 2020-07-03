package main

import "github.com/gin-gonic/gin"

var router *gin.Engine
func init(){
	// 初始化路由
	router = gin.Default()
	router.GET("/getString", OnGetStringRequest)
	router.POST("/loginForm", OnLoginRequestForForm)
	router.POST("/loginJson", OnLoginRequestForJson)
	// router.LoadHTMLGlob("E:/mygo/resources/pages/*") //定义模板文件路径
	// router.GET("/practice", OnPracticeRequest)
}
