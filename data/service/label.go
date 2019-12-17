package service

import (
	"github.com/tsbxmw/datasource/common"
	"github.com/tsbxmw/datasource/data/models"
	"time"
)

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

func (ds *DataSourceService) LabelGetDetail(req *LabelGetDetailRequest) *LabelGetDetailResponse {
	var (
		err error
		res = LabelGetDetailResponse{
			LabelSummary: models.LabelSummaryModel{},
		}
	)

	if err = common.DB.Table(res.LabelSummary.TableName()).Where("id=?", req.LabelId).Find(&res.LabelSummary).Error; err != nil {
		panic(err)
	}
	return &res
}
