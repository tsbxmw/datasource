package main

import (
    "datasource/common"
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
    "github.com/spf13/viper"
)

type ServiceConfig struct {
    ServiceName string
    LogFile     string
    JaegerAddr  string
    ConsulAddr  string
    HttpAddr    string
    DbUri       string
    Port        int
}

func configFromFileName(config string) (serviceConfig ServiceConfig, err error) {
    viper.SetConfigFile(config)
    if err = viper.ReadInConfig(); err != nil {
        return
    }

    serviceConfig = ServiceConfig{
        ServiceName: viper.GetString("service_name"),
        LogFile:     viper.GetString("log_file"),
        JaegerAddr:  viper.GetString("jaeger_addr"),
        ConsulAddr:  viper.GetString("consul_addr"),
        HttpAddr:    viper.GetString("address"),
        DbUri:       viper.GetString("db_uri"),
        Port:        viper.GetInt("port"),
    }
    return
}


func (serviceConfig ServiceConfig) GetLogger() (logger gin.HandlerFunc, err error) {
    logger, err = common.Logger(serviceConfig.LogFile)
    if err != nil {
        return
    }
    return

}


func (serviceConfig ServiceConfig) GetDB() (db *gorm.DB, err error) {
    db, err = gorm.Open("mysql", serviceConfig.DbUri)
    if err!=nil {
        fmt.Println(err)
        return
    }
    db.SingularTable(true)
    db.Lock(true)
}