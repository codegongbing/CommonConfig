package v1

import (
	"CommonConfig/utils/upload"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseApi struct{}

// @Tags aaa
// @Summary 测试test
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Router /base/test [get]
func (*BaseApi) Test(c *gin.Context) {
	_, err := c.FormFile("file")
	if err != nil {
		fmt.Println(err)
	}
	oss := upload.NewOSS()
	oss.DeleteFile("https://codegongbing-blog.oss-cn-beijing.aliyuncs.com/test/2021/8/15/1629024968668922457TgJFESLvGQ.jpg")
	c.JSON(
		http.StatusOK, gin.H{
			"msg": "test",
		},
	)
}

func (*BaseApi) Test2(c *gin.Context) {

	c.JSON(
		http.StatusOK, gin.H{
			"msg": "test2",
		},
	)
}
