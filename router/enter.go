package router

type RouterGroup struct {
	InitRouter
	AuthorityRouter
	BaseRouter
}

var RouterGroupApp = new(RouterGroup)