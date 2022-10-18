package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func LoadApi(engine *gin.Engine) {
	api := engine.Group("/api/")
	addUserRoutes(api)
	addPostRoutes(api)

}

// addUserRoutes adds the user routes to the router.
func addUserRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")
	users.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "users")
	})

	users.GET("/1", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"success": 1,
			"data": gin.H{
				"id":   1,
				"name": "friday",
				"time": time.Now(),
				"path": c.Request.URL.Path,
			},
		})
	})
}

// addUserRoutes adds the user routes to the router.
func addPostRoutes(rg *gin.RouterGroup) {
	posts := rg.Group("/posts")

	posts.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "posts")
	})
}
