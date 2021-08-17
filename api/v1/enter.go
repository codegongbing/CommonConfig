package v1

type ApiGroup struct {
	AdminApi
	AuthorityApi
	BaseApi
}

var ApiGroupApp = new(ApiGroup)
