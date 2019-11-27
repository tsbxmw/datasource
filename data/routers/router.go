package routers

import (
    "datasource/common/middleware"
    v1 "datasource/data/routers/v1"
    "github.com/gin-gonic/gin"
)

func InitRouter(e *gin.Engine) {
    GroupV1 := e.Group("/v1", middleware.AuthMiddleware())
    {
        GroupV1.GET("/data/", v1.DataInit)
        GroupV1.POST("/data/upload", v1.DataUpload)
        GroupV1.POST("/data/task", v1.TaskInit)
        GroupV1.POST("/data/label", v1.LabelInit)
    }

    GroupHealth := e.Group("/v1/health")
    {
        GroupHealth.GET("", v1.HealthCheck)
    }
}
