package v1

import (
	"datasource/common"
	"github.com/gin-gonic/gin"
)

func AuthCheck(c *gin.Context) {
	common.LogrusLogger.Error("test error")
	common.LogrusLogger.Info("test Info")
	var username int
	if err := common.DB.Table("user").Where("name=?", "").Count(&username).Error; err != nil {
		panic(err)
	}
	c.JSON(200, gin.H{
		"v1":       "test",
		"username": username,
	})
}
