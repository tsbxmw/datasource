package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/tsbxmw/datasource/common/middleware"
	v1 "github.com/tsbxmw/datasource/data/routers/v1"
)

func InitRouter(e *gin.Engine) {
	GroupV1 := e.Group("/v1", middleware.AuthMiddleware())
	{
		// data upload tewst
		GroupV1.GET("/data/", v1.DataInit)
		/**
		  * @api {post} /v1/data/upload 上传 数据
		  * @apiName DataUpload
		  * @apiGroup Data
		  *
		  * @apiParam {int} task_id task id
		  * @apiParam {int} label_id label id
		  * @apiParam {string} label_name label name
		  * @apiParam {float32} fps fps
		  * @apiParam {float32} cpu_total cpu_total
		  * @apiParam {float32} cpu_app cpu_app
		  * @apiParam {float32} memory_total memory_total
		  * @apiParam {float32} memory_virtual memory_virtual
		  * @apiParam {float32} memory_real memory_real
		  * @apiParam {float32} network_send network_send
		  * @apiParam {float32} network_receive network_receive
		  * @apiParam {float32} gpu_rendor gpu_rendor
		  * @apiParam {float32} gpu_tiler gpu_tiler
		  * @apiParam {float32} gpu_device gpu_device
		  * @apiParam {float32} c_switch c_switch
		  * @apiParam {float32} battery_current battery_current
		  * @apiParam {float32} battery_power battery_power
		  * @apiParam {float32} battery_voltage battery_voltage
		  * @apiParam {float32} screen_shot screen_shot
		  * @apiSuccessExample {json} Success-Response:
			HTTP/1.1 200 OK
		    {
				"code": 200,
				"msg": "success",
				"data": {}
			}
		*/
		GroupV1.POST("/data/upload", v1.DataUpload)

		task := GroupV1.Group("/data/task")
		// task group
		{
			// 初始化 task
			/**
			  * @api {post} /v1/data/task 创建 task
			  * @apiName TaskInit
			  * @apiGroup Data
			  *
			  * @apiParam {int} user_id user id
			  * @apiParam {string} task_name task_name
			  * @apiParam {string} [sdk_version] sdk version
			  * @apiSuccessExample {json} Success-Response:
				HTTP/1.1 200 OK
			    {
					"code": 200,
					"msg": "success",
					"data": {
						"task_id": 34,
						"task_name": "test2"
					}
				}
			*/
			task.POST("/", v1.TaskInit)
			// 获取 task 列表
			/**
			  * @api {get} /v1/data/task/list 获取 task 列表
			  * @apiName TaskListGet
			  * @apiGroup Data
			  *
			  * @apiParam {int} user_id user id
			  * @apiParam {int} page_size page
			  * @apiParam {int} page_index page
			  * @apiSuccessExample {json} Success-Response:
				HTTP/1.1 200 OK
			    {
					"code": 200,
					"msg": "success",
					"data": [
						{
							"app_name": "",
							"app_picture": "",
							"app_version": "",
							"app_package": "",
							"device_name": "",
							"name": "test2",
							"avg_fps": "",
							"upload_time": "2019-12-26T03:15:38Z",
							"creator_id": 5,
							"creator_name": "",
							"duration": "",
							"sdk_version": ""
						},
						{
							"app_name": "",
							"app_picture": "",
							"app_version": "",
							"app_package": "",
							"device_name": "",
							"name": "test_5",
							"avg_fps": "",
							"upload_time": "2019-12-26T03:19:17Z",
							"creator_id": 5,
							"creator_name": "",
							"duration": "",
							"sdk_version": "0.0.5"
						}
					]
				}
			*/
			task.GET("/list", v1.TaskGetList)
			// 获取 task 详细信息
			/**
			  * @api {get} /v1/data/task/detail 获取 task 详细信息
			  * @apiName TaskDetailGet
			  * @apiGroup Data
			  *
			  * @apiParam {int} task_id task's id
			  * @apiSuccessExample {json} Success-Response:
				HTTP/1.1 200 OK
			    {
					"code": 200,
					"msg": "success",
					"data": {
						"summary": {
							"id": 0,
							"creation_time": "0001-01-01T00:00:00Z",
							"modified_time": "0001-01-01T00:00:00Z",
							"task_id": 0,
							"fps_avg": 0,
							"fps_var": 0,
							"fps_ge_18": 0,
							"fps_ge_25": 0,
							"fps_drop": 0,
							"jank_avg": 0,
							"jank_per_hour": 0,
							"big_jank_avg": 0,
							"big_jank_per_hour": 0,
							"ftime_avg": 0,
							"ftime_ge_100": 0,
							"ftime_delta": 0,
							"cpu_avg": 0,
							"cpu_le_50_avg": 0,
							"cpu_le_80_avg": 0,
							"cpu_app_avg": 0,
							"cpu_app_le_60_avg": 0,
							"cpu_app_le_80_avg": 0,
							"cpu_temp_avg": 0,
							"mem_use_avg": 0,
							"mem_use_max": 0,
							"mem_swap_avg": 0,
							"mem_swap_max": 0,
							"mem_native_pss_avg": 0,
							"mem_gfx_avg": 0,
							"mem_gl_avg": 0,
							"mem_unknown_avg": 0,
							"mem_xcode_avg": 0,
							"mem_xcode_max": 0,
							"mem_real_avg": 0,
							"mem_real_max": 0,
							"mem_virtual_avg": 0,
							"mem_virtual_max": 0,
							"gpu_use_avg": 0,
							"gpu_clock_avg": 0,
							"gpu_render_avg": 0,
							"gpu_tiler_avg": 0,
							"gpu_device_avg": 0,
							"switch_avg": 0,
							"wakeup_avg": 0,
							"net_send_avg": 0,
							"net_send_sum": 0,
							"net_send_pre_10m_avg": 0,
							"net_recv_avg": 0,
							"net_recv_sum": 0,
							"net_recv_pre_10m_avg": 0,
							"battery_current_avg": 0,
							"battery_power_avg": 0,
							"battery_voltage_avg": 0,
							"cpu_energy": 0,
							"gpu_energy": 0,
							"net_energy": 0,
							"location_energy": 0,
							"display_energy": 0,
							"overhead_energy": 0
						},
						"task_detail": {
							"app_name": "",
							"app_picture": "",
							"app_version": "",
							"app_package": "",
							"device_name": "",
							"name": "test2",
							"avg_fps": "",
							"upload_time": "2019-12-26T03:15:38Z",
							"creator_id": 5,
							"creator_name": "",
							"duration": "",
							"sdk_version": ""
						},
						"label_infos": {
							"label": []
						}
					}
				}
			*/
			task.GET("/detail", v1.TaskGetDetail)
			// 获取 task 报告
			/**
			  * @api {get} /v1/data/task/report 获取 task report
			  * @apiName TaskReportGet
			  * @apiGroup Data
			  *
			  * @apiParam {int} user_id user id
			  * @apiParam {int} page_size page
			  * @apiParam {int} page_index page
			  * @apiSuccessExample {json} Success-Response:
				HTTP/1.1 200 OK
			    {
					"code": 200,
					"msg": "success",
					"data": {
						"summary": {
							"id": 0,
							"creation_time": "0001-01-01T00:00:00Z",
							"modified_time": "0001-01-01T00:00:00Z",
							"task_id": 0,
							"fps_avg": 0,
							"fps_var": 0,
							"fps_ge_18": 0,
							"fps_ge_25": 0,
							"fps_drop": 0,
							"jank_avg": 0,
							"jank_per_hour": 0,
							"big_jank_avg": 0,
							"big_jank_per_hour": 0,
							"ftime_avg": 0,
							"ftime_ge_100": 0,
							"ftime_delta": 0,
							"cpu_avg": 0,
							"cpu_le_50_avg": 0,
							"cpu_le_80_avg": 0,
							"cpu_app_avg": 0,
							"cpu_app_le_60_avg": 0,
							"cpu_app_le_80_avg": 0,
							"cpu_temp_avg": 0,
							"mem_use_avg": 0,
							"mem_use_max": 0,
							"mem_swap_avg": 0,
							"mem_swap_max": 0,
							"mem_native_pss_avg": 0,
							"mem_gfx_avg": 0,
							"mem_gl_avg": 0,
							"mem_unknown_avg": 0,
							"mem_xcode_avg": 0,
							"mem_xcode_max": 0,
							"mem_real_avg": 0,
							"mem_real_max": 0,
							"mem_virtual_avg": 0,
							"mem_virtual_max": 0,
							"gpu_use_avg": 0,
							"gpu_clock_avg": 0,
							"gpu_render_avg": 0,
							"gpu_tiler_avg": 0,
							"gpu_device_avg": 0,
							"switch_avg": 0,
							"wakeup_avg": 0,
							"net_send_avg": 0,
							"net_send_sum": 0,
							"net_send_pre_10m_avg": 0,
							"net_recv_avg": 0,
							"net_recv_sum": 0,
							"net_recv_pre_10m_avg": 0,
							"battery_current_avg": 0,
							"battery_power_avg": 0,
							"battery_voltage_avg": 0,
							"cpu_energy": 0,
							"gpu_energy": 0,
							"net_energy": 0,
							"location_energy": 0,
							"display_energy": 0,
							"overhead_energy": 0
						},
						"task_detail": {
							"app_name": "",
							"app_picture": "",
							"app_version": "",
							"app_package": "",
							"device_name": "",
							"name": "",
							"avg_fps": "",
							"upload_time": "0001-01-01T00:00:00Z",
							"creator_id": 0,
							"creator_name": "",
							"duration": "",
							"sdk_version": ""
						},
						"label_infos": {
							"label": null
						}
					}
				}
			*/
			task.GET("/report", v1.TaskGetReport)
			// 创建 app
			/**
			  * @api {post} /v1/data/task/app 创建 task app 实例
			  * @apiName TaskAppInit
			  * @apiGroup Data
			  *
			  * @apiParamExample {json} Request-Example:
				{
					"name": "qutoutiao",
					"version": "1.0.0",
					"package": "com.qutoutiao.Main",
					"extention": "",
					"remark": "test here",
					"task_id": 23
				}
			  * @apiSuccessExample {json} Success-Response:
				HTTP/1.1 200 OK
			    {
					"code": 200,
					"msg": "success",
					"data": {
						"task_id": 34,
						"task_name": "test2"
					}
				}
			*/
			task.POST("/app", v1.AppInit)
			// 通过 id、task_id 获取 app 信息
			/**
			  * @api {get} /v1/data/task/app 获取 task app
			  * @apiName TaskAppGet
			  * @apiGroup Data
			  *
			  * @apiParam {int} app_id app's id,if app_id==-1,using task_id
			  * @apiParam {int} [task_id] task id for app_id==-1
			  * @apiSuccessExample {json} Success-Response:
				HTTP/1.1 200 OK
				{
					"code": 200,
					"data": {
						"id": 7,
						"creation_time": "2019-12-10T17:27:41+08:00",
						"name": "qutoutiao",
						"version": "1.0.0",
						"package": "",
						"extention": "",
						"remark": "test here",
						"task_id": 23
					},
					"msg": "success"
				}
			*/
			task.GET("/app", v1.AppGetById)
			// 创建 device
			/**
			  * @api {post} /v1/data/task/device 创建 task device 实例
			  * @apiName TaskDeviceInit
			  * @apiGroup Data
			  * @apiParam {int} task_id task id
			  * @apiParam {string} name device name
			  * @apiParam {string} cpu device cpu
			  * @apiParam {string} gpu device gpu
			  * @apiParam {string} type device type
			  * @apiParam {string} os device os
			  * @apiParam {string} cpu_type device cpu_type
			  * @apiParam {string} cpu_arch device cpu_arch
			  * @apiParam {string} cpu_core_number device cpu_core_number
			  * @apiParam {string} ram device ram
			  * @apiParam {string} rom device rom
			  *
			  * @apiParamExample {json} Request-Example:
				{
					"name": "Iphone6sPlus",
					"task_id": 23
				}
			  * @apiSuccessExample {json} Success-Response:
				HTTP/1.1 200 OK
			    {
					"code": 200,
					"data": {
						"device_name": "Iphone6sPlus",
						"device_id": 1
					},
					"msg": "success"
				}
			*/
			task.POST("/device", v1.DeviceInit)
			// 通过 id、task_id 获取 device 信息
			/**
			  * @api {get} /v1/data/task/device 获取 task device
			  * @apiName TaskDeviceGet
			  * @apiGroup Data
			  *
			  * @apiParam {int} device_id device's id,if device_id==-1,using task_id
			  * @apiParam {int} [task_id] task id for app_id==-1
			  * @apiSuccessExample {json} Success-Response:
				HTTP/1.1 200 OK
				{
					"code": 200,
					"msg": "success",
					"data": {
						"id": 1,
						"creation_time": "2019-12-10T17:27:47+08:00",
						"modified_time": "2019-12-10T17:27:47+08:00",
						"name": "Iphone6sPlus",
						"gpu": "",
						"os": "",
						"cpu_arch": "",
						"cpu_core_number": 0,
						"cpu_frequency": "",
						"ram": "",
						"rom": "",
						"type": "",
						"task_id": 23
					}
				}
			*/
			task.GET("/device", v1.DeviceGetById)
			// 通过 task_id 获取 label 信息
			/**
			  * @api {get} /v1/data/task/label 获取 task label 列表
			  * @apiName TaskLabelListGet
			  * @apiGroup Data
			  *
			  * @apiParam {int} task_id task's id
			  * @apiSuccessExample {json} Success-Response:
				HTTP/1.1 200 OK
				{
					"code": 0,
					"msg": "success",
					"data": {
						"label": [
							{
								"summary": {
									"id": 6,
									"creation_time": "2019-12-20T11:39:10+08:00",
									"modified_time": "2019-12-24T16:10:57+08:00",
									"label_id": 18,
									"jank_avg": 0,
									"jank_per_hour": 0,
									"big_jank_avg": 0,
									"big_jank_per_hour": 0,
									"ftime_avg": "2019-12-20T11:39:10+08:00",
									"ftime_ge_100": "2019-12-20T11:39:10+08:00",
									"ftime_delta": "2019-12-20T11:39:10+08:00",
									"fps_avg": 30.629108,
									"fps_var": 0,
									"fps_ge_18": 0,
									"fps_ge_25": 0,
									"fps_drop": 0,
									"cpu_avg": 70.05157,
									"cpu_le_50_avg": 0,
									"cpu_le_80_avg": 0,
									"cpu_app_avg": 20.45274,
									"cpu_app_le_60_avg": 0,
									"cpu_app_le_80_avg": 0,
									"cpu_temp_avg": 0,
									"mem_use_avg": 0,
									"mem_use_max": 0,
									"mem_swap_avg": 0,
									"mem_swap_max": 0,
									"mem_native_pss_avg": 0,
									"mem_gfx_avg": 0,
									"mem_gl_avg": 0,
									"mem_unknown_avg": 0,
									"mem_xcode_avg": 0,
									"mem_xcode_max": 0,
									"mem_real_avg": 0,
									"mem_real_max": 0,
									"mem_virtual_avg": 0,
									"mem_virtual_max": 0,
									"gpu_use_avg": 0,
									"gpu_clock_avg": 0,
									"gpu_render_avg": 69.968864,
									"gpu_tiler_avg": 20.520794,
									"gpu_device_avg": 20.462141,
									"switch_avg": 0,
									"wakeup_avg": 0,
									"net_send_avg": 20077.973,
									"net_send_sum": 0,
									"net_send_pre_10m_avg": 0,
									"net_recv_avg": 20011.871,
									"net_recv_sum": 0,
									"net_recv_pre_10m_avg": 0,
									"battery_current_avg": 0,
									"battery_power_avg": 0,
									"battery_voltage_avg": 0,
									"cpu_energy": 0,
									"gpu_energy": 0,
									"net_energy": 0,
									"location_energy": 0,
									"display_energy": 0,
									"overhead_energy": 0
								},
								"info": {
									"id": 18,
									"creation_time": "2019-12-20T11:33:21+08:00",
									"modified_time": "2019-12-20T11:33:21+08:00",
									"task_id": 29,
									"name": "test_5",
									"begin_data_id": 1,
									"end_data_id": 29359
								}
							},
							{
								"summary": {
									"id": 21,
									"creation_time": "2019-12-20T14:18:42+08:00",
									"modified_time": "2019-12-24T16:11:34+08:00",
									"label_id": 32,
									"jank_avg": 0,
									"jank_per_hour": 0,
									"big_jank_avg": 0,
									"big_jank_per_hour": 0,
									"ftime_avg": "2019-12-20T14:18:42+08:00",
									"ftime_ge_100": "2019-12-20T14:18:42+08:00",
									"ftime_delta": "2019-12-20T14:18:42+08:00",
									"fps_avg": 30.451536,
									"fps_var": 0,
									"fps_ge_18": 0,
									"fps_ge_25": 0,
									"fps_drop": 0,
									"cpu_avg": 70.02469,
									"cpu_le_50_avg": 0,
									"cpu_le_80_avg": 0,
									"cpu_app_avg": 20.510868,
									"cpu_app_le_60_avg": 0,
									"cpu_app_le_80_avg": 0,
									"cpu_temp_avg": 0,
									"mem_use_avg": 0,
									"mem_use_max": 0,
									"mem_swap_avg": 0,
									"mem_swap_max": 0,
									"mem_native_pss_avg": 0,
									"mem_gfx_avg": 0,
									"mem_gl_avg": 0,
									"mem_unknown_avg": 0,
									"mem_xcode_avg": 0,
									"mem_xcode_max": 0,
									"mem_real_avg": 0,
									"mem_real_max": 0,
									"mem_virtual_avg": 0,
									"mem_virtual_max": 0,
									"gpu_use_avg": 0,
									"gpu_clock_avg": 0,
									"gpu_render_avg": 69.9793,
									"gpu_tiler_avg": 20.50709,
									"gpu_device_avg": 20.452662,
									"switch_avg": 0,
									"wakeup_avg": 0,
									"net_send_avg": 19972.244,
									"net_send_sum": 0,
									"net_send_pre_10m_avg": 0,
									"net_recv_avg": 20014.521,
									"net_recv_sum": 0,
									"net_recv_pre_10m_avg": 0,
									"battery_current_avg": 0,
									"battery_power_avg": 0,
									"battery_voltage_avg": 0,
									"cpu_energy": 0,
									"gpu_energy": 0,
									"net_energy": 0,
									"location_energy": 0,
									"display_energy": 0,
									"overhead_energy": 0
								},
								"info": {
									"id": 32,
									"creation_time": "2019-12-20T14:18:35+08:00",
									"modified_time": "2019-12-20T14:18:35+08:00",
									"task_id": 29,
									"name": "test_6",
									"begin_data_id": 29360,
									"end_data_id": 129659
								}
							}
						]
					}
				}
			*/
			task.GET("/label", v1.LabelGetByTaskId)
			// 计算 task summary
			/**
			  * @api {post} /v1/data/task/calsummary 计算 task summary
			  * @apiName TaskCalSummary
			  * @apiGroup Data
			  * @apiParam {int} task_id task id
			  *
			  * @apiParamExample {json} Request-Example:
				{
					"task_id":32
				}
			  * @apiSuccessExample {json} Success-Response:
				HTTP/1.1 200 OK
			    {
					"code": 200,
					"msg": "success",
					"data": {}
				}
			*/
			task.POST("/calsummary", v1.TaskCalSummary)

		}

		//label group
		//GroupV1.POST("/data/task/")
		label := GroupV1.Group("/data/label")
		{
			/**
			  * @api {post} /v1/data/label 创建 lable
			  * @apiName LabelInit
			  * @apiGroup Data
			  *
			  * @apiParam {int} task_id task'id has this label
			  * @apiParam {string} label_name label name
			  * @apiSuccessExample {json} Success-Response:
				HTTP/1.1 200 OK
			    {
					"code": 200,
					"msg": "success",
					"data": {
						"label_id": 13,
						"label_name": "test2"
					}
				}
			*/
			label.POST("/", v1.LabelInit)
			// 通过

			/**
			  * @api {get} /v1/data/label/detail 获取 task 详细信息
			  * @apiName LabelDetailGet
			  * @apiGroup Data
			  *
			  * @apiParam {int} label_id label's id
			  * @apiSuccessExample {json} Success-Response:
				HTTP/1.1 200 OK
			    {
					"code": 0,
					"msg": "success",
					"data": {
						"summary": {
							"id": 13,
							"creation_time": "2019-12-20T12:23:58+08:00",
							"modified_time": "2019-12-24T14:35:14+08:00",
							"label_id": 31,
							"jank_avg": 0,
							"jank_per_hour": 0,
							"big_jank_avg": 0,
							"big_jank_per_hour": 0,
							"ftime_avg": "2019-12-20T12:23:58+08:00",
							"ftime_ge_100": "2019-12-20T12:23:58+08:00",
							"ftime_delta": "2019-12-20T12:23:58+08:00",
							"fps_avg": 30.3075,
							"fps_var": 0,
							"fps_ge_18": 0,
							"fps_ge_25": 0,
							"fps_drop": 0,
							"cpu_avg": 70.0048,
							"cpu_le_50_avg": 0,
							"cpu_le_80_avg": 0,
							"cpu_app_avg": 20.5544,
							"cpu_app_le_60_avg": 0,
							"cpu_app_le_80_avg": 0,
							"cpu_temp_avg": 0,
							"mem_use_avg": 0,
							"mem_use_max": 0,
							"mem_swap_avg": 0,
							"mem_swap_max": 0,
							"mem_native_pss_avg": 0,
							"mem_gfx_avg": 0,
							"mem_gl_avg": 0,
							"mem_unknown_avg": 0,
							"mem_xcode_avg": 0,
							"mem_xcode_max": 0,
							"mem_real_avg": 0,
							"mem_real_max": 0,
							"mem_virtual_avg": 0,
							"mem_virtual_max": 0,
							"gpu_use_avg": 0,
							"gpu_clock_avg": 0,
							"gpu_render_avg": 69.943,
							"gpu_tiler_avg": 20.3808,
							"gpu_device_avg": 20.5829,
							"switch_avg": 0,
							"wakeup_avg": 0,
							"net_send_avg": 19935.879,
							"net_send_sum": 0,
							"net_send_pre_10m_avg": 0,
							"net_recv_avg": 19856.094,
							"net_recv_sum": 0,
							"net_recv_pre_10m_avg": 0,
							"battery_current_avg": 0,
							"battery_power_avg": 0,
							"battery_voltage_avg": 0,
							"cpu_energy": 0,
							"gpu_energy": 0,
							"net_energy": 0,
							"location_energy": 0,
							"display_energy": 0,
							"overhead_energy": 0
						},
						"info": {
							"id": 31,
							"creation_time": "2019-12-20T12:23:09+08:00",
							"modified_time": "2019-12-20T12:23:09+08:00",
							"task_id": 33,
							"name": "2019-12-20-12:23:09",
							"begin_data_id": 5377,
							"end_data_id": 15376
						}
					}
				}
			*/
			label.GET("/", v1.LabelGetDetailById)
			// 计算 label summary
			/**
			  * @api {post} /v1/data/label/calsummary 计算 label summary
			  * @apiName LabelSummaryCal
			  * @apiGroup Data
			  * @apiParam {int} label_id label's id
			  *
			  * @apiParamExample {json} Request-Example:
				{
					"label_id":32
				}
			  * @apiSuccessExample {json} Success-Response:
				HTTP/1.1 200 OK
			    {
					"code": 200,
					"msg": "success",
					"data": {}
				}
			*/
			label.POST("/calsummary", v1.LabelCalLabelSummary)
		}
	}

	GroupHealth := e.Group("/v1/health")
	{
		/**
		   * @api {get} /v1/health 健康监测
		   * @apiName Health_Check
		   * @apiGroup Data
		   *
		   * @apiSuccessExample {json} Success-Response:
		     HTTP/1.1 200 OK
		  	{
		  		"result": "health"
		  	}
		*/
		GroupHealth.GET("", v1.HealthCheck)
	}
}
