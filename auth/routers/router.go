package routers

import (
	v1 "datasource/auth/routers/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	GroupV1 := r.Group("/v1")
	{
		GroupV1.GET("/auth/token", v1.Token)
		GroupV1.GET("/auth/token/refresh", v1.RefreshToken)
		GroupV1.GET("/health", v1.HealthCheck)
	}
}
