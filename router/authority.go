package router

import "github.com/gin-gonic/gin"

type AuthorityRouter struct{}

func (AuthorityRouter) AuthorityRouterGroup(r *gin.RouterGroup) {
	authorityGroup := r.Group("authority")
	{
		authorityGroup.GET("")
	}
}
