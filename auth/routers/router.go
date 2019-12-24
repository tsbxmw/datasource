package routers

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/tsbxmw/datasource/auth/routers/v1"
	"github.com/tsbxmw/datasource/common/middleware"
)

func InitRouter(r *gin.Engine) {
	GroupV1 := r.Group("/v1")
	{
		GroupV1.POST("/auth/token", v1.Token)
		GroupV1.GET("/health", v1.HealthCheck)
	}

	GroupV1Auth := r.Group("/v1", middleware.AuthMiddleware())
	{
        GroupV1Auth.POST("/auth/token/refresh", v1.RefreshToken)
	}
}
