package http

import (
    "datasource/middleware"
    "datasource/routers"
    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
    "github.com/sirupsen/logrus"
)

type (
    HttpServer struct {
        DB            *gorm.DB
        SvcName       string
        Address       string
        Port          int
        GrpcPort      string
        DbUri         string
        ConsulAddr    string
        JaegerAddr    string
        RedisHost     string
        RedisPort     string
        RedisPassword string
        RedisDB       int
    }
)

func (h HttpServer) Serve() {
    logrus.Info("test")
    engin := gin.New()
    middleware.LoggerInit(engin, "/Users/mengwei/workspace/mine/go_data/src/datasource/log/datasource.log")

    routers.InitRouter(engin)
    engin.Run()
}

func getting(context *gin.Context) {
    context.Writer.Write([]byte("test"))

}
