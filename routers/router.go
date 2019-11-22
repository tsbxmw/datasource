package routers

import (
    v1 "datasource/routers/v1"
    "datasource/transport/http"
    "github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine, h *http.HttpServer) {
    GroupV1 := r.Group("/v1")
    {
        GroupV1.GET("/auth", v1.AuthCheck)
    }
}
