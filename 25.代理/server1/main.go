package main

import (
	"github.com/gin-gonic/gin"
)
type User struct {
	Name string `json:"name"`
}
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code":200,
			"msg":"PONG",
			"server": "server1",
		})
	})
	r.POST("/user", func(c *gin.Context) {
		user := User{}
		if err := c.ShouldBindJSON(&user);err != nil {
			c.JSON(400, gin.H{
				"err": err,
				"server": "server1",
			})
			return
		}
		c.JSON(200, gin.H{
			"server": "server1",
			"code":200,
			"name": user.Name,
		})
	})
	r.Run(":7001")
}
