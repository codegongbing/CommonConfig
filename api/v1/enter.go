package v1

type ApiGroup struct {
	AuthorityApi
	BaseApi
}

var ApiGroupApp = new(ApiGroup)
