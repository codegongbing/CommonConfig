package v1

import (
	"CommonConfig/model/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseApi struct{}

// @Tags Test
// @Summary 测试test
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Router /base/test [get]
func (*BaseApi) Test(c *gin.Context) {
	c.JSON(
		http.StatusOK, response.Result{
			Code: response.SUCCESS,
			Data: nil,
			Msg:  "成功",
		},
	)
}

func (*BaseApi) Test2(c *gin.Context) {

	c.JSON(
		http.StatusOK, response.Result{
			Code: response.SUCCESS,
			Data: nil,
			Msg:  "成功",
		},
	)
}
