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

func (ds *DataSourceService) TaskGet(taskId int) *TaskGetListResponse {

	task := models.TaskModel{}
	taskInfo := TaskGetListResponse{}

	if err := common.DB.Table(task.TableName()).Where("id=?", taskId).First(&task).Error; err != nil {
		if err.Error() != "record not found" {
			common.LogrusLogger.Error(err)
			panic(err)
		} else {
			return &taskInfo
		}
	}

	deviceInfo := models.DeviceModel{}
	appInfo := models.AppModel{}

	if err := common.DB.Table(deviceInfo.TableName()).Where("task_id=?", taskId).First(&deviceInfo).Error; err != nil {
		if err.Error() != "record not found" {
			common.LogrusLogger.Error(err)
			panic(err)
		}
	}

	if err := common.DB.Table(appInfo.TableName()).Where("task_id=?", taskId).First(&appInfo).Error; err != nil {
		if err.Error() != "record not found" {
			common.LogrusLogger.Error(err)
			panic(err)
		}
	}

	taskInfo = TaskGetListResponse{
		Name:       task.Name,
		DeviceName: deviceInfo.Name,
		AppName:    appInfo.Name,
		AppPackage: appInfo.Package,
		AppPicture: appInfo.Extention,
		AppVersion: appInfo.Version,
		CreatorId:  task.UserId,
		UploadTime: task.CreationTime,
		SDKVersion: task.SdkVersion,
	}

	return &taskInfo
}

func (ds *DataSourceService) TaskGetList(req *TaskGetListRequest) *[]TaskGetListResponse {
	var (
		err error
		res = make([]TaskGetListResponse, 0)
	)
	taskList := make([]models.TaskModel, 0)
	if err = common.DB.Table(models.TaskModel{}.TableName()).Where("user_id=?", req.UserId).Limit(req.PageSize).Offset((req.PageIndex - 1) * req.PageSize).Find(&taskList).Error; err != nil {
		common.LogrusLogger.Error(err)
		panic(err)
	}
	common.LogrusLogger.Debug(req)
	for _, value := range taskList {
		temp, _ := json.Marshal(value)
		common.LogrusLogger.Info(string(temp))
		resOne := ds.TaskGet(value.ID)
		res = append(res, *resOne)
	}
	return &res
}

func (ds *DataSourceService) TaskGetDetail(req *TaskGetDetailRequest) *TaskGetDetailResponse {
	var (
		res = TaskGetDetailResponse{
			TaskSummary: models.TaskSummaryModel{},
			TaskDetail:  TaskGetListResponse{},
			LabelInfos:  LabelGetListResponse{},
		}
	)
	if err := common.DB.Table(res.TaskSummary.TableName()).Where("task_id=?", req.TaskId).First(&res.TaskSummary).Error; err != nil {
		if err.Error() != "record not found" {
			common.LogrusLogger.Error(err)
			panic(err)
		}
	}

	res.TaskDetail = *ds.TaskGet(req.TaskId)
	res.LabelInfos = *ds.LabelGetListByTaskId(&LabelGetListByTaskIdRequest{TaskId: req.TaskId})

	return &res
}

func (ds *DataSourceService) TaskGetReort(req *TaskGetReportRequest) *TaskGetDetailResponse {
	var (
		res = TaskGetDetailResponse{
			TaskSummary: models.TaskSummaryModel{},
			TaskDetail:  TaskGetListResponse{},
		}
	)
	if err := common.DB.Table(res.TaskSummary.TableName()).Where("task_id=?", req.TaskId).First(&res.TaskSummary).Error; err != nil {
		if err.Error() != "record not found" {
			common.LogrusLogger.Error(err)
			panic(err)
		}
	}

	res.TaskDetail = *ds.TaskGet(req.TaskId)

	return &res
}


// Task Get Summary by Task id
func (ds *DataSourceService) TaskSummaryGet(taskId int) *models.TaskSummaryModel {
	var (
		err          error
		taskSummary = models.TaskSummaryModel{}
	)

	ds.TaskSummaryInit(taskId)
	if err = common.DB.Table(taskSummary.TableName()).Where("task_id=?", taskId).Find(&taskSummary).Error; err != nil {
		panic(err)
	}
	return &taskSummary
}

