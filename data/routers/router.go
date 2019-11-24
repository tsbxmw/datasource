package routers

import (
    "datasource/common/middleware"
    v1 "datasource/data/routers/v1"
    "github.com/gin-gonic/gin"
)

func InitRouter(e *gin.Engine) {
    GroupV1 := e.Group("/v1", middleware.AuthMiddleware())
    {
        GroupV1.GET("/auth", v1.AuthCheck)
    }

    GroupHealth := e.Group("/v1/health")
    {
        GroupHealth.GET("", v1.HealthCheck)
    }
}
