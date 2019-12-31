package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tsbxmw/datasource/common"
	"github.com/tsbxmw/datasource/common/consul"
	"github.com/tsbxmw/datasource/common/handler"
	"github.com/tsbxmw/datasource/common/middleware"
	"github.com/tsbxmw/datasource/common/mq"
	"github.com/tsbxmw/datasource/data/routers"
	"github.com/tsbxmw/datasource/data/workers"
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
	gin.SetMode(gin.ReleaseMode)
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
	// init header middleware
	middleware.HeaderMiddlewareInit(engin)
	middleware.ResponseMiddlewareInit(engin)
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
		DeregisterCriticalServiceAfter: time.Second * 60,
		Interval:                       time.Second * 60,
	}
	consulRegister.RegisterHTTP()

	if err := engin.Run("0.0.0.0:" + strconv.Itoa(httpServer.Port)); err != nil {
		panic(err)
	}
}

func (httpServer HttpServer) Shutdown() {
	consulName := httpServer.SvcName + "-" + common.LocalIP()
	common.LogrusLogger.Info("Consul Deregister Now ", consulName)
	// init consul
	consulRegister := consul.ConsulRegister{
		Address:                        httpServer.Address,
		Port:                           httpServer.Port,
		ConsulAddress:                  httpServer.ConsulAddr,
		ConsulPort:                     httpServer.ConsulPort,
		Service:                        httpServer.SvcName,
		Tag:                            []string{httpServer.SvcName},
		DeregisterCriticalServiceAfter: time.Second * 60,
		Interval:                       time.Second * 60,
	}
	consulClient := consulRegister.NewConsulClient()
	if err:=consulClient.Agent().ServiceDeregister(consulName); err != nil {
		common.LogrusLogger.Error(err)
		panic(err)
	}
}

func (httpServer HttpServer) ServeWorker() {
	fmt.Println("test on httpserver", httpServer.SvcName)
	engin := gin.New()
	common.InitDB(httpServer.DbUri)
	// init logger
	middleware.LoggerInit(engin, "./log/datasource.log")
	// init redis pool
	common.InitRedisPool("tcp", httpServer.RedisHost+":"+httpServer.RedisPort, httpServer.RedisPassword, httpServer.RedisDB)
	// init mq
	mq.MQInit(httpServer.MqUri)
	// init exception
	middleware.ExceptionInit(engin)
	// init router
	routers.InitRouter(engin)
	// init worker
	workers.WorkerInit(httpServer.MqUri)

	if err := engin.Run("0.0.0.0:" + strconv.Itoa(httpServer.Port+1)); err != nil {
		panic(err)
	}
}

func (httpServer HttpServer) Init(config common.ServiceConfig, configPath string) common.HttpServer {
	configReal := config.ConfigFromFileName(configPath).(common.ServiceConfigImpl)
	httpServer.SvcName = configReal.ServiceName
	httpServer.Address = configReal.HttpAddr
	httpServer.Port = configReal.Port
	httpServer.DbUri = configReal.DbUri
	httpServer.ConsulAddr = configReal.ConsulAddr
	httpServer.JaegerAddr = configReal.JaegerAddr
	httpServer.ConsulPort = configReal.ConsulPort
	httpServer.RedisDB = configReal.RedisDB
	httpServer.RedisHost = configReal.RedisHost
	httpServer.RedisPassword = configReal.RedisPassword
	httpServer.RedisPort = configReal.RedisPort
	httpServer.MqUri = configReal.MqUri
	return httpServer
}
