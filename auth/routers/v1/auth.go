package v1

import (
	"datasource/auth/service"
	"datasource/common"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Token(c *gin.Context) {
	userId := c.Query("user_id")
	if userId == "" {
		c.JSON(200, gin.H{
			"code":    common.HTTP_MISS_PARAMS,
			"message": "missing user_id",
		})
		return
	}
	var (
		userIdInt int
		err       error
	)
	if userIdInt, err = strconv.Atoi(userId); err != nil {
		c.JSON(200, gin.H{
			"code":    common.HTTP_MISS_PARAMS,
			"message": "missing user_id",
		})
		return
	}
	as := service.AuthService{}
	key, secret, err := as.TokenGenerate(userIdInt)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    common.HTTP_INTERNAL_ERROR,
			"message": "error:" + err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"data": gin.H{
			"key":    key,
			"secret": secret,
		},
		"messag": "",
	})
}

func RefreshToken(c *gin.Context) {
	userId := c.Query("user_id")
	if userId == "" {
		c.JSON(200, gin.H{
			"code":    common.HTTP_MISS_PARAMS,
			"message": "missing user_id",
		})
		return
	}
	var (
		userIdInt int
		err       error
	)
	if userIdInt, err = strconv.Atoi(userId); err != nil {
		c.JSON(200, gin.H{
			"code":    common.HTTP_PARAMS_ERROR,
			"message": "missing user_id",
		})
		return
	}
	as := service.AuthService{}
	key, secret, err := as.RefreshToken(userIdInt)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    common.HTTP_INTERNAL_ERROR,
			"message": "error:" + err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"data": gin.H{
			"key":    key,
			"secret": secret,
		},
		"messag": "",
	})

}
