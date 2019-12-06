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
    TaskId         int    `json:"task_id" binding:"required"`
    LabelId        int    `json:"label_id" binding:"required"`
    LabelName      string `json:"label_name" binding:"required"`
    Fps            string `json:"fps"`
    CpuTotal       string `json:"cpu_total"`
    CpuApp         string `json:"cpu_app"`
    MemoryTotal    string `json:"memory_total"`
    MemoryVirtual  string `json:"memory_virtual"`
    MemoryReal     string `json:"memory_real"`
    NetworkSend    string `json:"network_send"`
    NetworkReceive string `json:"network_receive"`
    GpuRendor      string `json:"gpu_rendor"`
    GpuTiler       string `json:"gpu_tiler"`
    GpuDevice      string `json:"gpu_device"`
    CSwitch        string `json:"c_switch"`
    BatteryCurrent string `json:"battery_current"`
    BatteryPower   string `json:"battery_power"`
    BatteryVoltage string `json:"battery_voltage"`
    ScreenShot     string `json:"screen_shot"`
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
}

func (DeviceModel) TableName() string {
    return "device"
}

type LabelBatteryModel struct {
    common.BaseModelCreate
    TaskId     int    `json:"task_id"`
    LabelId    int    `json:"label_id"`
    CurrentAvg string `json:"current_avg"`
    PowerAvg   string `json:"power_avg"`
    VoltageAvg string `json:"voltage_avg"`
}

func (LabelBatteryModel) TableName() string {
    return "label_battery"
}

type LabelCupModel struct {
    common.BaseModelCreate
    TaskId      int    `json:"task_id"`
    LabelId     int    `json:"label_id"`
    CpuTotalAvg string `json:"cpu_total_avg"`
    CpuTotalMax string `json:"cpu_total_max"`
    CpuTotal50  string `json:"cpu_total_50"`
    CpuTotal90  string `json:"cpu_total_90"`
    CpuTotal95  string `json:"cpu_total_95"`
    CpuAppAvg   string `json:"cpu_app_avg"`
    CpuAppMax   string `json:"cpu_app_max"`
    CpuApp50    string `json:"cpu_app_50"`
    CpuApp90    string `json:"cpu_app_90"`
    CpuApp95    string `json:"cpu_app_95"`
    Remark      string `json:"remark"`
}

func (LabelCupModel) TableName() string {
    return "label_cpu"
}

type LabelFpsModel struct {
    common.BaseModelCreate
    TaskId  int    `json:"task_id"`
    LabelId int    `json:"label_id"`
    FpsAvg  string `json:"fps_avg"`
    FpsVar  string `json:"fps_var"`
    Fps18   string `json:"fps_18"`
    Fps25   string `json:"fps_25"`
    FpsDrop string `json:"fps_drop"`
}

func (LabelFpsModel) TableName() string {
    return "label_fps"
}

type LabelGpuModel struct {
    common.BaseModelCreate
    TaskId    int    `json:"task_id"`
    LabelId   int    `json:"label_id"`
    RendorAvg string `json:"rendor_avg"`
    TilerAvg  string `json:"tiler_avg"`
    DeviceAvg string `json:"device_avg"`
}

func (LabelGpuModel) TableName() string {
    return "label_gpu"
}

type LabelMemoryModel struct {
    common.BaseModelCreate
    TaskId           int    `json:"task_id"`
    LabelId          int    `json:"label_id"`
    MemoryPeak       string `json:"memory_peak"`
    MemoryTotalAvg   string `json:"memory_total_avg"`
    MemoryTotalMax   string `json:"memory_total_max"`
    MemoryTotal50    string `json:"memory_total_50"`
    MemoryTotal90    string `json:"memory_total_90"`
    MemoryTotal95    string `json:"memory_total_95"`
    MemoryRealAvg    string `json:"memory_real_avg"`
    MemoryRealMax    string `json:"memory_real_max"`
    MemoryReal50     string `json:"memory_real_50"`
    MemoryReal90     string `json:"memory_real_90"`
    MemoryReal95     string `json:"memory_real_95"`
    MemoryVirtualAvg string `json:"memory_virtual_avg"`
    MemoryVirtualMax string `json:"memory_virtual_max"`
    MemoryVirtual50  string `json:"memory_virtual_50"`
    MemoryVirtual90  string `json:"memory_virtual_90"`
    MemoryVirtual95  string `json:"memory_virtual_95"`
}

func (LabelMemoryModel) TableName() string {
    return "label_memory"
}

type LabelNetworkModel struct {
    common.BaseModelCreate
    Name       string `json:"name"`
    TaskId     int    `json:"task_id"`
    LabelId    int    `json:"label_id"`
    SendAvg    string `json:"send_avg"`
    SendSum    string `json:"send_sum"`
    ReceiveAvg string `json:"receive_avg"`
    ReceiveSum string `json:"receive_sum"`
}

func (LabelNetworkModel) TableName() string {
    return "label_network"
}
