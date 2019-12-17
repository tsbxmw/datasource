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
	c.JSON(common.HTTP_STATUS_OK, common.Response{
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
	if err = c.ShouldBindJSON(&taskReq); err != nil {
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
	c.JSON(common.HTTP_STATUS_OK, common.Response{
		Code:    common.HTTP_STATUS_OK,
		Message: "success",
		Data:    taskRes,
	})
}

func TaskGetReport(c *gin.Context) {
	common.LogrusLogger.Info("Task Get Report")
}

func TaskGetDetail(c *gin.Context) {
	common.LogrusLogger.Info("Task Get Detail")
	var (
		err error
	)
	taskDetailReq := service.TaskGetDetailRequest{}
	if err = c.ShouldBindJSON(&taskDetailReq); err != nil {
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

	c.JSON(common.HTTP_STATUS_OK, common.Response{
		Code:    common.HTTP_STATUS_OK,
		Message: "success",
		Data:    &res,
	},
	)

}
