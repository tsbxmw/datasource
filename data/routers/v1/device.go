package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/tsbxmw/datasource/common"
	"github.com/tsbxmw/datasource/data/service"
)

func DeviceInit(c *gin.Context) {
	deviceReq := service.DeviceInitRequest{}
	var err error

	if err := c.ShouldBindJSON(&deviceReq); err != nil {
		common.LogrusLogger.Error(err)
		common.InitKey(c)
		c.Keys["code"] = common.HTTP_PARAMS_ERROR
		panic(err)
	}
	var (
		ds *service.DataSourceService
	)
	ds, err = service.NewDataSourceMgr(c)
	if err != nil {
		common.LogrusLogger.Error(err)
		panic(err)
	}

	res := ds.DeviceInit(&deviceReq)
	c.JSON(200, gin.H{
		"code": common.HTTP_STATUS_OK,
		"msg":  "success",
		"data": &res,
	})
}

func DeviceGetById(c *gin.Context) {
	deviceReq := service.DeviceGetByIdRequest{}
	var err error
	if err = c.ShouldBind(&deviceReq); err != nil {
		common.LogrusLogger.Error(err)
		common.InitKey(c)
		c.Keys["code"] = common.HTTP_PARAMS_ERROR
		panic(err)
	}
	var (
		ds *service.DataSourceService
	)

	ds, err = service.NewDataSourceMgr(c)
	if err != nil {
		common.LogrusLogger.Error(err)
		panic(err)
	}
	res := ds.DeviceGetById(&deviceReq)
	c.JSON(common.HTTP_STATUS_OK, common.Response{
		Code:    common.HTTP_STATUS_OK,
		Message: "success",
		Data:    &res,
	})
}
