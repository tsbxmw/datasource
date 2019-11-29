package v1

import (
	"github.com/tsbxmw/datasource/auth/service"
	"github.com/tsbxmw/datasource/common"
	"github.com/gin-gonic/gin"
)

func Token(c *gin.Context) {
	common.InitKey(c)
	tokenReq := service.TokenRequest{}
	if err := c.ShouldBind(&tokenReq); err != nil {
		c.Keys["code"] = common.HTTP_MISS_PARAMS
		panic(err)
	}
	as := service.NewAuthMgr(c)
	res := as.TokenGenerate(&tokenReq)

	c.JSON(200, gin.H{
		"code":   200,
		"data":   &res,
		"messag": "",
	})
}

func RefreshToken(c *gin.Context) {
	common.InitKey(c)
	tokenReq := service.TokenRequest{}
	if err := c.ShouldBind(&tokenReq); err != nil {
		c.Keys["code"] = common.HTTP_MISS_PARAMS
		panic(err)
	}
	as := service.NewAuthMgr(c)
	res := as.RefreshToken(&tokenReq)

	c.JSON(200, gin.H{
		"code":   200,
		"data":   &res,
		"messag": "",
	})

}
