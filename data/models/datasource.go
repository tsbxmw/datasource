package models

import (
    "github.com/tsbxmw/datasource/common"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

type DataSourceModel struct {
    common.BaseModeNormal
    RecId     int `json:"rec_id"`
    ContentId int `json:"content_id"`
    Status    int `json:"status";gorm:"DEFAULT:0"`
}

func (DataSourceModel) TableName() string {
    return "message"
}

type TaskModel struct {
    common.BaseModel
    Name       string `json:"name"`
    DeviceId   int    `json:"device_id"`
    UserId     int    `json:"user_id"`
    TimeUse    int    `json:"time_use"`
    SdkVersion string `json:"sdk_version"`
    AppId      int    `json:"app_id"`
    Remark     string `json:"remark"`
}

func (TaskModel) TableName() string {
    return "task"
}

type TaskUserModel struct {
    common.BaseModel
    UserId int    `json:"user_id"`
    TaskId int    `json:"task_id"`
    Remark string `json:"remark"`
}

func (TaskUserModel) TableName() string {
    return "task_user"
}

type LabelModel struct {
    common.BaseModel
    TaskId int    `json:"task_id"`
    Name   string `json:"name"`
}

func (LabelModel) TableName() string {
    return "label"
}

type DataUploadModel struct {
    common.BaseModel
    TaskId          int    `json:"task_id" binding:"required"`
    LabelId         int    `json:"label_id" binding:"required"`
    LabelName       string `json:"label_name" binding:"required"`
    Fps             string `json:"fps"`
    CpuTotal        string `json:"cpu_total"`
    CpuApp          string `json:"cpu_app"`
    MemoryTotal     string `json:"memory_total"`
    MemoryVirtual   string `json:"memory_virtual"`
    MemoryReal      string `json:"memory_real"`
    NetworkSend     string `json:"network_send"`
    NetworkReceive  string `json:"network_receive"`
    GpuRendor       string `json:"gpu_rendor"`
    GpuTiler        string `json:"gpu_tiler"`
    GpuDevice       string `json:"gpu_device"`
    CSwitch         string `json:"c_switch"`
    BatteryCurrent  string `json:"battery_current"`
    BatteryPower    string `json:"battery_power"`
    BatteryVoltage string `json:"battery_voltage"`
    ScreenShot     string `json:"screen_shot"`
}


func (DataUploadModel) TableName() string {
    return "data"
}