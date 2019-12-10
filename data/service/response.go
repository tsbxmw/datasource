package service

import (
    "time"
)

type (
    TaskInitResponse struct {
        TaskId   int    `json:"task_id"`
        TaskName string `json:"task_name"`
    }

    TaskGetListResponse struct {
        AppName     string    `json:"app_name"`
        AppPicture  string    `json:"app_picture"`
        AppVersion  string    `json:"app_version"`
        AppPackage  string    `json:"app_package"`
        DeviceName  string    `json:"device_name"`
        Name        string    `json:"name"`
        AvgFps      string    `json:"avg_fps"`
        UploadTime  time.Time `json:"upload_time"`
        CreatorId   int       `json:"creator_id"`
        CreatorName string    `json:"creator_name"`
        Duration    string    `json:"duration"`
        SDKVersion  string    `json:"sdk_version"`
    }
)

type (
    LabelInitResponse struct {
        LabelId   int    `json:"label_id"`
        LabelName string `json:"label_name"`
    }

    DataUploadResponse struct {
    }
)

type (
    DeviceInitResponse struct {
        DeviceName string `json:"device_name"`
        DeviceId int `json:"device_id"`
    }
)

type (
    AppInitResponse struct {
        AppId   int    `json:"app_id"`
        AppName string `json:"app_name"`
    }
)
