package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/tsbxmw/datasource/common/middleware"
	v1 "github.com/tsbxmw/datasource/data/routers/v1"
)

func InitRouter(e *gin.Engine) {
	GroupV1 := e.Group("/v1", middleware.AuthMiddleware())
	{
		// data upload
		GroupV1.GET("/data/", v1.DataInit)
		GroupV1.POST("/data/upload", v1.DataUpload)

		// task group
		// 初始化 task
		GroupV1.POST("/data/task", v1.TaskInit)
		// 获取 task 列表
		GroupV1.GET("/data/task/list", v1.TaskGetList)
		// 获取 task 详细信息
		GroupV1.GET("/data/task/detail", v1.TaskGetDetail)
		// 获取 task 报告
		GroupV1.GET("/data/task/report", v1.TaskGetReport)
		// 创建 app
		GroupV1.POST("/data/task/app", v1.AppInit)
		// 通过 id、task_id 获取 app 信息
		GroupV1.GET("/data/task/app", v1.AppGetById)
		// 创建 device
		GroupV1.POST("/data/task/device", v1.DeviceInit)
		// 通过 id、task_id 获取 device 信息
		GroupV1.GET("/data/task/device", v1.DeviceGetById)
		// 通过 task_id 获取 label 信息
		GroupV1.GET("/data/task/label", v1.LabelGetByTaskId)

		//label group
		//GroupV1.POST("/data/task/")
		GroupV1.POST("/data/label", v1.LabelInit)
		// 通过
		GroupV1.GET("/data/label", v1.LabelGetDetailById)
	}

	GroupHealth := e.Group("/v1/health")
	{
		GroupHealth.GET("", v1.HealthCheck)
	}
}
