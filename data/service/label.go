package service

import (
	"github.com/tsbxmw/datasource/common"
	"github.com/tsbxmw/datasource/data/models"
	"time"
)

// Label Get Summasry by Label id
func (ds *DataSourceService) LabelGetSummary(labelId int) *models.LabelSummaryModel {
	var (
		err          error
		labelSummary = models.LabelSummaryModel{}
	)

	ds.LabelSummaryInit(labelId)
	if err = common.DB.Table(labelSummary.TableName()).Where("label_id=?", labelId).Find(&labelSummary).Error; err != nil {
		panic(err)
	}
	return &labelSummary
}

// Label Summary Init
func (ds *DataSourceService) LabelSummaryInit(labelId int) {
	var (
		err          error
		labelSummary = models.LabelSummaryModel{}
	)
	if err = common.DB.Table(labelSummary.TableName()).Where("label_id=?", labelId).Find(&labelSummary).Error; err != nil {
		if err.Error() != "record not found" {
			common.LogrusLogger.Error(err)
			common.InitKey(ds.Ctx)
			ds.Ctx.Keys["code"] = common.MYSQL_QUERY_ERROR
			panic(err)
		} else {
			labelSummary.LabelId = labelId
			labelSummary.CreationTime = time.Now()
			labelSummary.ModifiedTime = time.Now()
			labelSummary.FtimeAvg = time.Now()
			labelSummary.FtimeDelta = time.Now()
			labelSummary.FtimeGe_100 = time.Now()
			if err = common.DB.Table(labelSummary.TableName()).Create(&labelSummary).Error; err != nil {
				panic(err)
			}
		}
	}
}

// labelSummary Update
func (ds *DataSourceService) LabelSummaryUpdate(labelId int, labelSummary *models.LabelSummaryModel) {

	var (
		err error
	)
	labelCurrent := ds.LabelGetSummary(labelId)
	labelSummary.LabelId = labelId
	labelSummary.ID = labelCurrent.ID
	labelSummary.CreationTime = labelCurrent.CreationTime
	labelSummary.ModifiedTime = time.Now()
	if err = common.DB.Model(&labelCurrent).Update(&labelSummary).Error; err != nil {
		common.LogrusLogger.Error(err)
		common.InitKey(ds.Ctx)
		ds.Ctx.Keys["code"] = common.MYSQL_UPDATE_ERROR
		panic(err)
	}
}

// Label init handler
func (ds *DataSourceService) LabelInit(req *LabelInitRequest) *LabelInitResponse {
	var (
		err error
		res = LabelInitResponse{}
	)
	labelModel := models.LabelModel{}
	if err = common.DB.Table(labelModel.TableName()).Where("task_id=? and name=?", req.TaskId, req.LabelName).First(&labelModel).Error; err != nil {
		if err.Error() != "record not found" {
			common.LogrusLogger.Error(err)
			common.InitKey(ds.Ctx)
			ds.Ctx.Keys["code"] = common.MYSQL_QUERY_ERROR
			panic(err)
		}
	}
	if labelModel.ID > 0 {
		res.LabelId = labelModel.ID
		res.LabelName = labelModel.Name
	} else {
		labelModel.TaskId = req.TaskId
		labelModel.Name = req.LabelName
		labelModel.CreationTime = time.Now()
		labelModel.ModifiedTime = time.Now()
		if err = common.DB.Table(labelModel.TableName()).Create(&labelModel).Error; err != nil {
			common.LogrusLogger.Error(err)
			panic(err)
		}

		res.LabelId = labelModel.ID
		res.LabelName = labelModel.Name
	}

	return &res
}

// 获取 label 根据 label_id
func (ds *DataSourceService) LabelGetByLabelId(labelId int) *models.LabelModel {
	var (
		err   error
		label = models.LabelModel{}
	)

	if err = common.DB.Table(label.TableName()).Where("id=?", labelId).Find(&label).Error; err != nil {
		panic(err)
	}

	return &label
}

// 获取 label 详细信息
func (ds *DataSourceService) LabelGetDetail(req *LabelGetDetailRequest) *LabelResponse {
	var (
		res = LabelResponse{
			Summary: models.LabelSummaryModel{},
			Info:    models.LabelModel{},
		}
		err error
	)
	if err = common.DB.Table(models.LabelModel{}.TableName()).Where("id=?", req.LabelId).First(&res.Info).Error; err != nil {
		panic(err)
	}
	res.Summary = *ds.LabelGetSummary(req.LabelId)
	ds.CalDateOfLabelBeginEnd(&res.Info)
	return &res
}

