package workers

import (
    "datasource/common"
    "datasource/data/models"
    "datasource/data/service"
    "encoding/json"
    "time"
)

type AReceiver struct {
    Name string
}

func (ar AReceiver) QueueName() string {
    return "data_1"
}

func (ar AReceiver) RouterKey() string {
    return "data_1"
}

func (ar AReceiver) OnError(error) {

}

func (ar AReceiver) OnReceive(body []byte) bool {
    common.LogrusLogger.Info(ar.Name, " -------> receiver")
    req := service.DataUploadRequest{}
    err := json.Unmarshal(body, &req)
    if err != nil {
       common.LogrusLogger.Error(err)
       panic(err)
    }
    common.LogrusLogger.Info(string(body))
    dataModel := models.DataUploadModel{
       BaseModel: common.BaseModel{
           CreationTime: time.Now(),
           ModifiedTime: time.Now(),
       },
       TaskId: req.TaskId,
       LabelId: req.LabelId,
       LabelName: req.LabelName,
       Fps: req.Fps,
       CpuTotal: req.CpuTotal,
       CpuApp: req.CpuApp,
       MemoryTotal: req.MemoryTotal,
       MemoryReal: req.MemoryReal,
       MemoryVirtual: req.MemoryVirtual,
       NetworkReceive: req.NetworkReceive,
       NetworkSend: req.NetworkSend,
       GpuDevice: req.GpuDevice,
       GpuRendor: req.GpuRendor,
       GpuTiler: req.GpuTiler,
       CSwitch: req.CSwitch,
       BatteryCurrent: req.BatteryCurrent,
       BatteryPower: req.BatteryPower,
       BatteryVoltage: req.BatteryVoltage,
       ScreenShot: req.ScreenShot,
    }
    if err = common.DB.Table(dataModel.TableName()).Create(&dataModel).Error; err != nil {
       common.LogrusLogger.Error(err)
    }
    return true
}
