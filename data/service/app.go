package service

import (
    "github.com/tsbxmw/datasource/common"
    "github.com/tsbxmw/datasource/data/models"
    "time"
)

func (ds *DataSourceService) AppInit(req *AppInitRequest) *AppInitResponse {
    var (
        err error
        res = AppInitResponse{}
    )
    app := models.AppModel{}
    if err = common.DB.Table(app.TableName()).Where("name=? and version=? and package=?", req.Name, req.Version, req.Package).First(&app).Error; err != nil {
        if err.Error() != "record not found" {
            common.LogrusLogger.Error(err)
            common.InitKey(ds.Ctx)
            ds.Ctx.Keys["code"] = common.MYSQL_QUERY_ERROR
            panic(err)
        }
    }
    if app.ID > 0 {
        res.AppId = app.ID
        res.AppName = app.Name
    } else {
        app.Name = req.Name
        app.Version = req.Version
        app.Extention = req.Extention
        app.Remark = req.Remark
        app.CreationTime = time.Now()
        app.TaskId = req.TaskId
        if err = common.DB.Table(app.TableName()).Create(&app).Error; err != nil {
            common.LogrusLogger.Error(err)
            panic(err)
        }
        res.AppId = app.ID
        res.AppName = app.Name
    }
    return &res
}