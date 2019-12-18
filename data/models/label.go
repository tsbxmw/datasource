package models

import (
	"github.com/tsbxmw/datasource/common"
	"time"
)

type LabelModel struct {
	common.BaseModel
	TaskId      int    `json:"task_id"`
	Name        string `json:"name"`
	BeginDataId int    `json:"begin_data_id"`
	EndDataId   int    `json:"end_data_id"`
}

func (LabelModel) TableName() string {
	return "label"
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

type LabelSummaryModel struct {
	common.BaseModel
	LabelId           int       `json:"label_id"`
	JankAvg           string    `json:"jank_avg"`
	JankPerHour       string    `json:"jank_per_hour"`
	BigJankAvg        string    `json:"big_jank_avg"`
	BigJankPerHour    string    `json:"big_jank_per_hour"`
	FtimeAvg          time.Time `json:"ftime_avg"`
	FtimeGe_100       time.Time `json:"ftime_ge_100"`
	FtimeDelta        time.Time `json:"ftime_delta"`
	FpsAvg            string    `json:"fps_avg"`
	FpsVar            string    `json:"fps_var"`
	FpsGe_18          string    `json:"fps_ge_18"`
	FpsGe_25          string    `json:"fps_ge_25"`
	FpsDrop           string    `json:"fps_drop"`
	CpuAvg            string    `json:"cpu_avg"`
	CpuLe_50Avg       string    `json:"cpu_le_50_avg"`
	CpuLe_90Avg       string    `json:"cpu_le_80_avg"`
	CpuAppAvg         string    `json:"cpu_app_avg"`
	CpuAppLe_50Avg    string    `json:"cpu_app_le_60_avg"`
	CpuAppLe_90Avg    string    `json:"cpu_app_le_80_avg"`
	CpuTempAvg        string    `json:"cpu_temp_avg"`
	MemUseAvg         string    `json:"mem_use_avg"`
	MemUseMax         string    `json:"mem_use_max"`
	MemSwapAvg        string    `json:"mem_swap_avg"`
	MemSwapMax        string    `json:"mem_swap_max"`
	MemNativePssAvg   string    `json:"mem_native_pss_avg"`
	MemGfxAvg         string    `json:"mem_gfx_avg"`
	MemGlAvg          string    `json:"mem_gl_avg"`
	MemUnknownAvg     string    `json:"mem_unknown_avg"`
	MemXcodeAvg       string    `json:"mem_xcode_avg"`
	MemXcodeMax       string    `json:"mem_xcode_max"`
	MemRealAvg        string    `json:"mem_real_avg"`
	MemRealMax        string    `json:"mem_real_max"`
	MemVirtualAvg     string    `json:"mem_virtual_avg"`
	MemVirtualMax     string    `json:"mem_virtual_max"`
	GpuUseAvg         string    `json:"gpu_use_avg"`
	GpuClockAvg       string    `json:"gpu_clock_avg"`
	GpuRenderAvg      string    `json:"gpu_render_avg"`
	GpuTilerAvg       string    `json:"gpu_tiler_avg"`
	GpuDeviceAvg      string    `json:"gpu_device_avg"`
	SwitchAvg         string    `json:"switch_avg"`
	WakeupAvg         string    `json:"wakeup_avg"`
	NetSendAvg        string    `json:"net_send_avg"`
	NetSendSum        string    `json:"net_send_sum"`
	NetSendPre_10mAvg string    `json:"net_send_pre_10m_avg"`
	NetRecvAvg        string    `json:"net_recv_avg"`
	NetRecvSum        string    `json:"net_recv_sum"`
	NetRecvPre_10mAvg string    `json:"net_recv_pre_10m_avg"`
	BatteryCurrentAvg string    `json:"battery_current_avg"`
	BatteryPowerAvg   string    `json:"battery_power_avg"`
	BatteryVoltageAvg string    `json:"battery_voltage_avg"`
	CpuEnergy         string    `json:"cpu_energy"`
	GpuEnergy         string    `json:"gpu_energy"`
	NetEnergy         string    `json:"net_energy"`
	LocationEnergy    string    `json:"location_energy"`
	DisplayEnergy     string    `json:"display_energy"`
	OverheadEnergy    string    `json:"overhead_energy"`
}

func (LabelSummaryModel) TableName() string {
	return "label_summary"
}
