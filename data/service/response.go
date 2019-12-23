package service

import (
	"github.com/tsbxmw/datasource/data/models"
	"time"
)

type (
	DataUploadResponse struct {
	}

	DataGetResponse struct {
		CpuUsage float32 ``
		//{
		//	"TimeStamp": "0",
		//	"CpuUsage": {
		//	"AppUsage": 10.0,
		//	"TotalUsage": 22.0
		//},
		//	"AndroidFps": {
		//	"fps": 16.0,
		//	"Jank": 0.0
		//},
		//	"AndroidMemoryUsage": {
		//	"Memory": 42.0,
		//	"SwapMemory": 0.0
		//},
		//	"NetworkUsage": {
		//	"UpSpeed": 0.0,
		//	"DownSpeed": 0.0
		//},
		//	"CpuTemperature": {
		//	"CpuTemperature": 35.0
		//},
		//	"IsDelete": false,
		//	"BigJank": {
		//	"BigJank": 0.0
		//},
		//	"VirtualMemory": {
		//	"VirtualMemory": 1545.0
		//}
		//},
	}
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

	TaskGetResponse struct {
	}

	TaskGetDetailResponse struct {
		TaskSummary models.TaskSummaryModel `json:"summary"`
		TaskDetail  TaskGetListResponse     `json:"task_detail"`
		LabelInfos  LabelGetListResponse    `json:"label_infos"`
	}

	TaskGetReportResponse struct {
		TaskDetail TaskGetListResponse  `json:"task_detail"`
		LabelInfos LabelGetListResponse `json:"label_infos"`
	}

	TaskCalSummaryResponse struct {
	}
)

type (
	LabelInitResponse struct {
		LabelId   int    `json:"label_id"`
		LabelName string `json:"label_name"`
	}

	LabelResponse struct {
		Summary models.LabelSummaryModel `json:"summary"`
		Info    models.LabelModel        `json:"info"`
	}

	LabelGetListResponse struct {
		Label []LabelResponse `json:"label"`
	}

	LabelCalSummaryResponse struct {
	}
)

type (
	DeviceInitResponse struct {
		DeviceName string `json:"device_name"`
		DeviceId   int    `json:"device_id"`
	}

	DeviceGetResponse struct {
		models.DeviceModel
	}
)

type (
	AppInitResponse struct {
		AppId   int    `json:"app_id"`
		AppName string `json:"app_name"`
	}

	AppGetResponse struct {
		models.AppModel
	}
)
