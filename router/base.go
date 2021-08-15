package router

import (
	v1 "CommonConfig/api/v1"
	"github.com/gin-gonic/gin"
)

type BaseRouter struct{}

func (*BaseRouter) BaseRouterGroup(r *gin.RouterGroup) {
	baseGroup := r.Group("base")
	var baseApi = v1.ApiGroupApp.BaseApi
	{
		baseGroup.POST("/test", baseApi.Test)
		baseGroup.GET("", baseApi.Test2)
	}
}
