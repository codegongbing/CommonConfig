package initialize

import (
	_ "CommonConfig/docs"
	"CommonConfig/global"
	"CommonConfig/router"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(cor()) //跨域
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.LOG.Info("register swagger handler")

	routerGroup := router.RouterGroupApp
	// 公共路由,不需要鉴权
	publicGroup := r.Group(global.CONFIG.Project.RelativePath)
	{
		routerGroup.BaseRouterGroup(publicGroup)
	}
	//私有路由,需要鉴权
	privateGroup := r.Group(global.CONFIG.Project.RelativePath)
	{
		routerGroup.AuthorityRouterGroup(privateGroup)
	}
	return r
}

func cor() gin.HandlerFunc {
	return cors.New(
		cors.Config{
			//准许跨域请求网站,多个使用,分开,限制使用*
			AllowOrigins: []string{"*"},
			//准许使用的请求方式
			AllowMethods: []string{"PUT", "PATCH", "POST", "GET", "DELETE", "OPTION"},
			//准许使用的请求表头
			AllowHeaders: []string{"Origin", "Authorization", "Content-Type"},
			//显示的请求表头
			ExposeHeaders: []string{"Content-Type"},
			//凭证共享,确定共享
			AllowCredentials: true,
			//容许跨域的原点网站,可以直接return true就万事大吉了
			AllowOriginFunc: func(origin string) bool {
				return true
			},
		},
	)
}
