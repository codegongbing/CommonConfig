package router

import (
	v1 "CommonConfig/api/v1"
	"github.com/gin-gonic/gin"
)

type AdminRouter struct{}

func (*AdminRouter) AdminRouterGroup(r *gin.RouterGroup) {
	adminGroup := r.Group("admin")
	adminApi := v1.ApiGroupApp.AdminApi
	{
		adminGroup.GET("", adminApi.TestAdmin)
	}
}
