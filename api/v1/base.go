package v1

import (
	"CommonConfig/model/response"
	"CommonConfig/utils/upload"
	"fmt"
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

	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.JSON(
			http.StatusInternalServerError, gin.H{
				"msg": "error",
			})
	}
	oss := upload.NewOSS()
	fileName, err := oss.UploadFile(fileHeader)
	if err != nil {
		c.JSON(
			http.StatusOK, response.Result{
				Code: response.ERROR,
				Data: nil,
				Msg:  "上传图片失败",
			})
	}

	fmt.Println(fileName)

	c.JSON(
		http.StatusOK, response.Result{
			Code: response.SUCCESS,
			Data: nil,
			Msg:  "成功",
		},
	)
}
