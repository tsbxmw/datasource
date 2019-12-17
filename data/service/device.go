package service

import (
	"github.com/tsbxmw/datasource/common"
	"github.com/tsbxmw/datasource/data/models"
	"time"
)

func (ds *DataSourceService) DeviceInit(req *DeviceInitRequest) *DeviceInitResponse {
	var (
		err error
		res = DeviceInitResponse{}
	)
	device := models.DeviceModel{}
	if err = common.DB.Table(device.TableName()).Where("task_id=?", req.TaskId).First(&device).Error; err != nil {
		if err.Error() != "record not found" {
			common.LogrusLogger.Error(err)
			common.InitKey(ds.Ctx)
			ds.Ctx.Keys["code"] = common.MYSQL_QUERY_ERROR
			panic(err)
		}
	}
	if device.ID > 0 {
		res.DeviceId = device.ID
		res.DeviceName = device.Name
	} else {
		device.Name = req.Name
		device.TaskId = req.TaskId
		device.Cpu = req.Cpu
		device.CpuArch = req.CpuArch
		device.CreationTime = time.Now()
		device.Gpu = req.Gpu
		device.Type = req.Type
		device.Os = req.Os
		device.CpuType = req.CpuType
		device.CpuCoreNumber = req.CpuCoreNumber
		device.CpuFrequency = req.CpuFrequency
		device.Ram = req.Ram
		device.Rom = req.Rom
		device.ModifiedTime = time.Now()
		if err = common.DB.Table(device.TableName()).Create(&device).Error; err != nil {
			common.LogrusLogger.Error(err)
			panic(err)
		}
		res.DeviceId = device.ID
		res.DeviceName = device.Name
	}
	return &res
}

func (ds *DataSourceService) DeviceGetById(req *DeviceGetByIdRequest) *DeviceGetResponse {
	var (
		err error
		res = DeviceGetResponse{
			DeviceModel: models.DeviceModel{},
		}
	)

	device := models.DeviceModel{}

	if req.TaskId > 0 && req.DeviceId <= 0 {
		if err = common.DB.Table(device.TableName()).Where("task_id=?", req.TaskId).First(&device).Error; err != nil {
			panic(err)
		}
	} else if req.DeviceId > 0 {
		if err = common.DB.Table(device.TableName()).Where("id=?", req.DeviceId).First(&device).Error; err != nil {
			panic(err)
		}
	}

	res.DeviceModel = device
	return &res
}
