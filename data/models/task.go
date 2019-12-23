package models

import "github.com/tsbxmw/datasource/common"

type TaskModel struct {
    common.BaseModel
    Name       string `json:"name"`
    UserId     int    `json:"user_id"`
    TimeUse    int    `json:"time_use"`
    SdkVersion string `json:"sdk_version"`
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

type TaskSummaryModel struct {
    common.BaseModel
    TaskId            int `json:"task_id"`
    FpsAvg            float32 `json:"fps_avg"`
    FpsVar            float32 `json:"fps_var"`
    FpsGe_18           float32 `json:"fps_ge_18"`
    FpsGe_25           float32 `json:"fps_ge_25"`
    FpsDrop           float32 `json:"fps_drop"`
    JankAvg            float32 `json:"jank_avg"`
    JankPerHour       float32 `json:"jank_per_hour"`
    BigJankAvg        float32 `json:"big_jank_avg"`
    BigJankPerHour    float32 `json:"big_jank_per_hour"`
    FtimeAvg          float32 `json:"ftime_avg"`
    FtimeGe_100        float32 `json:"ftime_ge_100"`
    FtimeDelta        float32 `json:"ftime_delta"`
    CpuAvg            float32 `json:"cpu_avg"`
    CpuLe_50Avg        float32 `json:"cpu_le_50_avg"`
    CpuLe_90Avg        float32 `json:"cpu_le_80_avg"`
    CpuAppAvg         float32 `json:"cpu_app_avg"`
    CpuAppLe_50Avg     float32 `json:"cpu_app_le_60_avg"`
    CpuAppLe_90Avg     float32 `json:"cpu_app_le_80_avg"`
    CpuTempAvg        float32 `json:"cpu_temp_avg"`
    MemUseAvg         float32 `json:"mem_use_avg"`
    MemUseMax         float32 `json:"mem_use_max"`
    MemSwapAvg        float32 `json:"mem_swap_avg"`
    MemSwapMax        float32 `json:"mem_swap_max"`
    MemNativePssAvg   float32 `json:"mem_native_pss_avg"`
    MemGfxAvg         float32 `json:"mem_gfx_avg"`
    MemGlAvg          float32 `json:"mem_gl_avg"`
    MemUnknownAvg     float32 `json:"mem_unknown_avg"`
    MemXcodeAvg       float32 `json:"mem_xcode_avg"`
    MemXcodeMax       float32 `json:"mem_xcode_max"`
    MemRealAvg        float32 `json:"mem_real_avg"`
    MemRealMax        float32 `json:"mem_real_max"`
    MemVirtualAvg     float32 `json:"mem_virtual_avg"`
    MemVirtualMax     float32 `json:"mem_virtual_max"`
    GpuUseAvg         float32 `json:"gpu_use_avg"`
    GpuClockAvg       float32 `json:"gpu_clock_avg"`
    GpuRenderAvg      float32 `json:"gpu_render_avg"`
    GpuTilerAvg       float32 `json:"gpu_tiler_avg"`
    GpuDeviceAvg      float32 `json:"gpu_device_avg"`
    SwitchAvg         float32 `json:"switch_avg"`
    WakeupAvg         float32 `json:"wakeup_avg"`
    NetSendAvg        float32 `json:"net_send_avg"`
    NetSendSum        float32 `json:"net_send_sum"`
    NetSendPre_10mAvg  float32 `json:"net_send_pre_10m_avg"`
    NetRecvAvg        float32 `json:"net_recv_avg"`
    NetRecvSum        float32 `json:"net_recv_sum"`
    NetRecvPre_10mAvg  float32 `json:"net_recv_pre_10m_avg"`
    BatteryCurrentAvg float32 `json:"battery_current_avg"`
    BatteryPowerAvg   float32 `json:"battery_power_avg"`
    BatteryVoltageAvg float32 `json:"battery_voltage_avg"`
    CpuEnergy         float32 `json:"cpu_energy"`
    GpuEnergy         float32 `json:"gpu_energy"`
    NetEnergy         float32 `json:"net_energy"`
    LocationEnergy    float32 `json:"location_energy"`
    DisplayEnergy     float32 `json:"display_energy"`
    OverheadEnergy     float32 `json:"overhead_energy"`
}

func (TaskSummaryModel) TableName() string {
    return "task_summary"
}
