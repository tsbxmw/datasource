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
    c.JSON(200, common.Response{
        Code:    200,
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
    c.JSON(200, common.Response{
        Code: 200,
        Message: "success",
        Data: taskRes,
    })
}
