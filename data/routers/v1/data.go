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

}


func TaskInit(c *gin.Context) {
    common.LogrusLogger.Info("Task Init")
    var (
        //taskName string
        //sdkVersion string
        //userId int
        err error
    )
    task := service.TaskInitRequest{}
    if err:=c.ShouldBindJSON(&task); err!=nil{
        common.LogrusLogger.Error(err)
        common.InitKey(c)
        c.Keys["code"] = common.HTTP_PARAMS_ERROR
        panic(err)
    }
    //taskName = c.PostForm("task_name")
    //sdkVersion = c.PostForm("sdk_version")
    //if taskName == "" {
    //    panic(errors.New("task_name should not be null"))
    //}
    authGlobal := c.Keys["auth"].(*common.AuthGlobal)
    task.UserId = authGlobal.UserId

    var (
        ds  *service.DataSourceService
    )
    ds, err = service.NewDataSourceMgr(c)
    if err != nil {
        common.LogrusLogger.Error(err)
        panic(err)
    }
    taskRes := ds.TaskInit(&task)
    c.JSON(200, common.Response{
        Code: 200,
        Message: "success",
        Data: taskRes,
    })
}

func LabelInit(c *gin.Context) {
    common.LogrusLogger.Info("Label Init")
    var (
        //taskName string
        //sdkVersion string
        //userId int
        err error
    )
    label := service.LabelInitRequest{}
    if err:=c.ShouldBindJSON(&label); err!=nil{
        common.LogrusLogger.Error(err)
        common.InitKey(c)
        c.Keys["code"] = common.HTTP_PARAMS_ERROR
        panic(err)
    }
    //taskName = c.PostForm("task_name")
    //sdkVersion = c.PostForm("sdk_version")
    //if taskName == "" {
    //    panic(errors.New("task_name should not be null"))
    //}

    var (
        ds  *service.DataSourceService
    )
    ds, err = service.NewDataSourceMgr(c)
    if err != nil {
        common.LogrusLogger.Error(err)
        panic(err)
    }

    labelRes := ds.LabelInit(&label)
    c.JSON(200, common.Response{
        Code: 200,
        Message: "success",
        Data: labelRes,
    })
}

