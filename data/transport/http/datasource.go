package http

import (
    "datasource/common"
    "datasource/common/consul"
    "datasource/common/middleware"
    "datasource/common/mq"
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
    common.InitDB(httpServer.DbUri)
    // init logger
    middleware.LoggerInit(engin, "./log/datasource.log")
    common.InitRedisPool("tcp", httpServer.RedisHost+":"+httpServer.RedisPort, httpServer.RedisPassword, httpServer.RedisDB)
    mq.MQInit(httpServer.MqUri)
    // init exception
    middleware.ExceptionInit(engin)
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

    if err := engin.Run("0.0.0.0:" + strconv.Itoa(httpServer.Port)); err != nil {
        panic(err)
    }
}

func (httpServer HttpServer) Init(config *common.ServiceConfig) (common.HttpServer) {
    httpServer.SvcName = config.ServiceName
    httpServer.Address = config.HttpAddr
    httpServer.Port = config.Port
    httpServer.DbUri = config.DbUri
    httpServer.ConsulAddr = config.ConsulAddr
    httpServer.JaegerAddr = config.JaegerAddr
    httpServer.ConsulPort = config.ConsulPort
    httpServer.RedisDB = config.RedisDB
    httpServer.RedisHost = config.RedisHost
    httpServer.RedisPassword = config.RedisPassword
    httpServer.RedisPort = config.RedisPort
    httpServer.MqUri = config.MqUri
    return httpServer
}
