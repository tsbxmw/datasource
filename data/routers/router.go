package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/tsbxmw/datasource/common/middleware"
	v1 "github.com/tsbxmw/datasource/data/routers/v1"
)

func InitRouter(e *gin.Engine) {
	GroupV1 := e.Group("/v1", middleware.AuthMiddleware())
	{
		GroupV1.GET("/data/", v1.DataInit)
		GroupV1.POST("/data/upload", v1.DataUpload)

		// task group
		GroupV1.POST("/data/task", v1.TaskInit)
		GroupV1.POST("/data/task/list", v1.TaskGetList)
		GroupV1.POST("/data/task/detail", v1.TaskGetDetail)
		GroupV1.POST("/data/task/report", v1.TaskGetReport)
		GroupV1.POST("/data/task/app", v1.AppInit)
		GroupV1.GET("/data/task/app", v1.AppGetById)
		GroupV1.POST("/data/task/device", v1.DeviceInit)
		GroupV1.GET("/data/task/device", v1.DeviceGetById)

		//label group
		//GroupV1.POST("/data/task/")
		GroupV1.POST("/data/label", v1.LabelInit)
		GroupV1.GET("/data/label", v1.LabelGetDetailById)
	}

	GroupHealth := e.Group("/v1/health")
	{
		GroupHealth.GET("", v1.HealthCheck)
	}
}
