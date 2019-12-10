package v1

import (
    "github.com/gin-gonic/gin"
    "github.com/tsbxmw/datasource/common"
    "github.com/tsbxmw/datasource/data/service"
)

func DataInit(c *gin.Context) {
    common.LogrusLogger.Error("test error")
    common.LogrusLogger.Info("test Info")
    var (
        ds  *service.DataSourceService
        err error
    )
    ds, err = service.NewDataSourceMgr(c)
    if err != nil {
        common.LogrusLogger.Error(err)
        panic(err)
    }
    ds.AuthCheck("", "")
    c.JSON(200, gin.H{
        "v1":       "test",
        "username": "test",
    })
}

func DataUpload(c *gin.Context) {
    common.LogrusLogger.Info("Data Upload")

    var (
        ds  *service.DataSourceService
        err error
    )

    req := service.DataUploadRequest{}
    if err := c.ShouldBindJSON(&req); err != nil {
        common.LogrusLogger.Error(err)
        common.InitKey(c)
        c.Keys["code"] = common.HTTP_PARAMS_ERROR
        panic(err)
    }
    ds, err = service.NewDataSourceMgr(c)
    if err != nil {
        common.LogrusLogger.Error(err)
        panic(err)
    }
    res := ds.DataUpload(&req)

    c.JSON(200, common.Response{
        Code:    200,
        Message: "success",
        Data:    res,
    })
}
