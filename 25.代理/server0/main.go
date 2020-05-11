package main

import (
	"time"

	"github.com/gin-gonic/gin"
)
type User struct {
	Name string `json:"name"`
}
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"server": "server0",
			"time": time.Now(),
		})
	})
	r.POST("/user", func(c *gin.Context) {
		user := User{}
		if err := c.ShouldBindJSON(&user);err != nil {
			c.JSON(400, gin.H{
				"err": err,
				"server": "server0",
			})
			return
		}
		c.JSON(200, gin.H{
			"code":200,
			"name": user.Name,
			"server": "server0",
		})
	})
	r.Run(":7000")
}
