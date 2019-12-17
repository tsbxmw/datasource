package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/tsbxmw/datasource/common"
	"github.com/tsbxmw/datasource/data/service"
)

func AppInit(c *gin.Context) {
	appReq := service.AppInitRequest{}
	var err error

	if err := c.ShouldBindJSON(&appReq); err != nil {
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

	res := ds.AppInit(&appReq)
	c.JSON(200, gin.H{
		"code": common.HTTP_STATUS_OK,
		"msg":  "success",
		"data": &res,
	})
}

func AppGetById(c *gin.Context) {
	appReq := service.AppGetByIdRequest{}
	var err error

	if err = c.ShouldBind(&appReq); err != nil {
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

	res := ds.AppGetById(&appReq)
	c.JSON(200, gin.H{
		"code": common.HTTP_STATUS_OK,
		"msg":  "success",
		"data": &res,
	})
}


