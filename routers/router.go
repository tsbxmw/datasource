package routers

import (
    v1 "datasource/routers/v1"
    "github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
    GroupV1 := r.Group("/v1")
    {
        GroupV1.Any("/test", v1.Test)
    }
}