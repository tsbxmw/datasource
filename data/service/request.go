package service

import "github.com/tsbxmw/datasource/common"

type (
    TaskInitRequest struct {
        UserId     int    `json:"user_id"`
        TaskName   string `json:"task_name" binding:"required"`
        SdkVersion string `json:"sdk_version"`
    }

    TaskGetListRequest struct {
        common.PageBaseRequst
        UserId int `json:"user_id"`
    }
)

type (
    LabelInitRequest struct {
        TaskId    int    `json:"task_id" binding:"required"`
        LabelName string `json:"label_name"`
    }
)

type (
    DataUploadRequest struct {
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
)
