package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/tsbxmw/datasource/common"
)

func HealthCheck(c *gin.Context) {
	c.JSON(common.HTTP_STATUS_OK, common.Response{
		Code:    common.HTTP_RESPONSE_OK,
		Message: common.HTTP_MESSAGE_OK,
		Data:    []string{},
	})
}
