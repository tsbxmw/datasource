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
	common.LogrusLogger.Info(req.Name, req.Version, req.Package)
	app := models.AppModel{}
	if err = common.DB.Table(app.TableName()).Where("name=? and version=? and package=? and task_id=?", req.Name, req.Version, req.Package, req.TaskId).First(&app).Error; err != nil {
		if err.Error() != "record not found" {
			common.LogrusLogger.Error(err)
			common.InitKey(ds.Ctx)
			ds.Ctx.Keys["code"] = common.MYSQL_QUERY_ERROR
			panic(err)
		}
	}
	common.LogrusLogger.Info(app.TaskId)
	if app.ID > 0 {
		res.AppId = app.ID
		res.AppName = app.Name
	} else {
		app.Name = req.Name
		app.Version = req.Version
		app.Extention = req.Extention
		app.Remark = req.Remark
		app.CreationTime = time.Now()
		app.Package = req.Package
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

func (ds *DataSourceService) AppGetById(req *AppGetByIdRequest) *AppGetResponse {
	var (
		err error
		res = AppGetResponse{
			AppModel: models.AppModel{},
		}
	)

	app := models.AppModel{}

	if req.TaskId > 0 && req.AppId <= 0 {
		if err = common.DB.Table(app.TableName()).Where("task_id=?", req.TaskId).First(&app).Error; err != nil {
			panic(err)
		}
	} else if req.AppId > 0 {
		if err = common.DB.Table(app.TableName()).Where("id=?", req.AppId).First(&app).Error; err != nil {
			panic(err)
		}
	}

	res.AppModel = app
	return &res
}
