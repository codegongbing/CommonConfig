package router

type RouterGroup struct {
	AuthorityRouter
	BaseRouter
}

var RouterGroupApp = new(RouterGroup)