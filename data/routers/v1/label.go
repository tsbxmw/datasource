package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/tsbxmw/datasource/common"
	"github.com/tsbxmw/datasource/data/service"
)

func LabelInit(c *gin.Context) {
	common.LogrusLogger.Info("Label Init")
	var (
		//taskName string
		//sdkVersion string
		//userId int
		err error
	)
	label := service.LabelInitRequest{}
	if err := c.ShouldBindJSON(&label); err != nil {
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

	labelRes := ds.LabelInit(&label)
	c.JSON(200, common.Response{
		Code:    200,
		Message: "success",
		Data:    labelRes,
	})
}

func LabelGetDetailById(c *gin.Context) {
	common.LogrusLogger.Info("Label Get Detail")
	var (
		err error
	)
	labelReq := service.LabelGetDetailRequest{}
	if err := c.ShouldBind(&labelReq); err != nil {
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

	labelRes := ds.LabelGetDetail(&labelReq)
	c.JSON(200, common.Response{
		Code:    200,
		Message: "success",
		Data:    labelRes,
	})
}

func LabelGetByTaskId(c *gin.Context) {
	common.LogrusLogger.Info("Label Get Detail")
	var (
		err error
	)
	labelReq := service.LabelGetListByTaskIdRequest{}
	if err = c.ShouldBind(&labelReq); err != nil {
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
		panic(err)
	}
	labelRes := ds.LabelGetListByTaskId(&labelReq)
	c.JSON(common.HTTP_STATUS_OK, common.Response{
		Code:    common.HTTP_STATUS_OK,
		Message: common.HTTP_MESSAGE_OK,
		Data:    &labelRes,
	})
}