// 通过 task_id 获取 label 列表
func (ds *DataSourceService) LabelGetListByTaskId(req *LabelGetListByTaskIdRequest) *LabelGetListResponse {
	var (
		err error
		res = LabelGetListResponse{
			Label: make([]LabelResponse, 0),
		}
	)
	labels := make([]models.LabelModel, 0)
	if err = common.DB.Table(models.LabelModel{}.TableName()).Where("task_id=?", req.TaskId).Find(&labels).Error; err != nil {
		panic(err)
	}
	for _, label := range labels {
		resOne := LabelResponse{
			Info:    label,
			Summary: *ds.LabelGetSummary(label.ID),
		}
		res.Label = append(res.Label, resOne)
	}
	return &res
}

// 计算 label 包含信息的开始和结束 ID
func (ds *DataSourceService) CalDateOfLabelBeginEnd(label *models.LabelModel) {

	var (
		err            error
		dataTableIndex = common.GetDBIndex(label.TaskId)
		dataFirst      = models.DataUploadModel{}
		dataLast       = models.DataUploadModel{}
	)
	if err = common.DB.Table(dataFirst.TableName()+"_"+dataTableIndex).Where("task_id=? and label_id=?", label.TaskId, label.ID).First(&dataFirst).Error; err != nil {
		if err.Error() != "record not found" {
			common.LogrusLogger.Error(err)
			common.InitKey(ds.Ctx)
			ds.Ctx.Keys["code"] = common.MYSQL_QUERY_ERROR
			panic(err)
		}
	}
	if err = common.DB.Table(dataLast.TableName()+"_"+dataTableIndex).Where("task_id=? and label_id=?", label.TaskId, label.ID).Last(&dataLast).Error; err != nil {
		if err.Error() != "record not found" {
			common.LogrusLogger.Error(err)
			common.InitKey(ds.Ctx)
			ds.Ctx.Keys["code"] = common.MYSQL_QUERY_ERROR
			panic(err)
		}
	}
	label.BeginDataId = dataFirst.ID
	label.EndDataId = dataLast.ID

	if err = common.DB.Model(&label).Update(&label).Error; err != nil {
		panic(err)
	}
}

func (ds *DataSourceService) CalLabelSummary(req *LabelCalSummaryRequest) *LabelCalSummaryResponse {
	var (
		err            error
		label          = ds.LabelGetByLabelId(req.LabelId)
		dataTableIndex = common.GetDBIndex(req.TaskId)
		dataAll        = make([]models.DataUploadModel, 0)
		labelSum       = models.LabelSummaryModel{}
		res            = LabelCalSummaryResponse{}
	)
	ds.CalDateOfLabelBeginEnd(label)
	if err = common.DB.Table(models.DataUploadModel{}.TableName()+"_"+dataTableIndex).Where("id between ? and ?", label.BeginDataId, label.EndDataId).Find(&dataAll).Error; err != nil {
		if err.Error() != "record not found" {
			common.LogrusLogger.Error(err)
			common.InitKey(ds.Ctx)
			ds.Ctx.Keys["code"] = common.MYSQL_QUERY_ERROR
			panic(err)
		}
	}

	dataCount := len(dataAll)

	if dataCount == 0 {
		return &res
	}

	for _, data := range dataAll {
		labelSum.BatteryCurrentAvg += data.BatteryCurrent
		labelSum.BatteryPowerAvg += data.BatteryPower
		labelSum.BatteryVoltageAvg += data.BatteryVoltage

		labelSum.FpsAvg += data.Fps

		labelSum.CpuAppAvg += data.CpuApp
		labelSum.CpuAvg += data.CpuTotal

		labelSum.GpuDeviceAvg += data.GpuDevice
		labelSum.GpuTilerAvg += data.GpuTiler
		labelSum.GpuRenderAvg += data.GpuRendor

		labelSum.NetRecvAvg += data.NetworkReceive
		labelSum.NetSendAvg += data.NetworkSend

	}

	labelSum.BatteryCurrentAvg = labelSum.BatteryCurrentAvg / float32(dataCount)
	labelSum.BatteryPowerAvg = labelSum.BatteryPowerAvg / float32(dataCount)
	labelSum.BatteryVoltageAvg = labelSum.BatteryVoltageAvg / float32(dataCount)

	labelSum.FpsAvg = labelSum.FpsAvg / float32(dataCount)

	labelSum.CpuAppAvg = labelSum.CpuAppAvg / float32(dataCount)
	labelSum.CpuAvg = labelSum.CpuAvg / float32(dataCount)

	labelSum.GpuDeviceAvg = labelSum.GpuDeviceAvg / float32(dataCount)
	labelSum.GpuTilerAvg = labelSum.GpuTilerAvg / float32(dataCount)
	labelSum.GpuRenderAvg = labelSum.GpuRenderAvg / float32(dataCount)

	labelSum.NetRecvAvg = labelSum.NetRecvAvg / float32(dataCount)
	labelSum.NetSendAvg = labelSum.NetSendAvg / float32(dataCount)

	ds.LabelSummaryUpdate(label.ID, &labelSum)
	return &res
}
