package v1

import (
    "datasource/common"
    "datasource/data/service"
    "github.com/gin-gonic/gin"
)

func DataInit(c *gin.Context) {
    common.LogrusLogger.Error("test error")
    common.LogrusLogger.Info("test Info")
    var (
        ds *service.DataSourceService
        err error
    )
    ds, err = service.NewDataSourceMgr()
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


func Data