package service

import (
    "encoding/json"
    "github.com/tsbxmw/datasource/common"
    "github.com/tsbxmw/datasource/data/models"
    "time"
)

func (ds *DataSourceService) TaskInit(req *TaskInitRequest) *TaskInitResponse {
    var (
        err error
        res = TaskInitResponse{}
    )
    taskModel := models.TaskModel{}
    if err = common.DB.Table(taskModel.TableName()).Where("user_id=? and name=?", req.UserId, req.TaskName).First(&taskModel).Error; err != nil {
        if err.Error() != "record not found" {
            common.LogrusLogger.Error(err)
            common.InitKey(ds.Ctx)
            ds.Ctx.Keys["code"] = common.MYSQL_QUERY_ERROR
            panic(err)
        }
    }

    common.LogrusLogger.Error("TASK init")
    if taskModel.ID > 0 {
        res.TaskId = taskModel.ID
        res.TaskName = taskModel.Name
    } else {
        taskModel.UserId = req.UserId
        taskModel.Name = req.TaskName
        taskModel.SdkVersion = req.SdkVersion
        taskModel.CreationTime = time.Now()
        taskModel.ModifiedTime = time.Now()
        if err = common.DB.Table(taskModel.TableName()).Create(&taskModel).Error; err != nil {
            common.DB.Rollback()
            common.LogrusLogger.Error(err)
            panic(err)
        }
        taskUserModel := models.TaskUserModel{
            TaskId: taskModel.ID,
            UserId: taskModel.UserId,
            BaseModel: common.BaseModel{
                CreationTime: taskModel.CreationTime,
                ModifiedTime: taskModel.ModifiedTime,
            },
            Remark: "",
        }
        if err = common.DB.Table(taskUserModel.TableName()).Create(&taskUserModel).Error; err != nil {
            common.LogrusLogger.Error(err)
            panic(err)
        }
        res.TaskId = taskModel.ID
        res.TaskName = taskModel.Name
    }
    return &res
}

func (ds *DataSourceService) TaskGetList(req *TaskGetListRequest) *[]TaskGetListResponse {
    var (
        err error
        res = []TaskGetListResponse{}
    )
    taskList := []models.TaskModel{}
    if err = common.DB.Table(models.TaskModel{}.TableName()).Where("user_id=?", req.UserId).Limit(req.PageSize).Offset((req.PageIndex - 1) * req.PageSize).Find(&taskList).Error; err != nil {
        common.LogrusLogger.Error(err)
        panic(err)
    }
    for _, value := range taskList {
        temp, _ := json.Marshal(value)
        common.LogrusLogger.Info(string(temp))
        deviceInfo := models.DeviceModel{}
        appInfo := models.AppModel{}

        if err = common.DB.Table(deviceInfo.TableName()).Where("task_id=?", value.ID).First(&deviceInfo).Error; err != nil {
            if err.Error() != "record not found" {
                common.LogrusLogger.Error(err)
                panic(err)
            }
        }

        if err = common.DB.Table(appInfo.TableName()).Where("task_id=?", value.ID).First(&appInfo).Error; err != nil {
            if err.Error() != "record not found" {
                common.LogrusLogger.Error(err)
                panic(err)
            }
        }

        res_one := TaskGetListResponse{
            Name:       value.Name,
            DeviceName: deviceInfo.Name,
            AppName:    appInfo.Name,
            AppPackage: appInfo.Package,
            AppPicture: appInfo.Extention,
            AppVersion: appInfo.Version,
            CreatorId:  value.UserId,
            UploadTime: value.CreationTime,
            SDKVersion: value.SdkVersion,
        }
        res = append(res, res_one)
    }
    return &res
}

func (ds *DataSourceService) TaskGetDetail(req *TaskGetDetailRequest) *TaskGetDetailResponse {
    var (
        err error
        res = TaskGetDetailResponse{}
    )

    return &res
}