package v1

import (
	"CommonConfig/model/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AdminApi struct{}

func (*AdminApi) TestAdmin(c *gin.Context) {
	c.JSON(
		http.StatusOK, response.Result{
			Code: response.SUCCESS,
			Data: nil,
			Msg:  "",
		},
	)
}
