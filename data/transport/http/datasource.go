package http

import (
    "datasource/common"
    "datasource/common/consul"
    "datasource/common/middleware"
    "datasource/data/models"
    "datasource/data/routers"
    "fmt"
    "github.com/gin-gonic/gin"
    "strconv"
    "time"
)

type (
    HttpServer struct {
        common.HttpServerImpl
    }
)

func (httpServer HttpServer) Serve() {
    fmt.Println("test on httpserver", httpServer.SvcName)
    engin := gin.New()
    // init logger
    middleware.LoggerInit(engin, "./log/datasource.log")
    // init router
    routers.InitRouter(engin)
    // init consul
    consulRegister := consul.ConsulRegister{
        Address:                        httpServer.Address,
        Port:                           httpServer.Port,
        ConsulAddress:                  httpServer.ConsulAddr,
        ConsulPort:                     httpServer.ConsulPort,
        Service:                        httpServer.SvcName,
        Tag:                            []string{httpServer.SvcName},
        DeregisterCriticalServiceAfter: time.Second * 10,
        Interval:                       time.Second * 5,
    }

    consulRegister.RegisterHTTP()

    models.InitDB(httpServer.DbUri)

    if err := engin.Run(httpServer.Address + ":" + strconv.Itoa(httpServer.Port)); err != nil {
        panic(err)
    }
}

func (httpServer HttpServer) Init(config *common.ServiceConfig) (common.HttpServer){
    httpServer.SvcName = config.ServiceName
    httpServer.Address = config.HttpAddr
    httpServer.Port = config.Port
    httpServer.DbUri = config.DbUri
    httpServer.ConsulAddr = config.ConsulAddr
    httpServer.JaegerAddr = config.JaegerAddr
    httpServer.ConsulPort = config.ConsulPort
    return httpServer
}
