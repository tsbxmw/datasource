package models

import "github.com/tsbxmw/datasource/common"

type LabelModel struct {
    common.BaseModel
    TaskId int    `json:"task_id"`
    Name   string `json:"name"`
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
    LabelId           float32 `json:"label_id"`
    FpsAvg            float32 `json:"fps_avg"`
    FpsVar            float32 `json:"fps_var"`
    FpsGe18           float32 `json:"fps_ge_18"`
    FpsGe25           float32 `json:"fps_ge_25"`
    FpsDrop           float32 `json:"fps_drop"`
    JanAvg            float32 `json:"jank_avg"`
    JankPerHour       float32 `json:"jank_per_hour"`
    BigJankAvg        float32 `json:"big_jank_avg"`
    BigJankPerHour    float32 `json:"big_jank_per_hour"`
    FTimeAvg          float32 `json:"ftime_avg"`
    FTimeGe100        float32 `json:"ftime_ge_100"`
    FTimeDelta        float32 `json:"ftime_delta"`
    CpuAvg            float32 `json:"cpu_avg"`
    CpuLe50Avg        float32 `json:"cpu_le_50_avg"`
    CpuLe90Avg        float32 `json:"cpu_le_80_avg"`
    CpuAppAvg         float32 `json:"cpu_app_avg"`
    CpuAppLe50Avg     float32 `json:"cpu_app_le_60_avg"`
    CpuAppLe90Avg     float32 `json:"cpu_app_le_80_avg"`
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
    SeitchAvg         float32 `json:"switch_avg"`
    WakeupAvg         float32 `json:"wakeup_avg"`
    NetSendAvg        float32 `json:"net_send_avg"`
    NetSendSum        float32 `json:"net_send_sum"`
    NetSendPre10mAvg  float32 `json:"net_send_pre_10m_avg"`
    NetRecvAvg        float32 `json:"net_recv_avg"`
    NetRecvSum        float32 `json:"net_recv_sum"`
    NetRecvPre10mAvg  float32 `json:"net_recv_pre_10m_avg"`
    BatteryCurrentAvg float32 `json:"battery_current_avg"`
    BatteryPowerAvg   float32 `json:"battery_power_avg"`
    BatteryVoltageAvg float32 `json:"battery_voltage_avg"`
    CpuEnergy         float32 `json:"cpu_energy"`
    GpuEnergy         float32 `json:"gpu_energy"`
    NetEnergy         float32 `json:"net_energy"`
    LocationEnergy    float32 `json:"location_energy"`
    DisplayEnergy     float32 `json:"display_energy"`
    OverheadEnery     float32 `json:"overhead_energy"`
}

func (LabelSummaryModel) TableName() string {
    return "label_summary"
}
