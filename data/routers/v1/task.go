package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/tsbxmw/datasource/common"
	"github.com/tsbxmw/datasource/data/service"
)

func TaskInit(c *gin.Context) {
	common.LogrusLogger.Info("Task Init")
	var (
		err error
	)
	task := service.TaskInitRequest{}
	if err := c.ShouldBindJSON(&task); err != nil {
		common.LogrusLogger.Error(err)
		common.InitKey(c)
		c.Keys["code"] = common.HTTP_PARAMS_ERROR
		panic(err)
	}
	authGlobal := c.Keys["auth"].(*common.AuthGlobal)
	task.UserId = authGlobal.UserId
	task.SdkVersion = c.Keys["tcsdk_version"].(string)
	var (
		ds *service.DataSourceService
	)
	ds, err = service.NewDataSourceMgr(c)
	if err != nil {
		common.LogrusLogger.Error(err)
		panic(err)
	}
	taskRes := ds.TaskInit(&task)
	c.JSON(common.HTTP_RESPONSE_OK, common.Response{
		Code:    common.HTTP_STATUS_OK,
		Message: "success",
		Data:    taskRes,
	})
}

func TaskGetList(c *gin.Context) {
	common.LogrusLogger.Info("Task Get List")
	var (
		err error
	)
	taskReq := service.TaskGetListRequest{}
	if err = c.ShouldBind(&taskReq); err != nil {
		common.LogrusLogger.Error(err)
		common.InitKey(c)
		c.Keys["code"] = common.HTTP_PARAMS_ERROR
		panic(err)
	}

	var ds *service.DataSourceService

	ds, err = service.NewDataSourceMgr(c)
	if err != nil {
		common.LogrusLogger.Error(err)
		panic(err)
	}
	taskRes := ds.TaskGetList(&taskReq)
	c.JSON(common.HTTP_RESPONSE_OK, common.Response{
		Code:    common.HTTP_STATUS_OK,
		Message: "success",
		Data:    taskRes,
	})
}

func TaskGetReport(c *gin.Context) {
	common.LogrusLogger.Info("Task Get Report")
	var (
		err error
	)
	taskReportReq := service.TaskGetReportRequest{}

	if err = c.ShouldBind(&taskReportReq); err != nil {
		common.LogrusLogger.Error(err)
		common.InitKey(c)
		c.Keys["code"] = common.HTTP_PARAMS_ERROR
		panic(err)
	}
	var ds *service.DataSourceService

	ds, err = service.NewDataSourceMgr(c)
	if err != nil {
		common.LogrusLogger.Error(err)
		panic(err)
	}

	res := ds.TaskGetReort(&taskReportReq)

	c.JSON(common.HTTP_RESPONSE_OK, common.Response{
		Code:    common.HTTP_STATUS_OK,
		Message: "success",
		Data:    &res,
	})
}

func TaskGetDetail(c *gin.Context) {
	common.LogrusLogger.Info("Task Get Detail")
	var (
		err error
	)
	taskDetailReq := service.TaskGetDetailRequest{}
	if err = c.ShouldBind(&taskDetailReq); err != nil {
		common.LogrusLogger.Error(err)
		common.InitKey(c)
		c.Keys["code"] = common.HTTP_PARAMS_ERROR
		panic(err)
	}

	var ds *service.DataSourceService

	ds, err = service.NewDataSourceMgr(c)
	if err != nil {
		common.LogrusLogger.Error(err)
		panic(err)
	}

	res := ds.TaskGetDetail(&taskDetailReq)

	c.JSON(common.HTTP_RESPONSE_OK, common.Response{
		Code:    common.HTTP_STATUS_OK,
		Message: "success",
		Data:    &res,
	})
}

func TaskCalSummary(c *gin.Context) {
	common.LogrusLogger.Info("Task Cal Summary")
	var (
		err error
	)
	req := service.TaskCalSummaryRequest{}
	if err = c.ShouldBindJSON(&req); err != nil {
		common.LogrusLogger.Error(err)
		common.InitKey(c)
		c.Keys["code"] = common.HTTP_PARAMS_ERROR
		panic(err)
	}
	var ds *service.DataSourceService

	ds, err = service.NewDataSourceMgr(c)
	if err != nil {
		common.LogrusLogger.Error(err)
		panic(err)
	}

	res := ds.TaskCalSummary(&req)

	c.JSON(common.HTTP_RESPONSE_OK, common.Response{
		Code:    common.HTTP_STATUS_OK,
		Message: "success",
		Data:    &res,
	})
}
