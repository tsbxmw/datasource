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
		if err.Error() != "record not found" {
		}
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

func (ds *DataSourceService) CalLabelSummary(label *models.LabelModel) {

}
