package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log"
	"net/http"
)

// OnGetStringRequest 返回success字符串的接口
func OnGetStringRequest(c *gin.Context) {
	c.String(http.StatusOK, "success")
}
// OnPracticeRequest 返回practice.html页面的接口
// func OnPracticeRequest(c *gin.Context) {
// 	c.HTML(http.StatusOK,"practice.html",gin.H{})
// }
// OnLoginRequestForForm 以表单形式传递参数的登录接口
func OnLoginRequestForForm(c *gin.Context) {
	req := &User{}
	if err := c.ShouldBindWith(req, binding.Form); err != nil {
		log.Printf("err:%v",err)
		c.JSON(http.StatusOK, gin.H{
			"errno":"1",
			"errmsg":"参数不匹配",
			"data":"",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"errno":"0",
		"errmsg":"",
		"data":req,
	})
}
// OnLoginRequestForJson 以Json形式传递参数的登录接口
func OnLoginRequestForJson(c *gin.Context) {
	req := &User{}
	if err := c.ShouldBindWith(req, binding.JSON); err != nil {
		log.Printf("err:%v",err)
		c.JSON(http.StatusOK, gin.H{
			"errno":"1",
			"errmsg":"参数不匹配",
			"data":"",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"errno":"0",
		"errmsg":"",
		"data":req,
	})
}