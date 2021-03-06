package models

import (
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "github.com/tsbxmw/datasource/common"
)

type DataSourceModel struct {
    common.BaseModelNormal
    RecId     int `json:"rec_id"`
    ContentId int `json:"content_id"`
    Status    int `json:"status";gorm:"DEFAULT:0"`
}

func (DataSourceModel) TableName() string {
    return "message"
}


type DataUploadModel struct {
    common.BaseModel
    TaskId         int    `json:"task_id" binding:"required"`
    LabelId        int    `json:"label_id" binding:"required"`
    LabelName      string `json:"label_name" binding:"required"`
    Fps            float32 `json:"fps"`
    CpuTotal       float32 `json:"cpu_total"`
    CpuApp         float32 `json:"cpu_app"`
    MemoryTotal    float32 `json:"memory_total"`
    MemoryVirtual  float32 `json:"memory_virtual"`
    MemoryReal     float32 `json:"memory_real"`
    NetworkSend    float32 `json:"network_send"`
    NetworkReceive float32 `json:"network_receive"`
    GpuRendor      float32 `json:"gpu_rendor"`
    GpuTiler       float32 `json:"gpu_tiler"`
    GpuDevice      float32 `json:"gpu_device"`
    CSwitch        float32 `json:"c_switch"`
    BatteryCurrent float32 `json:"battery_current"`
    BatteryPower   float32 `json:"battery_power"`
    BatteryVoltage float32 `json:"battery_voltage"`
    ScreenShot     float32 `json:"screen_shot"`
}

func (DataUploadModel) TableName() string {
    return "data"
}

type AuthModel struct {
    common.BaseModel
    UserId    int    `json:"user_id"`
    AppKey    string `json:"app_key"`
    AppSecret string `json:"app_secret"`
    Status    int    `json:"status"`
}

func (AuthModel) TableName() string {
    return "auth"
}

type AppModel struct {
    common.BaseModelCreate
    Name      string `json:"name"`
    Version   string `json:"version"`
    Package   string `json:"package"`
    Extention string `json:"extention"`
    Remark    string `json:"remark"`
    TaskId    int    `json:"task_id"`
}

func (AppModel) TableName() string {
    return "app"
}

type DeviceModel struct {
    common.BaseModel
    Name          string `json:"name"`
    Cpu           string `json:"cpu"`
    Gpu           string `json:"gpu"`
    Os            string `json:"os"`
    CpuType       string `json:"cpu"`
    CpuArch       string `json:"cpu_arch"`
    CpuCoreNumber int    `json:"cpu_core_number"`
    CpuFrequency  string `json:"cpu_frequency"`
    Ram           string `json:"ram"`
    Rom           string `json:"rom"`
    Type          string `json:"type"`
    TaskId        int    `json:"task_id"`
}

func (DeviceModel) TableName() string {
    return "device"
}
