package service

import "github.com/tsbxmw/datasource/common"

type (
	TaskInitRequest struct {
		UserId     int    `json:"user_id"`
		TaskName   string `json:"task_name" binding:"required"`
		SdkVersion string `json:"sdk_version"`
	}

	TaskGetListRequest struct {
		common.PageFormBaseRequest
		UserId int `form:"user_id" binding:"required"`
	}

	TaskGetDetailRequest struct {
		TaskId int `form:"task_id"`
	}

	TaskGetReportRequest struct {
		TaskId int `form:"task_id"`
	}
)

type (
	LabelInitRequest struct {
		TaskId    int    `json:"task_id" binding:"required"`
		LabelName string `json:"label_name"`
	}

	LabelGetDetailRequest struct {
		LabelId int `form:"label_id"`
	}

	LabelGetListByTaskIdRequest struct {
		TaskId int `form:"task_id" binding:"required"`
	}

	LabelCalSummaryRequest struct {
		LabelId int `json:"label_id" binding:"required"`
		TaskId  int `json:"task_id" binding:"required"`
	}
)

type (
	DataUploadRequest struct {
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
)

type (
	DeviceInitRequest struct {
		Name          string `json:"name"`
		Cpu           string `json:"cpu"`
		Gpu           string `json:"gpu"`
		Type          string `json:"type"`
		Os            string `json:"os"`
		CpuType       string `json:"cpu_type"`
		CpuArch       string `json:"cpu_arch"`
		CpuCoreNumber int    `json:"cpu_core_number"`
		CpuFrequency  string `json:"cpu_frequency"`
		Ram           string `json:"ram"`
		Rom           string `json:"rom"`
		TaskId        int    `json:"task_id" binding:"required"`
	}

	DeviceGetByIdRequest struct {
		DeviceId int `form:"device_id" binding:"required"`
		TaskId   int `form:"task_id" binding:"required"`
	}
)

type (
	AppInitRequest struct {
		Name      string `json:"name"`
		Version   string `json:"version"`
		Package   string `json:"package"`
		Extention string `json:"extention"`
		Remark    string `json:"remark"`
		TaskId    int    `json:"task_id" binding:"required"`
	}

	AppGetByIdRequest struct {
		AppId  int `form:"app_id" binding:"required"`
		TaskId int `form:"task_id" binding:"required"`
	}
)
