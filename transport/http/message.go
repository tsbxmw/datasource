package http

import (
    "datasource/middleware"
    "datasource/routers"
    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
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

func (httpServer HttpServer) Serve() {
    engin := gin.New()
    // init logger
    middleware.LoggerInit(engin, "/Users/mengwei/workspace/mine/go_data/src/datasource/log/datasource.log")
    // init router
    routers.InitRouter(engin)
    // init consul
    //consulRegister := consul.ConsulRegister{
    //    Address:                        httpServer.Address,
    //    Port:                           httpServer.Port,
    //    ConsulAddress:                  httpServer.ConsulAddr,
    //    ConsulPort:                     80,
    //    Service:                        httpServer.SvcName,
    //    Tag:                            []string{httpServer.SvcName},
    //    DeregisterCriticalServiceAfter: time.Second * 10,
    //    Interval:                       time.Second * 5,
    //}
    //
    //consulRegister.Register()

    if err:=engin.Run();err!=nil{
        panic(err)
    }
}

