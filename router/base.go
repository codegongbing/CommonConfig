package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseRouter struct{}

func (*BaseRouter) BaseRouterGroup(r *gin.RouterGroup) {
	baseGroup := r.Group("base")
	{
		baseGroup.GET(
			"/test", func(c *gin.Context) {
				c.JSON(
					http.StatusOK, gin.H{
						"test": "base",
					},
				)
			},
		)
	}
}
