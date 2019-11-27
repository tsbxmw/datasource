package http

import (
	"datasource/auth/routers"
	"datasource/common"
	"datasource/common/consul"
	"datasource/common/handler"
	"datasource/common/middleware"
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

func (httpServer HttpServer) ServeWorker() {

}

func (httpServer HttpServer) Serve() {
	fmt.Println("test on httpserver", httpServer.SvcName)
	gin.SetMode(gin.ReleaseMode)
	engin := gin.New()
	// init logger
	middleware.LoggerInit(engin, "./log/auth.log")
	// init redis
	common.InitRedisPool("tcp", httpServer.RedisHost+":"+string(httpServer.RedisPort), httpServer.RedisPassword, httpServer.RedisDB)
	// init exception
	middleware.ExceptionInit(engin)
	handler.HandlerInit(engin)
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

	common.InitDB(httpServer.DbUri)
	common.LogrusLogger.Info("serve on " + strconv.Itoa(httpServer.Port))
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
	return httpServer
}