// Task Summary Init
func (ds *DataSourceService) TaskSummaryInit(taskId int) {
	var (
		err          error
		taskSummary = models.TaskSummaryModel{}
	)
	if err = common.DB.Table(taskSummary.TableName()).Where("task_id=?", taskId).Find(&taskSummary).Error; err != nil {
		if err.Error() != "record not found" {
			common.LogrusLogger.Error(err)
			common.InitKey(ds.Ctx)
			ds.Ctx.Keys["code"] = common.MYSQL_QUERY_ERROR
			panic(err)
		} else {
			taskSummary.TaskId = taskId
			taskSummary.CreationTime = time.Now()
			taskSummary.ModifiedTime = time.Now()
			if err = common.DB.Table(taskSummary.TableName()).Create(&taskSummary).Error; err != nil {
				panic(err)
			}
		}
	}
}


func (ds *DataSourceService) TaskSummaryUpdate(taskId int, taskSummary models.TaskSummaryModel) {

	var (
		err error
	)
	taskCurrent := ds.TaskSummaryGet(taskId)
	taskSummary.TaskId = taskCurrent.TaskId
	taskSummary.ID = taskCurrent.ID
	taskSummary.CreationTime = taskCurrent.CreationTime
	taskSummary.ModifiedTime = time.Now()
	if err = common.DB.Model(&taskCurrent).Update(&taskSummary).Error; err != nil {
		common.LogrusLogger.Error(err)
		common.InitKey(ds.Ctx)
		ds.Ctx.Keys["code"] = common.MYSQL_UPDATE_ERROR
		panic(err)
	}
}


func (ds *DataSourceService) TaskCalSummary(req *TaskCalSummaryRequest) *TaskCalSummaryResponse {
	var (
		res = TaskCalSummaryResponse{}
		labels = ds.LabelGetListByTaskId(&LabelGetListByTaskIdRequest{TaskId: req.TaskId}).Label
		labelCount = float32(len(labels))
		taskSum = models.TaskSummaryModel{}
	)

	for _, labelInfo := range labels {
		ds.CalLabelSummary(&LabelCalSummaryRequest{LabelId: labelInfo.Info.ID, TaskId: labelInfo.Info.TaskId})
		label := labelInfo.Summary
		taskSum.BatteryCurrentAvg += label.BatteryCurrentAvg
		taskSum.BatteryPowerAvg += label.BatteryPowerAvg
		taskSum.BatteryVoltageAvg += label.BatteryVoltageAvg

		taskSum.FpsAvg += label.FpsAvg

		taskSum.CpuAppAvg += label.CpuAppAvg
		taskSum.CpuAvg += label.CpuAvg

		taskSum.GpuDeviceAvg += label.GpuDeviceAvg
		taskSum.GpuTilerAvg += label.GpuTilerAvg
		taskSum.GpuRenderAvg += label.GpuRenderAvg

		taskSum.NetRecvAvg += label.NetRecvAvg
		taskSum.NetSendAvg += label.NetSendAvg
	}

	taskSum.BatteryCurrentAvg += taskSum.BatteryCurrentAvg / labelCount
	taskSum.BatteryPowerAvg += taskSum.BatteryPowerAvg / labelCount
	taskSum.BatteryVoltageAvg += taskSum.BatteryVoltageAvg / labelCount

	taskSum.FpsAvg += taskSum.FpsAvg / labelCount

	taskSum.CpuAppAvg += taskSum.CpuAppAvg / labelCount
	taskSum.CpuAvg += taskSum.CpuAvg / labelCount

	taskSum.GpuDeviceAvg += taskSum.GpuDeviceAvg / labelCount
	taskSum.GpuTilerAvg += taskSum.GpuTilerAvg / labelCount
	taskSum.GpuRenderAvg += taskSum.GpuRenderAvg / labelCount

	taskSum.NetRecvAvg += taskSum.NetRecvAvg / labelCount
	taskSum.NetSendAvg += taskSum.NetSendAvg / labelCount

	ds.TaskSummaryUpdate(req.TaskId, taskSum)

	return &res
}