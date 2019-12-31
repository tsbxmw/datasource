package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/tsbxmw/datasource/auth/routers"
	"github.com/tsbxmw/datasource/common"
	"github.com/tsbxmw/datasource/common/consul"
	"github.com/tsbxmw/datasource/common/handler"
	"github.com/tsbxmw/datasource/common/middleware"
	"strconv"
	"time"
)

type (
	HttpServer struct {
		common.HttpServerImpl
		AppId     string
		AppSecret string
		GrantType string
	}

	ConfigServer struct {
		common.ServiceConfigImpl
		AppId     string
		AppSecret string
		GrantType string
	}
)

func (configServer ConfigServer) ConfigFromFileName(config string) common.ServiceConfig {
	fmt.Println("Config from file : ", config)
	viper.SetConfigFile(config)
	if err := viper.ReadInConfig(); err != nil {
		return configServer
	}
	configServer = ConfigServer{
		ServiceConfigImpl: common.ServiceConfigImpl{}.ConfigFromFileName(config).(common.ServiceConfigImpl),
		AppId:     viper.GetString("app_id"),
		AppSecret: viper.GetString("app_secret"),
		GrantType: viper.GetString("grant_type"),
	}
	return configServer
}

func (httpServer HttpServer) ServeWorker() {

}

func (httpServer HttpServer) Serve() {
	fmt.Println("test on httpserver", httpServer.SvcName)
	gin.SetMode(gin.ReleaseMode)
	engin := gin.New()
	// init logger
	middleware.LoggerInit(engin, "./log/auth.log")
	// init Apidoc
	//middleware.ApidocMiddlewareInit(engin)
	middleware.TracerInit(engin, httpServer.JaegerAddr, httpServer.SvcName)
	defer (*middleware.TracerCloser).Close()
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
		DeregisterCriticalServiceAfter: time.Second * 60,
		Interval:                       time.Second * 60,
	}

	consulRegister.RegisterHTTP()

	common.InitDB(httpServer.DbUri)

	common.LogrusLogger.Info("serve on " + strconv.Itoa(httpServer.Port))
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
	if err := consulClient.Agent().ServiceDeregister(consulName); err != nil {
		common.LogrusLogger.Error(err)
		panic(err)
	}
}

func (httpServer HttpServer) Init(conf common.ServiceConfig, configPath string) common.HttpServer {
	configReal := conf.ConfigFromFileName(configPath).(ConfigServer)
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
	httpServer.AppId = configReal.AppId
	httpServer.AppSecret = configReal.AppSecret
	httpServer.GrantType = configReal.GrantType
	return httpServer
}
