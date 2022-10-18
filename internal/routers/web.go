package routers

import "github.com/gin-gonic/gin"

func LoadWeb(engine *gin.Engine) {
	engine.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
