package http

import (
    "github.com/tsbxmw/datasource/common"
    "github.com/tsbxmw/datasource/common/consul"
    "github.com/tsbxmw/datasource/common/handler"
    "github.com/tsbxmw/datasource/common/middleware"
    "github.com/tsbxmw/datasource/common/mq"
    "github.com/tsbxmw/datasource/data/routers"
    "github.com/tsbxmw/datasource/data/workers"
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
    handler.HandlerInit(engin)

    common.InitRedisPool("tcp", httpServer.RedisHost+":"+httpServer.RedisPort, httpServer.RedisPassword, httpServer.RedisDB)
    mq.MQInit(httpServer.MqUri)
    defer mq.MqConn.Close()
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

func (httpServer HttpServer) ServeWorker() {
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

    workers.WorkerInit(httpServer.MqUri)

    if err := engin.Run("0.0.0.0:" + strconv.Itoa(httpServer.Port+1)); err != nil {
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
